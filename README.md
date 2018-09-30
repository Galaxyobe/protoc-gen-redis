# protoc-gen-redis

Generate redis load and store function for protobuffer message.
use redis string save the message proto data or hash
based on github.com/gomodule/redigo/redis

## using the following extensions:

* message options

    - enabled

        enable generate redis load and store function

        default: false

    - ttl

        enable generate  store function with expire ttl

        default: true

    - storage_type

        store to redis type ,only support string and hash

        default: string

## Installing and using

> go get -u github.com/galaxyobe/protoc-gen-redis

> protoc -I=$GOPATH/src -I=. --go_out=. --redis_out=. *.proto


## example

See [test.proto](https://github.com/Galaxyobe/protoc-gen-redis/blob/master/test/test.proto)
```protobuffer
syntax = "proto3";
package test;

import "protoc-gen-redis/proto/redis.proto";

message EnabledWithTTL {
    option (redis.enabled) = true;
    option (redis.ttl) = true;
    int32 some_integer = 1;
}

message EnabledWithTTL2 {
    option (redis.enabled) = true;
    int32 some_integer = 1;
}

message EnabledWithoutTTL {
    option (redis.enabled) = true;
    option (redis.ttl) = false;
    int32 some_integer = 1;
}

message Disabled {
    option (redis.enabled) = false;
    option (redis.ttl) = false;
    int32 some_integer = 1;
}

message Disabled2 {
    option (redis.enabled) = false;
    option (redis.ttl) = true;
    int32 some_integer = 1;
}

message Disabled3 {
    option (redis.ttl) = true;
    int32 some_integer = 1;
}

message Disabled4 {
    int32 some_integer = 1;
}
```
See [test.redis.go](https://github.com/Galaxyobe/protoc-gen-redis/blob/master/test/test.redis.go)
```golang
// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: test.proto

package test

import context "context"
import github_com_gomodule_redigo_redis "github.com/gomodule/redigo/redis"
import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "protoc-gen-redis/proto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// load EnabledWithTTL from redis
func (m *EnabledWithTTL) LoadFromRedis(ctx context.Context, conn github_com_gomodule_redigo_redis.Conn, key string) error {
	// load data from redis string
	data, err := github_com_gomodule_redigo_redis.Bytes(conn.Do("GET", key))
	if err != nil {
		return err
	}
	// unmarshal data to EnabledWithTTL
	err = proto.Unmarshal(data, m)
	if err != nil {
		return err
	}

	return nil
}

// store EnabledWithTTL to redis
// EnabledWithTTL will not expire when ttl is 0
func (m *EnabledWithTTL) StoreToRedis(ctx context.Context, conn github_com_gomodule_redigo_redis.Conn, key string, ttl uint64) error {
	// marshal EnabledWithTTL to []byte
	data, err := proto.Marshal(m)
	if err != nil {
		return err
	}

	// use redis string store the EnabledWithTTL data with expire second
	_, err = conn.Do("SETEX", key, ttl, data)

	if err != nil {
		return err
	}

	return nil
}

// load EnabledWithTTL2 from redis
func (m *EnabledWithTTL2) LoadFromRedis(ctx context.Context, conn github_com_gomodule_redigo_redis.Conn, key string) error {
	// load data from redis string
	data, err := github_com_gomodule_redigo_redis.Bytes(conn.Do("GET", key))
	if err != nil {
		return err
	}
	// unmarshal data to EnabledWithTTL2
	err = proto.Unmarshal(data, m)
	if err != nil {
		return err
	}

	return nil
}

// store EnabledWithTTL2 to redis
// EnabledWithTTL2 will not expire when ttl is 0
func (m *EnabledWithTTL2) StoreToRedis(ctx context.Context, conn github_com_gomodule_redigo_redis.Conn, key string, ttl uint64) error {
	// marshal EnabledWithTTL2 to []byte
	data, err := proto.Marshal(m)
	if err != nil {
		return err
	}

	// use redis string store the EnabledWithTTL2 data with expire second
	_, err = conn.Do("SETEX", key, ttl, data)

	if err != nil {
		return err
	}

	return nil
}

// load EnabledWithoutTTL from redis
func (m *EnabledWithoutTTL) LoadFromRedis(ctx context.Context, conn github_com_gomodule_redigo_redis.Conn, key string) error {
	// load data from redis string
	data, err := github_com_gomodule_redigo_redis.Bytes(conn.Do("GET", key))
	if err != nil {
		return err
	}
	// unmarshal data to EnabledWithoutTTL
	err = proto.Unmarshal(data, m)
	if err != nil {
		return err
	}

	return nil
}

// store EnabledWithoutTTL to redis
// EnabledWithoutTTL will not expire when ttl is 0
func (m *EnabledWithoutTTL) StoreToRedis(ctx context.Context, conn github_com_gomodule_redigo_redis.Conn, key string) error {
	// marshal EnabledWithoutTTL to []byte
	data, err := proto.Marshal(m)
	if err != nil {
		return err
	}

	// use redis string store the EnabledWithoutTTL data
	_, err = conn.Do("SET", key, data)

	if err != nil {
		return err
	}

	return nil
}
```

// demo
> protoc -I=. -I=$GOPATH/src -I=proto -I=examples --go_out=. --plugin=protoc-gen-redis=./protoc-gen-redis examples/*.proto