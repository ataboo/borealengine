package models

import "github.com/ataboo/borealengine/mathies"

type Transform struct {
	Position Position
	Heading  Heading
}

type Position mathies.Vector2

type Velocity mathies.Vector2

type Heading struct {
	Angle float32
}

type Radius struct {
	R float32
}