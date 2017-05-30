package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "initial"
	app.Usage = "just here for bootstrp"
	app.Action = func(c *cli.Context) error {
		fmt.Println("bootstrap cli")
		return nil
	}

	app.Run(os.Args)
}
