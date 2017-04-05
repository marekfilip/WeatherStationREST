package Config

import (
	"testing"
)

func TestGetBongoConfig(t *testing.T) {
	if GetBongoConfig() != config {
		t.Fatalf("Expected '%+v', got '%v'", config, GetBongoConfig())
	}
}

func TestGetSetCollectionName(t *testing.T) {
	environment = "prod"

	if GetSetCollectionName() != productionSetCollectionName {
		t.Fatalf("Expected '%s', got '%s'", productionSetCollectionName, GetSetCollectionName())
	}
}

func TestGetSetTestCollectionName(t *testing.T) {
	environment = "test"

	if GetSetCollectionName() != testingSetCollectionName {
		t.Fatalf("Expected '%s', got '%s'", testingSetCollectionName, GetSetCollectionName())
	}
}
