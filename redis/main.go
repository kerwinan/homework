package main

import (
	"fmt"
	"homework/redis/dao"
	"homework/redis/internal"
	"sync"
)

var (
	countNum = 100000
	goNum    = 100
	ValueSize = 10

	keyNum = 0
)

func main() {
	keyNum = countNum / goNum
	mulSet(ValueSize, 0, countNum)
	//mulDel(0, countNum)
}

func mulSet(valueLength int, index, count int) {
	wg := sync.WaitGroup{}
	t := 0
	wg.Add(goNum)
	for i := 0; i < goNum; i++ {
		go func(valueLength int, in, num int) {
			doSet(valueLength, in, num)
			wg.Done()
		}(valueLength, t, keyNum)
		t += keyNum
	}
	wg.Wait()
}

func doSet(valueLength int, index, count int) {
	cli := dao.NewRedisClient("192.168.0.105:6379")
	fmt.Println(valueLength, index, count)
	for i := 0; i < count; i++ {
		key := fmt.Sprintf("%06d", index)
		err := cli.Set(key, internal.RandomString(valueLength))
		if err != nil {
			fmt.Printf("set err: %v", err)
			break
		}
		index++
	}
}

func mulDel(index, count int) {
	wg := sync.WaitGroup{}
	t := 0
	wg.Add(goNum)
	for i := 0; i < goNum; i++ {
		go func(in, num int) {
			//doSet(valueLength, in, num)
			doDel(in, num)
			wg.Done()
		}(t, keyNum)
		t += keyNum
	}
	wg.Wait()
}

func doDel(index, count int) {
	cli := dao.NewRedisClient("192.168.0.105:6379")
	fmt.Println(index, count)
	for i := 0; i < count; i++ {
		key := fmt.Sprintf("%06d", index)
		err := cli.Del(key)
		if err != nil {
			fmt.Printf("set err: %v", err)
			break
		}
		index++
	}
}
