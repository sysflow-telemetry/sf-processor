module github.ibm.com/sysflow/sf-processor/plugins/example

go 1.14

require (
	github.com/sysflow-telemetry/sf-apis/go v0.0.0-20200422203822-89baf17b2999
	github.ibm.com/sysflow/goutils v0.0.0-20200528201643-85683bbabbe4
)

replace github.ibm.com/sysflow/goutils => ../../../goutils
