package main

import (
	"fmt"
	"strings"
)

type Position int

const (
	Up Position = iota
	Down
)

func trianglePosition(position Position) {
	const end = 10
	const cellWidth = 10

	switch position {
	case Up:
		for i := 1; i < end; i++ {
			for j := 1; j < i; j++ {
				fmt.Printf("%*s", cellWidth, "")
			}

			for j := i; j < end; j++ {
				expression := fmt.Sprintf("%d×%d=%d", i, j, i*j)
				fmt.Printf("%-*s", cellWidth, expression)
			}

			fmt.Println()
		}

	case Down:
		for i := 1; i < end; i++ {
			for j := 1; j <= i; j++ {
				expression := fmt.Sprintf("%d×%d=%d", j, i, i*j)
				fmt.Printf("%-*s", cellWidth, expression)
			}

			fmt.Println()
		}
	}
}

func main() {
	trianglePosition(Down)
	fmt.Println(strings.Repeat("=", 90))
	trianglePosition(Up)
}
