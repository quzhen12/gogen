module gogen

go 1.14

require (
	github.com/fsnotify/fsnotify v1.4.9 // indirect
	github.com/quzhen12/plugins v0.0.0-20210528082326-0d28a8471f6d
	github.com/smartystreets/goconvey v1.6.4
	github.com/spf13/cobra v1.2.1
	github.com/spf13/viper v1.8.1
	github.com/stretchr/testify v1.7.0 // indirect
	go.uber.org/zap v1.17.0
	golang.org/x/sys v0.0.0-20210112080510-489259a85091 // indirect
	golang.org/x/text v0.3.3 // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
)

replace github.com/quzhen12/plugins => ../plugins

replace github.com/spf13/viper v1.8.1 => github.com/spf13/viper v1.7.1
