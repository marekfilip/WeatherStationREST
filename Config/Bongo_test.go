package Config

import (
	"os"
	"os/exec"
	"testing"
)

func TestGetBongoConfig(t *testing.T) {
	if GetBongoConfig() != config {
		t.Fatalf("Expected '%+v', got '%v'", config, GetBongoConfig())
	}
}

func TestGetTemperatureCollectionName(t *testing.T) {
	environment = "prod"

	if GetTemperatureCollectionName() != PRODUCTION_TEMPERATURE_COLLECTION_NAME {
		t.Fatalf("Expected '%s', got '%s'", PRODUCTION_TEMPERATURE_COLLECTION_NAME, GetTemperatureCollectionName())
	}
}
func TestGetTemperatureTestCollectionName(t *testing.T) {
	environment = "test"

	if GetTemperatureCollectionName() != TESTING_TEMPERATURE_COLLECTION_NAME {
		t.Fatalf("Expected '%s', got '%s'", TESTING_TEMPERATURE_COLLECTION_NAME, GetTemperatureCollectionName())
	}
}
func TestGetBrightnessCollectionName(t *testing.T) {
	environment = "prod"

	if GetBrightnessCollectionName() != PRODUCTION_BRIGHTNESS_COLLECTION_NAME {
		t.Fatalf("Expected '%s', got '%s'", PRODUCTION_BRIGHTNESS_COLLECTION_NAME, GetBrightnessCollectionName())
	}
}
func TestGetBrightnessTestCollectionName(t *testing.T) {
	environment = "test"

	if GetBrightnessCollectionName() != TESTING_BRIGHTNESS_COLLECTION_NAME {
		t.Fatalf("Expected '%s', got '%s'", TESTING_BRIGHTNESS_COLLECTION_NAME, GetBrightnessCollectionName())
	}
}

func TestCheckEnvironmentSetNoCrash(t *testing.T) {
	environment = "prod"

	checkEnvironmentSet()
	t.Logf("Test did not crashed, it is ok")
}

func TestCheckEnvironmentSetCrash(t *testing.T) {
	environment = ""
	if os.Getenv("BE_CRASHER") == "1" {
		checkEnvironmentSet()
		return
	}

	cmd := exec.Command(os.Args[0], "-test.run=TestCheckEnvironmentSetCrash")
	cmd.Env = append(os.Environ(), "BE_CRASHER=1")

	err := cmd.Run()
	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		return
	}

	t.Fatalf("Process ran with error '%v', want exit status 1", err)
}
