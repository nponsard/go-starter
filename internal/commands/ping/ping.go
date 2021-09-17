package ping

import (
	cli "github.com/jawher/mow.cli"
	"github.com/nilsponsard/go-starter/pkg/verbosity"
)

//-- Fonction : configure la commande `ping`
//-- Author : Ponsard Nils
//-- Last update : 11/08/2021
func Ping(job *cli.Cmd) {

	//- les arguments

	pong := job.BoolOpt("p pong", false, "Answer ping")

	//-- La fonction à exécuter

	job.Action = func() {

		if *pong {
			verbosity.Info("Ping !")
		} else {
			verbosity.Info("Pong !")
		}
	}
}
