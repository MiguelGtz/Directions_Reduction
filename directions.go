package main

import (
	"fmt"
)

func dirReduc(arr []string) []string {
	for true {
		aux := len(arr)
		for i := 0; i<len(arr)-1; i++ {
			switch arr[i]{
			case "NORTH":
				if arr[i+1] == "SOUTH" {
					arr = append(arr[:i], arr[i+2:]...)
				}
			case "SOUTH":
				if arr[i+1] == "NORTH" {
					arr = append(arr[:i], arr[i+2:]...)
				}
			case "EAST":
				if arr[i+1] == "WEST" {
					arr = append(arr[:i], arr[i+2:]...)
				}
			case "WEST":
				if arr[i+1] == "EAST" {
					arr = append(arr[:i], arr[i+2:]...)
				}
			}
		}
		if aux == len(arr) {
			break
		}
	}
	return arr
}

func main() {
	fmt.Println(dirReduc([]string{"NORTH", "SOUTH", "EAST", "WEST"}))
	fmt.Println(dirReduc([]string{"NORTH", "SOUTH", "SOUTH", "EAST", "WEST", "NORTH", "WEST"}))
	fmt.Println(dirReduc([]string{"NORTH", "WEST", "SOUTH", "EAST"}))
	fmt.Println(dirReduc([]string{"NORTH", "SOUTH", "SOUTH", "EAST", "WEST", "NORTH", "NORTH"}))
}
