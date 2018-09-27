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