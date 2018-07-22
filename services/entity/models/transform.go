package models

import "github.com/ataboo/borealengine/mathies"

type Transform struct {
	Position Position `json:"position"`
	Heading  Heading `json:"heading"`
}

type Position mathies.Vector2

type Velocity mathies.Vector2

type Heading float32

type Radius float32