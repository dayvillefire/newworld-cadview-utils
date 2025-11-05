module github.com/dayvillefire/newworld-cadview-utils/cmd/cadpublish

go 1.25

replace (
	github.com/dayvillefire/newworld-cadview-agent/agent => ../../../newworld-cadview-agent/agent
	github.com/dayvillefire/newworld-cadview-utils => ../../
	github.com/dayvillefire/newworld-cadview-utils/util => ../../util
)

require (
	github.com/bwmarrin/discordgo v0.29.0
	github.com/dayvillefire/newworld-cadview-agent/agent v0.0.0-20251105145945-527f847ee336
	github.com/dayvillefire/newworld-cadview-utils/util v0.0.0-20251029171226-95b01170ec71
	github.com/jbuchbinder/shims v0.0.0-20251029164657-6c80f5d6bc01
	gopkg.in/yaml.v2 v2.4.0
)

require (
	github.com/chromedp/cdproto v0.0.0-20250803210736-d308e07a266d // indirect
	github.com/chromedp/chromedp v0.14.2 // indirect
	github.com/chromedp/sysutil v1.1.0 // indirect
	github.com/go-json-experiment/json v0.0.0-20251027170946-4849db3c2f7e // indirect
	github.com/gobwas/httphead v0.1.0 // indirect
	github.com/gobwas/pool v0.2.1 // indirect
	github.com/gobwas/ws v1.4.0 // indirect
	github.com/gorilla/websocket v1.5.3 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	golang.org/x/crypto v0.43.0 // indirect
	golang.org/x/sys v0.37.0 // indirect
	golang.org/x/text v0.30.0 // indirect
	gorm.io/gorm v1.31.1 // indirect
)
