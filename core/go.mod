module github.ibm.com/sysflow/sf-processor/core

go 1.14

require (
	github.com/RackSec/srslog v0.0.0-20180709174129-a4725f04ec91
	github.com/actgardner/gogen-avro/v7 v7.1.1 // indirect
	github.com/antlr/antlr4 v0.0.0-20200417160354-8c50731894e0
	github.com/enriquebris/goconcurrentqueue v0.6.0
	github.com/orcaman/concurrent-map v0.0.0-20190826125027-8c72a8bb44f6
	github.com/stretchr/testify v1.6.1
	github.com/sysflow-telemetry/sf-apis/go v0.0.0-20200618213240-a59f3a148871
	github.ibm.com/sysflow/goutils v0.0.0-00010101000000-000000000000
)

replace github.ibm.com/sysflow/goutils => ../modules/goutils

//replace github.com/sysflow-telemetry/sf-apis/go => ../modules/sf-apis/go
replace github.com/sysflow-telemetry/sf-apis/go => ../../sf-apis/go
