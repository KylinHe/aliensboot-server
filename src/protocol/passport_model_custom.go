package protocol

import "github.com/KylinHe/aliensboot-core/database"

func (m *User) Copy() database.IData {
	data := *m
	return &data
}

func (m *User) GetDataId() interface{} {
	return m.GetId()
}
