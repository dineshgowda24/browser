package platforms

type WatchOS struct {
	base
}

var (
	watchOSName          = "Apple Watch OS"
	watchOSVersionRegexp = []string{`(?i)Watch\s*OS[ ,/]([\d.]+)`, `Watch[^/]+/([\d.]+)`}
	watchOSMatchRegexp   = []string{`(?i)Watch\s*OS`, `Watch[\d+]`}
)

func NewWatchOS(userAgent string) *WatchOS {
	return &WatchOS{
		base{
			userAgent: userAgent,
		},
	}
}

func (w *WatchOS) Name() string {
	return watchOSName
}

func (w *WatchOS) Version() string {
	return w.version(watchOSVersionRegexp, 1, "0.0")
}

func (w *WatchOS) Match() bool {
	return w.match(watchOSMatchRegexp)
}
