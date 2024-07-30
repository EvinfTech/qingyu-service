package logKit

import (
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"log"
	"time"
)

var Log *logrus.Logger

func init() {
	Log = logrus.New()
}
func Init() {
	//设置日志级别
	Log.SetLevel(logrus.DebugLevel)
	//设置日志的格式
	Log.SetReportCaller(true)
	Log.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	//增加Hook
	Log.AddHook(newLfsHook())
	Log.Info("=======日志模块初始化=======")
}
func newLfsHook() logrus.Hook {
	//设置日志存放位置路径
	var path = "logger/log/"
	writer, err := rotatelogs.New(
		path+"%Y%m%d.log",
		//指定日志名
		rotatelogs.WithLinkName(path+"log.log"),
		//rotatelogs.WithMaxAge(5),
		rotatelogs.WithRotationCount(30),
		rotatelogs.WithRotationTime(time.Hour*24),
	)
	log.SetOutput(writer)
	if err != nil {
		logrus.Errorf("config local file system for logger error: %v", err)
	}
	lfsHook := lfshook.NewHook(lfshook.WriterMap{
		logrus.DebugLevel: writer,
		logrus.InfoLevel:  writer,
		logrus.WarnLevel:  writer,
		logrus.ErrorLevel: writer,
		logrus.FatalLevel: writer,
		logrus.PanicLevel: writer,
	}, &logrus.TextFormatter{DisableColors: true, TimestampFormat: "2006-01-02 15:04:05"})
	return lfsHook
}
