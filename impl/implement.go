/*-------------- Copyright (c) Shenzhen BB Team. -------------------

 File    : implement.go
 Time    : 2018/9/28 14:01
 Author  : yanue
 
 - 
 
------------------------------- go ---------------------------------*/

package impl

import (
	"context"
	"errors"
	"github.com/golang/protobuf/proto"
	"github.com/gomodule/redigo/redis"
	"protoc-gen-redis/test"
	"reflect"
	"strings"
)

type DemoRedis struct {
	*test.Demo
	rp *redis.Pool
}

// load DemoRedis from redis
func (this *DemoRedis) LoadFromRedis(ctx context.Context, key string) error {
	// redis conn
	conn := this.rp.Get()
	defer conn.Close()

	// load data from redis string
	data, err := redis.Bytes(conn.Do("GET", key))
	if err != nil {
		return err
	}

	// unmarshal data to Demo
	err = proto.Unmarshal(data, this.Demo)
	if err != nil {
		return err
	}

	return nil
}

// load DemoRedis from redis
func (this *DemoRedis) LoadFromRedisByHash(ctx context.Context, key string) error {
	// redis conn
	conn := this.rp.Get()
	defer conn.Close()

	// load data from redis string
	data, err := redis.Values(conn.Do("HGETALL", key))
	if err != nil {
		return err
	}

	err = redis.ScanStruct(data, this.Demo)

	return err
}

// store DemoRedis to redis
// DemoRedis will not expire when ttl is 0
func (this *DemoRedis) StoreToRedis(ctx context.Context, key string, ttl int) error {
	// redis conn
	conn := this.rp.Get()
	defer conn.Close()

	// marshal DemoRedis to []byte
	data, err := proto.Marshal(this.Demo)
	if err != nil {
		return err
	}

	if ttl > 0 {
		// use redis string store the DemoRedis data with expire second
		_, err = conn.Do("SETEX", key, ttl, data)
	} else {
		// use redis string store the DemoRedis data with expire second
		_, err = conn.Do("SET", key, data)
	}

	return err
}

// store DemoRedis to redis by one key
// key param is full match with original struct name
// DemoRedis will not expire when ttl is 0
func (this *DemoRedis) StoreToRedisByHash(ctx context.Context, key string, ttl int) error {
	// redis conn
	conn := this.rp.Get()
	defer conn.Close()

	var err error

	fields := make([]interface{}, 0)

	v := reflect.ValueOf(this.Demo)
	if v.IsValid() == false {
		return errors.New("reflect is InValid")
	}

	//找到最后指向的值，或者空指针，空指针是需要进行初始化的
	for v.Kind() == reflect.Ptr && !v.IsNil() {
		v = v.Elem()
	}

	st := v.Type()
	for i := 0; i < v.NumField(); i++ {
		// 判断是否为可导出字段
		if v.Field(i).CanInterface() {
			name := st.FieldByIndex([]int{i}).Name
			if strings.Contains(name, "XXX_") {
				continue
			}
			fields = append(fields, name, v.Field(i).Interface())
		}
	}

	if len(fields) == 0 {
		return errors.New("struct has no fields can export")
	}

	args := make([]interface{}, 0)
	args = append(args, key)
	args = append(args, fields...)

	// set expire
	if ttl > 0 {
		conn.Send("MULTI")
		conn.Send("HMSET", args...)
		conn.Send("EXPIRE", key, ttl)
		_, err = conn.Do("EXEC")
	} else {
		_, err = conn.Do("HMSET", args...)
	}

	return err
}

// store DemoRedis to redis by one hash key
// field param is full match with original struct name
func (this *DemoRedis) LoadFromRedisByKey(ctx context.Context, key string, field string) (interface{}, error) {
	// redis conn
	conn := this.rp.Get()
	defer conn.Close()

	// get by key
	return conn.Do("HGET", key, field)
}

// store DemoRedis to redis by one key
// field param is full match with original struct name
// DemoRedis will not expire when ttl is 0
func (this *DemoRedis) StoreToRedisByKey(ctx context.Context, key string, field string, val interface{}, ttl int) error {
	// redis conn
	conn := this.rp.Get()
	defer conn.Close()

	var err error

	// set expire
	if ttl > 0 {
		conn.Send("MULTI")
		conn.Send("HSET", key, field, val)
		conn.Send("EXPIRE", key, ttl)
		_, err = conn.Do("EXEC")
	} else {
		_, err = conn.Do("HSET", key, field, val)
	}

	return err
}
