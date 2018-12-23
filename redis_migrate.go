package main

import (
	"github.com/redis_migrate/redisUtil"
	"fmt"
)

func main() {
	sourceAddr := "192.168.244.10:6379"
	//destAddr := "192.168.244.20:6379"
	//redisUtil.BgSaveAndCheck(sourceAddr)
	fmt.Println(redisUtil.GetRDBPath(sourceAddr))

	//redisUtil.CopySlotInfo(sourceAddr,destAddr)
}
