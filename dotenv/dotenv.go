package dotenv

import (
	env "github.com/joho/godotenv"
	"os"
)

// Env maps environment variable names to values
type Env map[string]string

// Exists returns a boolean indicating whether the error is known to report that
// a file or directory does not exist. It is satisfied by ErrNotExist as well as
// some syscall errors.
// If paths is not specified, the default is ".env".
// For the first path that does not exist, immediately returns false, path.
// If paths exists, returns true, "".
func Exists(paths ...string) (bool, string) {
	if len(paths) == 0 {
		paths = []string{".env"}
	}

	for _, p := range paths {
		if _, err := os.Stat(p); os.IsNotExist(err) {
			return false, p
		}
	}

	return true, ""
}

// Read returns a map of names to values read from the specified files.
func Read(files ...string) (Env, error) {
	return env.Read(files...)
}

// Unmarshal returns a map of names to values read from the specified string.
func Unmarshal(file string) (Env, error) {
	return env.Unmarshal(file)
}
