package customlog

import (
	"fmt"
	"github.com/rs/zerolog"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"strconv"
)

func Init(infoPath string, infoSize, infoBackup int, errPath string, errSize, errBackup int) zerolog.Logger {
	fmt.Printf("infopath: %s errpath: %s", infoPath, errPath)
	info := &lumberjack.Logger{
		Filename:   infoPath,
		MaxSize:    infoSize,
		MaxBackups: infoBackup,
		Compress:   true,
	}

	err := &lumberjack.Logger{
		Filename:   errPath,
		MaxSize:    errSize,
		MaxBackups: errBackup,
		Compress:   true,
	}

	errWriter := zerolog.MultiLevelWriter(err)
	infoWriter := zerolog.MultiLevelWriter(info)
	console := zerolog.MultiLevelWriter(os.Stdout)

	zerolog.CallerMarshalFunc = func(pc uintptr, file string, line int) string {
		short := file
		for i := len(file) - 1; i > 0; i-- {
			if file[i] == 'c' && file[i-1] == 'r' && file[i-2] == 's' {
				short = file[i+1:]
				break
			}
		}
		file = short
		return file + ":" + strconv.Itoa(line)
	}

	writer := zerolog.MultiLevelWriter(&Hook{infoWriter, zerolog.InfoLevel}, &Hook{errWriter, zerolog.ErrorLevel}, console)
	return zerolog.New(writer).With().Timestamp().Caller().Logger()

}

type Hook struct {
	w      zerolog.LevelWriter
	Levels zerolog.Level
}

func (h *Hook) Write(p []byte) (n int, err error) {
	return h.w.Write(p)
}

func (h *Hook) WriteLevel(level zerolog.Level, p []byte) (n int, err error) {
	if h.Levels == level {
		return h.w.WriteLevel(level, p)
	}

	return len(p), nil
}
