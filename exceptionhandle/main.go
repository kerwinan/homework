package main

/*
	Q: 我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？
*/

/*
	A:
		1. 当使用特定条件去 mysql 查询信息时，若 mysql 返回 sql.ErrNoRows 信息，不使用 Wrap 进行 error 的封装，此时 error 返回 nil 即可，上层业务判断返回结果时
		发现即没有错误也没有信息，则此时可以进行降级再次查询。
*/

func main() {
	// todo something

	if err != nil {
		return
	}
}
