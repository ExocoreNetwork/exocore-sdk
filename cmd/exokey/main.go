package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "exokey"
	app.Description = "Exocore batch keys manager"
	app.Commands = []*cli.Command{
		commandGenerate,
		commandStore,
	}

	app.Usage = "Used to manage batch keys for testing"

	if err := app.Run(os.Args); err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

}
