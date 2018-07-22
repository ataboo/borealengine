package models

type BehaviorNode interface {

}

type NPC struct {
	Creature *Creature
	Behavior BehaviorNode
}