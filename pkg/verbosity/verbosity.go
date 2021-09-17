package verbosity

import (
	"fmt"
	"log"
	"os"

	"github.com/TwinProduction/go-color"
)

var (
	logToFile log.Logger
	verbose   = false
	saveLog   = true
)

// Set verbosity level and log file
func SetupLog(VerbosityActive bool, logPath string, SaveLog bool) {
	saveLog = SaveLog
	verbose = VerbosityActive

	// Opens the log file in append mode
	if saveLog {

		logFile, err := os.OpenFile(logPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

		// can’t open file, exit

		if err != nil {
			log.Fatal(err)
		}

		// create logger

		logToFile = *log.New(logFile, "", log.LstdFlags)

		logToFile.Println("------- New execution")
		logToFile.Println(os.Args)
	}
}

// Writes to terminal an log file
func doubleLog(chosenColor string, level string, v ...interface{}) {

	// show debug to terminal only when verbose

	if level != "Debug : " || verbose {

		// no timestamp in the terminal

		fmt.Print(chosenColor)
		fmt.Print(level)
		fmt.Print(v...)
		fmt.Print(color.Reset)
		fmt.Println()
	}

	if saveLog {

		// Save to file

		a := append([]interface{}{level}, v...)
		logToFile.Println(a...)
	}
}

// Debug message, shown when verbose mode is on. Always logged to file
func Debug(v ...interface{}) {
	doubleLog(color.Blue, "Debug : ", v...)
}

// Send message to user
func Info(v ...interface{}) {
	doubleLog("", "", v...)
}

// Error message
func Error(v ...interface{}) {
	doubleLog(color.Red, "Error : ", v...)
}

// Show error then exits
func Fatal(v ...interface{}) {
	doubleLog(color.Red, "Fatal : ", v...)
	os.Exit(1)
}

// Show a warning
func Warning(v ...interface{}) {
	doubleLog(color.Yellow, "Warning : ", v...)
}
