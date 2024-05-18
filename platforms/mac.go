package platforms

import (
	"strings"

	"github.com/dineshgowda24/browser/utils"
)

type Mac struct {
	p Parser
}

var (
	macName          = "macOS"
	macXName         = "Mac OS X"
	macVersionRegexp = []string{`(?:Mac|MAC) OS X\s*([0-9_.]+)?`}
	macMatchRegexp   = []string{`M(ac|AC)|macOS`}
)

func NewMac(p Parser) *Mac {
	return &Mac{
		p: p,
	}
}

func (m *Mac) Name() string {
	if utils.VersionGTE(m.Version(), "10.12") {
		return macName
	}
	return macXName
}

func (m *Mac) Version() string {
	version := m.p.Version(macVersionRegexp, 1, "0")
	return strings.Replace(version, "_", ".", -1)
}

func (m *Mac) Match() bool {
	return m.p.Match(macMatchRegexp)
}
