# Simple GO logger

Simple output log with UUID according by context.Context parameter. 

When logger is attached it's generate random global UUID.\
This approach is used for services performing parallel tasks in order to separate processes.\
When functions are called without a context.Context parameter, a global UUID value is used.

# Example
Usage:
```
package main

import (
	"context"

	"github.com/google/uuid"
	"github.com/ra-company/logging"
)

func main() {
	ctx := context.Background()

	title := "Sample"

	logging.Logs.Infof(ctx, "Service %s was started.", title)

	id := uuid.New().String()
	ctx = context.WithValue(ctx, logging.CtxKeyUUID, id)

	logging.Logs.Debugf(ctx, "Log data with UUID: %s.", id)

	logging.Logs.Error("CTX isn't used.")
}
```

Sample output:
```
2025/06/17 18:17:42.016 INF     [f4d14d28-ae09-4aed-958a-c6dcb6da2a89]  Service Sample was started.
2025/06/17 18:17:42.017 DBG     [3d861cf8-ab1c-4d6d-b91e-ba17027a0045]  Log data with UUID: 3d861cf8-ab1c-4d6d-b91e-ba17027a0045.
2025/06/17 18:17:42.018 ERR     [f4d14d28-ae09-4aed-958a-c6dcb6da2a89]  CTX isn't used.
```

# Using Graylog as log server

The system can send log data to a Graylog server. To enable logging to Graylog, the following global variables must be configured:

```
GraylogAddr string = ""               // Graylog address (when empty, Graylog is disabled)
Application string = "go-application" // Application name
Host        string = "localhost"      // Host name
Facility    string = "facility"       // Facility name
```

# Staying up to date
To update library to the latest version, use go get -u github.com/ra-company/logging.

# Supported go versions
We currently support the most recent major Go versions from 1.25.5 onward.

# License
This project is licensed under the terms of the GPL-3.0 license.

## Used libraries
* [gopkg.in/Graylog2/go-gelf.v2/gelf](https://gopkg.in/Graylog2/go-gelf.v2/gelf/) -- GELF Library and Writer for Go (MIT License)
* [github.com/google/uuid](https://github.com/google/uuid/) -- Go package for UUIDs based on RFC 4122 and DCE 1.1: Authentication and Security Services. (BSD-3 License)
* [github.com/stretchr/testify](https://github.com/stretchr/testify/) - Thou Shalt Write Tests (MIT License)
