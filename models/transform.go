package models

type Transform struct {
	Position Position
	Heading  Heading
}

type Position struct {
	X float32
	Y float32
}

type Heading struct {
	Angle float32
}

type Radius struct {
	R float32
}