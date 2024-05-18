package matchers

type Opera struct {
	p Parser
}

var (
	operaName = "Opera"
	// order matters for version detection
	operaVersionRegexp = []string{`Opera Mini/([\d.]+)`, `OP(?:R|iOS|T)/([\d.]+)`, `Opera/([\d.]+)`, `Version/([\d.]+)`}
	operaMatchRegexp   = []string{`(Opera|OP(R|iOS|T)/)`}
)

func NewOpera(p Parser) *Opera {
	return &Opera{
		p: p,
	}
}

func (o *Opera) Name() string {
	return operaName
}

func (o *Opera) Version() string {
	return o.p.Version(operaVersionRegexp, 1)
}

func (o *Opera) Match() bool {
	return o.p.Match(operaMatchRegexp)
}
