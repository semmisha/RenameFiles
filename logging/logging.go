package logging

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"path"
	"runtime"
)

func Init() *logrus.Logger {

	SmLog := logrus.New()
	SmLog.ReportCaller = true
	SmLog.Formatter = &logrus.TextFormatter{
		CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
			return fmt.Sprintf("%s()", frame.Func), fmt.Sprintf("%v", path.Base(frame.File))

		},
		ForceColors:   true,
		FullTimestamp: true,
	}
	logfile, err := os.OpenFile("./Logs_all_entry.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0655)
	if err != nil {
		panic(err)

	}
	multisource := io.MultiWriter(logfile, os.Stdout)
	SmLog.SetOutput(multisource)
	return SmLog
}
