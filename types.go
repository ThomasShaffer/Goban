package main

const margin = 3

type status int

const (
	todo status = iota
	doing
	did
	description
)

type todoModel struct {
	project string
	todo    []map[string]string
	doing   []map[string]string
	did     []map[string]string
}
