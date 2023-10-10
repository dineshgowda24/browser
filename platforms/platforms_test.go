package platforms

import (
	"fmt"
	"log"
	"os"
	"strings"
	"testing"

	"gopkg.in/yaml.v2"
)

// testPlatforms is a map of user agent strings for each platform from the platforms.yml file
var testPlatforms map[string]string

// initTestPlatforms reads the platforms.yml file and
// unmarshals it into the testPlatforms map.
func initTestPlatforms() {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatalf("failed to get working directory: %v", err)
	}

	wd = strings.Split(wd, "/browser")[0]
	yamlFile, err := os.ReadFile(fmt.Sprintf("%s/browser/assets/test/platforms.yml", wd))
	if err != nil {
		log.Fatalf("failed to read file: %v", err)
	}

	var data map[string]string
	if err := yaml.Unmarshal(yamlFile, &data); err != nil {
		log.Fatalf("failed to unmarshal: %v", err)
	}

	testPlatforms = data
}

// TestMain is the entry point for running tests in this package.
func TestMain(m *testing.M) {
	initTestPlatforms()
	exitCode := m.Run()
	os.Exit(exitCode)
}
