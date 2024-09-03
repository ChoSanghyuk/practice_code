package main

import (
	"fmt"
	"time"
)

func main() {

	a := time.Now()
	b := a.Add(time.Duration(-10) * time.Minute)
	fmt.Println(b)
}
