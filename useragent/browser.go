package useragent

type BrowserType string

type regexM map[BrowserType]*regexExpect

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
)

const (
	And Operator = "&"
	Or  Operator = "|"
)
