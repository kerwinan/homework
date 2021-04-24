package dao

import (
	xorm "github.com/go-xorm/xorm"
	"homework/exceptionhandle/dao/do"
)

type MysqlDb struct {
	syncEngine *xorm.Engine
}

func newMysqlDb() {
	// todo init mysql db connect
}

func (m *MysqlDb) GetUserInfoByName(userName string, clientType int32, info *[]do.UserInfo) (err error) {
	session := m.syncEngine.NewSession()
	if clientType != 0 {
		session.Where("client_type = ?", clientType)
	}
	if len(userName) != 0 {
		session.Where("user_name = ?", userName)
	}
	findSession := session.Clone()
	return findSession.Find(info)
}
