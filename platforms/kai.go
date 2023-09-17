package platforms

type KaiOS struct {
	base
}

var (
	kaiOSName          = "Kai OS"
	kaiOSVersionRegexp = []string{`KaiOS/([\d.]+)`}
	kaiOSMatchRegexp   = []string{`KaiOS`}
)

func NewKaiOS(userAgent string) *KaiOS {
	return &KaiOS{
		base{
			userAgent: userAgent,
		},
	}
}

func (k *KaiOS) Name() string {
	return kaiOSName
}

func (k *KaiOS) Version() string {
	return k.version(kaiOSVersionRegexp, 1, "")
}

func (k *KaiOS) Match() bool {
	return k.match(kaiOSMatchRegexp)
}
