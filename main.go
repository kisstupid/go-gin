package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/kisstupid/go-gin/configs"
	"github.com/kisstupid/go-gin/internal/database"
	"github.com/kisstupid/go-gin/internal/server"
)

func main() {
	environment := flag.String("e", "development", "")
	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}
	flag.Parse()

	configs.Init(*environment)
	database.Init()
	server.Init()
}
