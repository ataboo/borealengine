package models

type Species struct {
	Base   BaseStat
	Attack AttackCollider
}

type BaseStat struct {
	Speed float32
	Size  Radius
}

type AttackCollider struct {
	Damage float32
	Offset Position
	Radius Radius
}
