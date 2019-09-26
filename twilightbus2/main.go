package main

import "fmt"

// TwilightBus makes passengers vanish
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

// remove item from slice, move last into gap
// from GOPL, sec. 4.2.2, p. 93
func remove(slice []string, i int) []string {
	slice[i] = slice[len(slice)-1]
	return slice[:len(slice)-1]
}

// Drop drops a named passenger
func (bus *TwilightBus) Drop(name string) {
	foundPos := -1
	for i, current := range bus.passengers {
		if current == name {
			foundPos = i
			break
		}
	}
	if foundPos >= 0 {
		bus.passengers = remove(bus.passengers, foundPos)
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
