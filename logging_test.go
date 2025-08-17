package logging

import (
	"context"
	"testing"

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
