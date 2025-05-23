package ncp_solver

import (
	"fmt"
	"strings"
)

type Solution struct {
	solution     [9]int
	orientations [9]int
}

type Registry struct {
	solutions   []Solution
	comparisons int
	placements  int
	removals    int
}

func (r *Registry) add(solution, orientations [9]int) {
	r.solutions = append(r.solutions, Solution{solution, orientations})
}

func (r *Registry) format() string {
	result := []string{
		fmt.Sprintf("solutions: %d", len(r.solutions)),
		fmt.Sprintf("comparisons: %d", r.comparisons),
		fmt.Sprintf("placements: %d", r.placements),
		fmt.Sprintf("removals: %d", r.removals),
		"",
		"Tile number by position, including the orientation of North/top edge of tile",
		"",
		"P: T O",
	}

	for _, s := range r.solutions {
		result = append(result, s.format())
	}

	return strings.Join(result, "\n")
}

func (r *Registry) Count() int {
	return len(r.solutions)
}

func (s *Solution) format() string {
	lines := []string{
		"––––––",
	}
	for position, tile := range s.solution {
		orientation := formatO9n(s.orientations[tile])
		lines = append(lines, fmt.Sprintf("%d: %d %c", position, tile+1, orientation))
	}
	return strings.Join(lines, "\n")
}

func (s *Solution) rotate() Solution {
	result := Solution{
		solution: [9]int{
			s.solution[6], s.solution[3], s.solution[0],
			s.solution[7], s.solution[4], s.solution[1],
			s.solution[8], s.solution[5], s.solution[2],
		},
	}

	for tile, orientation := range s.orientations {
		result.orientations[tile] = (orientation + 1) % 4
	}

	return result
}
