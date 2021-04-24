package dao

import (
	"database/sql"
	xorm "github.com/go-xorm/xorm"
	"github.com/pkg/errors"
	"homework/exceptionhandle/dao/do"
)

type MysqlDb struct {
	syncEngine *xorm.Engine
}

func newMysqlDb() {
	// todo init mysql db connect
}

func (m *MysqlDb) GetUserInfoByGuid(guid string, info *do.UserInfo) (err error) {
	_, err = m.syncEngine.Where("user_guid = ?", guid).Get(info)
	if err == sql.ErrNoRows {
		return errors.Wrap(err, "get user info")
	}
	return
}
