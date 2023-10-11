package bots

import (
	"fmt"
	"log"
	"os"
	"strings"
	"testing"

	"gopkg.in/yaml.v2"
)

// testBots is a map of user agent strings for each bot from the bots.yml file
var testBots map[string]string

// initTestBots reads the bots.yml file and
// unmarshals it into the testPlatforms map.
func initTestBots() {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatalf("failed to get working directory: %v", err)
	}

	wd = strings.Split(wd, "/browser")[0]
	yamlFile, err := os.ReadFile(fmt.Sprintf("%s/browser/assets/test/bots.yml", wd))
	if err != nil {
		log.Fatalf("failed to read file: %v", err)
	}

	var data map[string]string
	if err := yaml.Unmarshal(yamlFile, &data); err != nil {
		log.Fatalf("failed to unmarshal: %v", err)
	}

	testBots = data
}

// TestMain is the entry point for running tests in this package.
func TestMain(m *testing.M) {
	initTestBots()
	exitCode := m.Run()
	os.Exit(exitCode)
}
