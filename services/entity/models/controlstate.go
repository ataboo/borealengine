package models

type ControlState interface {
	Update(interface{}) bool
}