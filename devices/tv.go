package devices

type TV struct {
	base
}

var (
	tvName       = "TV"
	tvMatchRegex = []string{`(?i)(\btv|Android.*?ADT-1|Nexus Player)`}
)

func NewTV(userAgent string) *TV {
	return &TV{
		base{
			userAgent: userAgent,
		},
	}
}

func (t *TV) Name() string {
	return tvName
}

func (t *TV) Match() bool {
	return t.match(tvMatchRegex)
}
