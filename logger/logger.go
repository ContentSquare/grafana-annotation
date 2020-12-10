package logger

import (
	"github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
	"os"
)

func init() {

	logrus.SetFormatter(&logrus.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	logrus.SetOutput(os.Stderr)

	// logging level
	lvl, ok := os.LookupEnv("LOGLEVEL")

	if !ok || lvl == "" {
		logrus.SetLevel(logrus.InfoLevel)
	}
	if lvl == "DEBUG" {
		logrus.SetLevel(logrus.DebugLevel)
	}

	// LogFormat
	format, ok := os.LookupEnv("LOGFORMAT")
	if format == "json" {
		logrus.SetFormatter(&logrus.JSONFormatter{})

	} else {
		logrus.SetFormatter(&prefixed.TextFormatter{
			DisableColors:   false,
			TimestampFormat: "2006-01-02 15:04:05",
			FullTimestamp:   true,
			ForceFormatting: true,
		})
	}
}
