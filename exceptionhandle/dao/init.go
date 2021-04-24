package dao

var (
	MysqlEngine *MysqlDb
)

func Init()  {
	// init db
	newMysqlDb()
}