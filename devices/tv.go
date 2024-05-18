package devices

type TV struct {
	p Parser
}

var (
	tvName       = "TV"
	tvMatchRegex = []string{`(?i)(\btv|Android.*?ADT-1|Nexus Player)`}
)

func NewTV(p Parser) *TV {
	return &TV{
		p: p,
	}
}

func (t *TV) Name() string {
	return tvName
}

func (t *TV) Match() bool {
	return t.p.Match(tvMatchRegex)
}
