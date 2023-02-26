package main

import (
	"fmt"

	"time"
)

type Snippet struct {
	ID      int
	Title   string
	Content string
	Create  time.Time
	Expires time.Time
}

func main() {
	fmt.Println("vim-go")
}
