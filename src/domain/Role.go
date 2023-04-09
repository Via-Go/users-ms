package domain

type Role int

const (
	Administrator Role = iota
	Moderator
	Common
)
