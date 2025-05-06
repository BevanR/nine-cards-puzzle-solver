package ncp_solver

import "fmt"

// type Puzzle [4]Tile
type Puzzle [9]Tile

type Tile [4]Edge

type Orientation int

const (
	North Orientation = iota
	East
	South
	West
)

type Edge struct {
	Position
	Color
	Piece
}

func (e Edge) String() string {
	pos := "First"
	if e.Position == Second {
		pos = "Second"
	}
	c := "White"
	if e.Color == Black {
		c = "Black"
	}
	p := "Top"
	if e.Piece == Bottom {
		p = "Bottom"
	}

	return fmt.Sprintf("%s %s %s", c, pos, p)
}

type Position int8

const (
	First Position = iota
	Second
)

type Color int8

const (
	White Color = iota
	Black
)

type Piece int8

const (
	Top Piece = iota
	Bottom
)
