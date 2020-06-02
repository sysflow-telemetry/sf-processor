module github.ibm.com/sysflow/sf-processor/driver

go 1.14

require (
	github.com/actgardner/gogen-avro v6.5.0+incompatible
	github.com/antlr/antlr4 v0.0.0-20200417160354-8c50731894e0
	github.com/dgraph-io/dgo v1.0.0 // indirect
	github.com/dgraph-io/dgo/v2 v2.2.0
	github.com/fsnotify/fsnotify v1.4.9 // indirect
	github.com/golang/snappy v0.0.1 // indirect
	github.com/mattn/go-runewidth v0.0.9 // indirect
	github.com/mitchellh/mapstructure v1.2.2 // indirect
	github.com/nsf/termbox-go v0.0.0-20200204031403-4d2b513ad8be // indirect
	github.com/orcaman/concurrent-map v0.0.0-20190826125027-8c72a8bb44f6
	github.com/pelletier/go-toml v1.7.0 // indirect
	github.com/spf13/afero v1.2.2 // indirect
	github.com/spf13/cast v1.3.1 // indirect
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/spf13/viper v1.6.3
	github.com/sysflow-telemetry/sf-apis/go v0.0.0-20200602024344-4a71552fc529
	github.ibm.com/sysflow/goutils v0.0.0-20200528201643-85683bbabbe4
	github.ibm.com/sysflow/sf-processor/core v0.0.0-20200417193244-61d8d9d5918f
	golang.org/x/sys v0.0.0-20200413165638-669c56c373c4 // indirect
	google.golang.org/grpc v1.29.1
	gopkg.in/ini.v1 v1.55.0 // indirect
)

replace github.ibm.com/sysflow/sf-processor/core => ../core
replace github.ibm.com/sysflow/goutils => ../../goutils
