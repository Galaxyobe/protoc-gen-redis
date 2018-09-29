/*-------------- Copyright (c) Shenzhen BB Team. -------------------

 File    : implement.go
 Time    : 2018/9/28 14:01
 Author  : yanue
 
 - 
 
------------------------------- go ---------------------------------*/

package impl

import (
	"context"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"protoc-gen-redis/test"
	"testing"
	"time"
)

var tpl DemoRedis

func init() {
	// 建立连接
	client := &redis.Pool{
		// 从配置文件获取maxidle以及maxactive，取不到则用后面的默认值
		MaxIdle:     10,
		MaxActive:   10,
		IdleTimeout: 180 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", "192.168.5.201:6379")
			if err != nil {
				return nil, err
			}
			// 选择db
			c.Do("SELECT", 0)
			return c, nil
		},
	}

	tpl.Demo = new(test.Demo)
	tpl.rp = client

	tpl.Demo.SomeInteger = int32(time.Now().Unix())
	fmt.Println("tpl.Demo", tpl.Demo)
}

func TestDemoRedis_LoadFromRedis(t *testing.T) {
	err := tpl.StoreToRedis(context.Background(), "test1", 10)
	fmt.Println("StoreToRedis err:", err)
	if err != nil {
		return
	}
	err = tpl.LoadFromRedis(context.Background(), "test1")
	fmt.Println("LoadFromRedis", err, tpl.Demo)

}

func TestDemoRedis_LoadFromRedisByHash(t *testing.T) {
	err := tpl.StoreToRedisByHash(context.Background(), "test", 0)
	fmt.Println("StoreToRedis", err)
	if err != nil {
		return
	}
	err = tpl.LoadFromRedisByHash(context.Background(), "test")
	fmt.Println("LoadFromRedis", err, tpl.Demo)

}

func TestDemoRedis_LoadFromRedisByKey(t *testing.T) {
	err := tpl.StoreToRedisByKey(context.Background(), "test", "SomeInteger", "123231231", 0)
	fmt.Println("StoreToRedis", err)
	if err != nil {
		return
	}
	reply, err := tpl.LoadFromRedisByKey(context.Background(), "test", "SomeInteger")
	fmt.Println("LoadFromRedis", err, tpl.Demo, fmt.Sprintf("%s", reply))

}
