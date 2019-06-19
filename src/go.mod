module github.com/KylinHe/aliensboot-server

go 1.12

replace github.com/KylinHe/aliensboot-core => ../../aliensboot-core

require (
	github.com/KylinHe/aliensboot-core v0.1.2
	github.com/eapache/queue v1.1.0
	github.com/gin-gonic/gin v1.4.0
	github.com/gogo/protobuf v1.2.1
	github.com/magicsea/behavior3go v0.0.0-20190506074935-238cc4d4d098
	go.uber.org/multierr v1.1.0 // indirect
)
