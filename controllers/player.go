package controllers

import (
	"time"
)

type PlayerController struct {
	eventQueue EventQueue
}

func (c *PlayerController) Update (delta time.Duration, ctx interface{})  {
	//
}