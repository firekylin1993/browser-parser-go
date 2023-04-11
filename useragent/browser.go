package useragent

type BrowserType string

type regexM map[BrowserType]*regexExpect

var mobileType = []string{
	"Mozilla.*Mobile",
	"Windows.*Phone",
	"iPhone Android.*Mobile",
	"BlackBerry.*Mobile",
	"MeeGo",
	"SymbianOS",
	"Opera.*Mini",
	"Chrome.*Mobile",
	"UCBrowser.*Mobile",
}

const (
	IsMobile    BrowserType = "mobile"
	IsBuiltIn   BrowserType = "builtIn"
	IsFireFox   BrowserType = "firefox"
	IsOpera     BrowserType = "opera"
	IsSogou     BrowserType = "sogou"
	IsEdge      BrowserType = "edge"
	IsQQbrowser BrowserType = "qqbrowser"
	IsBaidu     BrowserType = "baidu"
	IsQuark     BrowserType = "quack"
	IsUC        BrowserType = "uc"
	IsSafari    BrowserType = "safari"
	IsChrome    BrowserType = "chrome"
	IsIE        BrowserType = "ie"
	IsWx        BrowserType = "wx"
	IsDingDing  BrowserType = "dingding"
	IsFeiShu    BrowserType = "feishu"
)

const (
	And Operator = "&"
	Or  Operator = "|"
)
