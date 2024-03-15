module github.com/dayvillefire/newworld-cadview-utils/cmd/cadpublish

go 1.22

replace (
	github.com/dayvillefire/newworld-cadview-agent/agent => ../../../newworld-cadview-agent/agent
	github.com/dayvillefire/newworld-cadview-utils/util => ../../util
)

require (
	github.com/bwmarrin/discordgo v0.27.1
	github.com/dayvillefire/newworld-cadview-agent/agent v0.0.0-20240315224611-a4e6e0837d6b
	github.com/dayvillefire/newworld-cadview-utils/util v0.0.0-20240107151630-34a4eca519d3
	github.com/jbuchbinder/shims v0.0.0-20240127163204-18a2ea0be2dc
	gopkg.in/yaml.v2 v2.4.0
)

require (
	github.com/chromedp/cdproto v0.0.0-20240312231614-1e5096e63154 // indirect
	github.com/chromedp/chromedp v0.9.5 // indirect
	github.com/chromedp/sysutil v1.0.0 // indirect
	github.com/gobwas/httphead v0.1.0 // indirect
	github.com/gobwas/pool v0.2.1 // indirect
	github.com/gobwas/ws v1.3.2 // indirect
	github.com/gorilla/websocket v1.5.1 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	golang.org/x/crypto v0.21.0 // indirect
	golang.org/x/net v0.22.0 // indirect
	golang.org/x/sys v0.18.0 // indirect
	gorm.io/gorm v1.25.7 // indirect
)
