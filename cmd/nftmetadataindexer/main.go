// cmd/nftmetadataindexer/main.go
package main

import (
	"flag"
	"log"
	"os"
	
	"nftmetadataindexer/internal/nftmetadataindexer"
)

func main() {
	verbose := flag.Bool("verbose", false, "Enable verbose logging")
	flag.Parse()
	
	app := nftmetadataindexer.NewApp(*verbose)
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
