package main

import (
	"fmt"

	"github.com/rokoga/filas-backend/web"
)

func main() {

	done := make(chan string)
	go web.Run(done)

	fmt.Print(<-done)
}
