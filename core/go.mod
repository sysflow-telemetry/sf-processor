module github.ibm.com/sysflow/sf-processor/core

go 1.14

require (
	github.com/RackSec/srslog v0.0.0-20180709174129-a4725f04ec91
	github.com/antlr/antlr4 v0.0.0-20200417160354-8c50731894e0
	github.com/containerd/containerd v1.3.4 // indirect
	github.com/docker/distribution v2.7.1+incompatible // indirect
	github.com/docker/docker v17.12.0-ce-rc1.0.20200427224914-45369c61a48c+incompatible
	github.com/docker/go-connections v0.4.0 // indirect
	github.com/docker/go-units v0.4.0 // indirect
	github.com/enriquebris/goconcurrentqueue v0.6.0
	github.com/gogo/protobuf v1.3.1 // indirect
	github.com/opencontainers/go-digest v1.0.0-rc1 // indirect
	github.com/opencontainers/image-spec v1.0.1 // indirect
	github.com/orcaman/concurrent-map v0.0.0-20190826125027-8c72a8bb44f6
	github.com/pkg/errors v0.9.1 // indirect
	github.com/stretchr/testify v1.5.1
	github.com/sysflow-telemetry/sf-apis/go v0.0.0-20200422203822-89baf17b2999
	github.ibm.com/sysflow/goutils v0.0.0-20200528201643-85683bbabbe4
	golang.org/x/net v0.0.0-20200425230154-ff2c4b7c35a0 // indirect
	google.golang.org/grpc v1.29.1 // indirect
)

replace github.ibm.com/sysflow/goutils => ../../goutils
