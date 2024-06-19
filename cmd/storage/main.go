package main

import (
	"fmt"

	"github.com/kirmorozov1992/storage/internal/storage"
)

func main() {
	st := storage.NewStorage()

	fmt.Println("it works", st)
}
