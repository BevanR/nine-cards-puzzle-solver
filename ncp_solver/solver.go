package ncp_solver

import (
	"fmt"
)

type Solver struct {
	puzzle       [9]Tile
	solution     [9]int // tile index, keyed by position.
	orientations [9]int // the orientation of each tile, keyed by tile index
	debug        bool
	findAll      bool
	registry     Registry
}

func Solve(puzzle [9]Tile) string {
	s := &Solver{
		debug:        false,
		findAll:      true,
		puzzle:       puzzle,
		solution:     [9]int{-1, -1, -1, -1, -1, -1, -1, -1, -1},
		orientations: [9]int{-1, -1, -1, -1, -1, -1, -1, -1, -1},
		registry:     Registry{},
	}

	s.findTileThatFits(0, 0, s.neighboringEdges(0))

	if len(s.registry) > 0 {
		// End log with an empty line
		s.logger("")
		return s.registry.format()
	} else {
		return "no solution"
	}
}

func (s *Solver) solve() {
	s.findTileThatFits(0, 0, s.neighboringEdges(0))
}

func (s *Solver) findTileThatFits(position, orientation int, neighboringEdges [4]*Edge) bool {
	for tile := range 9 {
		if !s.alreadyPlaced(tile) && s.fits(tile, orientation, neighboringEdges) {
			s.place(tile, position, orientation)

			// Recursion! Move onto next position.
			if s.rotateAndSolve(position + 1) {
				return true
			}

			s.remove(tile, position, orientation)
		}
	}

	return false
}

func (s *Solver) rotateAndSolve(position int) bool {
	if position == 9 {
		s.registry.add(s.solution, s.orientations)
		return !s.findAll
	}

	neighboringEdges := s.neighboringEdges(position)

	s.logger("Position %d:", position)
	s.logNeighbors(neighboringEdges)

	// Try fitting each tile in each orientation. I.e. turn the tile clockwise 90º, four times.
	// The 1st loop checks N edge with S edge of N tile, E edge with W edge of E tile, etc.
	// The 2nd loop checks N edge with W edge of E tile, E edge with N edge of S tile, etc.
	// Etcetera
	for orientation := range 4 {
		if s.findTileThatFits(position, orientation, neighboringEdges) {
			return true
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
		if tile := s.solution[position]; tile >= 0 {
			// Factor in the tile's current orientation.
			orientation := s.orientations[tile]
			o := (4 - orientation + int(side)) % 4
			return &s.puzzle[tile][o]
		}
	}
	return nil
}

func (s *Solver) alreadyPlaced(tile int) bool {
	return s.orientations[tile] >= 0
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
	s.logger("✅ tile %d fits at position %d facing %c", tile+1, position, formatO9n(orientation))
	s.solution[position] = tile
	s.orientations[tile] = orientation
}

func (s *Solver) remove(tile, position, orientation int) {
	s.logger("❌ tile %d does not fit at position %d facing %c", tile+1, position, formatO9n(orientation))
	s.solution[position] = -1
	s.orientations[tile] = -1
}

func (s *Solver) logNeighbors(neighboringEdges [4]*Edge) {
	for i, edge := range neighboringEdges {
		if edge != nil {
			s.logger("Neighbor edge %c: %s", formatO9n(i), edge.String())
		}
	}
}

func (s *Solver) logger(format string, args ...any) {
	if s.debug {
		fmt.Printf(format, args...)
		fmt.Println()
	}
}

func formatO9n(orientation int) uint8 {
	return "NESW"[orientation]
}
