package main

import "github.com/suppayami/pettanko/handlers"

func main() {
	handlers.SetupRoute()
	handlers.Serve(":3000")
}
