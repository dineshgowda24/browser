package matchers

import (
	"regexp"
	"strings"
)

type InternetExplorer struct {
	p Parser
}

var (
	ieName       = "Internet Explorer"
	ieMatchRegex = []string{`Trident/.*?; rv:(.*?)`}
	// https://en.wikipedia.org/wiki/Trident_(layout_engine)
	// A map of trident versions and their respective IE versions
	tridentMapping = map[string]string{
		"4.0": "8",  // IE8
		"5.0": "9",  // IE9
		"6.0": "10", // IE10
		"7.0": "11", // IE11
		"8.0": "12", // IE12
	}
)

func NewInternetExplorer(p Parser) *InternetExplorer {
	return &InternetExplorer{
		p: p,
	}
}

func (i *InternetExplorer) Name() string {
	return ieName
}

// Version returns the version of Internet Explorer
// It first tries to get the version from the Trident version
// If that fails, it tries to get the version from the MSIE version
// If that fails, it returns 0.0
func (i *InternetExplorer) Version() string {
	t := i.tridentVersion()
	if v, ok := tridentMapping[t]; ok {
		return v
	}

	return i.msieVersion()
}

// tridentVersion returns the version of Trident
func (i *InternetExplorer) tridentVersion() string {
	re := regexp.MustCompile(`Trident/([0-9.]+)`)
	matches := re.FindStringSubmatch(i.p.String())
	if len(matches) > 1 {
		return matches[1]
	}

	return "0.0"
}

func (i *InternetExplorer) msieVersion() string {
	return strings.Split(i.msieFullVersion(), ".")[0]
}

// msieFullVersion returns the full version of MSIE
func (i *InternetExplorer) msieFullVersion() string {
	re := regexp.MustCompile(`MSIE ([\d.]+)|Trident/.*?; rv:([\d.]+)`)

	matches := re.FindStringSubmatch(i.p.String())
	if len(matches) > 2 {
		if matches[1] != "" {
			return matches[1]
		} else if matches[2] != "" {
			return matches[2]
		}
	}

	return "0.0"
}

// Match returns true if the user agent matches one of the
// Internet Explorer matchers
// It also checks if the user agent contains MSIE and does not contain Opera
func (i *InternetExplorer) Match() bool {
	return i.p.Match(ieMatchRegex) || (strings.Contains(i.p.String(), "MSIE") &&
		!strings.Contains(i.p.String(), "Opera"))
}
