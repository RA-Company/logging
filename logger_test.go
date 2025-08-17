package logging

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

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

	Logs.Debug(ctx, "Standard Logger with ctx and variable: %s", world)
	Logs.Debug(ctx, "Standard Logger with ctx without variable")
	Logs.Debug("Standard Logger without ctx and variable: %s", world)
	Logs.Debug("Standard Logger without ctx without variable")

	logs.Debug(ctx, "CustomerLogger with ctx and variable: %s", world)
	logs.Debug(ctx, "CustomerLogger with ctx without variable")
	logs.Debug("CustomerLogger without ctx with variable: %s", world)
	logs.Debug("CustomerLogger without ctx without variable")

	logs.Info(ctx, "CustomerLogger with ctx and variable: %s", world)
	logs.Info(ctx, "CustomerLogger with ctx without variable")
	logs.Info("CustomerLogger without ctx with variable: %s", world)
	logs.Info("CustomerLogger without ctx without variable")

	logs.Warn(ctx, "CustomerLogger with ctx and variable: %s", world)
	logs.Warn(ctx, "CustomerLogger with ctx without variable")
	logs.Warn("CustomerLogger without ctx with variable: %s", world)
	logs.Warn("CustomerLogger without ctx without variable")

	logs.Error(ctx, "CustomerLogger with ctx and variable: %s", world)
	logs.Error(ctx, "CustomerLogger with ctx without variable")
	logs.Error("CustomerLogger without ctx with variable: %s", world)
	logs.Error("CustomerLogger without ctx without variable")

	logs.Fatal(ctx, "CustomerLogger with ctx and variable: %s", world)

	// Unordered output:
	// INF	[b846c7ab-9bc3-4c3a-b9e9-c65ae7bdd050]	test service is starting...
	// INF	[b846c7ab-9bc3-4c3a-b9e9-c65ae7bdd050]	test service is stopping...
	// DBG	[4577c272-e9b8-4a19-a9d0-4ec0bde6063f]	Standard Logger with ctx and variable: World
	// DBG	[4577c272-e9b8-4a19-a9d0-4ec0bde6063f]	Standard Logger with ctx without variable
	// DBG	[b846c7ab-9bc3-4c3a-b9e9-c65ae7bdd050]	Standard Logger without ctx and variable: World
	// DBG	[b846c7ab-9bc3-4c3a-b9e9-c65ae7bdd050]	Standard Logger without ctx without variable
	// DBG	[4577c272-e9b8-4a19-a9d0-4ec0bde6063f]	CustomerLogger with ctx and variable: World
	// DBG	[4577c272-e9b8-4a19-a9d0-4ec0bde6063f]	CustomerLogger with ctx without variable
	// DBG	[b846c7ab-9bc3-4c3a-b9e9-c65ae7bdd049]	CustomerLogger without ctx with variable: World
	// DBG	[b846c7ab-9bc3-4c3a-b9e9-c65ae7bdd049]	CustomerLogger without ctx without variable
	// INF	[4577c272-e9b8-4a19-a9d0-4ec0bde6063f]	CustomerLogger with ctx and variable: World
	// INF	[4577c272-e9b8-4a19-a9d0-4ec0bde6063f]	CustomerLogger with ctx without variable
	// INF	[b846c7ab-9bc3-4c3a-b9e9-c65ae7bdd049]	CustomerLogger without ctx with variable: World
	// INF	[b846c7ab-9bc3-4c3a-b9e9-c65ae7bdd049]	CustomerLogger without ctx without variable
	// WRN	[4577c272-e9b8-4a19-a9d0-4ec0bde6063f]	CustomerLogger with ctx and variable: World
	// WRN	[4577c272-e9b8-4a19-a9d0-4ec0bde6063f]	CustomerLogger with ctx without variable
	// WRN	[b846c7ab-9bc3-4c3a-b9e9-c65ae7bdd049]	CustomerLogger without ctx with variable: World
	// WRN	[b846c7ab-9bc3-4c3a-b9e9-c65ae7bdd049]	CustomerLogger without ctx without variable
	// ERR	[4577c272-e9b8-4a19-a9d0-4ec0bde6063f]	CustomerLogger with ctx and variable: World
	// ERR	[4577c272-e9b8-4a19-a9d0-4ec0bde6063f]	CustomerLogger with ctx without variable
	// ERR	[b846c7ab-9bc3-4c3a-b9e9-c65ae7bdd049]	CustomerLogger without ctx with variable: World
	// ERR	[b846c7ab-9bc3-4c3a-b9e9-c65ae7bdd049]	CustomerLogger without ctx without variable
	// FTL	[4577c272-e9b8-4a19-a9d0-4ec0bde6063f]	CustomerLogger with ctx and variable: World
}
