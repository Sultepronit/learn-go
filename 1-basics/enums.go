package main

import "fmt"

// Go doesn’t have an enum type as a distinct language feature, but enums are simple to implement using existing language idioms.

// Our enum type ServerState has an underlying int type.
type ServerState int

// The possible values for ServerState are defined as constants. The special keyword iota generates successive constant values automatically; in this case 0, 1, 2 and so on.
const (
	StateIdle ServerState = iota
	StateConnected
	StateError
	StateRetrying
)

var stateName = map[ServerState]string{
	StateIdle:      "idle",
	StateConnected: "connected",
	StateError:     "error",
	StateRetrying:  "retrying",
}

// method String() for stateName map, that return string
func (ss ServerState) String() string {
	return stateName[ss]
}

func main() {
	ns := transition(StateIdle)
	fmt.Println(ns)

	ns2 := transition(ns)
	fmt.Println(ns2)
}

func transition(s ServerState) ServerState {
	switch s {
	case StateIdle:
		return StateConnected
	case StateConnected, StateRetrying:
		return StateIdle
	case StateError:
		return StateError
	default:
		panic(fmt.Errorf("unknown state: %s", s))
	}
}
