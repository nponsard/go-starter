package commands

import (
	cli "github.com/jawher/mow.cli"
	"github.com/nilsponsard/go-starter/internal/commands/ping"
)

var devMode = true

//-- Author : Ponsard Nils
//-- Last update :27/08/2021

// configure toutes les sous-commandes
func SetupCommands(app *cli.Cli) {

	app.Command("ping", "ping", ping.Ping)
}
