package devices

type Kindle struct {
	p Parser
}

var (
	kindleName       = "Kindle"
	kindleMatchRegex = []string{`Kindle`}
)

func NewKindle(p Parser) *Kindle {
	return &Kindle{
		p: p,
	}
}

func (k *Kindle) Name() string {
	return kindleName
}

func (k *Kindle) Match() bool {
	return k.p.Match(kindleMatchRegex)
}
