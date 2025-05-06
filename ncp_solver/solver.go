package ncp_solver

import (
	"fmt"
	"strings"
)

type Solver struct {
	size         int
	puzzle       [9]Tile
	positions    [9]int // map of tile index to its position in the current solution
	tiles        [9]int // the reverse map; tile index by position.
	orientations [9]int // the orientation of each tile, by position
	debug        bool
}

func Solve(puzzle [9]Tile) string {
	s := &Solver{
		debug:        false,
		puzzle:       (puzzle),
		size:         9,
		tiles:        [9]int{-1, -1, -1, -1, -1, -1, -1, -1, -1},
		positions:    [9]int{-1, -1, -1, -1, -1, -1, -1, -1, -1},
		orientations: [9]int{-1, -1, -1, -1, -1, -1, -1, -1, -1},
	}

	if s.solve(0) {
		// End log with an empty line
		s.logger("")
		return s.format()
	} else {
		return "no solution"
	}
}

func (s *Solver) solve(position int) bool {
	if position == s.size {
		return true
	}

	neighboringEdges := s.neighboringEdges(position)

	s.logger("Position %d:", position)
	s.logNeighbors(neighboringEdges)

	// Try fitting each tile in each orientation. I.e. turn the tile clockwise 90º, four times.
	// The 1st loop checks N edge with S edge of N tile, E edge with W edge of E tile, etc.
	// The 2nd loop checks N edge with W edge of E tile, E edge with N edge of S tile, etc.
	// Etcetera
	for orientation := range 4 {
		for tile := range s.size {
			if !s.alreadyPlaced(tile) && s.fits(tile, orientation, neighboringEdges) {
				s.place(tile, position, orientation)

				// Recursion! Move onto next position.
				if s.solve(position + 1) {
					return true
				}

				s.remove(tile, position, orientation)
			}
		}
	}

	return false
}

func (s *Solver) neighboringEdges(position int) [4]*Edge {
	x := position % 3
	y := position / 3
	return [4]*Edge{
		s.getEdge(x, y-1, South), // South edge of North tile
		nil,                      // There is no tile on the East side yet.
		nil,                      // There is no tile on the South side yet.
		s.getEdge(x-1, y, East),  // East edge of West tile
	}
}

func (s *Solver) getEdge(x, y int, side Orientation) *Edge {
	if 0 <= x && x < 3 && 0 <= y && y < 3 {
		position := x + y*3
		if index := s.tiles[position]; index >= 0 {
			// Factor in the tile's current orientation.
			orientation := s.orientations[position]
			o := (4 - orientation + int(side)) % 4
			return &s.puzzle[index][o]
		}
	}
	return nil
}

func (s *Solver) alreadyPlaced(tile int) bool {
	return s.positions[tile] >= 0
}

func (s *Solver) fits(tile, orientation int, neighbouringEdges [4]*Edge) bool {
	for i, edge := range s.puzzle[tile] {
		i = (i + orientation) % 4
		otherEdge := neighbouringEdges[i]
		if otherEdge != nil && !s.edgeFits(&edge, otherEdge) {
			return false
		}
	}
	return true
}

func (s *Solver) edgeFits(a, b *Edge) bool {
	return a.Color == b.Color && a.Piece != b.Piece && a.Position != b.Position
}

func (s *Solver) place(tile, position, orientation int) {
	s.logger("✅ tile %d fits at position %d facing %c", tile+1, position, s.orientationStr(orientation))
	s.positions[tile] = position
	s.tiles[position] = tile
	s.orientations[position] = orientation
}

func (s *Solver) remove(tile, position, orientation int) {
	s.logger("❌ tile %d does not fit at position %d facing %c", tile+1, position, s.orientationStr(orientation))
	s.tiles[position] = -1
	s.positions[tile] = -1
	s.orientations[position] = -1
}

func (s *Solver) logNeighbors(neighboringEdges [4]*Edge) {
	for i, edge := range neighboringEdges {
		if edge != nil {
			s.logger("Neighbor edge %c: %s", s.orientationStr(i), edge.String())
		}
	}
}

func (s *Solver) logger(format string, args ...any) {
	if s.debug {
		fmt.Printf(format, args...)
		fmt.Println()
	}
}

func (s *Solver) format() string {
	lines := []string{
		"Tile number by position, including the orientation of North/top edge of tile",
		"",
		"P: T O",
		"––––––",
	}
	for p := range s.size {
		t := s.tiles[p]
		o := s.orientationStr(s.orientations[p])
		lines = append(lines, fmt.Sprintf("%d: %d %c", p, t+1, o))
	}
	return strings.Join(lines, "\n")
}

func (s *Solver) orientationStr(orientation int) uint8 {
	orientations := "NESW"
	return orientations[orientation]
}
