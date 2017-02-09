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

	if GetSetCollectionName() != PRODUCTION_SET_COLLECTION_NAME {
		t.Fatalf("Expected '%s', got '%s'", PRODUCTION_SET_COLLECTION_NAME, GetSetCollectionName())
	}
}

func TestGetSetTestCollectionName(t *testing.T) {
	environment = "test"

	if GetSetCollectionName() != TESTING_SET_COLLECTION_NAME {
		t.Fatalf("Expected '%s', got '%s'", TESTING_SET_COLLECTION_NAME, GetSetCollectionName())
	}
}
