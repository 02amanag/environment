# environment

```Bash
go get github.com/02amanag/environment
```

```Go
import "github.com/02amanag/environment"
```

## Documentation

See [godoc](http://godoc.org/github.com/kelseyhightower/envconfig)


## Usage

Set some environment variables in .env:

```
PORT=8080
USER=testUser
TIMEOUT_IN_HOUR=1
HOST=root
```

## Example

A very basic example:

```go
package main

import (
	"fmt"
	"time"

	"github.com/02amanag/environment"
)

type config struct {
	User         string        `env:"USER"`
	Port         int           `env:"PORT"`
	Host         string        `env:"HOST"`
	TimeOut      int           `env:"TIMEOUT_IN_HOUR"`
}

func main() {
	cfg := config{}
	if err := environment.Unmarshal(&cfg); err != nil {
		fmt.Printf("%+v\n", err)
	}

	fmt.Printf("%+v\n", cfg.User) //output -> testUser

        val, err := environment.Getenv("HOST")
        if err != nil {
		fmt.Printf("%+v\n", err)
	}

        fmt.Printf("%+v\n", val) //output -> root
}
```


## Supported types and defaults

Out of the box all built-in types are supported, plus a few others that
are commonly used.

Complete list:

- `string`
- `bool`
- `int`
- `int8`
- `int16`
- `int32`
- `int64`