package sysinit

import (
	"encoding/json"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

var (
	consoleLogs *logs.BeeLogger
	prodLogs    *logs.BeeLogger
	runmode     string
)

type adapterMultiFileConfig struct {
	filename string
	maxlines int64
	maxsize  int64
	daily    bool
	maxdays  int64
	level    string
}

func InitLogs() {
	consoleLogs = logs.NewLogger(1)
	consoleLogs.SetLogger(logs.AdapterConsole)
	consoleLogs.Async()
	consoleLogs.EnableFuncCallDepth(true)

	filename := beego.AppConfig.String("logs::level")
	maxlines, _ := beego.AppConfig.Int64("logs::maxlines")
	maxsize, _ := beego.AppConfig.Int64("logs::maxsize")
	daily, _ := beego.AppConfig.Bool("logs::daily")
	maxdays, _ := beego.AppConfig.Int64("logs::maxdays")
	level := beego.AppConfig.String("logs::level")

	logConfig := adapterMultiFileConfig{
		filename: filename,
		maxlines: maxlines,
		maxsize:  maxsize,
		daily:    daily,
		maxdays:  maxdays,
		level:    level,
	}

	jsonbyte, _ := json.MarshalIndent(logConfig, "", "")
	prodLogs = logs.NewLogger(10000)

	prodLogs.SetLogger(logs.AdapterMultiFile, string(jsonbyte))
	prodLogs.Async()
	runmode = strings.TrimSpace(strings.ToLower(beego.AppConfig.String("runmode")))
	if runmode == "" {
		runmode = "dev"
	}
}

func LogEmergency(v interface{}) {
	log("emergency", v)
}

func LogAlert(v interface{}) {
	log("alert", v)
}

func LogCritical(v interface{}) {
	log("critical", v)
}

func LogError(v interface{}) {
	log("error", v)
}

func LogWarning(v interface{}) {
	log("warning", v)
}

func LogNotice(v interface{}) {
	log("notice", v)
}

func LogInfo(v interface{}) {
	log("info", v)
}

func LogDebug(v interface{}) {
	log("debug", v)
}

func LogTrace(v interface{}) {
	log("trace", v)
}

func log(level, v interface{}) {
	format := "%s"
	if level == "" {
		level = "debug"
	}
	if runmode == "dev" {
		switch level {
		case "emergency":
			prodLogs.Emergency(format, v)
		case "alert":
			prodLogs.Alert(format, v)
		case "critical":
			prodLogs.Critical(format, v)
		case "error":
			prodLogs.Error(format, v)
		case "warning":
			prodLogs.Warning(format, v)
		case "notice":
			prodLogs.Notice(format, v)
		case "info":
			prodLogs.Info(format, v)
		case "debug":
			prodLogs.Debug(format, v)
		case "trace":
			prodLogs.Trace(format, v)
		default:
			prodLogs.Debug(format, v)
		}
	}

	switch level {
	case "emergency":
		consoleLogs.Emergency(format, v)
	case "alert":
		consoleLogs.Alert(format, v)
	case "critical":
		consoleLogs.Critical(format, v)
	case "error":
		consoleLogs.Error(format, v)
	case "warning":
		consoleLogs.Warning(format, v)
	case "notice":
		consoleLogs.Notice(format, v)
	case "info":
		consoleLogs.Info(format, v)
	case "debug":
		consoleLogs.Debug(format, v)
	case "trace":
		consoleLogs.Trace(format, v)
	default:
		consoleLogs.Debug(format, v)
	}
}
