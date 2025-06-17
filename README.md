# Simple GO logger

Simple output log with UUID according by context.Context parameter.

# Example
Usage:
```
package main

import (
	"context"
	
	"github.com/ra-company/logging"
)

func main() {
	ctx := context.Background()

	title := "Sample"

	logging.Logs.Info(ctx, "Service %s was started.", title)
}
```

Sample output:
```
2025/06/17 17:40:05.399	INF	[cd668e17-138e-4841-9c44-28b4558be374]	Service Sample was started.
```

# Staying up to date
To update library to the latest version, use go get -u github.com/ra-company/logging.

# Supported go versions
We currently support the most recent major Go versions from 1.24.3 onward.

# License
This project is licensed under the terms of the GPL-3.0 license.