package matchers

type DuckDuckGo struct {
	base
}

var (
	duckDuckGoName          = "DuckDuckGo"
	duckDuckGoVersionRegexp = []string{`DuckDuck(?:Go|GoKite)?/([\d.]+)`}
	duckDuckGoMatchRegex    = []string{`DuckDuck(Go|GoKite)?`}
)

func NewDuckDuckGo(userAgent string) *DuckDuckGo {
	return &DuckDuckGo{
		base: newBase(userAgent),
	}
}

func (ddg *DuckDuckGo) Name() string {
	return duckDuckGoName
}

func (ddg *DuckDuckGo) Version() string {
	return ddg.version(duckDuckGoVersionRegexp, 1)
}

func (ddg *DuckDuckGo) Match() bool {
	return ddg.match(duckDuckGoMatchRegex)
}
