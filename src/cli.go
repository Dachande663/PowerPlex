package main

import (
	"encoding/json"
	"github.com/codegangsta/cli"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

// Launch the command-line interface
func startCli() {

	app := cli.NewApp()
	app.Name = "powerplex"
	app.Usage = "Unlock the power of your Plex library"
	app.Version = "0.1.0 - Keyser Soze"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "plex-host",
			Value: "localhost",
			Usage: "Plex Media server hostname",
		},
		cli.IntFlag{
			Name:  "plex-port",
			Value: 32400,
			Usage: "Plex Media server port",
		},
	}

	// Default to opening the web interface
	// app.Action = webCommand

	app.Commands = []cli.Command{
		{
			Name:   "export",
			Usage:  "Export Plex library data",
			Action: exportCommand,
		},
		{
			Name:   "web",
			Usage:  "Launch the web interface",
			Action: webCommand,
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "disable-launch",
					Usage: "Disable opening a web browser on start",
				},
				cli.IntFlag{
					Name:  "app-port",
					Value: 32432,
					Usage: "Local PowerPlex port",
				},
			},
		},
	}

	app.Run(os.Args)

}

// Begin the exporter
func exportCommand(c *cli.Context) {

	exporter := makeExporter()
	exporter.Url = c.GlobalString("plex-url")
	exporter.Reporter = func(msg string) {
		println(msg)
	}

	library := exporter.Export()
	j, _ := json.MarshalIndent(library, "", "  ")

	println(string(j))

	cwd, _ := os.Getwd()
	dir, _ := filepath.Abs(cwd)
	file := dir + "/data.json"

	println(file)

	ioutil.WriteFile(file, j, 0644)

}

// Begin the web interface
func webCommand(c *cli.Context) {

	url := "http://localhost:" + c.String("app-port")

	if !c.Bool("disable-launch") {
		switch runtime.GOOS {
		case "linux":
			exec.Command("xdg-open", url).Start()
		case "windows", "darwin":
			exec.Command("open", url).Start()
		}
	}

	println("Now listening on " + url)
	startHttp(":" + c.String("app-port"))

}
