module github.ibm.com/sysflow/sf-processor/plugins/example

go 1.14

require (
	github.com/sysflow-telemetry/sf-apis/go v0.0.0-20200618213240-a59f3a148871
	github.ibm.com/sysflow/goutils v0.0.0-20200528201643-85683bbabbe4
)

replace github.ibm.com/sysflow/goutils => ../../modules/goutils

replace github.com/sysflow-telemetry/sf-apis/go => ../../modules/sf-apis/go