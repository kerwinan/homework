package main

import (
	"fmt"
	"homework/redis/dao"
)

func main() {
	cli := dao.NewRedisClient()
	value, _ := cli.Get("key1")
	fmt.Println(value)
}

func doSet()  {
	
}