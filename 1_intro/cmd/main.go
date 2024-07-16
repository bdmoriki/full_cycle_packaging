package main

import (
	"fmt"

	"github.com/bdmoriki/full_cycle_packaging/1_intro/math"
)

func main() {
	m := math.Math{A: 1, B: 2}
	fmt.Println(m.Add())
}
