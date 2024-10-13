package main

import (
	"fmt"
	"os"

	"github.com/ExocoreNetwork/exocore-sdk/cmd/exokey/generate"
	"github.com/ExocoreNetwork/exocore-sdk/cmd/exokey/store"
	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.NewApp()
	app.Name = "exokey"
	app.Description = "Exocore batch keys manager"
	app.Commands = []*cli.Command{
		generate.Command,
		store.Command,
	}

	app.Usage = "Used to manage batch keys for testing"

	if err := app.Run(os.Args); err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

}
