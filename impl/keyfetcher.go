package impl

import (
	"os"
)

func GetValveApiKey() string {
	return os.Getenv("VALVE_API_KEY")
}