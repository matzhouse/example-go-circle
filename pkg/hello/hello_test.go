package hello

import (
	"fmt"
	"testing"
)

func TestHello_Hello(t *testing.T) {

	name := "mat"

	goldenOutput := fmt.Sprintf("Hello %s", name)

	str := Hello(name)

	if str != goldenOutput {
		t.Log("Unexpected output")
		t.Log("Output: ", str)
		t.Log("Expected: ", goldenOutput)

		t.Fail()
	}

}
