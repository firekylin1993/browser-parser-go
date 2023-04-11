package useragent

import (
	"regexp"
	"strings"
)

type Operator string

// UserAgent ua对象
type UserAgent struct {
	ua      []byte
	regexs  regexM
	mobile  bool
	builtin bool
}

var gbRegex regexM

type regexExpect struct {
	Oper Operator
	epxs []expectParam
}

type expectParam struct {
	regex  *regexp.Regexp
	expect bool
}

// Init 绑定正则
func Init() {
	gbRegex = map[BrowserType]*regexExpect{
		IsMobile: registeRegexp(
			And,
			expectParam{regex: regexp.MustCompile(strings.Join(mobileType, "|")), expect: true},
		),
		IsBuiltIn: registeRegexp(
			And,
			expectParam{regex: regexp.MustCompile("MicroMessenger"), expect: true},
			expectParam{regex: regexp.MustCompile("wxwork"), expect: true},
		),
		IsFireFox: registeRegexp(And, expectParam{regex: regexp.MustCompile("Firefox"), expect: true}),
		IsOpera: registeRegexp(
			Or,
			expectParam{regex: regexp.MustCompile("OPR"), expect: true},
			expectParam{regex: regexp.MustCompile("Opera"), expect: true},
		),
		IsSogou: registeRegexp(
			Or,
			expectParam{regex: regexp.MustCompile("SE"), expect: true},
			expectParam{regex: regexp.MustCompile("MetaSr"), expect: true},
		),
		IsEdge: registeRegexp(
			Or,
			expectParam{regex: regexp.MustCompile("Edg"), expect: true},
			expectParam{regex: regexp.MustCompile("Edge"), expect: true},
		),
		IsQQbrowser: registeRegexp(And, expectParam{regex: regexp.MustCompile("QQBrowser"), expect: true}),
		IsBaidu:     registeRegexp(And, expectParam{regex: regexp.MustCompile("Baidu"), expect: true}),
		IsQuark:     registeRegexp(And, expectParam{regex: regexp.MustCompile("Quark"), expect: true}),
		IsUC: registeRegexp(
			Or,
			expectParam{regex: regexp.MustCompile("UCBrowser"), expect: true},
			expectParam{regex: regexp.MustCompile("UCWEB"), expect: true},
		),
		//safari，不可以有chrome关键字
		IsSafari: registeRegexp(
			And,
			expectParam{regex: regexp.MustCompile("Safari"), expect: true},
			expectParam{regex: regexp.MustCompile("Chrome"), expect: false},
		),
		//谷歌浏览器
		IsChrome: registeRegexp(
			And,
			expectParam{regex: regexp.MustCompile("Chrome"), expect: true},
			expectParam{regex: regexp.MustCompile("Quark"), expect: false},
			expectParam{regex: regexp.MustCompile("UCBrowser"), expect: false},
			expectParam{regex: regexp.MustCompile("UCWEB"), expect: false},
			expectParam{regex: regexp.MustCompile("QQBrowser"), expect: false},
			expectParam{regex: regexp.MustCompile("Edge"), expect: false},
			expectParam{regex: regexp.MustCompile("Edg"), expect: false},
			expectParam{regex: regexp.MustCompile("Baidu"), expect: false},
			expectParam{regex: regexp.MustCompile("DingTalk"), expect: false},
			expectParam{regex: regexp.MustCompile("Feishu"), expect: false},
			expectParam{regex: regexp.MustCompile("MicroMessenger"), expect: false},
			expectParam{regex: regexp.MustCompile("wxwork"), expect: false},
		),
		//ie浏览器。搜狗使用ie内核
		IsIE: registeRegexp(
			And,
			expectParam{regex: regexp.MustCompile("MSIE"), expect: true},
			expectParam{regex: regexp.MustCompile("SE"), expect: false},
			expectParam{regex: regexp.MustCompile("MetaSr"), expect: false},
		),
		//微信浏览器
		IsWx: registeRegexp(
			And,
			expectParam{regex: regexp.MustCompile("MicroMessenger"), expect: true},
			expectParam{regex: regexp.MustCompile("wxwork"), expect: false},
		),
		IsDingDing: registeRegexp(And, expectParam{regex: regexp.MustCompile("DingTalk"), expect: true}),
		IsFeiShu:   registeRegexp(And, expectParam{regex: regexp.MustCompile("Feishu"), expect: true}),
	}
}

func registeRegexp(operator Operator, rgxs ...expectParam) *regexExpect {
	re := new(regexExpect)
	re.Oper = operator
	for _, v := range rgxs {
		re.epxs = append(re.epxs, v)
	}
	return re
}

// NewUserAgent 初始化ua对象
func NewUserAgent(ua string) *UserAgent {
	return &UserAgent{
		ua:     []byte(ua),
		regexs: gbRegex,
	}
}

// Browser 判断ua类型
func (p *UserAgent) Browser(b BrowserType) bool {
	var ret bool
	if p.regexs[b].Oper == And {
		ret = true
		for _, v := range p.regexs[b].epxs {
			if v.regex.Match(p.ua) != v.expect {
				ret = false
				break
			}
		}
	} else {
		for _, v := range p.regexs[b].epxs {
			if v.regex.Match(p.ua) == v.expect {
				ret = true
				break
			}
		}
	}
	return ret
}
