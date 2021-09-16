package command

import (
	"fmt"
	"kubectl-gui/config"
	"kubectl-gui/server"
	"log"
	"os"

	flag "github.com/spf13/pflag"
)

const version = "1.0.0"

func init() {
	log.SetPrefix(fmt.Sprintf("[%d]-kubectl-ui ", os.Getpid()))
}

func printVersion() {
	fmt.Printf("%s\n", version)
}

func usage() {
	usage := `
Kubectlui

Usage: kubectlui <command> ...

Command: version | start

verison
....................
Gets the verison of kubectl ui

Start
....................
Starts kubectl ui

--port=<port>  Specify a port for kubectlui to run on
defaults to 8080
Example:
kubectlui start --port=3553

--filepath=<filepath> Place to look for kubernetes yaml files
defaults to $PATH/<kubectlUi>/example
Example:
kubectlui start --filepath=/Users/phirmware/backend-app

--version Show kubectlui version
Example:
kubectlui --version

	`

	fmt.Printf("%s\n", usage)
}

func Execute() {
	port := flag.String("port", "", "The port to run the application")
	configPath := flag.String("filepath", "", "The path kubectlui looks for your kubernetes config files")
	version := flag.BoolP("version", "v", false, "Show version")
	flag.Usage = usage

	flag.Parse()

	if len(flag.Args()) == 0 {
		if *version {
			printVersion()
			os.Exit(0)
		}

		flag.Usage()
		os.Exit(1)
	}

	cmd := flag.Arg(0)
	cfg := getConfig()

	cfg.ReplaceDefault(config.ApplicationPort, *port)
	cfg.ReplaceDefault(config.FilePath, *configPath)

	switch cmd {
	case "version":
		printVersion()
	case "start":
		startServer(cfg)
	default:
		flag.Usage()
		os.Exit(1)
	}
}

func getConfig() *config.KubectlUI {
	cfg := config.Load()
	return cfg
}

func startServer(cfg *config.KubectlUI) {
	server.Start(cfg)
}
