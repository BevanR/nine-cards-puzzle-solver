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
	found, ok := r[key]
	if ok && found {
		panic(fmt.Sprintf("solution already exists: %v", key))
	} else {
		r[key] = true
		if !ok {
			for range 3 {
				// Deduplicate solutions.
				// TODO Rotate 90º and get new key.
				// Once we're confident that each unique solution is found exactly four times, simplify the algorithm;
				// Never rotate position 0.
				r[key] = false
			}
		}
	}
}

func (r Registry) format() string {
	lines := []string{
		fmt.Sprintf("Solutions found: %d", len(r)),
		"",
		"Tile number by position, including the orientation of North/top edge of tile",
		"",
		"P: T O",
	}
	for solution := range r {
		lines = append(lines, solution.format())
	}
	return strings.Join(lines, "\n")
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
