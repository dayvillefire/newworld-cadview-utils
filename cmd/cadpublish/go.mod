module github.com/dayvillefire/newworld-cadview-utils/cmd/cadpublish

go 1.23

replace (
	github.com/dayvillefire/newworld-cadview-agent/agent => ../../../newworld-cadview-agent/agent
	github.com/dayvillefire/newworld-cadview-utils/util => ../../util
)

require (
	github.com/bwmarrin/discordgo v0.28.1
	github.com/dayvillefire/newworld-cadview-agent/agent v0.0.0-20240820164125-6c5ba7d7c5e1
	github.com/dayvillefire/newworld-cadview-utils/util v0.0.0-20240820164235-c54df6dcd1c6
	github.com/jbuchbinder/shims v0.0.0-20240506232043-4fac4ec97ccb
	gopkg.in/yaml.v2 v2.4.0
)

require (
	github.com/chromedp/cdproto v0.0.0-20241030022559-23c28aebe8cb // indirect
	github.com/chromedp/chromedp v0.11.1 // indirect
	github.com/chromedp/sysutil v1.1.0 // indirect
	github.com/gobwas/httphead v0.1.0 // indirect
	github.com/gobwas/pool v0.2.1 // indirect
	github.com/gobwas/ws v1.4.0 // indirect
	github.com/gorilla/websocket v1.5.3 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	golang.org/x/crypto v0.28.0 // indirect
	golang.org/x/net v0.26.0 // indirect
	golang.org/x/sys v0.26.0 // indirect
	golang.org/x/text v0.19.0 // indirect
	gorm.io/gorm v1.25.12 // indirect
)
