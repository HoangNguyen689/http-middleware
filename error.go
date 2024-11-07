package main

type Code string

const (
	Internal      Code = "Internal"
	Unavailable   Code = "Unavailable"
	Unimplemented Code = "Unimplemented"
	Unauthorized  Code = "Unauthorized"
	NotFound      Code = "NotFound"
	Invalid       Code = "Invalid"
)

type Error struct {
	error
	Code Code
}
