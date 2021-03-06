/*
  btcbot is a Bitcoin trading bot for HUOBI.com written
  in golang, it features multiple trading methods using
  technical analysis.

  Disclaimer:

  USE AT YOUR OWN RISK!

  The author of this project is NOT responsible for any damage or loss caused
  by this software. There can be bugs and the bot may not perform as expected
  or specified. Please consider testing it first with paper trading /
  backtesting on historical data. Also look at the code to see what how
  it's working.

  Weibo:http://weibo.com/bocaicfa
*/

package logger

import (
	"fmt"
	. "github.com/qshuai162/config"
	"io"
	"log"
	// "math"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

var (
	// 日志文件
	trade_file = "trade.csv"
	// override_file = "report"
	info_file     = "info.log"
	debug_file    = "debug.log"
	trace_file    = "trace.log"
	error_file    = "error.log"
	fatal_file    = "fatal.log"
	backtest_file string
)

func GetFileName(source string) string {
	t := time.Now()
	path := fmt.Sprintf("%s/log/%4d%0.2d%0.2d/", ROOT, t.Year(), t.Month(), t.Day())
	_, err := os.Stat(path)
	if err != nil {
		os.Mkdir(path, 0777)
	}
	if source == "trade.csv" || source == "error.log" || source == "fatal.log" {
		return fmt.Sprintf("%s%s", path, source)
	} else {
		return fmt.Sprintf("%s%2d_%s", path, t.Hour(), source)
	}
}

func init() {
	os.Mkdir(ROOT+"/log/", 0777)
	os.Mkdir(ROOT+"/cache/", 0777)

}

type logger struct {
	*log.Logger
}

func New(out io.Writer) *logger {
	return &logger{
		Logger: log.New(out, "", log.LstdFlags),
	}
}

func NewReport(out io.Writer) *logger {
	return &logger{
		Logger: log.New(out, "", log.LstdFlags),
	}
}

// func SetBacktestFile(filepath string) {
// 	backtest_file = filepath
// }

// func Backtestf(format string, args ...interface{}) {
// 	file, err := os.OpenFile(ROOT+backtest_file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
// 	if err != nil {
// 		fmt.Println("OpenFile error")
// 		return
// 	}
// 	defer file.Close()
// 	message := fmt.Sprintf(format, args...)
// 	file.Write([]byte(message))
// 	fmt.Printf(message)
// }

func Tradef(format string, args ...interface{}) {
	file, err := os.OpenFile(GetFileName(trade_file), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return
	}
	defer file.Close()
	New(file).Printf(format, args...)
	if Config["infoconsole"] == "1" {
		log.Printf(format, args...)
	}
}

func Tradeln(args ...interface{}) {
	file, err := os.OpenFile(GetFileName(trade_file), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return
	}
	defer file.Close()
	New(file).Println(args...)
	if Config["infoconsole"] == "1" {
		log.Println(args...)
	}
}

func Infof(format string, args ...interface{}) {
	file, err := os.OpenFile(GetFileName(info_file), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return
	}
	defer file.Close()
	New(file).Printf(format, args...)
	if Config["infoconsole"] == "1" {
		log.Printf(format, args...)
	}
}

func Infoln(args ...interface{}) {
	file, err := os.OpenFile(GetFileName(info_file), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return
	}
	defer file.Close()
	New(file).Println(args...)
	if Config["infoconsole"] == "1" {
		log.Println(args...)
	}
}

func Errorf(format string, args ...interface{}) {
	file, err := os.OpenFile(GetFileName(error_file), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return
	}
	defer file.Close()
	New(file).Printf(format, args...)
	if Config["errorconsole"] == "1" {
		log.Printf(format, args...)
	}
}

func Errorln(args ...interface{}) {
	file, err := os.OpenFile(GetFileName(error_file), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return
	}
	defer file.Close()
	// 加上文件调用和行号
	_, callerFile, line, ok := runtime.Caller(1)
	if ok {
		args = append([]interface{}{"[", filepath.Base(callerFile), "]", line}, args...)
	}
	New(file).Println(args...)
	if Config["errorconsole"] == "1" {
		log.Println(args...)
	}
}

func Fatalf(format string, args ...interface{}) {
	file, err := os.OpenFile(GetFileName(fatal_file), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return
	}
	defer file.Close()
	New(file).Printf(format, args...)
	if Config["fatalconsole"] == "1" {
		log.Printf(format, args...)
	}
}

func Fatalln(args ...interface{}) {
	file, err := os.OpenFile(GetFileName(fatal_file), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return
	}
	defer file.Close()
	// 加上文件调用和行号
	_, callerFile, line, ok := runtime.Caller(1)
	if ok {
		args = append([]interface{}{"[", filepath.Base(callerFile), "]", line}, args...)
	}
	New(file).Println(args...)
	if Config["fatalconsole"] == "1" {
		log.Println(args...)
	}
}

func Fatal(args ...interface{}) {
	file, err := os.OpenFile(GetFileName(fatal_file), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return
	}
	defer file.Close()
	// 加上文件调用和行号
	_, callerFile, line, ok := runtime.Caller(1)
	if ok {
		args = append([]interface{}{"[", filepath.Base(callerFile), "]", line}, args...)
	}
	New(file).Println(args...)
	if Config["fatalconsole"] == "1" {
		log.Println(args...)
	}
}

func Debugf(format string, args ...interface{}) {
	if Config["debug"] == "1" {
		file, err := os.OpenFile(GetFileName(debug_file), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
		if err != nil {
			return
		}
		defer file.Close()
		New(file).Printf(format, args...)

		if Config["debugconsole"] == "1" {
			log.Printf(format, args...)
		}
	}
}

func Debugln(args ...interface{}) {
	if Config["debug"] == "1" {
		file, err := os.OpenFile(GetFileName(debug_file), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
		if err != nil {
			return
		}
		defer file.Close()
		// 加上文件调用和行号
		_, callerFile, line, ok := runtime.Caller(1)
		if ok {
			args = append([]interface{}{"[", filepath.Base(callerFile), "]", line}, args...)
		}
		New(file).Println(args...)
		if Config["debugconsole"] == "1" {
			log.Println(args...)
		}
	}
}

func Tracef(format string, args ...interface{}) {
	file, err := os.OpenFile(GetFileName(trace_file), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return
	}
	defer file.Close()
	New(file).Printf(format, args...)
}

func Traceln(args ...interface{}) {
	file, err := os.OpenFile(GetFileName(trace_file), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return
	}
	defer file.Close()
	// 加上文件调用和行号
	_, callerFile, line, ok := runtime.Caller(1)
	if ok {
		args = append([]interface{}{"[", filepath.Base(callerFile), "]", line}, args...)
	}
	New(file).Println(args...)
}

// var what string

// func OverrideStart(Peroid int) {
// 	what = fmt.Sprintf("%03d", Peroid)
// 	file, err := os.OpenFile(GetFileName(override_file)+what+".log", os.O_CREATE|os.O_TRUNC, 0644)
// 	if err != nil {
// 		return
// 	}
// 	defer file.Close()
// }

// func Overridef(format string, args ...interface{}) {
// 	file, err := os.OpenFile(override_file+what+".log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
// 	if err != nil {
// 		return
// 	}
// 	defer file.Close()
// 	NewReport(file).Printf(format, args...)
// }

// func Overrideln(args ...interface{}) {
// 	file, err := os.OpenFile(override_file+what+".log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
// 	if err != nil {
// 		return
// 	}
// 	defer file.Close()

// 	NewReport(file).Println(args...)
// }
