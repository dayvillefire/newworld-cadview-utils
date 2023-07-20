module github.com/dayvillefire/newworld-cadview-utils/cmd/cadpublish

go 1.20

replace (
	github.com/dayvillefire/newworld-cadview-agent/agent => ../../../newworld-cadview-agent/agent
	github.com/dayvillefire/newworld-cadview-utils/util => ../../util
)

require (
	github.com/bwmarrin/discordgo v0.27.1
	github.com/dayvillefire/newworld-cadview-agent/agent v0.0.0-20230531144245-2e2a6797601a
	github.com/dayvillefire/newworld-cadview-utils/util v0.0.0-20230301122230-b60519495264
	github.com/jbuchbinder/shims v0.0.0-20230627161102-16aacaee8147
	gopkg.in/yaml.v2 v2.4.0
)

require (
	github.com/chromedp/cdproto v0.0.0-20230716001748-3ed7c525ec8b // indirect
	github.com/chromedp/chromedp v0.9.1 // indirect
	github.com/chromedp/sysutil v1.0.0 // indirect
	github.com/gobwas/httphead v0.1.0 // indirect
	github.com/gobwas/pool v0.2.1 // indirect
	github.com/gobwas/ws v1.2.1 // indirect
	github.com/gorilla/websocket v1.5.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	golang.org/x/crypto v0.11.0 // indirect
	golang.org/x/sys v0.10.0 // indirect
	gorm.io/gorm v1.25.2 // indirect
)
