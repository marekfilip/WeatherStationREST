package Config

import (
	"bytes"
	"testing"
)

func TestGetSecret(t *testing.T) {
	if bytes.Compare(GetSecret(), []byte(_SECRET_KEY)) != 0 {
		t.Fatalf("Expected '%v', got '%v'", []byte(_SECRET_KEY), GetSecret())
	}
}
