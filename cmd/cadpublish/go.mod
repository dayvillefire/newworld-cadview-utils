module github.com/dayvillefire/newworld-cadview-utils/cmd/cadpublish

go 1.23

replace (
	github.com/dayvillefire/newworld-cadview-agent/agent => ../../../newworld-cadview-agent/agent
	github.com/dayvillefire/newworld-cadview-utils/util => ../../util
)

require (
	github.com/bwmarrin/discordgo v0.28.1
	github.com/dayvillefire/newworld-cadview-agent/agent v0.0.0-20241220200249-1e4d2b17b36e
	github.com/dayvillefire/newworld-cadview-utils/util v0.0.0-20241220200326-ad94b5bfee75
	github.com/jbuchbinder/shims v0.0.0-20240506232043-4fac4ec97ccb
	gopkg.in/yaml.v2 v2.4.0
)

require (
	github.com/chromedp/cdproto v0.0.0-20241222144035-c16d098c0fb6 // indirect
	github.com/chromedp/chromedp v0.11.2 // indirect
	github.com/chromedp/sysutil v1.1.0 // indirect
	github.com/gobwas/httphead v0.1.0 // indirect
	github.com/gobwas/pool v0.2.1 // indirect
	github.com/gobwas/ws v1.4.0 // indirect
	github.com/gorilla/websocket v1.5.3 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/mailru/easyjson v0.9.0 // indirect
	golang.org/x/crypto v0.31.0 // indirect
	golang.org/x/net v0.26.0 // indirect
	golang.org/x/sys v0.28.0 // indirect
	golang.org/x/text v0.21.0 // indirect
	gorm.io/gorm v1.25.12 // indirect
)
