package ncp_solver

import (
	"fmt"
	"strings"
)

type Solution struct {
	solution     [9]int
	orientations [9]int
}

type Registry map[Solution]bool

func (r Registry) add(solution, orientations [9]int) {
	key := Solution{solution, orientations}
	_, ok := r[key]
	if !ok {
		r[key] = true
		for range 3 {
			// Deduplicate solutions.
			// TODO Rotate 90º and get new key.
			// Once we're confident that each unique solution is found exactly four times, simplify the algorithm;
			// Never rotate position 0.
			key = key.rotate()
			r[key] = false
		}
	}
}

func (r Registry) format() string {
	solutions := []string{}
	for solution, unique := range r {
		if unique {
			solutions = append(solutions, solution.format())
		}
	}
	result := []string{
		fmt.Sprintf("Solutions found: %d", len(solutions)),
		"",
		"Tile number by position, including the orientation of North/top edge of tile",
		"",
		"P: T O",
	}
	result = append(result, solutions...)
	return strings.Join(result, "\n")
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
