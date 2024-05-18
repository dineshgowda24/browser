package devices

type KindleFire struct {
	p Parser
}

var (
	kindleFireName       = "Kindle Fire"
	kindleFireMatchRegex = []string{`Kindle Fire|KFTT`}
)

func NewKindleFire(p Parser) *KindleFire {
	return &KindleFire{
		p: p,
	}
}

func (k *KindleFire) Name() string {
	return kindleFireName
}

func (k *KindleFire) Match() bool {
	return k.p.Match(kindleFireMatchRegex)
}
