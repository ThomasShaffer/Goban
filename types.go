package main

const margin = 3

type status int

const (
	todo status = iota
	doing
	did
	footer
)

func StatusToString(s status) string {
    switch s {
    case 0: 
        return "todo"
    case 1: 
        return "doing"
    case 2: 
        return "did"
    case 3: 
        return "footer"
    default:
        return "OH NO UWU"
    }
}

type todoModel struct {
	project string
	todo    []map[string]string
	doing   []map[string]string
	did     []map[string]string
}
