package platforms

import (
	"fmt"
	"regexp"
	"strings"
)

type IOS struct {
	base
}

var (
	iOSName          = "iOS"
	iOSVersionRegexp = `OS (\d+)_(\d+)_?(\d+)?`
	iOSMatchRegexp   = []string{`(iPhone|iPad|iPod)`}
)

func NewIOS(userAgent string) *IOS {
	return &IOS{
		base{
			userAgent: userAgent,
		},
	}
}

func (i *IOS) Name() string {
	return fmt.Sprintf("%s (%s)", iOSName, i.device())
}

func (i *IOS) device() string {
	re := regexp.MustCompile(iOSMatchRegexp[0])
	matches := re.FindStringSubmatch(i.userAgent)
	if len(matches) > 1 {
		return matches[1]
	}

	return ""
}

func (i *IOS) Version() string {
	re := regexp.MustCompile(iOSVersionRegexp)
	matches := re.FindStringSubmatch(i.userAgent)
	if len(matches) == 0 {
		return "0"
	}

	versions := []string{matches[1]}
	if matches[3] != "" {
		versions = append(versions, matches[2], matches[3])
	} else if matches[2] != "0" {
		versions = append(versions, matches[2])
	}

	return strings.Join(versions, ".")
}

func (i *IOS) Match() bool {
	return i.match(iOSMatchRegexp)
}
