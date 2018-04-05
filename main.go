package main

import "github.com/suppayami/pettanko.go/handlers"

func main() {
	handlers.SetupRoute()
	handlers.Serve(":3000")
}
