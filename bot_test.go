package browser

import (
	"fmt"
	"log"
	"os"
	"strings"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
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

func TestNewBot(t *testing.T) {
	Convey("Subject: #NewBot", t, func() {
		Convey("It returns a new Bot", func() {
			b, err := NewBot("")

			So(err, ShouldBeNil)
			So(b, ShouldHaveSameTypeAs, &Bot{})
		})
	})
}

func TestBotIsBot(t *testing.T) {
	Convey("Subject: #IsBot", t, func() {
		Convey("When the user agent is a bot", func() {
			Convey("When the user agent is empty", func() {
				Convey("It returns true", func() {
					b, _ := NewBot("")

					So(b.IsBot(), ShouldBeTrue)
				})
			})

			Convey("When the user agent has bot keywords", func() {
				Convey("It returns true", func() {
					b, _ := NewBot(testBots["monitoring"])

					So(b.IsBot(), ShouldBeTrue)
				})
			})

			Convey("When the user agent is a known bot", func() {
				Convey("It returns true", func() {
					b, _ := NewBot(testBots["apis-google"])

					So(b.IsBot(), ShouldBeTrue)
				})
			})
		})
	})
}

func TestBotName(t *testing.T) {
	Convey("Subject: #Name", t, func() {
		Convey("When the user agent is a bot", func() {
			Convey("When the user agent is empty", func() {
				Convey("It returns Generic Bot", func() {
					b, _ := NewBot("")

					So(b.Name(), ShouldEqual, "Generic Bot")
				})
			})

			Convey("When the user agent is a known bot", func() {
				Convey("It returns the name of the bot", func() {
					b, _ := NewBot(testBots["apis-google"])

					So(b.Name(), ShouldEqual, "APIs-Google")
				})
			})
		})
	})
}

func TestBotWhy(t *testing.T) {
	Convey("Subject: #Why", t, func() {
		Convey("When the user agent is a bot", func() {
			Convey("When the user agent is empty", func() {
				Convey("It returns an empty string", func() {
					b, _ := NewBot("")

					So(b.Why(), ShouldEqual, "*bots.Empty")
				})
			})

			Convey("When the user agent is a known bot", func() {
				Convey("It returns the type of bot", func() {
					b, _ := NewBot(testBots["apis-google"])

					So(b.Why(), ShouldEqual, "*bots.Known")
				})
			})
		})
	})
}
