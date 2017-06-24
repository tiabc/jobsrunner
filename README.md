# jobrunner

Sometimes, there is a need for a cron-like jobs runner which run jobs (commands) and waits until
the job finishes before starting it again in some pre-specified interval.

## Project status

The tool is under active development, don't use it in production now. The library will be ready
for production in the following few days.

## Running

Installation is performed via a simple command:

    go install github.com/tiabc/jobrunner
    
Now, specify jobs in the config file (see below) and run the application:

    $GOPATH/bin/jobrunner config.json 
    
After launch, jobrunner immediately starts all the specified jobs. Information about jobs schedule
is not stored anywhere and execution of every job will be triggered upon restart.

### Configuration file

Jobs are specified as follows:

```json
{
  "version": 1,
  "jobs": [
    {
      "cmd": "yourapp check-statuses",
      "interval": "5 seconds"
    }
  ]
}
```

Currently, `version` must be `1` and the list of jobs consists of two fields:
1. `cmd` - a command to run.
1. `interval` - time to wait before running the command again. A positive integer and modifier
(`seconds`, `minutes`, `hours` and may also be singular) are expected. Bigger intervals are
not supported as the current version does not store information about jobs schedule anywhere
and triggers their execution on restart. 

## Using as a library

Jobrunner can also be used as a library, for example:

```go
package main

import (
	"context"
	"log"

	"github.com/tiabc/jobsrunner/state"
)

func main() {
	r, err := state.NewFromFile("your-config.json")
	if err != nil {
		log.Fatal(err)
	}
	
	// The second variable is the cancel function which can finish jobrunner.
	ctx, _ := context.WithCancel(context.Background())
	
	// This call blocks the execution until the context is cancelled.
	r.Run(ctx)
}
```

## Contributing

Contributions are welcome. Feel free to propose changes, create pull requests or request new functionality.

## License

MIT
