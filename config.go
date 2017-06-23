package jobrunner

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

// Config is jobrunner application configuration.
type Config struct {
	// Version is used for future incompatible config changes.
	Version int `json:"version"`

	// Jobs list to run.
	Jobs []ConfigJob `json:"jobs"`
}

func NewConfigFromFile(filename string) (Config, error) {
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		return Config{}, fmt.Errorf("can't read config: %s", err)
	}
	c := Config{}
	if err := json.Unmarshal(contents, &c); err != nil {
		return Config{}, fmt.Errorf("bad config file: %s", err)
	}
	if c.Version != 1 {
		return Config{}, fmt.Errorf("unsupported version: %d", c.Version)
	}
	return c, nil
}

// ConfigJob is a single job to run with the given interval.
type ConfigJob struct {
	Cmd      string            `json:"cmd"`
	Interval ConfigJobInterval `json:"interval"`
}

// ConfigJobInterval is a parsed interval.
type ConfigJobInterval int64

// json.Unmarshaler implementation.
func (c *ConfigJobInterval) UnmarshalJSON(vByte []byte) error {
	v := strings.Trim(string(vByte), `"`)

	// no-op by convention
	if v == "null" {
		return nil
	}

	vals := strings.Split(v, " ")
	if len(vals) != 2 {
		return fmt.Errorf("too many spaces")
	}
	n, err := strconv.ParseInt(vals[0], 10, 32)
	if err != nil {
		return fmt.Errorf("can't parse number")
	}
	if n <= 0 {
		return fmt.Errorf("number must be positive")
	}
	modifier := vals[1]
	switch modifier {
	case "second", "seconds":
		*c = ConfigJobInterval(n * int64(time.Second))
	case "minute", "minutes":
		*c = ConfigJobInterval(n * int64(time.Minute))
	case "hour", "hours":
		*c = ConfigJobInterval(n * int64(time.Hour))
	default:
		return fmt.Errorf("modifier must be one of: seconds, minutes, hours")
	}

	return nil
}
