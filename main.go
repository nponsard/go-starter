package main

import (
	"os"
	"path"

	cli "github.com/jawher/mow.cli"
	"github.com/nilsponsard/go-starter/internal/commands"
	"github.com/nilsponsard/go-starter/pkg/files"
	"github.com/nilsponsard/go-starter/pkg/verbosity"
)

// La version sera insérée par le script build.sh
var version string

func main() {

	//-- Création de l’app cli

	app := cli.App("go-starter", "starter project")
	app.Version("v version", version)

	defaultPath := files.ParsePath("~/.go-starter/")

	//-- arguments

	var (
		verbose     = app.BoolOpt("d debug", false, "Debug mode, more verbose operations")
		appPath     = app.StringOpt("c config", path.Join(defaultPath, "config.json"), "Path to the config file")
		disableLogs = app.BoolOpt("disable-logs", false, "Disable the saving of logs")
	)

	//-- Exécute cette fonction avant l’exécution d’une sous-commande

	app.Before = func() {

		//-- On traite le dossier de stockage de l’app

		parsedConfigPath := *appPath

		//-- On crée le dossier si il n’existe pas

		files.EnsureFolder(path.Join(defaultPath, "test"))
		files.EnsureFolder(parsedConfigPath)

		//-- Configuration des logs

		verbosity.SetupLog(*verbose, path.Join(defaultPath, "logs.txt"), !*disableLogs)

	}

	//-- enregistrer les sous-commandes

	commands.SetupCommands(app)

	//-- passe les arguments à app

	app.Run(os.Args)
}
