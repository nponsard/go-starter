package verbosity

import (
	"fmt"
	"log"
	"os"

	"github.com/TwinProduction/go-color"
)

var (
	//-- La struct permettant de log dans un fichier, à configurer avec SetupLog()
	logToFile log.Logger
	verbose   = false
	saveLog   = true
)

//-- Author : Ponsard Nils
//-- Last update : 17/08/2021

// Fonction : Définie le niveau de verbosité et ouvre le ficher log
func SetupLog(VerbosityActive bool, logPath string, SaveLog bool) {
	saveLog = SaveLog
	verbose = VerbosityActive

	//-- Ouvre le fichier de sortie en mode Append (crée si le fichier n’existe pas)
	if saveLog {

		logFile, err := os.OpenFile(logPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

		//-- on peut pas continuer sans fichier log, donc on crash

		if err != nil {
			log.Fatal(err)
		}

		//-- Crée la struct permettant de log dans le fichier

		logToFile = *log.New(logFile, "", log.LstdFlags)

		logToFile.Println("------- New execution")

		logToFile.Println(os.Args)
	}
}

//-- Author : Ponsard Nils
//-- Last update : 11/08/2021

// Fonction : Écrit dans la sortie Standard et le fichier de log en même temps
func doubleLog(chosenColor string, level string, v ...interface{}) {

	//-- on affiche dans les deux sorties

	//-- On affiche le debug qu’en mode verbose

	if level != "Debug : " || verbose {

		//-- pas de timestamp sur la première sortie

		fmt.Print(chosenColor)
		fmt.Print(level)
		fmt.Print(v...)
		fmt.Print(color.Reset)
		fmt.Println()
	}

	if saveLog {

		//-- Sauvegarde dans le fichier

		a := append([]interface{}{level}, v...)
		logToFile.Println(a...)
	}
}

//-- Author : Ponsard Nils
//-- Last update : 12/08/2021

// Fonction : Affiche le message si la verbosité est activée
func Debug(v ...interface{}) {
	doubleLog(color.Blue, "Debug : ", v...)
}

//-- Author : Ponsard Nils
//-- Last update : 23/04/2021

// Fonction : Affiche le message même si la verbosité est activée
func Info(v ...interface{}) {
	doubleLog("", "", v...)
}

//-- Author : Ponsard Nils
//-- Last update : 23/04/2021

// Fonction : Affiche le message d’erreur
func Error(v ...interface{}) {
	doubleLog(color.Red, "Error : ", v...)
}

//-- Author : Ponsard Nils
//-- Last update : 23/04/2021

// Affiche l’erreur et ferme le programme
func Fatal(v ...interface{}) {
	doubleLog(color.Red, "Fatal : ", v...)

	//-- Arrête le programme avec le code d’erreur 1

	os.Exit(1)
}

//-- Author : Ponsard Nils
//-- Last update : 23/04/2021

// Fonction : Affiche le message de warning
func Warning(v ...interface{}) {
	doubleLog(color.Yellow, "Warning : ", v...)
}
