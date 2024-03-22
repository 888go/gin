package ginzap

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
	
	"github.com/888go/gin"
	
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap/zaptest/observer"
)

func init() {
	gin类.X设置运行模式(gin类.X常量_运行模式_测试)
}

func buildDummyLogger() (*zap.Logger, *observer.ObservedLogs) {
	core, obs := observer.New(zap.InfoLevel)
	logger := zap.New(core)
	return logger, obs
}

func timestampLocationCheck(t *testing.T, timestampStr string, location *time.Location) error {
	timestamp, err := time.Parse(time.RFC3339, timestampStr)
	if err != nil {
		return err
	}
	if timestamp.Location() != location {
		return fmt.Errorf("timestamp should be utc but %v", timestamp.Location())
	}

	return nil
}

func TestGinzap(t *testing.T) {
	r := gin类.X创建()

	utcLogger, utcLoggerObserved := buildDummyLogger()
	r.X中间件(Ginzap(utcLogger, time.RFC3339, true))

	localLogger, localLoggerObserved := buildDummyLogger()
	r.X中间件(Ginzap(localLogger, time.RFC3339, false))

	r.X绑定GET("/test", func(c *gin类.Context) {
		c.X输出JSON(204, nil)
	})

	res1 := httptest.NewRecorder()
	req1, _ := http.NewRequest("GET", "/test", nil)
	r.ServeHTTP(res1, req1)

	if len(utcLoggerObserved.All()) != 1 {
		t.Fatalf("Log should be 1 line but there're %d", len(utcLoggerObserved.All()))
	}

	logLine := utcLoggerObserved.All()[0]
	pathStr := logLine.Context[2].String
	if pathStr != "/test" {
		t.Fatalf("logged path should be /test but %s", pathStr)
	}

	err := timestampLocationCheck(t, logLine.Context[7].String, time.UTC)
	if err != nil {
		t.Fatal(err)
	}

	if len(localLoggerObserved.All()) != 1 {
		t.Fatalf("Log should be 1 line but there're %d", len(utcLoggerObserved.All()))
	}

	logLine = localLoggerObserved.All()[0]
	pathStr = logLine.Context[2].String
	if pathStr != "/test" {
		t.Fatalf("logged path should be /test but %s", pathStr)
	}
}

func TestGinzapWithConfig(t *testing.T) {
	r := gin类.X创建()

	utcLogger, utcLoggerObserved := buildDummyLogger()
	r.X中间件(GinzapWithConfig(utcLogger, &Config{
		TimeFormat:   time.RFC3339,
		UTC:          true,
		SkipPaths:    []string{"/no_log"},
		DefaultLevel: zapcore.WarnLevel,
	}))

	r.X绑定GET("/test", func(c *gin类.Context) {
		c.X输出JSON(204, nil)
	})

	r.X绑定GET("/no_log", func(c *gin类.Context) {
		c.X输出JSON(204, nil)
	})

	res1 := httptest.NewRecorder()
	req1, _ := http.NewRequest("GET", "/test", nil)
	r.ServeHTTP(res1, req1)

	res2 := httptest.NewRecorder()
	req2, _ := http.NewRequest("GET", "/no_log", nil)
	r.ServeHTTP(res2, req2)

	if res2.Code != 204 {
		t.Fatalf("request /no_log is failed (%d)", res2.Code)
	}

	if len(utcLoggerObserved.All()) != 1 {
		t.Fatalf("Log should be 1 line but there're %d", len(utcLoggerObserved.All()))
	}

	logLine := utcLoggerObserved.All()[0]
	pathStr := logLine.Context[2].String
	if pathStr != "/test" {
		t.Fatalf("logged path should be /test but %s", pathStr)
	}

	err := timestampLocationCheck(t, logLine.Context[7].String, time.UTC)
	if err != nil {
		t.Fatal(err)
	}

	if logLine.Level != zapcore.WarnLevel {
		t.Fatalf("log level should be warn but was %s", logLine.Level.String())
	}
}
