package log

import (
	"angrymiao-ai/pkg"
	"fmt"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"time"
)

const (
	basePath        = "./logs"
	infoPath        = "./logs/info.log"
	errorPath       = "./logs/error.log"
	rotationHour    = 7 * 24
	maxRotationTime = 24
)

//全局log
var Log = logrus.New()

func Init() {
	logFile, err := os.OpenFile(os.DevNull, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println("log file open fail", err)
	}

	// 以JSON格式为输出，代替默认的ASCII格式
	logrus.SetFormatter(&logrus.JSONFormatter{})

	// 输出到stdout 和 logfile
	multiWrites := io.MultiWriter(os.Stdout, logFile)
	logrus.SetOutput(multiWrites)

	// 设置日志等级
	Log.SetLevel(logrus.DebugLevel)

	if ok := pkg.PathIsExist(basePath); !ok {
		fmt.Println("create log directory")

		os.Mkdir("log", os.ModePerm)
	}

	infoWriter, err := rotatelogs.New(
		infoPath+".%Y-%m-%d.log",
		rotatelogs.WithLinkName(infoPath),
		rotatelogs.WithMaxAge(rotationHour*time.Hour),
		rotatelogs.WithRotationTime(maxRotationTime*time.Hour),
	)
	if err != nil {
		panic(err)
	}

	errorWriter, err := rotatelogs.New(
		errorPath+".%Y-%m-%d",
		rotatelogs.WithLinkName(errorPath),
		rotatelogs.WithMaxAge(rotationHour*time.Hour),
		rotatelogs.WithRotationTime(maxRotationTime*time.Hour),
	)
	if err != nil {
		panic(err)
	}

	// 写入不同文件中
	writerMap := lfshook.WriterMap{
		logrus.InfoLevel:  infoWriter,
		logrus.DebugLevel: infoWriter,
		logrus.FatalLevel: errorWriter,
		logrus.ErrorLevel: errorWriter,
	}
	hook := lfshook.NewHook(writerMap, &logrus.JSONFormatter{})
	Log.AddHook(hook)
}
