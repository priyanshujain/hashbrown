package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/user"
	"time"

	"github.com/priyanshujain/hashbrown/hashbrown"
	"github.com/urfave/cli"
)

func init() {

	cli.HelpFlag = cli.BoolFlag{Name: "help"}
	cli.BashCompletionFlag = cli.BoolFlag{Name: "compgen", Hidden: true}
	cli.VersionFlag = cli.BoolFlag{Name: "print-version, V"}

	cli.HelpPrinter = func(w io.Writer, templ string, data interface{}) {
		fmt.Fprintf(w, "add -h flag for help\n")
	}
	cli.VersionPrinter = func(c *cli.Context) {
		fmt.Fprintf(c.App.Writer, "version=%s\n", c.App.Version)
	}
}

func main() {

	var service, config string
	var length int

	user, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	app := cli.NewApp()
	//app.EnableBashCompletion = true
	app.Name = "hasbrown"
	app.Usage = "Generate a safe and healthy password"
	app.Version = "0.0.1"
	app.Compiled = time.Now()
	app.Authors = []cli.Author{
		cli.Author{
			Name:  "Priyanshu Jain",
			Email: "priyanshudeveloper@gmail.com",
		},
	}
	app.Copyright = "(c) 2018 Priyanshu Jain"

	app.HideHelp = false
	app.HideVersion = false

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "service, s",
			Usage:       "Service name for generating password  ex. github",
			Destination: &service,
		},
		cli.StringFlag{
			Name:        "config, c",
			Usage:       "Load configuration from `FILE`",
			Value:       user.HomeDir + "/.hashbrown",
			Destination: &config,
		},
		cli.IntFlag{
			Name:        "length, l",
			Usage:       "length of password",
			Value:       10,
			Destination: &length,
		},
	}

	app.CommandNotFound = func(c *cli.Context, command string) {
		fmt.Fprintf(c.App.Writer, "Thar be no %q here.\n", command)
	}

	app.Action = func(c *cli.Context) error {
		if c.NumFlags() > 0 {
			_, err := os.Stat(config)
			if err == nil {
				fmt.Printf("file %s exists", config)
				hashbrown.Generate(config, service, user.Uid, length)
			} else if os.IsNotExist(err) {
				hashbrown.Setup(config)
				hashbrown.Generate(config, service, user.Uid, length)
			} else {
				fmt.Printf("file %s stat error: %v\n", config, err)
			}
		} else {
			cli.ShowAppHelp(c)
		}
		return nil
	}

	app.Run(os.Args)
}
