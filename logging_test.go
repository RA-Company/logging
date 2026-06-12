package logging

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLogger_GetLevel(t *testing.T) {

	testCases := []struct {
		logLevel int
		level    int
		want     string
	}{
		{2, -1, "INF"},
		{2, 5, "INF"},
		{2, 4, "INF"},
		{2, 3, "FTL"},
		{2, 2, "ERR"},
		{2, 1, ""},
		{2, 0, ""},
		{0, 1, "WRN"},
		{0, 0, "DBG"},
	}

	for _, tc := range testCases {
		Logs.LogLevel = tc.logLevel

		got, _, _ := Logs.GetLevel(tc.level, context.Background())
		require.Equal(t, tc.want, got, "GetLevel(%d) = %q; want: %q", tc.level, got, tc.want)
	}
}

func ExampleLogging_Print() {
	Logs.LogLevel = 0
	Logs.UUID = "b846c7ab-9bc3-4c3a-b9e9-c65ae7bdd049"
	Logs.ShowTime = false
	Logs.Starting("test")
	defer Logs.Stopping()

	ctx := context.Background()
	ctx = context.WithValue(ctx, CtxKeyUUID, "4577c272-e9b8-4a19-a9d0-4ec0bde6063f")

	testCases := []int{-1, 0, 1, 2, 3, 4, 5}

	for _, tc := range testCases {
		Logs.Print(tc, "Hello World")
		Logs.Printf(tc, "Hello %s", "Universe")
		Logs.Print(tc, ctx, "Hello World")
		Logs.Printf(tc, ctx, "Hello %s", "Universe")
	}

	// Unordered output:
	// INF	[b846c7ab-9bc3-4c3a-b9e9-c65ae7bdd049]	test service is starting...
	// INF	[b846c7ab-9bc3-4c3a-b9e9-c65ae7bdd049]	test service is stopping...
	// DBG	[b846c7ab-9bc3-4c3a-b9e9-c65ae7bdd049]	Hello World
	// WRN	[b846c7ab-9bc3-4c3a-b9e9-c65ae7bdd049]	Hello World
	// ERR	[b846c7ab-9bc3-4c3a-b9e9-c65ae7bdd049]	Hello World
	// FTL	[b846c7ab-9bc3-4c3a-b9e9-c65ae7bdd049]	Hello World
	// INF	[b846c7ab-9bc3-4c3a-b9e9-c65ae7bdd049]	Hello World
	// INF	[b846c7ab-9bc3-4c3a-b9e9-c65ae7bdd049]	Hello World
	// INF	[b846c7ab-9bc3-4c3a-b9e9-c65ae7bdd049]	Hello World
	// DBG	[b846c7ab-9bc3-4c3a-b9e9-c65ae7bdd049]	Hello Universe
	// WRN	[b846c7ab-9bc3-4c3a-b9e9-c65ae7bdd049]	Hello Universe
	// ERR	[b846c7ab-9bc3-4c3a-b9e9-c65ae7bdd049]	Hello Universe
	// FTL	[b846c7ab-9bc3-4c3a-b9e9-c65ae7bdd049]	Hello Universe
	// INF	[b846c7ab-9bc3-4c3a-b9e9-c65ae7bdd049]	Hello Universe
	// INF	[b846c7ab-9bc3-4c3a-b9e9-c65ae7bdd049]	Hello Universe
	// INF	[b846c7ab-9bc3-4c3a-b9e9-c65ae7bdd049]	Hello Universe
	// DBG	[4577c272-e9b8-4a19-a9d0-4ec0bde6063f]	Hello World
	// WRN	[4577c272-e9b8-4a19-a9d0-4ec0bde6063f]	Hello World
	// ERR	[4577c272-e9b8-4a19-a9d0-4ec0bde6063f]	Hello World
	// FTL	[4577c272-e9b8-4a19-a9d0-4ec0bde6063f]	Hello World
	// INF	[4577c272-e9b8-4a19-a9d0-4ec0bde6063f]	Hello World
	// INF	[4577c272-e9b8-4a19-a9d0-4ec0bde6063f]	Hello World
	// INF	[4577c272-e9b8-4a19-a9d0-4ec0bde6063f]	Hello World
	// DBG	[4577c272-e9b8-4a19-a9d0-4ec0bde6063f]	Hello Universe
	// WRN	[4577c272-e9b8-4a19-a9d0-4ec0bde6063f]	Hello Universe
	// ERR	[4577c272-e9b8-4a19-a9d0-4ec0bde6063f]	Hello Universe
	// FTL	[4577c272-e9b8-4a19-a9d0-4ec0bde6063f]	Hello Universe
	// INF	[4577c272-e9b8-4a19-a9d0-4ec0bde6063f]	Hello Universe
	// INF	[4577c272-e9b8-4a19-a9d0-4ec0bde6063f]	Hello Universe
	// INF	[4577c272-e9b8-4a19-a9d0-4ec0bde6063f]	Hello Universe
}

func ExampleLogging_Info() {
	Logs.UUID = "b846c7ab-9bc3-4c3a-b9e9-c65ae7bdd049"
	Logs.Starting("test")
	Logs.ShowTime = false
	defer Logs.Stopping()

	ctx := context.Background()
	ctx = context.WithValue(ctx, CtxKeyUUID, "4577c272-e9b8-4a19-a9d0-4ec0bde6063f")

	Logs.Info("Hello World")
	Logs.Infof("Hello %s", "Universe")
	Logs.Info(ctx, "Hello World")
	Logs.Infof(ctx, "Hello %s", "Universe")
	// Unordered output:
	// INF	[b846c7ab-9bc3-4c3a-b9e9-c65ae7bdd049]	test service is starting...
	// INF	[b846c7ab-9bc3-4c3a-b9e9-c65ae7bdd049]	test service is stopping...
	// INF	[b846c7ab-9bc3-4c3a-b9e9-c65ae7bdd049]	Hello World
	// INF	[b846c7ab-9bc3-4c3a-b9e9-c65ae7bdd049]	Hello Universe
	// INF	[4577c272-e9b8-4a19-a9d0-4ec0bde6063f]	Hello World
	// INF	[4577c272-e9b8-4a19-a9d0-4ec0bde6063f]	Hello Universe
}

func ExampleLogging_Debug() {
	Logs.UUID = "b846c7ab-9bc3-4c3a-b9e9-c65ae7bdd049"
	Logs.Starting("test")
	Logs.ShowTime = false
	defer Logs.Stopping()

	ctx := context.Background()
	ctx = context.WithValue(ctx, CtxKeyUUID, "4577c272-e9b8-4a19-a9d0-4ec0bde6063f")

	testCases := []int{0, 1, 2, 3}
	for _, tc := range testCases {
		Logs.LogLevel = tc

		Logs.Debug("Hello World")
		Logs.Debugf("Hello %s", "Universe")
		Logs.Debug(ctx, "Hello World")
		Logs.Debugf(ctx, "Hello %s", "Universe")
	}

	// Unordered output:
	// INF	[b846c7ab-9bc3-4c3a-b9e9-c65ae7bdd049]	test service is starting...
	// INF	[b846c7ab-9bc3-4c3a-b9e9-c65ae7bdd049]	test service is stopping...
	// DBG	[b846c7ab-9bc3-4c3a-b9e9-c65ae7bdd049]	Hello World
	// DBG	[b846c7ab-9bc3-4c3a-b9e9-c65ae7bdd049]	Hello Universe
	// DBG	[4577c272-e9b8-4a19-a9d0-4ec0bde6063f]	Hello World
	// DBG	[4577c272-e9b8-4a19-a9d0-4ec0bde6063f]	Hello Universe
}

func ExampleLogging_Warn() {
	Logs.UUID = "b846c7ab-9bc3-4c3a-b9e9-c65ae7bdd049"
	Logs.Starting("test")
	Logs.ShowTime = false
	defer Logs.Stopping()

	ctx := context.Background()
	ctx = context.WithValue(ctx, CtxKeyUUID, "4577c272-e9b8-4a19-a9d0-4ec0bde6063f")

	testCases := []int{0, 1, 2, 3}
	for _, tc := range testCases {
		Logs.LogLevel = tc

		Logs.Warn("Hello World")
		Logs.Warnf("Hello %s", "Universe")
		Logs.Warn(ctx, "Hello World")
		Logs.Warnf(ctx, "Hello %s", "Universe")
	}

	// Unordered output:
	// INF	[b846c7ab-9bc3-4c3a-b9e9-c65ae7bdd049]	test service is starting...
	// INF	[b846c7ab-9bc3-4c3a-b9e9-c65ae7bdd049]	test service is stopping...
	// WRN	[b846c7ab-9bc3-4c3a-b9e9-c65ae7bdd049]	Hello World
	// WRN	[b846c7ab-9bc3-4c3a-b9e9-c65ae7bdd049]	Hello Universe
	// WRN	[b846c7ab-9bc3-4c3a-b9e9-c65ae7bdd049]	Hello World
	// WRN	[b846c7ab-9bc3-4c3a-b9e9-c65ae7bdd049]	Hello Universe
	// WRN	[4577c272-e9b8-4a19-a9d0-4ec0bde6063f]	Hello World
	// WRN	[4577c272-e9b8-4a19-a9d0-4ec0bde6063f]	Hello Universe
	// WRN	[4577c272-e9b8-4a19-a9d0-4ec0bde6063f]	Hello World
	// WRN	[4577c272-e9b8-4a19-a9d0-4ec0bde6063f]	Hello Universe
}

func ExampleLogging_Error() {
	Logs.UUID = "b846c7ab-9bc3-4c3a-b9e9-c65ae7bdd049"
	Logs.Starting("test")
	Logs.ShowTime = false
	defer Logs.Stopping()

	ctx := context.Background()
	ctx = context.WithValue(ctx, CtxKeyUUID, "4577c272-e9b8-4a19-a9d0-4ec0bde6063f")

	testCases := []int{0, 1, 2, 3}
	for _, tc := range testCases {
		Logs.LogLevel = tc

		Logs.Error("Hello World")
		Logs.Errorf("Hello %s", "Universe")
		Logs.Error(ctx, "Hello World")
		Logs.Errorf(ctx, "Hello %s", "Universe")
	}

	// Unordered output:
	// INF	[b846c7ab-9bc3-4c3a-b9e9-c65ae7bdd049]	test service is starting...
	// INF	[b846c7ab-9bc3-4c3a-b9e9-c65ae7bdd049]	test service is stopping...
	// ERR	[b846c7ab-9bc3-4c3a-b9e9-c65ae7bdd049]	Hello World
	// ERR	[b846c7ab-9bc3-4c3a-b9e9-c65ae7bdd049]	Hello Universe
	// ERR	[b846c7ab-9bc3-4c3a-b9e9-c65ae7bdd049]	Hello World
	// ERR	[b846c7ab-9bc3-4c3a-b9e9-c65ae7bdd049]	Hello Universe
	// ERR	[b846c7ab-9bc3-4c3a-b9e9-c65ae7bdd049]	Hello World
	// ERR	[b846c7ab-9bc3-4c3a-b9e9-c65ae7bdd049]	Hello Universe
	// ERR	[4577c272-e9b8-4a19-a9d0-4ec0bde6063f]	Hello World
	// ERR	[4577c272-e9b8-4a19-a9d0-4ec0bde6063f]	Hello Universe
	// ERR	[4577c272-e9b8-4a19-a9d0-4ec0bde6063f]	Hello World
	// ERR	[4577c272-e9b8-4a19-a9d0-4ec0bde6063f]	Hello Universe
	// ERR	[4577c272-e9b8-4a19-a9d0-4ec0bde6063f]	Hello World
	// ERR	[4577c272-e9b8-4a19-a9d0-4ec0bde6063f]	Hello Universe
}

func ExampleLogging_Fatal() {
	Logs.UUID = "b846c7ab-9bc3-4c3a-b9e9-c65ae7bdd049"
	Logs.Starting("test")
	Logs.ShowTime = false
	Logs.DontStop = true // Prevent exit on fatal error
	defer Logs.Stopping()

	ctx := context.Background()
	ctx = context.WithValue(ctx, CtxKeyUUID, "4577c272-e9b8-4a19-a9d0-4ec0bde6063f")

	testCases := []int{0, 1, 2, 3}
	for _, tc := range testCases {
		Logs.LogLevel = tc

		Logs.Fatal("Hello World")
		Logs.Fatalf("Hello %s", "Universe")
		Logs.Fatal(ctx, "Hello World")
		Logs.Fatalf(ctx, "Hello %s", "Universe")
	}

	// Unordered output:
	// INF	[b846c7ab-9bc3-4c3a-b9e9-c65ae7bdd049]	test service is starting...
	// INF	[b846c7ab-9bc3-4c3a-b9e9-c65ae7bdd049]	test service is stopping...
	// FTL	[b846c7ab-9bc3-4c3a-b9e9-c65ae7bdd049]	Hello World
	// FTL	[b846c7ab-9bc3-4c3a-b9e9-c65ae7bdd049]	Hello Universe
	// FTL	[b846c7ab-9bc3-4c3a-b9e9-c65ae7bdd049]	Hello World
	// FTL	[b846c7ab-9bc3-4c3a-b9e9-c65ae7bdd049]	Hello Universe
	// FTL	[b846c7ab-9bc3-4c3a-b9e9-c65ae7bdd049]	Hello World
	// FTL	[b846c7ab-9bc3-4c3a-b9e9-c65ae7bdd049]	Hello Universe
	// FTL	[b846c7ab-9bc3-4c3a-b9e9-c65ae7bdd049]	Hello World
	// FTL	[b846c7ab-9bc3-4c3a-b9e9-c65ae7bdd049]	Hello Universe
	// FTL	[4577c272-e9b8-4a19-a9d0-4ec0bde6063f]	Hello World
	// FTL	[4577c272-e9b8-4a19-a9d0-4ec0bde6063f]	Hello Universe
	// FTL	[4577c272-e9b8-4a19-a9d0-4ec0bde6063f]	Hello World
	// FTL	[4577c272-e9b8-4a19-a9d0-4ec0bde6063f]	Hello Universe
	// FTL	[4577c272-e9b8-4a19-a9d0-4ec0bde6063f]	Hello World
	// FTL	[4577c272-e9b8-4a19-a9d0-4ec0bde6063f]	Hello Universe
	// FTL	[4577c272-e9b8-4a19-a9d0-4ec0bde6063f]	Hello World
	// FTL	[4577c272-e9b8-4a19-a9d0-4ec0bde6063f]	Hello Universe
}

func TestGelf(t *testing.T) {
	str := os.Getenv("GRAYLOG_URL")
	if str == "" {
		t.Skip("Set GRAYLOG_URL environment variable to test Graylog connection")
	}

	GraylogAddr = str
	Host = "current-host"
	Logs.ConsoleApp = false
	Logs.UUID = "b846c7ab-9bc3-4c3a-b9e9-c65ae7bdd049"
	Logs.Starting("test")
	defer Logs.Stopping()

	ctx := context.Background()
	ctx = context.WithValue(ctx, CtxKeyUUID, "4577c272-e9b8-4a19-a9d0-4ec0bde6063f")

	Logs.Info(ctx, "Testing Graylog connection")
	Logs.Infof(ctx, "Testing Graylog connection with %s", "formatting")
	Logs.Info(ctx, "Testing Graylog connection with context")
	Logs.Infof(ctx, "Testing Graylog connection with context and %s", "formatting")
	Logs.Debug(ctx, "Debug message")
	Logs.Warn(ctx, "Warning message")
	Logs.Error(ctx, "Error message")
	Logs.Info(ctx, "\033[1m\033[36mColored message\033[0m")
}

// captureStdout redirects os.Stdout for the duration of fn and returns what was written.
func captureStdout(fn func()) string {
	r, w, _ := os.Pipe()
	old := os.Stdout
	os.Stdout = w

	fn()

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	io.Copy(&buf, r)
	return buf.String()
}

func TestGetLevel_NoContext(t *testing.T) {
	logger := &Logging{LogLevel: 0}

	lev, _, withCtx := logger.GetLevel(0, "plain string")
	assert.Equal(t, "DBG", lev)
	assert.False(t, withCtx, "plain string arg should not set withContext")
}

func TestGetLevel_NilContext(t *testing.T) {
	logger := &Logging{LogLevel: 0}
	// nil is not a context.Context, so it falls to default branch
	lev, _, withCtx := logger.GetLevel(1, nil)
	assert.Equal(t, "WRN", lev)
	assert.False(t, withCtx)
}

func TestGetLevel_ContextWithoutUUID(t *testing.T) {
	logger := &Logging{LogLevel: 0, UUID: "global-uuid"}
	ctx := context.Background() // no CtxKeyUUID value

	_, uuid, withCtx := logger.GetLevel(0, ctx)
	assert.True(t, withCtx)
	assert.Equal(t, "global-uuid", uuid, "should fall back to logger.UUID when context has no UUID")
}

func TestGetLevel_ContextWithUUID(t *testing.T) {
	logger := &Logging{LogLevel: 0, UUID: "global-uuid"}
	ctx := context.WithValue(context.Background(), CtxKeyUUID, "ctx-uuid")

	_, uuid, withCtx := logger.GetLevel(0, ctx)
	assert.True(t, withCtx)
	assert.Equal(t, "ctx-uuid", uuid)
}

// --- ConsoleApp mode ---

func TestConsoleApp_SuppressesNonErrors(t *testing.T) {
	logger := &Logging{LogLevel: 0, ShowTime: false, ConsoleApp: true, UUID: "test-uuid"}

	output := captureStdout(func() {
		logger.Print(0, "debug message") // DBG — should be suppressed
		logger.Print(1, "warn message")  // WRN — should be suppressed
		logger.Print(4, "info message")  // INF — should be suppressed
	})

	assert.Empty(t, output, "ConsoleApp=true should suppress non-error levels")
}

func TestConsoleApp_PrintsErrors(t *testing.T) {
	logger := &Logging{LogLevel: 0, ShowTime: false, ConsoleApp: true, UUID: "test-uuid"}

	output := captureStdout(func() {
		logger.Print(2, "error message")
		logger.Print(3, "fatal message")
	})

	assert.Contains(t, output, "error message")
	assert.Contains(t, output, "fatal message")
}

func TestConsoleApp_PrintsErrorsWithContext(t *testing.T) {
	logger := &Logging{LogLevel: 0, ShowTime: false, ConsoleApp: true, UUID: "test-uuid"}
	ctx := context.WithValue(context.Background(), CtxKeyUUID, "ctx-uuid")

	output := captureStdout(func() {
		logger.Print(2, ctx, "context error")
	})

	assert.Contains(t, output, "context error")
	// UUID brackets should NOT appear in ConsoleApp mode
	assert.NotContains(t, output, "[ctx-uuid]")
}

// --- CustomLogger with nil inner logger (fallback to global Logs) ---

func TestCustomLogger_NilLogger_FallsBackToGlobal(t *testing.T) {
	origLogLevel := Logs.LogLevel
	origUUID := Logs.UUID
	origShowTime := Logs.ShowTime
	defer func() {
		Logs.LogLevel = origLogLevel
		Logs.UUID = origUUID
		Logs.ShowTime = origShowTime
	}()

	Logs = Logging{
		LogLevel: 0,
		UUID:     "global-fallback",
		ShowTime: false,
	}

	cl := &CustomLogger{} // logger == nil

	output := captureStdout(func() {
		cl.Info("fallback message")
	})

	assert.Contains(t, output, "INF")
	assert.Contains(t, output, "global-fallback")
	assert.Contains(t, output, "fallback message")
}

func TestCustomLogger_WithLogger_UsesProvided(t *testing.T) {
	inner := &Logging{
		LogLevel: 0,
		UUID:     "inner-uuid",
		ShowTime: false,
	}

	cl := &CustomLogger{}
	cl.SetLogger(inner)

	output := captureStdout(func() {
		cl.Debug("inner debug")
	})

	assert.Contains(t, output, "inner-uuid")
	assert.Contains(t, output, "inner debug")
}

// --- DontStop prevents exit ---

func TestFatal_DontStop(t *testing.T) {
	logger := &Logging{
		LogLevel: 0,
		UUID:     "test",
		ShowTime: false,
		DontStop: true,
	}

	// If DontStop=false this would call os.Exit and kill the test process.
	// With DontStop=true it must return normally.
	require.NotPanics(t, func() {
		captureStdout(func() {
			logger.Fatal("fatal but no exit")
		})
	})
}

// --- LogLevel filtering ---

func TestLogLevel_FiltersLowerLevels(t *testing.T) {
	cases := []struct {
		logLevel int
		msgLevel int
		wantOut  bool
	}{
		{2, 0, false}, // DBG suppressed when LogLevel=2
		{2, 1, false}, // WRN suppressed when LogLevel=2
		{2, 2, true},  // ERR passes
		{2, 3, true},  // FTL passes
		{0, 0, true},  // DBG passes when LogLevel=0
	}

	for _, tc := range cases {
		name := fmt.Sprintf("logLevel=%d_msgLevel=%d", tc.logLevel, tc.msgLevel)
		t.Run(name, func(t *testing.T) {
			logger := &Logging{LogLevel: tc.logLevel, ShowTime: false, UUID: "test", DontStop: true}
			output := captureStdout(func() {
				logger.Print(tc.msgLevel, "test message")
			})
			if tc.wantOut {
				assert.NotEmpty(t, output)
			} else {
				assert.Empty(t, output)
			}
		})
	}
}

// --- Printf format string behaviour ---

func TestPrintf_FormatsCorrectly(t *testing.T) {
	logger := &Logging{LogLevel: 0, ShowTime: false, UUID: "test-uuid"}

	output := captureStdout(func() {
		logger.Printf(4, "Hello %s, count=%d", "world", 42)
	})

	assert.Contains(t, output, "Hello world, count=42")
}

func TestPrintf_WithContext(t *testing.T) {
	logger := &Logging{LogLevel: 0, ShowTime: false, UUID: "test-uuid"}
	ctx := context.WithValue(context.Background(), CtxKeyUUID, "ctx-id")

	output := captureStdout(func() {
		logger.Printf(4, ctx, "Hello %s", "world")
	})

	assert.Contains(t, output, "[ctx-id]")
	assert.Contains(t, output, "Hello world")
}

// --- sendGelfMessage: skips when text is empty ---

func TestSendGelfMessage_SkipsEmptyText(t *testing.T) {
	// Verifies no panic and no connection attempt when text is empty.
	logger := &Logging{}
	GraylogAddr = "127.0.0.1:12201"
	defer func() { GraylogAddr = "" }()

	require.NotPanics(t, func() {
		logger.sendGelfMessage("", time.Now(), "INF", "test-uuid")
	})
	// graylog writer should not be initialised for empty text
	assert.Nil(t, logger.graylog, "graylog writer should not be created for empty message")
}

// --- Concurrent access ---

func TestConcurrentPrint(t *testing.T) {
	logger := &Logging{LogLevel: 0, ShowTime: false, UUID: "concurrent-test", DontStop: true}

	var wg sync.WaitGroup
	// Redirect stdout to discard so output doesn't pollute test results
	old := os.Stdout
	os.Stdout, _ = os.Open(os.DevNull)
	defer func() { os.Stdout = old }()

	for i := range 50 {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			logger.Print(0, fmt.Sprintf("goroutine %d", n))
		}(i)
	}
	wg.Wait()
}

// --- Starting / Stopping output ---

func TestStarting_Stopping_Output(t *testing.T) {
	logger := &Logging{LogLevel: 0, ShowTime: false, UUID: "svc-uuid"}

	output := captureStdout(func() {
		logger.Starting("MyService")
		logger.Stopping()
	})

	assert.Contains(t, output, "MyService service is starting...")
	assert.Contains(t, output, "MyService service is stopping...")
}

// --- TestGelf skips when GRAYLOG_URL is not set ---

func TestGelf_SkipIfNoURL(t *testing.T) {
	str := os.Getenv("GRAYLOG_URL")
	if str == "" {
		t.Skip("GRAYLOG_URL not set; skipping Graylog integration test")
	}

	GraylogAddr = str
	Host = "current-host"
	logger := &Logging{ConsoleApp: false, UUID: "b846c7ab-9bc3-4c3a-b9e9-c65ae7bdd049"}

	ctx := context.WithValue(context.Background(), CtxKeyUUID, "4577c272-e9b8-4a19-a9d0-4ec0bde6063f")
	logger.Info(ctx, "Testing Graylog from missing tests")
}

// Debug passes args through fmt.Sprint — no format substitution occurs.
// Two string args are concatenated without a separator.
func TestDebug_NoFormatSubstitution(t *testing.T) {
	logger := &Logging{LogLevel: 0, ShowTime: false, UUID: "test"}

	output := captureStdout(func() {
		logger.Debug("value:", "World")
	})

	assert.Contains(t, output, "value:World")
}
