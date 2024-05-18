package platforms

type KaiOS struct {
	p Parser
}

var (
	kaiOSName          = "Kai OS"
	kaiOSVersionRegexp = []string{`KaiOS/([\d.]+)`}
	kaiOSMatchRegexp   = []string{`KaiOS`}
)

func NewKaiOS(p Parser) *KaiOS {
	return &KaiOS{
		p: p,
	}
}

func (k *KaiOS) Name() string {
	return kaiOSName
}

func (k *KaiOS) Version() string {
	return k.p.Version(kaiOSVersionRegexp, 1, "")
}

func (k *KaiOS) Match() bool {
	return k.p.Match(kaiOSMatchRegexp)
}
