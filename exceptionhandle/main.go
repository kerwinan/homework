package main

import (
	"database/sql"
	"fmt"
	"homework/exceptionhandle/dao"
	"homework/exceptionhandle/dao/do"
)

/*
	Q: 我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？
*/

/*
	A:
		1. 每张 db 表存储着含义不同的信息，故当查询特定信息时，若 mysql 返回 sql.ErrNoRows 信息则应该 Wrap 这个 error，进行错误信息的封装，能够从打印的信息中快速
		定位是查询什么信息返回结果为空
		2. 在某些场景下，当 mysql 返回 sql.ErrNoRows 信息后，可以进行降级处理
*/

func main() {
	userGuid := "f8bf3cf0c3b37c8138b7aa22edb7a916"
	err := todoSomeThing(userGuid)
	if err != nil {
		fmt.Println(err)
		return
	}
	// todo ...
}

func todoSomeThing(userGuid string) (err error) {
	info := do.UserInfo{}
	err = dao.MysqlEngine.GetUserInfoByGuid(userGuid, &info)
	if err == sql.ErrNoRows {
		// todo 降级处理
	}
	if err != nil {
		return
	}
	// todo ...
	return nil
}
