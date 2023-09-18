package devices

type KindleFire struct {
	base
}

var (
	kindleFireName       = "Kindle Fire"
	kindleFireMatchRegex = []string{`Kindle Fire|KFTT`}
)

func NewKindleFire(userAgent string) *KindleFire {
	return &KindleFire{
		base{
			userAgent: userAgent,
		},
	}
}

func (k *KindleFire) Name() string {
	return kindleFireName
}

func (k *KindleFire) Match() bool {
	return k.match(kindleFireMatchRegex)
}
