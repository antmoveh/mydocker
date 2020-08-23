package main

import (
	"fmt"
	"github.com/urfave/cli"
	log "github.com/sirupsen/logrus"
	"mydocker/container"
)

var runCommand = cli.Command{
	Name: "run",
	Usage: "Create a container",
	Flags: []cli.Flag{
		cli.BoolFlag{
			Name:        "it",
			Usage:       "enable ttyp",
		},
	},
	Action: func(context *cli.Context) error {
		if len(context.Args()) < 1 {
			return fmt.Errorf("Missing container command")
		}
		cmd := context.Args().Get(0)
		tty := context.Bool("it")
		Run(tty, cmd)
		return nil
		
	},
}

var initCommand = cli.Command{
	Name:                   "init",
	Usage:              "Init container",

	Action: func(context *cli.Context) error {
		log.Infof("init come on")
		cmd := context.Args().Get(0)
		log.Infof("command %s", cmd)
		err := container.RunContainerInitProcess(cmd, nil)
		return err
	},
}