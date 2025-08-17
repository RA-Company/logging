package logging

import (
	"context"
	"testing"
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

	Logs.LogLevel = 0
	Logs.UUID = "b846c7ab-9bc3-4c3a-b9e9-c65ae7bdd050"
	Logs.ShowTime = false
	Logs.Starting("test")
	Logs.DontStop = true // Prevent exit on fatal error
	defer Logs.Stopping()

	ctx := context.Background()
	ctx = context.WithValue(ctx, CtxKeyUUID, "4577c272-e9b8-4a19-a9d0-4ec0bde6063f")

	world := "World"
	universe := "Universe"

	logs.Debug("Hello %s", world)
	logs.Debug("Hello World")
	logs.Debug(ctx, "Hello %s", world)
	logs.Debug(ctx, "Hello World")
	logs.Debug("Hello %s", universe)
	logs.Debug(ctx, "Hello %s", universe)
	logs.Info("Hello %s", world)
	logs.Info(ctx, "Hello %s", world)
	logs.Info("Hello %s", universe)
	logs.Info(ctx, "Hello %s", universe)
	logs.Warn("Hello %s", world)
	logs.Warn(ctx, "Hello %s", world)
	logs.Warn("Hello %s", universe)
	logs.Warn(ctx, "Hello %s", universe)
	logs.Error("Hello %s", world)
	logs.Error(ctx, "Hello %s", world)
	logs.Error("Hello %s", universe)
	logs.Error(ctx, "Hello %s", universe)
	logs.Fatal(ctx, "Hello %s", universe)

	// Unordered output:
	// INF	[b846c7ab-9bc3-4c3a-b9e9-c65ae7bdd050]	test service is starting...
	// INF	[b846c7ab-9bc3-4c3a-b9e9-c65ae7bdd050]	test service is stopping...
	// DBG	[b846c7ab-9bc3-4c3a-b9e9-c65ae7bdd049]	Hello World
	// DBG	[b846c7ab-9bc3-4c3a-b9e9-c65ae7bdd049]	Hello World
	// DBG	[b846c7ab-9bc3-4c3a-b9e9-c65ae7bdd049]	Hello Universe
	// DBG	[4577c272-e9b8-4a19-a9d0-4ec0bde6063f]	Hello World
	// DBG	[4577c272-e9b8-4a19-a9d0-4ec0bde6063f]	Hello World
	// DBG	[4577c272-e9b8-4a19-a9d0-4ec0bde6063f]	Hello Universe
	// INF	[b846c7ab-9bc3-4c3a-b9e9-c65ae7bdd049]	Hello World
	// INF	[b846c7ab-9bc3-4c3a-b9e9-c65ae7bdd049]	Hello Universe
	// INF	[4577c272-e9b8-4a19-a9d0-4ec0bde6063f]	Hello World
	// INF	[4577c272-e9b8-4a19-a9d0-4ec0bde6063f]	Hello Universe
	// WRN	[b846c7ab-9bc3-4c3a-b9e9-c65ae7bdd049]	Hello World
	// WRN	[b846c7ab-9bc3-4c3a-b9e9-c65ae7bdd049]	Hello Universe
	// WRN	[4577c272-e9b8-4a19-a9d0-4ec0bde6063f]	Hello World
	// WRN	[4577c272-e9b8-4a19-a9d0-4ec0bde6063f]	Hello Universe
	// ERR	[b846c7ab-9bc3-4c3a-b9e9-c65ae7bdd049]	Hello World
	// ERR	[b846c7ab-9bc3-4c3a-b9e9-c65ae7bdd049]	Hello Universe
	// ERR	[4577c272-e9b8-4a19-a9d0-4ec0bde6063f]	Hello World
	// ERR	[4577c272-e9b8-4a19-a9d0-4ec0bde6063f]	Hello Universe
	// FTL	[b846c7ab-9bc3-4c3a-b9e9-c65ae7bdd049]	Hello World
	// FTL	[b846c7ab-9bc3-4c3a-b9e9-c65ae7bdd049]	Hello Universe
	// FTL	[4577c272-e9b8-4a19-a9d0-4ec0bde6063f]	Hello World
	// FTL	[4577c272-e9b8-4a19-a9d0-4ec0bde6063f]	Hello Universe
}
