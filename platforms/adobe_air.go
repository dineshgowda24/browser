package platforms

type AdobeAir struct {
	base
}

var (
	adobeAirName          = "Adobe AIR"
	adobeAirVersionRegexp = []string{`AdobeAIR/([\d.]+)`}
	adobeAirMatchRegexp   = []string{`AdobeAIR`}
)

func NewAdobeAir(userAgent string) *AdobeAir {
	return &AdobeAir{
		base: base{
			userAgent: userAgent,
		},
	}
}

func (a *AdobeAir) Name() string {
	return adobeAirName
}

func (a *AdobeAir) Version() string {
	return a.version(adobeAirVersionRegexp, 1, "")
}

func (a *AdobeAir) Match() bool {
	return a.match(adobeAirMatchRegexp)
}
