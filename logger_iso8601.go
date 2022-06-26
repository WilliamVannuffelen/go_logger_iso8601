package logger_iso8601

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type customFormatter struct {
	log.TextFormatter
}

func (f *customFormatter) Format(entry *log.Entry) ([]byte, error) {
	// required for coloured output
	var levelColor int
	switch entry.Level {
	case log.DebugLevel, log.TraceLevel:
		levelColor = 31 // grey
	case log.WarnLevel:
		levelColor = 33 // yellow
	case log.ErrorLevel, log.FatalLevel, log.PanicLevel:
		levelColor = 31 // red
	default:
		levelColor = 36 // blue
	}

	return []byte(fmt.Sprintf("%s - %s:%s - [%s] - \x1b[%dm%s\x1b[0m - %s\n",
			entry.Time.Format(f.TimestampFormat),
			filepath.Base(entry.Caller.File),
			strconv.Itoa(entry.Caller.Line),
			entry.Caller.Function,
			levelColor,
			strings.ToUpper(entry.Level.String()),
			entry.Message)),
		nil
}

func InitLogger(filePath string) (*log.Logger, error) {
	f, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY, 0777)
	myCustomLogger := &log.Logger{
		Out:          io.MultiWriter(os.Stderr, f),
		Level:        log.InfoLevel,
		ReportCaller: true,
		Formatter: &customFormatter{log.TextFormatter{
			FullTimestamp:          true,
			TimestampFormat:        "2006-01-02T15:04:05-0700",
			ForceColors:            true,
			DisableLevelTruncation: true,
		},
		},
	}
	return myCustomLogger, err
}
