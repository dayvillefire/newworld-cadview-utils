module github.com/dayvillefire/newworld-cadview-utils/cmd/cadingest

go 1.22

replace github.com/dayvillefire/newworld-cadview-agent/agent => ../../../newworld-cadview-agent/agent

require (
	github.com/dayvillefire/newworld-cadview-agent/agent v0.0.0-20240703195347-1736dac6778d
	gorm.io/driver/mysql v1.5.7
	gorm.io/gorm v1.25.10
)

require (
	filippo.io/edwards25519 v1.1.0 // indirect
	github.com/chromedp/cdproto v0.0.0-20240626232640-f933b107c653 // indirect
	github.com/chromedp/chromedp v0.9.5 // indirect
	github.com/chromedp/sysutil v1.0.0 // indirect
	github.com/go-sql-driver/mysql v1.8.1 // indirect
	github.com/gobwas/httphead v0.1.0 // indirect
	github.com/gobwas/pool v0.2.1 // indirect
	github.com/gobwas/ws v1.4.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	golang.org/x/sys v0.21.0 // indirect
)
