package platforms

type AdobeAir struct {
	p Parser
}

var (
	adobeAirName          = "Adobe AIR"
	adobeAirVersionRegexp = []string{`AdobeAIR/([\d.]+)`}
	adobeAirMatchRegexp   = []string{`AdobeAIR`}
)

func NewAdobeAir(p Parser) *AdobeAir {
	return &AdobeAir{
		p: p,
	}
}

func (a *AdobeAir) Name() string {
	return adobeAirName
}

func (a *AdobeAir) Version() string {
	return a.p.Version(adobeAirVersionRegexp, 1, "")
}

func (a *AdobeAir) Match() bool {
	return a.p.Match(adobeAirMatchRegexp)
}
