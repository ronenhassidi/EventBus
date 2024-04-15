package logger

import (
	//"fmt"
	"os"

	config "github.com/gilbarco-ai/event-bus/common/configuration/viper"
	"github.com/gilbarco-ai/event-bus/common/constants"

	//"github.com/logrusorgru/aurora"
	"github.com/sirupsen/logrus"
)

var Logger *logrus.Logger

func Init() {
	generalVP := config.GetViper("general")
	logFilePath := generalVP.GetString("general.log_file_path")
	stdoutFlg := generalVP.GetString("general.stdout_flg")
	logLevel := generalVP.GetString("general.log_level")
	//ppxLog = generalVP.GetString("general.ppx_log")
	if len(logLevel) == 0 {
		logLevel = constants.Info
	}
	if stdoutFlg == "true" {
		initLog(logFilePath, true, logLevel)
	} else {
		initLog(logFilePath, false, logLevel)
	}
}

func initLog(logFilePath string, stdoutFlg bool, infoLevel string) *logrus.Logger {
	// Create a new Logrus logger instance
	Logger = logrus.New()

	// Set the logging level (options: Trace, Debug, Info, Warn, Error, Fatal, Panic)
	Logger.SetLevel(setLogLevelFromConfig(infoLevel))

	// Set the formatter for the logger
	Logger.SetFormatter(&logrus.TextFormatter{
		//FullTimestamp: true,
		DisableLevelTruncation: true,
		DisableTimestamp:       true,
	})

	// Set the output for the logger (default is os.Stderr)
	if stdoutFlg == true {
		Logger.SetOutput(os.Stdout)
	}
	// Add a hook for sending logs to a file
	file, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		Logger.SetOutput(file)
	} else {
		Logger.Info("Failed to log to file, using default stderr")
	}
	return Logger
}

func setLogLevelFromConfig(configLogLevel string) logrus.Level {
	switch configLogLevel {
	case "trace":
		return logrus.TraceLevel
	case "debug":
		return logrus.DebugLevel
	case "info":
		return logrus.InfoLevel
	case "warn":
		return logrus.WarnLevel
	case "error":
		return logrus.ErrorLevel
	case "fatal":
		return logrus.FatalLevel
	case "panic":
		return logrus.PanicLevel
	default:
		return logrus.InfoLevel
	}
}

/*
func main() {
	// Initialize the logger
	logger := Init()

	// Log a simple message
	logger.Info("This is an informational message")

	// Log a warning message
	logger.Warn("This is a warning message")

	// Log an error message
	logger.Error("This is an error message")

	// Log a message with fields
	logger.WithFields(logrus.Fields{
		"animal": "walrus",
		"size":   10,
	}).Info("A group of walrus emerges from the ocean")

	// Log a message with colors using aurora package
	logger.SetFormatter(&logrus.TextFormatter{
		ForceColors: true,
	})
	logger.Info(aurora.Green("This is a green message"))

	// Close the log file, if used
	if file, ok := logger.Out.(*os.File); ok {
		file.Close()
	}
}
*/
