package service

import (
	"context"
	"database/sql"
	"homework/exceptionhandle/dao"
	"homework/exceptionhandle/dao/do"
	"homework/exceptionhandle/idl"
)

type ExceptionHandle struct {
}

func (e ExceptionHandle) ToDoSomeThing(ctx context.Context, input *idl.ToDoSomeThingInput, output *idl.ToDoSomeThingOutput) (err error) {
	info := []do.UserInfo{}
	err = dao.MysqlEngine.GetUserInfoByName(input.UserName, input.ClientType, &info)
	if err != nil {
		return err
	}
	if err == sql.ErrNoRows {
		// 降级查找
		err = dao.MysqlEngine.GetUserInfoByName(input.UserName, 0, &info)
		if err != nil {
			return err
		}
	}
	if err == sql.ErrNoRows {
		// 降级查找
		// 通常需要多次降级查找的场景，比如通过查询各个公司下的产品信息、各个产品发布的固件版本信息等等
	}
	// todo ...
	body := make([]idl.ToDoSomeThingOutputBody, 0)
	for _, user := range info {
		u := idl.ToDoSomeThingOutputBody{
			Id:         user.Id,
			UserName:   user.UserName,
			UserGuid:   user.UserGuid,
			ClientType: user.ClientType,
			InsertTime: user.InsertTime,
			UpdateTime: user.UpdateTime,
		}
		body = append(body, u)
	}
	output.Body = body
	output.Code = 0
	output.Count = int64(len(body))
	return nil
}
