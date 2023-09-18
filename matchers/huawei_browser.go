package matchers

type HuaweiBrowser struct {
	base
}

var (
	huaweiBrowserName          = "Huawei Browser"
	huaweiBrowserVersionRegexp = []string{`(?i)(?:HuaweiBrowser|HBPC)/([\d.]+)`}
	huaweiBrowserMatchRegexp   = []string{`(?i)(HuaweiBrowser|HBPC)`}
)

func NewHuaweiBrowser(userAgent string) *HuaweiBrowser {
	return &HuaweiBrowser{
		base: newBase(userAgent),
	}
}

func (h *HuaweiBrowser) Name() string {
	return huaweiBrowserName
}

func (h *HuaweiBrowser) Version() string {
	return h.version(huaweiBrowserVersionRegexp, 1)
}

func (h *HuaweiBrowser) Match() bool {
	return h.match(huaweiBrowserMatchRegexp)
}
