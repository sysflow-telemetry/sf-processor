module github.ibm.com/sysflow/sf-processor/driver

go 1.14

require (
	github.com/actgardner/gogen-avro v6.5.0+incompatible
	github.com/actgardner/gogen-avro/v7 v7.1.1
	github.com/fsnotify/fsnotify v1.4.9 // indirect
	github.com/golang/snappy v0.0.1 // indirect
	github.com/mitchellh/mapstructure v1.2.2 // indirect
	github.com/pelletier/go-toml v1.7.0 // indirect
	github.com/spf13/afero v1.2.2 // indirect
	github.com/spf13/cast v1.3.1 // indirect
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/spf13/viper v1.6.3
	github.com/sysflow-telemetry/sf-apis/go v0.0.0-20200618213240-a59f3a148871
	github.ibm.com/sysflow/goutils v0.0.0-20200619144433-a13c12f45010
	github.ibm.com/sysflow/sf-processor/core v0.0.0-20200417193244-61d8d9d5918f
	golang.org/x/sys v0.0.0-20200413165638-669c56c373c4 // indirect
	gopkg.in/ini.v1 v1.55.0 // indirect
)

replace github.ibm.com/sysflow/sf-processor/core => ../core

replace github.ibm.com/sysflow/goutils => ../modules/goutils

replace github.com/sysflow-telemetry/sf-apis/go => ../modules/sf-apis/go