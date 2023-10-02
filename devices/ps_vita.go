package devices

type PSVita struct {
	base
}

var (
	psVitaName       = "PlayStation Vita"
	psVitaMatchRegex = []string{`(?i)PlayStation Vita`}
)

func NewPSVita(userAgent string) *PSVita {
	return &PSVita{
		base{
			userAgent: userAgent,
		},
	}
}

func (p *PSVita) Name() string {
	return psVitaName
}

func (p *PSVita) Match() bool {
	return p.match(psVitaMatchRegex)
}
