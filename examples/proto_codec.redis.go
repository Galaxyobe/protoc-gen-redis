// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto_codec.proto

package test

import context "context"
import github_com_gomodule_redigo_redis "github.com/gomodule/redigo/redis"
import github_com_mitchellh_mapstructure "github.com/mitchellh/mapstructure"
import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/galaxyobe/protoc-gen-redis/proto"
import _ "github.com/gogo/protobuf/gogoproto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// new StringProtoCodec redis controller with redis pool
func (m *StringProtoCodec) RedisController(pool *github_com_gomodule_redigo_redis.Pool) *StringProtoCodecRedisController {
	return &StringProtoCodecRedisController{
		pool: pool,
		m:    m,
	}
}

// StringProtoCodec redis controller
type StringProtoCodecRedisController struct {
	pool *github_com_gomodule_redigo_redis.Pool
	m    *StringProtoCodec
}

// new StringProtoCodec redis controller with redis pool
func NewStringProtoCodecRedisController(pool *github_com_gomodule_redigo_redis.Pool) *StringProtoCodecRedisController {
	return &StringProtoCodecRedisController{pool: pool}
}

// get StringProtoCodec
func (r *StringProtoCodecRedisController) StringProtoCodec() *StringProtoCodec {
	return r.m
}

// store StringProtoCodec to redis string with context and key
func (r *StringProtoCodecRedisController) Store(ctx context.Context, key string, ttl uint64) error {
	// redis conn
	conn := r.pool.Get()
	defer conn.Close()

	// marshal StringProtoCodec to []byte
	data, err := proto.Marshal(r.m)
	if err != nil {
		return err
	}

	// use redis string store StringProtoCodec data with expire second
	_, err = conn.Do("SETEX", key, ttl, data)

	return err
}

// load StringProtoCodec from redis string with context and key
func (r *StringProtoCodecRedisController) Load(ctx context.Context, key string) error {
	// redis conn
	conn := r.pool.Get()
	defer conn.Close()

	// load data from redis string
	data, err := github_com_gomodule_redigo_redis.Bytes(conn.Do("GET", key))
	if err != nil {
		return err
	}

	// unmarshal data to StringStorageType
	return proto.Unmarshal(data, r.m)
}

// new HashProtoCodec redis controller with redis pool
func (m *HashProtoCodec) RedisController(pool *github_com_gomodule_redigo_redis.Pool) *HashProtoCodecRedisController {
	return &HashProtoCodecRedisController{
		pool: pool,
		m:    m,
	}
}

// HashProtoCodec redis controller
type HashProtoCodecRedisController struct {
	pool *github_com_gomodule_redigo_redis.Pool
	m    *HashProtoCodec
}

// new HashProtoCodec redis controller with redis pool
func NewHashProtoCodecRedisController(pool *github_com_gomodule_redigo_redis.Pool) *HashProtoCodecRedisController {
	return &HashProtoCodecRedisController{pool: pool}
}

// get HashProtoCodec
func (r *HashProtoCodecRedisController) HashProtoCodec() *HashProtoCodec {
	return r.m
}

// load HashProtoCodec from redis hash with context and key
func (r *HashProtoCodecRedisController) Load(ctx context.Context, key string) error {
	// redis conn
	conn := r.pool.Get()
	defer conn.Close()

	// load data from redis hash
	data, err := github_com_gomodule_redigo_redis.ByteSlices(conn.Do("HGETALL", key))
	if err != nil {
		return err
	}

	// parse redis hash field name and value
	structure := make(map[string]interface{})
	for i := 0; i < len(data); i += 2 {
		switch string(data[i]) {
		case "HashProtoCodec":
			// unmarshal HashProtoCodec
			r.m.HashProtoCodec = new(HashProtoCodec)
			if err := proto.Unmarshal(data[i+1], r.m.HashProtoCodec); err != nil {
				return err
			}
		default:
			structure[string(data[i])] = string(data[i+1])
		}
	}

	// use mapstructure weak decode structure to HashProtoCodec
	return github_com_mitchellh_mapstructure.WeakDecode(structure, r.m)
}

// store HashProtoCodec to redis hash with context and key
func (r *HashProtoCodecRedisController) Store(ctx context.Context, key string, ttl uint64) error {
	// redis conn
	conn := r.pool.Get()
	defer conn.Close()

	// make args
	args := make([]interface{}, 0)

	// add redis key
	args = append(args, key)

	// add redis field and value
	args = append(args, "SomeString", r.m.SomeString)
	args = append(args, "SomeBool", r.m.SomeBool)
	args = append(args, "SomeInt32", r.m.SomeInt32)
	args = append(args, "SomeUint32", r.m.SomeUint32)
	args = append(args, "SomeInt64", r.m.SomeInt64)
	args = append(args, "SomeUint64", r.m.SomeUint64)
	args = append(args, "SomeFloat", r.m.SomeFloat)
	// marshal HashProtoCodec
	if r.m.HashProtoCodec != nil {
		HashProtoCodec, HashProtoCodecError := proto.Marshal(r.m.HashProtoCodec)
		if HashProtoCodecError != nil {
			return HashProtoCodecError
		}
		args = append(args, "HashProtoCodec", HashProtoCodec)
	}

	// use redis hash store HashProtoCodec data with expire second
	err := conn.Send("MULTI")
	if err != nil {
		return err
	}
	err = conn.Send("HMSET", args...)
	if err != nil {
		return err
	}
	err = conn.Send("EXPIRE", key, ttl)
	if err != nil {
		return err
	}
	_, err = conn.Do("EXEC")

	return err
}

// get HashProtoCodec SomeString field value with key
func (r *HashProtoCodecRedisController) GetSomeString(key string) (SomeString string, err error) {
	// redis conn
	conn := r.pool.Get()
	defer conn.Close()

	// get SomeString field
	if value, err := github_com_gomodule_redigo_redis.String(conn.Do("HGET", key, "SomeString")); err != nil {
		return SomeString, err
	} else {
		r.m.SomeString = value
	}

	return r.m.SomeString, nil
}

// set HashProtoCodec SomeString field with key and SomeString
func (r *HashProtoCodecRedisController) SetSomeString(key string, SomeString string) (err error) {
	// redis conn
	conn := r.pool.Get()
	defer conn.Close()

	// set SomeString field
	r.m.SomeString = SomeString
	_, err = conn.Do("HSET", key, "SomeString", SomeString)

	return
}

// get HashProtoCodec SomeBool field value with key
func (r *HashProtoCodecRedisController) GetSomeBool(key string) (SomeBool bool, err error) {
	// redis conn
	conn := r.pool.Get()
	defer conn.Close()

	// get SomeBool field
	if value, err := github_com_gomodule_redigo_redis.Bool(conn.Do("HGET", key, "SomeBool")); err != nil {
		return SomeBool, err
	} else {
		r.m.SomeBool = value
	}

	return r.m.SomeBool, nil
}

// set HashProtoCodec SomeBool field with key and SomeBool
func (r *HashProtoCodecRedisController) SetSomeBool(key string, SomeBool bool) (err error) {
	// redis conn
	conn := r.pool.Get()
	defer conn.Close()

	// set SomeBool field
	r.m.SomeBool = SomeBool
	_, err = conn.Do("HSET", key, "SomeBool", SomeBool)

	return
}

// get HashProtoCodec SomeInt32 field value with key
func (r *HashProtoCodecRedisController) GetSomeInt32(key string) (SomeInt32 int32, err error) {
	// redis conn
	conn := r.pool.Get()
	defer conn.Close()

	// get SomeInt32 field
	if value, err := github_com_gomodule_redigo_redis.Int64(conn.Do("HGET", key, "SomeInt32")); err != nil {
		return SomeInt32, err
	} else {
		r.m.SomeInt32 = int32(value)
	}

	return r.m.SomeInt32, nil
}

// set HashProtoCodec SomeInt32 field with key and SomeInt32
func (r *HashProtoCodecRedisController) SetSomeInt32(key string, SomeInt32 int32) (err error) {
	// redis conn
	conn := r.pool.Get()
	defer conn.Close()

	// set SomeInt32 field
	r.m.SomeInt32 = SomeInt32
	_, err = conn.Do("HSET", key, "SomeInt32", SomeInt32)

	return
}

// get HashProtoCodec SomeUint32 field value with key
func (r *HashProtoCodecRedisController) GetSomeUint32(key string) (SomeUint32 uint32, err error) {
	// redis conn
	conn := r.pool.Get()
	defer conn.Close()

	// get SomeUint32 field
	if value, err := github_com_gomodule_redigo_redis.Uint64(conn.Do("HGET", key, "SomeUint32")); err != nil {
		return SomeUint32, err
	} else {
		r.m.SomeUint32 = uint32(value)
	}

	return r.m.SomeUint32, nil
}

// set HashProtoCodec SomeUint32 field with key and SomeUint32
func (r *HashProtoCodecRedisController) SetSomeUint32(key string, SomeUint32 uint32) (err error) {
	// redis conn
	conn := r.pool.Get()
	defer conn.Close()

	// set SomeUint32 field
	r.m.SomeUint32 = SomeUint32
	_, err = conn.Do("HSET", key, "SomeUint32", SomeUint32)

	return
}

// get HashProtoCodec SomeInt64 field value with key
func (r *HashProtoCodecRedisController) GetSomeInt64(key string) (SomeInt64 int64, err error) {
	// redis conn
	conn := r.pool.Get()
	defer conn.Close()

	// get SomeInt64 field
	if value, err := github_com_gomodule_redigo_redis.Int64(conn.Do("HGET", key, "SomeInt64")); err != nil {
		return SomeInt64, err
	} else {
		r.m.SomeInt64 = value
	}

	return r.m.SomeInt64, nil
}

// set HashProtoCodec SomeInt64 field with key and SomeInt64
func (r *HashProtoCodecRedisController) SetSomeInt64(key string, SomeInt64 int64) (err error) {
	// redis conn
	conn := r.pool.Get()
	defer conn.Close()

	// set SomeInt64 field
	r.m.SomeInt64 = SomeInt64
	_, err = conn.Do("HSET", key, "SomeInt64", SomeInt64)

	return
}

// get HashProtoCodec SomeUint64 field value with key
func (r *HashProtoCodecRedisController) GetSomeUint64(key string) (SomeUint64 uint64, err error) {
	// redis conn
	conn := r.pool.Get()
	defer conn.Close()

	// get SomeUint64 field
	if value, err := github_com_gomodule_redigo_redis.Uint64(conn.Do("HGET", key, "SomeUint64")); err != nil {
		return SomeUint64, err
	} else {
		r.m.SomeUint64 = value
	}

	return r.m.SomeUint64, nil
}

// set HashProtoCodec SomeUint64 field with key and SomeUint64
func (r *HashProtoCodecRedisController) SetSomeUint64(key string, SomeUint64 uint64) (err error) {
	// redis conn
	conn := r.pool.Get()
	defer conn.Close()

	// set SomeUint64 field
	r.m.SomeUint64 = SomeUint64
	_, err = conn.Do("HSET", key, "SomeUint64", SomeUint64)

	return
}

// get HashProtoCodec SomeFloat field value with key
func (r *HashProtoCodecRedisController) GetSomeFloat(key string) (SomeFloat float32, err error) {
	// redis conn
	conn := r.pool.Get()
	defer conn.Close()

	// get SomeFloat field
	if value, err := github_com_gomodule_redigo_redis.Float64(conn.Do("HGET", key, "SomeFloat")); err != nil {
		return SomeFloat, err
	} else {
		r.m.SomeFloat = float32(value)
	}

	return r.m.SomeFloat, nil
}

// set HashProtoCodec SomeFloat field with key and SomeFloat
func (r *HashProtoCodecRedisController) SetSomeFloat(key string, SomeFloat float32) (err error) {
	// redis conn
	conn := r.pool.Get()
	defer conn.Close()

	// set SomeFloat field
	r.m.SomeFloat = SomeFloat
	_, err = conn.Do("HSET", key, "SomeFloat", SomeFloat)

	return
}

// get HashProtoCodec HashProtoCodec field value with key
func (r *HashProtoCodecRedisController) GetHashProtoCodec(key string) (HashProtoCodec *HashProtoCodec, err error) {
	// redis conn
	conn := r.pool.Get()
	defer conn.Close()

	// get HashProtoCodec field
	if value, err := github_com_gomodule_redigo_redis.Bytes(conn.Do("HGET", key, "HashProtoCodec")); err != nil {
		return HashProtoCodec, err
	} else {
		// unmarshal HashProtoCodec
		r.m.HashProtoCodec = new(HashProtoCodec)
		if err = proto.Unmarshal(value, r.m.HashProtoCodec); err != nil {
			return HashProtoCodec, err
		}
	}

	return r.m.HashProtoCodec, nil
}

// set HashProtoCodec HashProtoCodec field with key and HashProtoCodec
func (r *HashProtoCodecRedisController) SetHashProtoCodec(key string, HashProtoCodec *HashProtoCodec) error {
	// redis conn
	conn := r.pool.Get()
	defer conn.Close()

	// marshal HashProtoCodec
	if r.m.HashProtoCodec != nil {
		r.m.HashProtoCodec = HashProtoCodec
		if data, err := proto.Marshal(r.m.HashProtoCodec); err != nil {
			return err
		} else {
			// set HashProtoCodec field
			_, err = conn.Do("HSET", key, "HashProtoCodec", data)
			return err
		}
	}

	return nil
}
