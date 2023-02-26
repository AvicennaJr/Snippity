package main

import (
	"database/sql"
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

type SnippetModel struct {
	DB *sql.DB
}

func main() {
	fmt.Println("vim-go")
}
