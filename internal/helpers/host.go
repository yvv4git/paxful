package helpers

import "fmt"

// ServerAddr is used for parsing host and port to server address string.
func ServerAddr(host string, port int) string {
	return fmt.Sprintf("%s:%d", host, port)
}
