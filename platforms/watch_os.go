package platforms

type WatchOS struct {
	p Parser
}

var (
	watchOSName          = "Apple Watch OS"
	watchOSVersionRegexp = []string{`(?i)Watch\s*OS[ ,/]([\d.]+)`, `Watch[^/]+/([\d.]+)`}
	watchOSMatchRegexp   = []string{`(?i)Watch\s*OS`, `Watch[\d+]`}
)

func NewWatchOS(p Parser) *WatchOS {
	return &WatchOS{
		p: p,
	}
}

func (w *WatchOS) Name() string {
	return watchOSName
}

func (w *WatchOS) Version() string {
	return w.p.Version(watchOSVersionRegexp, 1, "0.0")
}

func (w *WatchOS) Match() bool {
	return w.p.Match(watchOSMatchRegexp)
}
