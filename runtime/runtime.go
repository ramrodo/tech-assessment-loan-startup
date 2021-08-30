package runtime

import (
	"os"
)

func IsLambdaEnvironment() bool {
	val := os.Getenv("AWS_EXECUTION_ENV")
	return val != ""
}
