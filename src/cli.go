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
	app.Name = "plex"
	app.Usage = "Explore your Plex library"
	app.Version = "0.1.0 - Keyser Soze"
	app.EnableBashCompletion = true

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "plex-url",
			Value: "http://localhost:32400",
			Usage: "The URL to your Plex Media server",
		},
	}

	app.Action = webCommand

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

	if false {
		switch runtime.GOOS {
		case "linux":
			exec.Command("xdg-open", "http://localhost:8000").Start()
		case "windows", "darwin":
			exec.Command("open", "http://localhost:8000").Start()
		}
	}

	println("Now listening on localhost:8000")
	startHttp()

}
