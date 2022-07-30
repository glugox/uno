package uno

import (
	"os"
)

// Get env variable from ENV, and return the passed default
// value if the env var was not present. If the var was present
// but empty, it will return "" (empty)
func Env(key string, alt string) (v string) {
	v, has := os.LookupEnv(key)
	if !has {
		return alt
	}
	return
}
