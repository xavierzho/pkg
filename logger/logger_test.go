package logger

import (
	"errors"
	"go.uber.org/zap"
	"testing"
	//"github.com/pkg/errors"
)

func TestJSONLogger(t *testing.T) {
	logger, err := New(
		WithField("defined_key", "defined_value"),
		WithLogLevel(zap.DebugLevel),
		WithLogPath("log/test.log"),
		WithFileRotation("log/test01.log"),
		WithTimeFormat("2006-01-02 15:04:05"),
		WithDisableConsole(),
	)
	if err != nil {
		t.Fatal(err)
	}
	defer logger.Sync()

	err = errors.New("pkg error")
	//logger.Error("err occurs", WrapMeta(nil, NewMeta("para1", "value1"), NewMeta("para2", "value2"))...)
	//logger.Error("err occurs", WrapMeta(err, NewMeta("para1", "value1"), NewMeta("para2", "value2"))...)

}

func BenchmarkJsonLogger(b *testing.B) {
	b.ResetTimer()
	logger, err := New(
		WithField("defined_key", "defined_value"),
	)
	if err != nil {
		b.Fatal(err)
	}

	defer logger.Sync()

}
