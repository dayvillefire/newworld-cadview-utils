module github.com/dayvillefire/newworld-cadview-utils/cmd/cadingest

go 1.20

replace github.com/dayvillefire/newworld-cadview-agent/agent => ../../../newworld-cadview-agent/agent

require (
	github.com/dayvillefire/newworld-cadview-agent/agent v0.0.0-20231028230913-f0bff1ccc92d
	gorm.io/driver/mysql v1.5.2
	gorm.io/gorm v1.25.5
)

require (
	github.com/chromedp/cdproto v0.0.0-20231126232103-8e31ff06e23b // indirect
	github.com/chromedp/chromedp v0.9.3 // indirect
	github.com/chromedp/sysutil v1.0.0 // indirect
	github.com/go-sql-driver/mysql v1.7.1 // indirect
	github.com/gobwas/httphead v0.1.0 // indirect
	github.com/gobwas/pool v0.2.1 // indirect
	github.com/gobwas/ws v1.3.1 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	golang.org/x/sys v0.15.0 // indirect
)
