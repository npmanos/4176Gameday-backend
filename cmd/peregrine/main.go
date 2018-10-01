package main

import (
	"flag"
	"fmt"
	"strconv"
	"time"

	"github.com/Pigmice2733/peregrine-backend/internal/store"
	"github.com/Pigmice2733/peregrine-backend/internal/tba"

	"github.com/Pigmice2733/peregrine-backend/internal/config"
	"github.com/Pigmice2733/peregrine-backend/internal/server"
)

func main() {
	var basePath = flag.String("basePath", ".", "Path to the etc directory where the config file is.")

	flag.Parse()

	c, err := config.Open(*basePath)
	if err != nil {
		fmt.Printf("Error: opening config: %v\n", err)
		return
	}

	tba := tba.Service{
		URL:    c.TBA.URL,
		APIKey: c.TBA.APIKey,
	}

	store, err := store.New(c.Database)
	if err != nil {
		fmt.Printf("Error: unable to connect to postgres server: %v\n", err)
		return
	}

	year, err := strconv.Atoi(c.Server.Year)
	if err != nil {
		year = time.Now().Year()
	}

	server := server.New(
		tba,
		store,
		c.Server.HTTPAddress,
		c.Server.HTTPSAddress,
		c.Server.CertFile,
		c.Server.KeyFile,
		c.Server.Origin,
		year,
	)

	if err := server.Run(); err != nil {
		fmt.Printf("Error: server.Run: %v\n", err)
	}
}
