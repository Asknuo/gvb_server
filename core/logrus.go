package core

import (
	"bytes"
	"fmt"
	"os"
	"path"

	"github.com/sirupsen/logrus"
)

const (
	red    = 31
	yellow = 33
	blue   = 34
	gray   = 37
)

type LogFormatter struct{}

func (t *LogFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var levelColor int
	switch entry.Level {
	case logrus.DebugLevel, logrus.TraceLevel:
		levelColor = gray
	case logrus.WarnLevel:
		levelColor = yellow
	case logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel:
		levelColor = red
	default:
		levelColor = blue
	}
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}
	//自定义日志格式
	timestamp := entry.Time.Format("2006-01-02 15:04:05")
	if entry.HasCaller() {
		//自定义文件路径
		funcVal := entry.Caller.Func
		filVal := fmt.Sprintf("%s:%d", path.Base(entry.Caller.File), entry.Caller.Line)
		//自定义输出格式
		fmt.Fprintf(b, "[%s] \x1b[%dm%s\x1b[0m [%s] [%s] %s\n", timestamp, levelColor, entry.Level, funcVal, filVal, entry.Message)
	} else {
		//自定义输出格式
		fmt.Fprintf(b, "[%s] \x1b[%dm%s\x1b[0m %s\n", timestamp, levelColor, entry.Level, entry.Message)
	}
	return b.Bytes(), nil
}

func InitLog() *logrus.Logger {
	mLog := logrus.New()
	//设置日志输出路径
	mLog.SetOutput(os.Stdout)
	//设置日志级别
	mLog.SetLevel(logrus.DebugLevel)
	//设置日志格式
	mLog.SetFormatter(&LogFormatter{})
	//开启返回函数名和行号
	mLog.SetReportCaller(true)
	return mLog
}

func InitDefaultLogger() {
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&LogFormatter{})
	logrus.SetReportCaller(true)
}
