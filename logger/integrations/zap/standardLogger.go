package leafZap

import (
	"io"

	"github.com/labstack/gommon/log"
	"go.uber.org/zap"
)

type standardLogger struct {
	instance *zap.SugaredLogger
	prefix   string
	level    log.Lvl
}

func (l standardLogger) Write(p []byte) (n int, err error) {
	return 0, nil
}

func (l standardLogger) Output() io.Writer {
	return l
}

func (l *standardLogger) SetOutput(w io.Writer) {

}

func (l standardLogger) Prefix() string {
	return l.prefix
}

func (l *standardLogger) SetPrefix(prefix string) {
	l.prefix = prefix
}

func (l standardLogger) Level() log.Lvl {
	return l.level
}

func (l *standardLogger) SetLevel(v log.Lvl) {
	l.level = v
}

func (l standardLogger) SetHeader(header string) {

}

func (l *standardLogger) Info(args ...interface{}) {
	if l.level != log.OFF {
		l.instance.Info(args...)
	}
}

func (l *standardLogger) Infof(format string, args ...interface{}) {
	if l.level != log.OFF {
		l.instance.Infof(format, args...)
	}
}

func (l *standardLogger) Infoj(j log.JSON) {
	if l.level != log.OFF {
		l.Infof("%+v\n", j)
	}
}

func (l *standardLogger) Debug(args ...interface{}) {
	if l.level != log.OFF {
		l.instance.Debug(args...)
	}
}

func (l *standardLogger) Debugf(format string, args ...interface{}) {
	if l.level != log.OFF {
		l.instance.Debugf(format, args...)
	}
}

func (l *standardLogger) Debugj(j log.JSON) {
	if l.level != log.OFF {
		l.Debugf("%+v\n", j)
	}
}

func (l *standardLogger) Error(args ...interface{}) {
	if l.level != log.OFF {
		l.instance.Error(args...)
	}
}

func (l *standardLogger) Errorf(format string, args ...interface{}) {
	if l.level != log.OFF {
		l.instance.Errorf(format, args...)
	}
}

func (l *standardLogger) Errorj(j log.JSON) {
	if l.level != log.OFF {
		l.Errorf("%+v\n", j)
	}
}

func (l *standardLogger) Warning(args ...interface{}) {
	if l.level != log.OFF {
		l.instance.Warn(args...)
	}
}

func (l *standardLogger) Warningf(format string, args ...interface{}) {
	if l.level != log.OFF {
		l.instance.Warnf(format, args...)
	}
}

func (l *standardLogger) Warningj(j log.JSON) {
	if l.level != log.OFF {
		l.Warningf("%+v\n", j)
	}
}

func (l *standardLogger) Fatal(args ...interface{}) {
	if l.level != log.OFF {
		l.instance.Fatal(args...)
	}
}

func (l *standardLogger) Fatalf(format string, args ...interface{}) {
	if l.level != log.OFF {
		l.instance.Fatalf(format, args...)
	}
}

func (l *standardLogger) Fatalj(j log.JSON) {
	if l.level != log.OFF {
		l.Fatalf("%+v\n", j)
	}
}

func (l *standardLogger) Print(args ...interface{}) {
	if l.level != log.OFF {
		l.instance.Info(args...)
	}
}

func (l *standardLogger) Println(args ...interface{}) {
	if l.level != log.OFF {
		l.instance.Info(args...)
	}
}

func (l *standardLogger) Printf(format string, args ...interface{}) {
	if l.level != log.OFF {
		l.instance.Infof(format, args...)
	}
}

func (l *standardLogger) Printj(j log.JSON) {
	if l.level != log.OFF {
		l.Printf("%+v\n", j)
	}
}

func (l *standardLogger) Warn(i ...interface{}) {
	if l.level != log.OFF {
		l.instance.Warn(i...)
	}
}

func (l *standardLogger) Warnf(format string, i ...interface{}) {
	if l.level != log.OFF {
		l.instance.Warnf(format, i...)
	}
}

func (l *standardLogger) Warnj(j log.JSON) {
	if l.level != log.OFF {
		l.Warnf("%+v\n", j)
	}
}

func (l *standardLogger) Panic(i ...interface{}) {
	if l.level != log.OFF {
		l.instance.Panic(i...)
	}
}

func (l *standardLogger) Panicf(format string, i ...interface{}) {
	if l.level != log.OFF {
		l.instance.Panicf(format, i...)
	}
}

func (l *standardLogger) Panicj(j log.JSON) {
	if l.level != log.OFF {
		l.Panicf("%+v\n", j)
	}
}

func (l standardLogger) Instance() interface{} {
	return l.instance
}

func (l standardLogger) Log(msg string) {
	if l.level != log.OFF {
		l.instance.Info(msg)
	}
}

