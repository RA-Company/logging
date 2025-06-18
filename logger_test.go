package logging

import (
	"context"
	"testing"
)

func TestCustomLogger(t *testing.T) {
	logs := &CustomLogger{}
	Logs.LogLevel = 0
	Logs.UUID = "b846c7ab-9bc3-4c3a-b9e9-c65ae7bdd049"
	Logs.Starting("test")
	Logs.ShowTime = false
	Logs.DontStop = true // Prevent exit on fatal error
	defer Logs.Stopping()

	ctx := context.Background()
	ctx = context.WithValue(ctx, CtxKeyUUID, "4577c272-e9b8-4a19-a9d0-4ec0bde6063f")

	logs.Debug("Hello World")
	logs.Debug(ctx, "Hello World")
	logs.Debug("Hello %s", "Universe")
	logs.Debug(ctx, "Hello %s", "Universe")
	logs.Info("Hello World")
	logs.Info(ctx, "Hello World")
	logs.Info("Hello %s", "Universe")
	logs.Info(ctx, "Hello %s", "Universe")
	logs.Warn("Hello World")
	logs.Warn(ctx, "Hello World")
	logs.Warn("Hello %s", "Universe")
	logs.Warn(ctx, "Hello %s", "Universe")
	logs.Error("Hello World")
	logs.Error(ctx, "Hello World")
	logs.Error("Hello %s", "Universe")
	logs.Error(ctx, "Hello %s", "Universe")
	logs.Fatal(ctx, "Hello %s", "Universe")

	// Unordered output:
	// INF	[b846c7ab-9bc3-4c3a-b9e9-c65ae7bdd049]	test service is starting...
	// INF	[b846c7ab-9bc3-4c3a-b9e9-c65ae7bdd049]	test service is stopping...
	// DBG	[b846c7ab-9bc3-4c3a-b9e9-c65ae7bdd049]	Hello World
	// DBG	[b846c7ab-9bc3-4c3a-b9e9-c65ae7bdd049]	Hello Universe
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
