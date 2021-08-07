package logger

import (
	"fmt"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
	"io"
	"microshop/infrastructure/config"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"time"
)

var log = logrus.New()

func init() {
	cfg := config.GetConfig()
	lv,err :=  logrus.ParseLevel(cfg.LogLevel)
	if err != nil {
		log.SetLevel(logrus.InfoLevel)
	}else{
		log.SetLevel(lv)
	}
	log.SetReportCaller(true)
	log.Formatter = &logrus.TextFormatter{
		CallerPrettyfier: func(f *runtime.Frame) (string,string) {
			filename := path.Base(f.File)
			return "", fmt.Sprintf("%s:%d", filename, f.Line)
		},
	}
	if cfg.LogFile != "" {
		fullPath, err := filepath.Abs(cfg.LogFile)
		if err != nil {
			fullPath = cfg.LogFile
		}
		writer, _ := rotatelogs.New(
			fullPath + ".%Y%m%d",
			rotatelogs.WithLinkName(cfg.LogFile),
			rotatelogs.WithRotationCount(cfg.LogCount),
			rotatelogs.WithRotationTime(time.Duration(24)*time.Hour),
		)
		writers := []io.Writer{
			writer,
			os.Stdout}
		fileAndStdoutWriter := io.MultiWriter(writers...)
		log.SetOutput(fileAndStdoutWriter)
	}else{
		log.SetOutput(os.Stdout)
	}
}

func GetLogger() *logrus.Logger {
	return log
}
