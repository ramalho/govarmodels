package main

import "fmt"

func Example() {
	team := []string{"Su", "Tina", "Mia", "Dea", "Pri"}
	fmt.Printf("team:\n  %#v\n", team)
	bus := New(team)
	fmt.Printf("bus:\n  %#v\n", bus)
	fmt.Println("# bus drops Tina")
	bus.Drop("Tina")
	fmt.Printf("bus:\n  %#v\n", bus)
	fmt.Printf("team:\n  %#v\n", team)
	// Output:
	// team:
	//   []string{"Su", "Tina", "Mia", "Dea", "Pri"}
	// bus:
	//   &main.Bus{passengers:[]string{"Su", "Tina", "Mia", "Dea", "Pri"}}
	// # bus drops Tina
	// bus:
	//   &main.Bus{passengers:[]string{"Su", "Mia", "Dea", "Pri"}}
	// team:
	//   []string{"Su", "Tina", "Mia", "Dea", "Pri"}
}
