package log

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/rs/zerolog"
	"gopkg.in/natefinch/lumberjack.v2"
)

type Logger struct {
	zerolog.Logger
}

var gl Logger

func init() {
	zerolog.TimeFieldFormat = TimeFormatUnix
	zerolog.SetGlobalLevel(zerolog.Level(zerolog.TraceLevel))
	gl.Logger = zerolog.New(os.Stdout).With().Timestamp().Logger()
}

func GetLogger(name string) Logger {
	if name != "" {
		return Logger{gl.With().Str(ModuleKey, name).Logger()}
	}
	return gl
}

func NewLogger(conf Config) Logger {
	var writers []io.Writer
	if conf.PrettyLog {
		writers = append(writers, zerolog.ConsoleWriter{Out: os.Stderr})
	} else {
		writers = append(writers, os.Stdout)
	}
	if conf.FileLogging {
		writers = append(writers, newLumberJack(conf.LumberJackConfig))
	}
	ctx := zerolog.New(zerolog.MultiLevelWriter(writers...)).
		Level(zerolog.Level(conf.LogLevel-1)).
		With().
		Timestamp().
		Str(ModuleKey, conf.Name)
	if conf.Caller {
		ctx = ctx.Caller()
	}

	return Logger{ctx.Logger()}
}

func Configure(conf Config) {
	var writers []io.Writer
	if conf.FileLogging {
		writers = append(writers, newLumberJack(conf.LumberJackConfig))
	}
	if conf.PrettyLog {
		writers = append(writers, zerolog.ConsoleWriter{Out: os.Stderr})
	} else {
		writers = append(writers, os.Stdout)
	}
	zerolog.TimeFieldFormat = conf.TimeFormat

	ctx := zerolog.New(zerolog.MultiLevelWriter(writers...)).
		Level(zerolog.Level(conf.LogLevel - 1)).
		With().
		Timestamp()
	if conf.Name != "" {
		ctx = ctx.Str(ModuleKey, conf.Name)
	}
	if conf.Caller {
		ctx = ctx.Caller()
	}
	gl.Logger = ctx.Logger()
}

func (l Logger) SetReqID(id string) Logger {
	l.Logger = l.Logger.With().Str(ReqID, id).Logger()
	return l
}

func newLumberJack(conf LumberJackConfig) *lumberjack.Logger {

	path := func() string {
		if conf.Filepath != "" {
			return conf.Filepath
		}
		return DefaultFilepath
	}()
	name := func() string {
		if conf.Filename != "" {
			return conf.Filename
		}
		return DefaultFilename
	}()

	filename := fmt.Sprintf("%v%v.log", path, name)

	fl := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    DefaultMaxSize,
		MaxBackups: DefaultMaxBackups,
		LocalTime:  DefaultLocalTime,
		Compress:   DefaultCompress,
	}

	if conf.MaxSize != 0 {
		fl.MaxSize = conf.MaxSize
	}
	if conf.MaxBackups != 0 {
		fl.MaxBackups = conf.MaxBackups
	}
	if conf.LocalTime {
		fl.LocalTime = conf.LocalTime
	}
	if conf.Compress {
		fl.Compress = conf.Compress
	}

	go func() {
		rotate := func() time.Duration {
			if conf.Rotate != 0 {
				return conf.Rotate
			}
			return DefaultRotate
		}()
		for {
			<-time.After(time.Second * rotate)
			fl.Rotate()
		}
	}()

	return fl
}
