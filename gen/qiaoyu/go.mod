module qiaoyu

go 1.14

require (
	github.com/gin-gonic/gin v1.7.2
	github.com/jinzhu/gorm v1.9.16
	github.com/natefinch/lumberjack v2.0.0+incompatible // indirect
	github.com/quzhen12/plugins v0.0.0-20210528082326-0d28a8471f6d
	github.com/spf13/viper v1.7.1 // indirect
	go.uber.org/zap v1.17.0
	gorm.io/driver/mysql v1.1.0 // indirect
)

replace github.com/quzhen12/plugins v0.0.0-20210528082326-0d28a8471f6d => ../plugins
