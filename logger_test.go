package logging

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// compile-time check that *CustomLogger satisfies Logger
var _ Logger = (*CustomLogger)(nil)

func TestCustomLogger(t *testing.T) {
	cLog := &Logging{
		LogLevel:   0,
		UUID:       "b846c7ab-9bc3-4c3a-b9e9-c65ae7bdd049",
		ShowTime:   false,
		ConsoleApp: false,
		DontStop:   true, // Prevent exit on fatal error
	}

	logs := &CustomLogger{}
	logs.SetLogger(cLog)

	require.NotNil(t, logs.logger, "CustomLogger should have a logger set")
}

func ExampleCustomLogger() {
	cLog := &Logging{
		LogLevel:   0,
		UUID:       "b846c7ab-9bc3-4c3a-b9e9-c65ae7bdd049",
		ShowTime:   false,
		ConsoleApp: false,
		DontStop:   true, // Prevent exit on fatal error
	}

	logs := &CustomLogger{}
	logs.SetLogger(cLog)

	Logs.LogLevel = 0
	Logs.UUID = "b846c7ab-9bc3-4c3a-b9e9-c65ae7bdd050"
	Logs.ShowTime = false
	Logs.Starting("test")
	Logs.DontStop = true // Prevent exit on fatal error
	defer Logs.Stopping()

	ctx := context.Background()
	ctx = context.WithValue(ctx, CtxKeyUUID, "4577c272-e9b8-4a19-a9d0-4ec0bde6063f")

	world := "World"

	Logs.Debugf(ctx, "Standard Logger with ctx and variable: %s", world)
	Logs.Debug(ctx, "Standard Logger with ctx without variable")
	Logs.Debugf("Standard Logger without ctx and variable: %s", world)
	Logs.Debug("Standard Logger without ctx without variable")

	logs.Debug(ctx, "CustomerLogger with ctx and variable: "+world)
	logs.Debug(ctx, "CustomerLogger with ctx without variable")
	logs.Debug("CustomerLogger without ctx without variable")

	logs.Info(ctx, "CustomerLogger with ctx and variable: "+world)
	logs.Info(ctx, "CustomerLogger with ctx without variable")
	logs.Info("CustomerLogger without ctx without variable")

	logs.Warn(ctx, "CustomerLogger with ctx and variable: "+world)
	logs.Warn(ctx, "CustomerLogger with ctx without variable")
	logs.Warn("CustomerLogger without ctx without variable")

	logs.Error(ctx, "CustomerLogger with ctx and variable: "+world)
	logs.Error(ctx, "CustomerLogger with ctx without variable")
	logs.Error("CustomerLogger without ctx without variable")

	logs.Fatal(ctx, "CustomerLogger with ctx and variable: "+world)

	logs.Debugf(ctx, "CustomerLogger Debugf with ctx: %s", world)
	logs.Debugf("CustomerLogger Debugf without ctx: %s", world)
	logs.Infof(ctx, "CustomerLogger Infof with ctx: %s", world)
	logs.Infof("CustomerLogger Infof without ctx: %s", world)
	logs.Warnf(ctx, "CustomerLogger Warnf with ctx: %s", world)
	logs.Warnf("CustomerLogger Warnf without ctx: %s", world)
	logs.Errorf(ctx, "CustomerLogger Errorf with ctx: %s", world)
	logs.Errorf("CustomerLogger Errorf without ctx: %s", world)
	logs.Fatalf(ctx, "CustomerLogger Fatalf with ctx: %s", world)

	// Unordered output:
	// INF	[b846c7ab-9bc3-4c3a-b9e9-c65ae7bdd050]	test service is starting...
	// INF	[b846c7ab-9bc3-4c3a-b9e9-c65ae7bdd050]	test service is stopping...
	// DBG	[4577c272-e9b8-4a19-a9d0-4ec0bde6063f]	Standard Logger with ctx and variable: World
	// DBG	[4577c272-e9b8-4a19-a9d0-4ec0bde6063f]	Standard Logger with ctx without variable
	// DBG	[b846c7ab-9bc3-4c3a-b9e9-c65ae7bdd050]	Standard Logger without ctx and variable: World
	// DBG	[b846c7ab-9bc3-4c3a-b9e9-c65ae7bdd050]	Standard Logger without ctx without variable
	// DBG	[4577c272-e9b8-4a19-a9d0-4ec0bde6063f]	CustomerLogger with ctx and variable: World
	// DBG	[4577c272-e9b8-4a19-a9d0-4ec0bde6063f]	CustomerLogger with ctx without variable
	// DBG	[b846c7ab-9bc3-4c3a-b9e9-c65ae7bdd049]	CustomerLogger without ctx without variable
	// INF	[4577c272-e9b8-4a19-a9d0-4ec0bde6063f]	CustomerLogger with ctx and variable: World
	// INF	[4577c272-e9b8-4a19-a9d0-4ec0bde6063f]	CustomerLogger with ctx without variable
	// INF	[b846c7ab-9bc3-4c3a-b9e9-c65ae7bdd049]	CustomerLogger without ctx without variable
	// WRN	[4577c272-e9b8-4a19-a9d0-4ec0bde6063f]	CustomerLogger with ctx and variable: World
	// WRN	[4577c272-e9b8-4a19-a9d0-4ec0bde6063f]	CustomerLogger with ctx without variable
	// WRN	[b846c7ab-9bc3-4c3a-b9e9-c65ae7bdd049]	CustomerLogger without ctx without variable
	// ERR	[4577c272-e9b8-4a19-a9d0-4ec0bde6063f]	CustomerLogger with ctx and variable: World
	// ERR	[4577c272-e9b8-4a19-a9d0-4ec0bde6063f]	CustomerLogger with ctx without variable
	// ERR	[b846c7ab-9bc3-4c3a-b9e9-c65ae7bdd049]	CustomerLogger without ctx without variable
	// FTL	[4577c272-e9b8-4a19-a9d0-4ec0bde6063f]	CustomerLogger with ctx and variable: World
	// DBG	[4577c272-e9b8-4a19-a9d0-4ec0bde6063f]	CustomerLogger Debugf with ctx: World
	// DBG	[b846c7ab-9bc3-4c3a-b9e9-c65ae7bdd049]	CustomerLogger Debugf without ctx: World
	// INF	[4577c272-e9b8-4a19-a9d0-4ec0bde6063f]	CustomerLogger Infof with ctx: World
	// INF	[b846c7ab-9bc3-4c3a-b9e9-c65ae7bdd049]	CustomerLogger Infof without ctx: World
	// WRN	[4577c272-e9b8-4a19-a9d0-4ec0bde6063f]	CustomerLogger Warnf with ctx: World
	// WRN	[b846c7ab-9bc3-4c3a-b9e9-c65ae7bdd049]	CustomerLogger Warnf without ctx: World
	// ERR	[4577c272-e9b8-4a19-a9d0-4ec0bde6063f]	CustomerLogger Errorf with ctx: World
	// ERR	[b846c7ab-9bc3-4c3a-b9e9-c65ae7bdd049]	CustomerLogger Errorf without ctx: World
	// FTL	[4577c272-e9b8-4a19-a9d0-4ec0bde6063f]	CustomerLogger Fatalf with ctx: World
}

// newInnerLogger returns a *Logging suitable for CustomLogger unit tests.
func newInnerLogger() *Logging {
	return &Logging{LogLevel: 0, UUID: "inner-uuid", ShowTime: false, DontStop: true}
}

func TestCustomLogger_Debugf_WithInnerLogger(t *testing.T) {
	cl := &CustomLogger{}
	cl.SetLogger(newInnerLogger())
	ctx := context.WithValue(context.Background(), CtxKeyUUID, "ctx-uuid")

	out := captureStdout(func() {
		cl.Debugf(ctx, "value: %s %d", "x", 42)
		cl.Debugf("value: %s %d", "x", 42)
	})
	assert.Contains(t, out, "[ctx-uuid]\tvalue: x 42")
	assert.Contains(t, out, "[inner-uuid]\tvalue: x 42")
}

func TestCustomLogger_Debugf_Fallback(t *testing.T) {
	origUUID, origLevel, origTime := Logs.UUID, Logs.LogLevel, Logs.ShowTime
	defer func() { Logs.UUID, Logs.LogLevel, Logs.ShowTime = origUUID, origLevel, origTime }()
	Logs.UUID, Logs.LogLevel, Logs.ShowTime = "global-uuid", 0, false

	cl := &CustomLogger{}
	out := captureStdout(func() { cl.Debugf("count: %d", 7) })
	assert.Contains(t, out, "[global-uuid]\tcount: 7")
}

func TestCustomLogger_Infof_WithInnerLogger(t *testing.T) {
	cl := &CustomLogger{}
	cl.SetLogger(newInnerLogger())
	ctx := context.WithValue(context.Background(), CtxKeyUUID, "ctx-uuid")

	out := captureStdout(func() {
		cl.Infof(ctx, "hello %s", "world")
		cl.Infof("hello %s", "world")
	})
	assert.Contains(t, out, "[ctx-uuid]\thello world")
	assert.Contains(t, out, "[inner-uuid]\thello world")
}

func TestCustomLogger_Infof_Fallback(t *testing.T) {
	origUUID, origLevel, origTime := Logs.UUID, Logs.LogLevel, Logs.ShowTime
	defer func() { Logs.UUID, Logs.LogLevel, Logs.ShowTime = origUUID, origLevel, origTime }()
	Logs.UUID, Logs.LogLevel, Logs.ShowTime = "global-uuid", 0, false

	cl := &CustomLogger{}
	out := captureStdout(func() { cl.Infof("hello %s", "world") })
	assert.Contains(t, out, "[global-uuid]\thello world")
}

func TestCustomLogger_Warnf_WithInnerLogger(t *testing.T) {
	cl := &CustomLogger{}
	cl.SetLogger(newInnerLogger())
	ctx := context.WithValue(context.Background(), CtxKeyUUID, "ctx-uuid")

	out := captureStdout(func() {
		cl.Warnf(ctx, "retry %d", 3)
		cl.Warnf("retry %d", 3)
	})
	assert.Contains(t, out, "[ctx-uuid]\tretry 3")
	assert.Contains(t, out, "[inner-uuid]\tretry 3")
}

func TestCustomLogger_Warnf_Fallback(t *testing.T) {
	origUUID, origLevel, origTime := Logs.UUID, Logs.LogLevel, Logs.ShowTime
	defer func() { Logs.UUID, Logs.LogLevel, Logs.ShowTime = origUUID, origLevel, origTime }()
	Logs.UUID, Logs.LogLevel, Logs.ShowTime = "global-uuid", 0, false

	cl := &CustomLogger{}
	out := captureStdout(func() { cl.Warnf("retry %d", 3) })
	assert.Contains(t, out, "[global-uuid]\tretry 3")
}

func TestCustomLogger_Errorf_WithInnerLogger(t *testing.T) {
	cl := &CustomLogger{}
	cl.SetLogger(newInnerLogger())
	ctx := context.WithValue(context.Background(), CtxKeyUUID, "ctx-uuid")

	out := captureStdout(func() {
		cl.Errorf(ctx, "failed: %s", "timeout")
		cl.Errorf("failed: %s", "timeout")
	})
	assert.Contains(t, out, "[ctx-uuid]\tfailed: timeout")
	assert.Contains(t, out, "[inner-uuid]\tfailed: timeout")
}

func TestCustomLogger_Errorf_Fallback(t *testing.T) {
	origUUID, origLevel, origTime := Logs.UUID, Logs.LogLevel, Logs.ShowTime
	defer func() { Logs.UUID, Logs.LogLevel, Logs.ShowTime = origUUID, origLevel, origTime }()
	Logs.UUID, Logs.LogLevel, Logs.ShowTime = "global-uuid", 0, false

	cl := &CustomLogger{}
	out := captureStdout(func() { cl.Errorf("failed: %s", "timeout") })
	assert.Contains(t, out, "[global-uuid]\tfailed: timeout")
}

func TestCustomLogger_Fatalf_WithInnerLogger(t *testing.T) {
	inner := &Logging{LogLevel: 0, UUID: "inner-uuid", ShowTime: false, DontStop: true}
	cl := &CustomLogger{}
	cl.SetLogger(inner)
	ctx := context.WithValue(context.Background(), CtxKeyUUID, "ctx-uuid")

	out := captureStdout(func() {
		cl.Fatalf(ctx, "critical: %s", "disk full")
		cl.Fatalf("critical: %s", "disk full")
	})
	assert.Contains(t, out, "[ctx-uuid]\tcritical: disk full")
	assert.Contains(t, out, "[inner-uuid]\tcritical: disk full")
}

func TestCustomLogger_Fatalf_Fallback(t *testing.T) {
	origUUID, origLevel, origTime, origStop := Logs.UUID, Logs.LogLevel, Logs.ShowTime, Logs.DontStop
	defer func() {
		Logs.UUID, Logs.LogLevel, Logs.ShowTime, Logs.DontStop = origUUID, origLevel, origTime, origStop
	}()
	Logs.UUID, Logs.LogLevel, Logs.ShowTime, Logs.DontStop = "global-uuid", 0, false, true

	cl := &CustomLogger{}
	out := captureStdout(func() { cl.Fatalf("critical: %s", "disk full") })
	assert.Contains(t, out, "[global-uuid]\tcritical: disk full")
}
