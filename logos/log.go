package logos

import (
	"encoding/json"
	"fmt"
	"github.com/Sheyiyuan/ChronoMind/config"
	"io"
	"log"
	"os"
	"runtime"
)

var LogLevel = 2
var LogFile *os.File

func InitLog() {
	file, err := os.OpenFile("./data/log/ChronoMind.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		Error("无法打开日志文件:", err)
	}
	LogFile = file                      // 设置日志输出到文件
	log.SetPrefix("[ChronoMind] ")      // 设置日志前缀
	log.SetOutput(os.Stdout)            // 设置日志输出到标准输出
	log.SetFlags(log.Ldate | log.Ltime) // 设置日志格式
	// 读取配置文件
	globalConfig := config.GlobalConfig{}
	// 直接读取./conf/config.json
	bytes, err := os.ReadFile("./conf/config.json")
	if err != nil {
		Error("无法读取配置文件./conf/config.json:", err)
	}
	err = json.Unmarshal(bytes, &globalConfig)
	if err != nil {
		Error("配置文件格式错误:", err)
	}
	LogLevel = globalConfig.LogConfig.LogLevel
}

// ANSI 颜色代码
const (
	ColorReset  = "\033[0m"
	ColorRed    = "\033[31m"
	ColorGreen  = "\033[32m"
	ColorYellow = "\033[33m"
	ColorBlue   = "\033[34m"
	ColorPurple = "\033[35m"
	ColorCyan   = "\033[36m"
)

func Trace(text string, msg ...interface{}) {
	if LogLevel > 1 {
		return
	}
	log.SetOutput(io.MultiWriter(os.Stdout, LogFile))
	var pc uintptr
	var file string
	var line int
	var ok bool
	// 从第 2 层开始循环查找调用栈，直到找不到调用信息为止
	for i := 2; ; i++ {
		pc, file, line, ok = runtime.Caller(i)
		if !ok || line <= 0 {
			pc, file, line, ok = runtime.Caller(i - 1)
			break
		}
	}
	funcName := runtime.FuncForPC(pc).Name()
	log.Printf("%s[Trace]  [%s:%d %s()] %s%s\n", ColorCyan, file, line, funcName, fmt.Sprintf(text, msg...), ColorReset)
}

func Debug(text string, msg ...interface{}) {
	if LogLevel > 2 {
		return
	}
	log.SetOutput(io.MultiWriter(os.Stdout, LogFile))
	log.Printf("%s[Debug]  %s%s\n", ColorBlue, fmt.Sprintf(text, msg...), ColorReset)
}

func Info(text string, msg ...interface{}) {
	if LogLevel > 3 {
		return
	}
	msgText := fmt.Sprintf(text, msg...)
	log.SetOutput(io.MultiWriter(os.Stdout, LogFile))
	log.Printf("%s[Info]  %s%s\n", ColorGreen, msgText, ColorReset)
}

func Notice(text string, msg ...interface{}) {
	if LogLevel > 4 {
		return
	}
	msgText := fmt.Sprintf(text, msg...)
	log.SetOutput(io.MultiWriter(os.Stdout, LogFile))
	log.Printf("%s[Notice] %s%s\n", ColorPurple, msgText, ColorReset)
}

func Warn(text string, msg ...interface{}) {
	if LogLevel > 5 {
		return
	}
	msgText := fmt.Sprintf(text, msg...)
	log.SetOutput(io.MultiWriter(os.Stdout, LogFile))
	log.Printf("%s[Warn]  %s%s\n", ColorYellow, msgText, ColorReset)
}

func Error(text string, msg ...interface{}) {
	if LogLevel > 6 {
		return
	}
	msgText := fmt.Sprintf(text, msg...)
	log.SetOutput(io.MultiWriter(os.Stdout, LogFile))
	log.Printf("%s[Error] %s%s\n", ColorRed, msgText, ColorReset)
}

func Fatal(text string, msg ...interface{}) {
	msgText := fmt.Sprintf(text, msg...)
	log.SetOutput(io.MultiWriter(os.Stdout, LogFile))
	log.Fatalf("%s[Fatal] %s%s\n", ColorRed, msgText, ColorReset)
}
