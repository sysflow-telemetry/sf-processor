module github.ibm.com/sysflow/sf-processor/processor

go 1.14

require (
	github.com/actgardner/gogen-avro v7.0.0+incompatible
	github.com/antlr/antlr4 v0.0.0-20200412020049-7a11432ede99 // indirect
	github.com/dgraph-io/dgo v1.0.0 // indirect
	github.com/dgraph-io/dgo/v2 v2.2.0
	github.com/golang/snappy v0.0.1 // indirect
	github.com/mattn/go-runewidth v0.0.9 // indirect
	github.com/nsf/termbox-go v0.0.0-20200204031403-4d2b513ad8be // indirect
	github.com/spf13/viper v1.6.3
	github.com/sysflow-telemetry/sf-apis/go v0.0.0-20200415202402-e5659ec14bfd
	github.ibm.com/sysflow/sf-processor/common v0.0.0-20200416031336-865e8ef0b5bb
	github.ibm.com/sysflow/sf-processor/plugins/flattener v0.0.0-20200415205203-9fa361cf78fd
	github.ibm.com/sysflow/sf-processor/plugins/processor v0.0.0-20200415205203-9fa361cf78fd
	github.ibm.com/sysflow/sf-processor/plugins/sfpe v0.0.0-20200416031336-865e8ef0b5bb
	google.golang.org/grpc v1.23.0
)
