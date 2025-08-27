module github.com/dayvillefire/newworld-cadview-utils/cmd/cadpublish

go 1.24

toolchain go1.24.3

replace (
	github.com/dayvillefire/newworld-cadview-agent/agent => ../../../newworld-cadview-agent/agent
	github.com/dayvillefire/newworld-cadview-utils/util => ../../util
)

require (
	github.com/bwmarrin/discordgo v0.29.0
	github.com/dayvillefire/newworld-cadview-agent/agent v0.0.0-20250819181530-fe66698c5d3c
	github.com/dayvillefire/newworld-cadview-utils/util v0.0.0-20250819192840-82fd8305ca8c
	github.com/jbuchbinder/shims v0.0.0-20250818154854-22c0ac83b788
	gopkg.in/yaml.v2 v2.4.0
)

require (
	github.com/chromedp/cdproto v0.0.0-20250803210736-d308e07a266d // indirect
	github.com/chromedp/chromedp v0.14.1 // indirect
	github.com/chromedp/sysutil v1.1.0 // indirect
	github.com/go-json-experiment/json v0.0.0-20250813233538-9b1f9ea2e11b // indirect
	github.com/gobwas/httphead v0.1.0 // indirect
	github.com/gobwas/pool v0.2.1 // indirect
	github.com/gobwas/ws v1.4.0 // indirect
	github.com/gorilla/websocket v1.5.3 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/mailru/easyjson v0.9.0 // indirect
	golang.org/x/crypto v0.41.0 // indirect
	golang.org/x/net v0.42.0 // indirect
	golang.org/x/sys v0.35.0 // indirect
	golang.org/x/text v0.28.0 // indirect
	gorm.io/gorm v1.30.1 // indirect
)
