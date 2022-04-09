package dry

import (
	"reflect"
	"testing"
)

// TestHandleError check error for test
func TestHandleError(t *testing.T, err error) {
	if err != nil {
		t.Error(err)
	}
}

// TestCheckEqual compare variables for test
func TestCheckEqual(t *testing.T, expected, testing interface{}) {
	if !reflect.DeepEqual(expected, testing) {
		t.Errorf("expect %v, got %v", expected, testing)
	}
}
