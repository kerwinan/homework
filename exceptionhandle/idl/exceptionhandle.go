package idl

import "context"

type ExceptionHandleIDL interface {
	ToDoSomeThing(ctx context.Context, input *ToDoSomeThingInput, output *ToDoSomeThingOutput) (err error)
}

type ToDoSomeThingInput struct {
	Input
	UserName   string
	ClientType int32
}

type ToDoSomeThingOutput struct {
	Output
	Body []ToDoSomeThingOutputBody
}

type ToDoSomeThingOutputBody struct {
	Id         int64
	UserName   string
	UserGuid   string
	ClientType int32
	InsertTime int64
	UpdateTime int64
}

type Input struct {
	RequestID string `json:"request_id"`
	Cmd       string `json:"cmd"`
	Page      int64  `json:"page"`
	PageSize  int64  `json:"page_size"`
}

type Output struct {
	RequestID string `json:"request_id"`
	Count     int64  `json:"count"`
	Code      int64  `json:"code"`
	CodeMsg   string `json:"code_msg"`
	TraceID   string `json:"trace_id"`
}
