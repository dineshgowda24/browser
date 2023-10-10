package devices

import (
	"fmt"
	"log"
	"os"
	"strings"
	"testing"

	"gopkg.in/yaml.v2"
)

// testDevices is a map of user agent strings for each device from the devices.yml file
var testDevices map[string]string

// initTestDevices reads the devices.yml file and
// unmarshals it into the testDevices map.
func initTestDevices() {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatalf("failed to get working directory: %v", err)
	}

	wd = strings.Split(wd, "/browser")[0]
	yamlFile, err := os.ReadFile(fmt.Sprintf("%s/browser/assets/test/devices.yml", wd))
	if err != nil {
		log.Fatalf("failed to read file: %v", err)
	}

	var data map[string]string
	if err := yaml.Unmarshal(yamlFile, &data); err != nil {
		log.Fatalf("failed to unmarshal: %v", err)
	}

	testDevices = data
}

// TestMain is the entry point for running tests in this package.
func TestMain(m *testing.M) {
	initTestDevices()
	exitCode := m.Run()
	os.Exit(exitCode)
}
