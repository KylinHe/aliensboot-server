module github.com/KylinHe/aliensboot-server

go 1.12

replace github.com/KylinHe/aliensboot-core => ../../aliensboot-core

require (
	github.com/KylinHe/aliensboot-core v1.0.18
	github.com/eapache/queue v1.1.0
	github.com/gin-gonic/gin v1.5.0
	github.com/gogo/protobuf v1.3.1
	github.com/magicsea/behavior3go v0.0.0-20190506074935-238cc4d4d098
	gopkg.in/mgo.v2 v2.0.0-20190816093944-a6b53ec6cb22
)
