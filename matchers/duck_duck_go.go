package matchers

type DuckDuckGo struct {
	p Parser
}

var (
	duckDuckGoName          = "DuckDuckGo"
	duckDuckGoVersionRegexp = []string{`DuckDuck(?:Go|GoKite)?/([\d.]+)`}
	duckDuckGoMatchRegex    = []string{`DuckDuck(Go|GoKite)?`}
)

func NewDuckDuckGo(p Parser) *DuckDuckGo {
	return &DuckDuckGo{
		p: p,
	}
}

func (ddg *DuckDuckGo) Name() string {
	return duckDuckGoName
}

func (ddg *DuckDuckGo) Version() string {
	return ddg.p.Version(duckDuckGoVersionRegexp, 1)
}

func (ddg *DuckDuckGo) Match() bool {
	return ddg.p.Match(duckDuckGoMatchRegex)
}
