package matchers

type Opera struct {
	base
}

var (
	operaName = "Opera"
	// order matters for version detection
	operaVersionRegexp = []string{`Opera Mini/([\d.]+)`, `OP(?:R|iOS|T)/([\d.]+)`, `Opera/([\d.]+)`, `Version/([\d.]+)`}
	operaMatchRegexp   = []string{`(Opera|OP(R|iOS|T)/)`}
)

func NewOpera(userAgent string) *Opera {
	return &Opera{
		base: newBase(userAgent),
	}
}

func (o *Opera) Name() string {
	return operaName
}

func (o *Opera) Version() string {
	return o.version(operaVersionRegexp, 1)
}

func (o *Opera) Match() bool {
	return o.match(operaMatchRegexp)
}
