package main

import "fmt"

// TwilightBus drops passengers to the 5th dimension
type TwilightBus struct {
	passengers []string
}

// New builds and returns a pointer to a TwilightBus
func New(passengers []string) *TwilightBus {
	bus := TwilightBus{}
	bus.passengers = passengers
	return &bus
}

// Pick picks a named passenger
func (bus *TwilightBus) Pick(name string) {
	bus.passengers = append(bus.passengers, name)
}

// find needle in slice, returns index, found bool
func find(slice []string, needle string) (int, bool) {
	for i, item := range slice {
		if item == needle {
			return i, true
		}
	}
	return -1, false
}

// remove item from slice, preserving order
// from GOPL, sec. 4.2.2, p. 93
func remove(slice []string, i int) []string {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}

// Drop drops a named passenger
func (bus *TwilightBus) Drop(name string) {
	i, found := find(bus.passengers, name)
	if found {
		bus.passengers = remove(bus.passengers, i)
	}
}

func main() {
	team := []string{"Su", "Tina", "Maya", "Diana", "Pri"}
	fmt.Printf("team: %#v\n", team)
	bus := New(team)
	fmt.Printf("bus: %#v\n", bus)
	fmt.Println("# bus drops Tina")
	bus.Drop("Tina")
	fmt.Printf("bus: %#v\n", bus)
	fmt.Printf("team: %#v\n", team)
}
