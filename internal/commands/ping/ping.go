package ping

import (
	"github.com/NilsPonsard/verbosity"
	cli "github.com/jawher/mow.cli"
)

// setup ping command
func Ping(job *cli.Cmd) {

	// arguments

	pong := job.BoolOpt("p pong", false, "Answer ping")

	// function to execute

	job.Action = func() {

		if *pong {
			verbosity.Info("Ping !")
		} else {
			verbosity.Info("Pong !")
		}
	}
}
