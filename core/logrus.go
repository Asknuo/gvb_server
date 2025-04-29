package core

import (
	"bytes"
	"fmt"
	"gvb_server/global"
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
	log := global.Config.Logger
	//自定义日志格式
	timestamp := entry.Time.Format("2006-01-02 15:04:05")
	if entry.HasCaller() {
		//自定义文件路径
		funcVal := entry.Caller.Func
		filVal := fmt.Sprintf("%s:%d", path.Base(entry.Caller.File), entry.Caller.Line)
		//自定义输出格式
		fmt.Fprintf(b, "%s[%s] \x1b[%dm%s\x1b[0m [%s] [%s] %s\n", log.Prefix, timestamp, levelColor, entry.Level, funcVal, filVal, entry.Message)
	} else {
		//自定义输出格式
		fmt.Fprintf(b, "%s[%s] \x1b[%dm%s\x1b[0m %s\n", log.Prefix, timestamp, levelColor, entry.Level, entry.Message)
	}
	return b.Bytes(), nil
}

func InitLog() *logrus.Logger {
	mLog := logrus.New()
	//设置日志输出路径
	mLog.SetOutput(os.Stdout)
	//设置日志级别
	level, err := logrus.ParseLevel(global.Config.Logger.Level)
	if err != nil {
		level = logrus.InfoLevel
	}
	mLog.SetLevel(level)
	//设置日志格式
	mLog.SetFormatter(&LogFormatter{})
	//开启返回函数名和行号
	mLog.SetReportCaller(global.Config.Logger.ShowLine)
	return mLog
}

func InitDefaultLogger() {
	logrus.SetOutput(os.Stdout)
	level, err := logrus.ParseLevel(global.Config.Logger.Level)
	if err != nil {
		level = logrus.InfoLevel
	}
	logrus.SetLevel(level)
	logrus.SetFormatter(&LogFormatter{})
	logrus.SetReportCaller(true)
}
