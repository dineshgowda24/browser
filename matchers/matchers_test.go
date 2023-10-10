package matchers

import (
	"fmt"
	"log"
	"os"
	"strings"
	"testing"

	"gopkg.in/yaml.v2"
)

// ua is a struct to hold the user agent strings for each platform from the matchers.yml file
// only used for testing
type ua struct {
	Android string `yaml:"android"`
	IOS     string `yaml:"ios"`
	Mac     string `yaml:"mac"`
	Windows string `yaml:"windows"`
	Linux   string `yaml:"linux"`
}

// testUserAgents is a map of user agent strings for each platform from the matchers.yml file
var testUserAgents map[string]ua

// initTestUserAgents reads the matchers.yml file and
// unmarshals it into the testUserAgents map.
func initTestUserAgents() {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatalf("failed to get working directory: %v", err)
	}

	wd = strings.Split(wd, "/browser")[0]
	yamlFile, err := os.ReadFile(fmt.Sprintf("%s/browser/assets/test/matchers.yml", wd))
	if err != nil {
		log.Fatalf("failed to read file: %v", err)
	}

	var data map[string]ua
	if err := yaml.Unmarshal(yamlFile, &data); err != nil {
		log.Fatalf("failed to unmarshal: %v", err)
	}

	testUserAgents = data
}

// TestMain is the entry point for running tests in this package.
func TestMain(m *testing.M) {
	initTestUserAgents()
	exitCode := m.Run()
	os.Exit(exitCode)
}
