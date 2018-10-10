// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: hash_type.proto

package test

import context "context"
import github_com_gomodule_redigo_redis "github.com/gomodule/redigo/redis"
import github_com_mitchellh_mapstructure "github.com/mitchellh/mapstructure"
import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/galaxyobe/protoc-gen-redis/proto"
import _ "github.com/gogo/protobuf/gogoproto"
import _ "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// new HashStorageType redis controller with redis pool
func (m *HashStorageType) RedisController(pool *github_com_gomodule_redigo_redis.Pool) *HashStorageTypeRedisController {
	return &HashStorageTypeRedisController{
		pool: pool,
		m:    m,
	}
}

// HashStorageType redis controller
type HashStorageTypeRedisController struct {
	pool *github_com_gomodule_redigo_redis.Pool
	m    *HashStorageType
}

// new HashStorageType redis controller with redis pool
func NewHashStorageTypeRedisController(pool *github_com_gomodule_redigo_redis.Pool) *HashStorageTypeRedisController {
	return &HashStorageTypeRedisController{pool: pool}
}

// get HashStorageType
func (r *HashStorageTypeRedisController) HashStorageType() *HashStorageType {
	return r.m
}

// load HashStorageType from redis hash with context and key
func (r *HashStorageTypeRedisController) Load(ctx context.Context, key string) error {
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
		default:
			structure[string(data[i])] = string(data[i+1])
		}
	}

	// use mapstructure weak decode structure to HashStorageType
	return github_com_mitchellh_mapstructure.WeakDecode(structure, r.m)
}

// store HashStorageType to redis hash with context and key
func (r *HashStorageTypeRedisController) Store(ctx context.Context, key string, ttl uint64) error {
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
	args = append(args, "SomeEnum", r.m.SomeEnum)

	// use redis hash store HashStorageType data with expire second
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

// get HashStorageType SomeString field value with key
func (r *HashStorageTypeRedisController) GetSomeString(key string) (SomeString string, err error) {
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

// set HashStorageType SomeString field with key and SomeString
func (r *HashStorageTypeRedisController) SetSomeString(key string, SomeString string) (err error) {
	// redis conn
	conn := r.pool.Get()
	defer conn.Close()

	// set SomeString field
	r.m.SomeString = SomeString
	_, err = conn.Do("HSET", key, "SomeString", SomeString)

	return
}

// get HashStorageType SomeBool field value with key
func (r *HashStorageTypeRedisController) GetSomeBool(key string) (SomeBool bool, err error) {
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

// set HashStorageType SomeBool field with key and SomeBool
func (r *HashStorageTypeRedisController) SetSomeBool(key string, SomeBool bool) (err error) {
	// redis conn
	conn := r.pool.Get()
	defer conn.Close()

	// set SomeBool field
	r.m.SomeBool = SomeBool
	_, err = conn.Do("HSET", key, "SomeBool", SomeBool)

	return
}

// get HashStorageType SomeInt32 field value with key
func (r *HashStorageTypeRedisController) GetSomeInt32(key string) (SomeInt32 int32, err error) {
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

// set HashStorageType SomeInt32 field with key and SomeInt32
func (r *HashStorageTypeRedisController) SetSomeInt32(key string, SomeInt32 int32) (err error) {
	// redis conn
	conn := r.pool.Get()
	defer conn.Close()

	// set SomeInt32 field
	r.m.SomeInt32 = SomeInt32
	_, err = conn.Do("HSET", key, "SomeInt32", SomeInt32)

	return
}

// get HashStorageType SomeUint32 field value with key
func (r *HashStorageTypeRedisController) GetSomeUint32(key string) (SomeUint32 uint32, err error) {
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

// set HashStorageType SomeUint32 field with key and SomeUint32
func (r *HashStorageTypeRedisController) SetSomeUint32(key string, SomeUint32 uint32) (err error) {
	// redis conn
	conn := r.pool.Get()
	defer conn.Close()

	// set SomeUint32 field
	r.m.SomeUint32 = SomeUint32
	_, err = conn.Do("HSET", key, "SomeUint32", SomeUint32)

	return
}

// get HashStorageType SomeInt64 field value with key
func (r *HashStorageTypeRedisController) GetSomeInt64(key string) (SomeInt64 int64, err error) {
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

// set HashStorageType SomeInt64 field with key and SomeInt64
func (r *HashStorageTypeRedisController) SetSomeInt64(key string, SomeInt64 int64) (err error) {
	// redis conn
	conn := r.pool.Get()
	defer conn.Close()

	// set SomeInt64 field
	r.m.SomeInt64 = SomeInt64
	_, err = conn.Do("HSET", key, "SomeInt64", SomeInt64)

	return
}

// get HashStorageType SomeUint64 field value with key
func (r *HashStorageTypeRedisController) GetSomeUint64(key string) (SomeUint64 uint64, err error) {
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

// set HashStorageType SomeUint64 field with key and SomeUint64
func (r *HashStorageTypeRedisController) SetSomeUint64(key string, SomeUint64 uint64) (err error) {
	// redis conn
	conn := r.pool.Get()
	defer conn.Close()

	// set SomeUint64 field
	r.m.SomeUint64 = SomeUint64
	_, err = conn.Do("HSET", key, "SomeUint64", SomeUint64)

	return
}

// get HashStorageType SomeFloat field value with key
func (r *HashStorageTypeRedisController) GetSomeFloat(key string) (SomeFloat float32, err error) {
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

// set HashStorageType SomeFloat field with key and SomeFloat
func (r *HashStorageTypeRedisController) SetSomeFloat(key string, SomeFloat float32) (err error) {
	// redis conn
	conn := r.pool.Get()
	defer conn.Close()

	// set SomeFloat field
	r.m.SomeFloat = SomeFloat
	_, err = conn.Do("HSET", key, "SomeFloat", SomeFloat)

	return
}

// get HashStorageType SomeEnum field value with key
func (r *HashStorageTypeRedisController) GetSomeEnum(key string) (SomeEnum HashStorageType_Enum, err error) {
	// redis conn
	conn := r.pool.Get()
	defer conn.Close()

	// get SomeEnum field
	if value, err := github_com_gomodule_redigo_redis.Int64(conn.Do("HGET", key, "SomeEnum")); err != nil {
		return SomeEnum, err
	} else {
		r.m.SomeEnum = HashStorageType_Enum(value)
	}

	return r.m.SomeEnum, nil
}

// set HashStorageType SomeEnum field with key and SomeEnum
func (r *HashStorageTypeRedisController) SetSomeEnum(key string, SomeEnum HashStorageType_Enum) (err error) {
	// redis conn
	conn := r.pool.Get()
	defer conn.Close()

	// set SomeEnum field
	r.m.SomeEnum = SomeEnum
	_, err = conn.Do("HSET", key, "SomeEnum", SomeEnum)

	return
}

// new HashStorageType2 redis controller with redis pool
func (m *HashStorageType2) RedisController(pool *github_com_gomodule_redigo_redis.Pool) *HashStorageType2RedisController {
	return &HashStorageType2RedisController{
		pool: pool,
		m:    m,
	}
}

// HashStorageType2 redis controller
type HashStorageType2RedisController struct {
	pool *github_com_gomodule_redigo_redis.Pool
	m    *HashStorageType2
}

// new HashStorageType2 redis controller with redis pool
func NewHashStorageType2RedisController(pool *github_com_gomodule_redigo_redis.Pool) *HashStorageType2RedisController {
	return &HashStorageType2RedisController{pool: pool}
}

// get HashStorageType2
func (r *HashStorageType2RedisController) HashStorageType2() *HashStorageType2 {
	return r.m
}

// load HashStorageType2 from redis hash with context and key
func (r *HashStorageType2RedisController) Load(ctx context.Context, key string) error {
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
		case "SomeMessage":
			// unmarshal SomeMessage
			r.m.SomeMessage = new(HashStorageType)
			if err := proto.Unmarshal(data[i+1], r.m.SomeMessage); err != nil {
				return err
			}
		case "Timestamp":
			// unmarshal Timestamp
			r.m.Timestamp = new(timestamp.Timestamp)
			if err := proto.Unmarshal(data[i+1], r.m.Timestamp); err != nil {
				return err
			}
		default:
			structure[string(data[i])] = string(data[i+1])
		}
	}

	// use mapstructure weak decode structure to HashStorageType2
	return github_com_mitchellh_mapstructure.WeakDecode(structure, r.m)
}

// store HashStorageType2 to redis hash with context and key
func (r *HashStorageType2RedisController) Store(ctx context.Context, key string, ttl uint64) error {
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
	// marshal SomeMessage
	if r.m.SomeMessage != nil {
		SomeMessage, SomeMessageError := proto.Marshal(r.m.SomeMessage)
		if SomeMessageError != nil {
			return SomeMessageError
		}
		args = append(args, "SomeMessage", SomeMessage)
	}
	// marshal Timestamp
	if r.m.Timestamp != nil {
		Timestamp, TimestampError := proto.Marshal(r.m.Timestamp)
		if TimestampError != nil {
			return TimestampError
		}
		args = append(args, "Timestamp", Timestamp)
	}

	// use redis hash store HashStorageType2 data with expire second
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

// get HashStorageType2 SomeString field value with key
func (r *HashStorageType2RedisController) GetSomeString(key string) (SomeString string, err error) {
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

// set HashStorageType2 SomeString field with key and SomeString
func (r *HashStorageType2RedisController) SetSomeString(key string, SomeString string) (err error) {
	// redis conn
	conn := r.pool.Get()
	defer conn.Close()

	// set SomeString field
	r.m.SomeString = SomeString
	_, err = conn.Do("HSET", key, "SomeString", SomeString)

	return
}

// get HashStorageType2 SomeBool field value with key
func (r *HashStorageType2RedisController) GetSomeBool(key string) (SomeBool bool, err error) {
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

// set HashStorageType2 SomeBool field with key and SomeBool
func (r *HashStorageType2RedisController) SetSomeBool(key string, SomeBool bool) (err error) {
	// redis conn
	conn := r.pool.Get()
	defer conn.Close()

	// set SomeBool field
	r.m.SomeBool = SomeBool
	_, err = conn.Do("HSET", key, "SomeBool", SomeBool)

	return
}

// get HashStorageType2 SomeInt32 field value with key
func (r *HashStorageType2RedisController) GetSomeInt32(key string) (SomeInt32 int32, err error) {
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

// set HashStorageType2 SomeInt32 field with key and SomeInt32
func (r *HashStorageType2RedisController) SetSomeInt32(key string, SomeInt32 int32) (err error) {
	// redis conn
	conn := r.pool.Get()
	defer conn.Close()

	// set SomeInt32 field
	r.m.SomeInt32 = SomeInt32
	_, err = conn.Do("HSET", key, "SomeInt32", SomeInt32)

	return
}

// get HashStorageType2 SomeUint32 field value with key
func (r *HashStorageType2RedisController) GetSomeUint32(key string) (SomeUint32 uint32, err error) {
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

// set HashStorageType2 SomeUint32 field with key and SomeUint32
func (r *HashStorageType2RedisController) SetSomeUint32(key string, SomeUint32 uint32) (err error) {
	// redis conn
	conn := r.pool.Get()
	defer conn.Close()

	// set SomeUint32 field
	r.m.SomeUint32 = SomeUint32
	_, err = conn.Do("HSET", key, "SomeUint32", SomeUint32)

	return
}

// get HashStorageType2 SomeInt64 field value with key
func (r *HashStorageType2RedisController) GetSomeInt64(key string) (SomeInt64 int64, err error) {
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

// set HashStorageType2 SomeInt64 field with key and SomeInt64
func (r *HashStorageType2RedisController) SetSomeInt64(key string, SomeInt64 int64) (err error) {
	// redis conn
	conn := r.pool.Get()
	defer conn.Close()

	// set SomeInt64 field
	r.m.SomeInt64 = SomeInt64
	_, err = conn.Do("HSET", key, "SomeInt64", SomeInt64)

	return
}

// get HashStorageType2 SomeUint64 field value with key
func (r *HashStorageType2RedisController) GetSomeUint64(key string) (SomeUint64 uint64, err error) {
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

// set HashStorageType2 SomeUint64 field with key and SomeUint64
func (r *HashStorageType2RedisController) SetSomeUint64(key string, SomeUint64 uint64) (err error) {
	// redis conn
	conn := r.pool.Get()
	defer conn.Close()

	// set SomeUint64 field
	r.m.SomeUint64 = SomeUint64
	_, err = conn.Do("HSET", key, "SomeUint64", SomeUint64)

	return
}

// get HashStorageType2 SomeFloat field value with key
func (r *HashStorageType2RedisController) GetSomeFloat(key string) (SomeFloat float32, err error) {
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

// set HashStorageType2 SomeFloat field with key and SomeFloat
func (r *HashStorageType2RedisController) SetSomeFloat(key string, SomeFloat float32) (err error) {
	// redis conn
	conn := r.pool.Get()
	defer conn.Close()

	// set SomeFloat field
	r.m.SomeFloat = SomeFloat
	_, err = conn.Do("HSET", key, "SomeFloat", SomeFloat)

	return
}

// get HashStorageType2 SomeMessage field value with key
func (r *HashStorageType2RedisController) GetSomeMessage(key string) (SomeMessage *HashStorageType, err error) {
	// redis conn
	conn := r.pool.Get()
	defer conn.Close()

	// get SomeMessage field
	if value, err := github_com_gomodule_redigo_redis.Bytes(conn.Do("HGET", key, "SomeMessage")); err != nil {
		return SomeMessage, err
	} else {
		// unmarshal SomeMessage
		r.m.SomeMessage = new(HashStorageType)
		if err = proto.Unmarshal(value, r.m.SomeMessage); err != nil {
			return SomeMessage, err
		}
	}

	return r.m.SomeMessage, nil
}

// set HashStorageType2 SomeMessage field with key and SomeMessage
func (r *HashStorageType2RedisController) SetSomeMessage(key string, SomeMessage *HashStorageType) error {
	// redis conn
	conn := r.pool.Get()
	defer conn.Close()

	// marshal SomeMessage
	if r.m.SomeMessage != nil {
		r.m.SomeMessage = SomeMessage
		if data, err := proto.Marshal(r.m.SomeMessage); err != nil {
			return err
		} else {
			// set SomeMessage field
			_, err = conn.Do("HSET", key, "SomeMessage", data)
			return err
		}
	}

	return nil
}

// get HashStorageType2 Timestamp field value with key
func (r *HashStorageType2RedisController) GetTimestamp(key string) (Timestamp *timestamp.Timestamp, err error) {
	// redis conn
	conn := r.pool.Get()
	defer conn.Close()

	// get Timestamp field
	if value, err := github_com_gomodule_redigo_redis.Bytes(conn.Do("HGET", key, "Timestamp")); err != nil {
		return Timestamp, err
	} else {
		// unmarshal Timestamp
		r.m.Timestamp = new(timestamp.Timestamp)
		if err = proto.Unmarshal(value, r.m.Timestamp); err != nil {
			return Timestamp, err
		}
	}

	return r.m.Timestamp, nil
}

// set HashStorageType2 Timestamp field with key and Timestamp
func (r *HashStorageType2RedisController) SetTimestamp(key string, Timestamp *timestamp.Timestamp) error {
	// redis conn
	conn := r.pool.Get()
	defer conn.Close()

	// marshal Timestamp
	if r.m.Timestamp != nil {
		r.m.Timestamp = Timestamp
		if data, err := proto.Marshal(r.m.Timestamp); err != nil {
			return err
		} else {
			// set Timestamp field
			_, err = conn.Do("HSET", key, "Timestamp", data)
			return err
		}
	}

	return nil
}