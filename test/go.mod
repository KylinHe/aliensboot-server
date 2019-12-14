module test

go 1.12

replace github.com/KylinHe/aliensboot-server => ../src

replace github.com/KylinHe/aliensboot-core => ../../aliensboot-core

require (
	github.com/KylinHe/aliensboot-core v1.0.23
	github.com/KylinHe/aliensboot-server v0.0.1
	github.com/gogo/protobuf v1.3.1
	gopkg.in/mgo.v2 v2.0.0-20190816093944-a6b53ec6cb22
)
