package which

import (
	"errors"
	"fmt"
	"testing"
)

// NOTE: this test may fail depending on the system `ls` binary is placed in a different location
func TestWhichLs(t *testing.T) {
	actual, err := Which("ls")
	expected := "/usr/bin/ls"

	if err != nil {
		t.Error(err)
		return
	}

	if actual != expected {
		t.Errorf("Path was incorrect, got: %s, wanted: %s", actual, expected)
	}
}

func TestWhichMissing(t *testing.T) {
	executable := "missingexecutable"
	actual, actualErr := Which(executable)
	expected := ""
	expectedErr := errors.New(fmt.Sprintf("%s not found", executable))

	if actualErr.Error() != expectedErr.Error() {
		t.Errorf("Error was incorrect, got: %s, wanted: %s", actualErr, expectedErr)
	}

	if actual != expected {
		t.Errorf("Path was incorrect, got: %s, wanted: %s", actual, expected)
	}
}

// TODO: add executability tests
