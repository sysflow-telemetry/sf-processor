//
// Copyright (C) 2020 IBM Corporation.
//
// Authors:
// Frederico Araujo <frederico.araujo@ibm.com>
// Teryl Taylor <terylt@ibm.com>
//
module github.ibm.com/sysflow/sf-processor/driver

go 1.14

require (
	github.com/Shopify/sarama v0.0.0-00010101000000-000000000000
	github.com/actgardner/gogen-avro v6.5.0+incompatible
	github.com/actgardner/gogen-avro/v7 v7.1.1
	github.com/elastic/beats/v7 v7.8.1
	github.com/fsnotify/fsnotify v1.4.9 // indirect
	github.com/linkedin/goavro v2.1.0+incompatible
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

replace github.ibm.com/sysflow/goutils => ../../goutils

replace github.com/sysflow-telemetry/sf-apis/go => ../../sf-apis/go

replace (
	github.com/Shopify/sarama => github.com/elastic/sarama v1.19.1-0.20200629123429-0e7b69039eec
	github.com/dop251/goja => github.com/andrewkroh/goja v0.0.0-20190128172624-dd2ac4456e20
	github.com/fsnotify/fsevents => github.com/elastic/fsevents v0.0.0-20181029231046-e1d381a4d270
	github.com/fsnotify/fsnotify => github.com/adriansr/fsnotify v0.0.0-20180417234312-c9bbe1f46f1d
)
