package hello

import (
	"fmt"
)

// Hello takes a string and returns a string with the
// greeting Hello in front of it
func Hello(name string) string {

	return fmt.Sprintf("Hello %s", name)

}
