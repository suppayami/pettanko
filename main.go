package main

import (
	"flag"
	"os"

	"github.com/suppayami/pettanko/handlers"
)

func main() {
	addr := flag.String("p", ":3000", "Listening address HOST:PORT")
	help := flag.Bool("h", false, "Help")

	flag.Parse()

	if *help {
		flag.PrintDefaults()
		os.Exit(1)
	}

	handlers.SetupRoute()
	handlers.Serve(*addr)
}
