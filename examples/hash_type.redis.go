// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: hash_type.proto

package test

import (
	fmt "fmt"
	math "math"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/galaxyobe/protoc-gen-redis/proto"
	_ "github.com/gogo/protobuf/gogoproto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	github_com_gomodule_redigo_redis "github.com/gomodule/redigo/redis"
	github_com_mitchellh_mapstructure "github.com/mitchellh/mapstructure"
	github_com_json_iterator_go "github.com/json-iterator/go"
)

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
	return &HashStorageTypeRedisController{pool: pool, m: new(HashStorageType)}
}

// get HashStorageType
func (r *HashStorageTypeRedisController) HashStorageType() *HashStorageType {
	return r.m
}

// set HashStorageType
func (r *HashStorageTypeRedisController) SetHashStorageType(m *HashStorageType) {
	r.m = m
}

// load HashStorageType from redis hash
func (r *HashStorageTypeRedisController) Load(key string) error {
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

// get HashStorageType field from redis hash return string value
func (r *HashStorageTypeRedisController) GetString(key string, field string) (value string, err error) {
	// redis conn
	conn := r.pool.Get()
	defer conn.Close()

	// get field
	return github_com_gomodule_redigo_redis.String(conn.Do("HGET", key, field))
}

// get HashStorageType field from redis hash return bool value
func (r *HashStorageTypeRedisController) GetBool(key string, field string) (value bool, err error) {
	// redis conn
	conn := r.pool.Get()
	defer conn.Close()

	// get field
	return github_com_gomodule_redigo_redis.Bool(conn.Do("HGET", key, field))
}

// get HashStorageType field from redis hash return int64 value
func (r *HashStorageTypeRedisController) GetInt64(key string, field string) (value int64, err error) {
	// redis conn
	conn := r.pool.Get()
	defer conn.Close()

	// get field
	return github_com_gomodule_redigo_redis.Int64(conn.Do("HGET", key, field))
}

// get HashStorageType field from redis hash return uint64 value
func (r *HashStorageTypeRedisController) GetUint64(key string, field string) (value uint64, err error) {
	// redis conn
	conn := r.pool.Get()
	defer conn.Close()

	// get field
	return github_com_gomodule_redigo_redis.Uint64(conn.Do("HGET", key, field))
}

// get HashStorageType field from redis hash return float64 value
func (r *HashStorageTypeRedisController) GetFloat64(key string, field string) (value float64, err error) {
	// redis conn
	conn := r.pool.Get()
	defer conn.Close()

	// get field
	return github_com_gomodule_redigo_redis.Float64(conn.Do("HGET", key, field))
}

// get HashStorageType field from redis hash return interface
func (r *HashStorageTypeRedisController) GetInterface(key string, field string) (value interface{}, err error) {
	// redis conn
	conn := r.pool.Get()
	defer conn.Close()

	// get field
	return conn.Do("HGET", key, field)
}

// store HashStorageType to redis hash
func (r *HashStorageTypeRedisController) Store(key string) error {
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
	args = append(args, "SomeEnum", int32(r.m.SomeEnum))

	// use redis hash store HashStorageType data
	_, err := conn.Do("HMSET", args...)

	return err
}

// store HashStorageType to redis hash with key and ttl expire second
func (r *HashStorageTypeRedisController) StoreWithTTL(key string, ttl uint64) error {
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
	args = append(args, "SomeEnum", int32(r.m.SomeEnum))

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

// set HashStorageType field value to redis hash
func (r *HashStorageTypeRedisController) SetFieldValue(key string, field string, value interface{}) (err error) {
	// redis conn
	conn := r.pool.Get()
	defer conn.Close()

	// set field
	_, err = conn.Do("HSET", key, field, value)

	return
}

// get HashStorageType SomeString field value with key
func (r *HashStorageTypeRedisController) GetSomeString(key string) (someString string, err error) {
	// redis conn
	conn := r.pool.Get()
	defer conn.Close()

	// get SomeString field
	if value, err := github_com_gomodule_redigo_redis.String(conn.Do("HGET", key, "SomeString")); err != nil {
		return someString, err
	} else {
		r.m.SomeString = value
	}

	return r.m.SomeString, nil
}

// set HashStorageType SomeString field with key and SomeString
func (r *HashStorageTypeRedisController) SetSomeString(key string, someString string) (err error) {
	// redis conn
	conn := r.pool.Get()
	defer conn.Close()

	// set SomeString field
	r.m.SomeString = someString
	_, err = conn.Do("HSET", key, "SomeString", someString)

	return
}

// get HashStorageType SomeBool field value with key
func (r *HashStorageTypeRedisController) GetSomeBool(key string) (someBool bool, err error) {
	// redis conn
	conn := r.pool.Get()
	defer conn.Close()

	// get SomeBool field
	if value, err := github_com_gomodule_redigo_redis.Bool(conn.Do("HGET", key, "SomeBool")); err != nil {
		return someBool, err
	} else {
		r.m.SomeBool = value
	}

	return r.m.SomeBool, nil
}

// set HashStorageType SomeBool field with key and SomeBool
func (r *HashStorageTypeRedisController) SetSomeBool(key string, someBool bool) (err error) {
	// redis conn
	conn := r.pool.Get()
	defer conn.Close()

	// set SomeBool field
	r.m.SomeBool = someBool
	_, err = conn.Do("HSET", key, "SomeBool", someBool)

	return
}

// get HashStorageType SomeInt32 field value with key
func (r *HashStorageTypeRedisController) GetSomeInt32(key string) (someInt32 int32, err error) {
	// redis conn
	conn := r.pool.Get()
	defer conn.Close()

	// get SomeInt32 field
	if value, err := github_com_gomodule_redigo_redis.Int64(conn.Do("HGET", key, "SomeInt32")); err != nil {
		return someInt32, err
	} else {
		r.m.SomeInt32 = int32(value)
	}

	return r.m.SomeInt32, nil
}

// set HashStorageType SomeInt32 field with key and SomeInt32
func (r *HashStorageTypeRedisController) SetSomeInt32(key string, someInt32 int32) (err error) {
	// redis conn
	conn := r.pool.Get()
	defer conn.Close()

	// set SomeInt32 field
	r.m.SomeInt32 = someInt32
	_, err = conn.Do("HSET", key, "SomeInt32", someInt32)

	return
}

// get HashStorageType SomeUint32 field value with key
func (r *HashStorageTypeRedisController) GetSomeUint32(key string) (someUint32 uint32, err error) {
	// redis conn
	conn := r.pool.Get()
	defer conn.Close()

	// get SomeUint32 field
	if value, err := github_com_gomodule_redigo_redis.Uint64(conn.Do("HGET", key, "SomeUint32")); err != nil {
		return someUint32, err
	} else {
		r.m.SomeUint32 = uint32(value)
	}

	return r.m.SomeUint32, nil
}

// set HashStorageType SomeUint32 field with key and SomeUint32
func (r *HashStorageTypeRedisController) SetSomeUint32(key string, someUint32 uint32) (err error) {
	// redis conn
	conn := r.pool.Get()
	defer conn.Close()

	// set SomeUint32 field
	r.m.SomeUint32 = someUint32
	_, err = conn.Do("HSET", key, "SomeUint32", someUint32)

	return
}

// get HashStorageType SomeInt64 field value with key
func (r *HashStorageTypeRedisController) GetSomeInt64(key string) (someInt64 int64, err error) {
	// redis conn
	conn := r.pool.Get()
	defer conn.Close()

	// get SomeInt64 field
	if value, err := github_com_gomodule_redigo_redis.Int64(conn.Do("HGET", key, "SomeInt64")); err != nil {
		return someInt64, err
	} else {
		r.m.SomeInt64 = value
	}

	return r.m.SomeInt64, nil
}

// set HashStorageType SomeInt64 field with key and SomeInt64
func (r *HashStorageTypeRedisController) SetSomeInt64(key string, someInt64 int64) (err error) {
	// redis conn
	conn := r.pool.Get()
	defer conn.Close()

	// set SomeInt64 field
	r.m.SomeInt64 = someInt64
	_, err = conn.Do("HSET", key, "SomeInt64", someInt64)

	return
}

// get HashStorageType SomeUint64 field value with key
func (r *HashStorageTypeRedisController) GetSomeUint64(key string) (someUint64 uint64, err error) {
	// redis conn
	conn := r.pool.Get()
	defer conn.Close()

	// get SomeUint64 field
	if value, err := github_com_gomodule_redigo_redis.Uint64(conn.Do("HGET", key, "SomeUint64")); err != nil {
		return someUint64, err
	} else {
		r.m.SomeUint64 = value
	}

	return r.m.SomeUint64, nil
}

// set HashStorageType SomeUint64 field with key and SomeUint64
func (r *HashStorageTypeRedisController) SetSomeUint64(key string, someUint64 uint64) (err error) {
	// redis conn
	conn := r.pool.Get()
	defer conn.Close()

	// set SomeUint64 field
	r.m.SomeUint64 = someUint64
	_, err = conn.Do("HSET", key, "SomeUint64", someUint64)

	return
}

// get HashStorageType SomeFloat field value with key
func (r *HashStorageTypeRedisController) GetSomeFloat(key string) (someFloat float32, err error) {
	// redis conn
	conn := r.pool.Get()
	defer conn.Close()

	// get SomeFloat field
	if value, err := github_com_gomodule_redigo_redis.Float64(conn.Do("HGET", key, "SomeFloat")); err != nil {
		return someFloat, err
	} else {
		r.m.SomeFloat = float32(value)
	}

	return r.m.SomeFloat, nil
}

// set HashStorageType SomeFloat field with key and SomeFloat
func (r *HashStorageTypeRedisController) SetSomeFloat(key string, someFloat float32) (err error) {
	// redis conn
	conn := r.pool.Get()
	defer conn.Close()

	// set SomeFloat field
	r.m.SomeFloat = someFloat
	_, err = conn.Do("HSET", key, "SomeFloat", someFloat)

	return
}

// get HashStorageType SomeEnum field value with key
func (r *HashStorageTypeRedisController) GetSomeEnum(key string) (someEnum HashStorageType_Enum, err error) {
	// redis conn
	conn := r.pool.Get()
	defer conn.Close()

	// get SomeEnum field
	if value, err := github_com_gomodule_redigo_redis.Int64(conn.Do("HGET", key, "SomeEnum")); err != nil {
		return someEnum, err
	} else {
		r.m.SomeEnum = HashStorageType_Enum(value)
	}

	return r.m.SomeEnum, nil
}

// set HashStorageType SomeEnum field with key and SomeEnum
func (r *HashStorageTypeRedisController) SetSomeEnum(key string, someEnum HashStorageType_Enum) (err error) {
	// redis conn
	conn := r.pool.Get()
	defer conn.Close()

	// set SomeEnum field
	r.m.SomeEnum = someEnum
	_, err = conn.Do("HSET", key, "SomeEnum", int32(someEnum))

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
	return &HashStorageType2RedisController{pool: pool, m: new(HashStorageType2)}
}

// get HashStorageType2
func (r *HashStorageType2RedisController) HashStorageType2() *HashStorageType2 {
	return r.m
}

// set HashStorageType2
func (r *HashStorageType2RedisController) SetHashStorageType2(m *HashStorageType2) {
	r.m = m
}

// load HashStorageType2 from redis hash
func (r *HashStorageType2RedisController) Load(key string) error {
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
			if r.m.SomeMessage == nil {
				r.m.SomeMessage = new(HashStorageType)
			}
			if err := proto.Unmarshal(data[i+1], r.m.SomeMessage); err != nil {
				return err
			}
		case "Timestamp":
			// unmarshal Timestamp
			if r.m.Timestamp == nil {
				r.m.Timestamp = new(timestamp.Timestamp)
			}
			if err := proto.Unmarshal(data[i+1], r.m.Timestamp); err != nil {
				return err
			}
		case "SomeMessages":
			// unmarshal SomeMessages
			if err := github_com_json_iterator_go.Unmarshal(data[i+1], &r.m.SomeMessages); err != nil {
				return err
			}
		default:
			structure[string(data[i])] = string(data[i+1])
		}
	}

	// use mapstructure weak decode structure to HashStorageType2
	return github_com_mitchellh_mapstructure.WeakDecode(structure, r.m)
}

// get HashStorageType2 field from redis hash return string value
func (r *HashStorageType2RedisController) GetString(key string, field string) (value string, err error) {
	// redis conn
	conn := r.pool.Get()
	defer conn.Close()

	// get field
	return github_com_gomodule_redigo_redis.String(conn.Do("HGET", key, field))
}

// get HashStorageType2 field from redis hash return bool value
func (r *HashStorageType2RedisController) GetBool(key string, field string) (value bool, err error) {
	// redis conn
	conn := r.pool.Get()
	defer conn.Close()

	// get field
	return github_com_gomodule_redigo_redis.Bool(conn.Do("HGET", key, field))
}

// get HashStorageType2 field from redis hash return int64 value
func (r *HashStorageType2RedisController) GetInt64(key string, field string) (value int64, err error) {
	// redis conn
	conn := r.pool.Get()
	defer conn.Close()

	// get field
	return github_com_gomodule_redigo_redis.Int64(conn.Do("HGET", key, field))
}

// get HashStorageType2 field from redis hash return uint64 value
func (r *HashStorageType2RedisController) GetUint64(key string, field string) (value uint64, err error) {
	// redis conn
	conn := r.pool.Get()
	defer conn.Close()

	// get field
	return github_com_gomodule_redigo_redis.Uint64(conn.Do("HGET", key, field))
}

// get HashStorageType2 field from redis hash return float64 value
func (r *HashStorageType2RedisController) GetFloat64(key string, field string) (value float64, err error) {
	// redis conn
	conn := r.pool.Get()
	defer conn.Close()

	// get field
	return github_com_gomodule_redigo_redis.Float64(conn.Do("HGET", key, field))
}

// get HashStorageType2 field from redis hash return interface
func (r *HashStorageType2RedisController) GetInterface(key string, field string) (value interface{}, err error) {
	// redis conn
	conn := r.pool.Get()
	defer conn.Close()

	// get field
	return conn.Do("HGET", key, field)
}

// store HashStorageType2 to redis hash
func (r *HashStorageType2RedisController) Store(key string) error {
	// redis conn
	conn := r.pool.Get()
	defer conn.Close()

	// make args
	args := make([]interface{}, 0)

	// add redis key
	args = append(args, key)

	// add redis field and value
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
	// marshal SomeMessages
	if r.m.SomeMessages != nil {
		SomeMessages, SomeMessagesError := github_com_json_iterator_go.Marshal(r.m.SomeMessages)
		if SomeMessagesError != nil {
			return SomeMessagesError
		}
		args = append(args, "SomeMessages", SomeMessages)
	}

	// use redis hash store HashStorageType2 data
	_, err := conn.Do("HMSET", args...)

	return err
}

// store HashStorageType2 to redis hash with key and ttl expire second
func (r *HashStorageType2RedisController) StoreWithTTL(key string, ttl uint64) error {
	// redis conn
	conn := r.pool.Get()
	defer conn.Close()

	// make args
	args := make([]interface{}, 0)

	// add redis key
	args = append(args, key)

	// add redis field and value
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
	// marshal SomeMessages
	if r.m.SomeMessages != nil {
		SomeMessages, SomeMessagesError := github_com_json_iterator_go.Marshal(r.m.SomeMessages)
		if SomeMessagesError != nil {
			return SomeMessagesError
		}
		args = append(args, "SomeMessages", SomeMessages)
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

// set HashStorageType2 field value to redis hash
func (r *HashStorageType2RedisController) SetFieldValue(key string, field string, value interface{}) (err error) {
	// redis conn
	conn := r.pool.Get()
	defer conn.Close()

	// set field
	_, err = conn.Do("HSET", key, field, value)

	return
}

// get HashStorageType2 SomeMessage field value with key
func (r *HashStorageType2RedisController) GetSomeMessage(key string) (ret *HashStorageType, err error) {
	// redis conn
	conn := r.pool.Get()
	defer conn.Close()

	// get SomeMessage field
	if value, err := github_com_gomodule_redigo_redis.Bytes(conn.Do("HGET", key, "SomeMessage")); err != nil {
		return ret, err
	} else {
		// unmarshal SomeMessage
		if r.m.SomeMessage == nil {
			r.m.SomeMessage = new(HashStorageType)
		}
		if err = proto.Unmarshal(value, r.m.SomeMessage); err != nil {
			return ret, err
		}
	}

	return r.m.SomeMessage, nil
}

// set HashStorageType2 SomeMessage field with key and SomeMessage
func (r *HashStorageType2RedisController) SetSomeMessage(key string, someMessage *HashStorageType) error {
	// redis conn
	conn := r.pool.Get()
	defer conn.Close()

	// marshal SomeMessage
	r.m.SomeMessage = someMessage
	if data, err := proto.Marshal(r.m.SomeMessage); err != nil {
		return err
	} else {
		// set SomeMessage field
		_, err = conn.Do("HSET", key, "SomeMessage", data)
		return err
	}

	return nil
}

// get HashStorageType2 Timestamp field value with key
func (r *HashStorageType2RedisController) GetTimestamp(key string) (ret *timestamp.Timestamp, err error) {
	// redis conn
	conn := r.pool.Get()
	defer conn.Close()

	// get Timestamp field
	if value, err := github_com_gomodule_redigo_redis.Bytes(conn.Do("HGET", key, "Timestamp")); err != nil {
		return ret, err
	} else {
		// unmarshal Timestamp
		if r.m.Timestamp == nil {
			r.m.Timestamp = new(timestamp.Timestamp)
		}
		if err = proto.Unmarshal(value, r.m.Timestamp); err != nil {
			return ret, err
		}
	}

	return r.m.Timestamp, nil
}

// set HashStorageType2 Timestamp field with key and Timestamp
func (r *HashStorageType2RedisController) SetTimestamp(key string, timestamp *timestamp.Timestamp) error {
	// redis conn
	conn := r.pool.Get()
	defer conn.Close()

	// marshal Timestamp
	r.m.Timestamp = timestamp
	if data, err := proto.Marshal(r.m.Timestamp); err != nil {
		return err
	} else {
		// set Timestamp field
		_, err = conn.Do("HSET", key, "Timestamp", data)
		return err
	}

	return nil
}

// get HashStorageType2 SomeMessages field value with key
func (r *HashStorageType2RedisController) GetSomeMessages(key string) (ret []*HashStorageType, err error) {
	// redis conn
	conn := r.pool.Get()
	defer conn.Close()

	// get SomeMessages field
	if value, err := github_com_gomodule_redigo_redis.Bytes(conn.Do("HGET", key, "SomeMessages")); err != nil {
		return ret, err
	} else {
		// unmarshal SomeMessages
		if err = github_com_json_iterator_go.Unmarshal(value, &r.m.SomeMessages); err != nil {
			return ret, err
		}
	}

	return r.m.SomeMessages, nil
}

// set HashStorageType2 SomeMessages field with key and SomeMessages
func (r *HashStorageType2RedisController) SetSomeMessages(key string, someMessages []*HashStorageType) error {
	// redis conn
	conn := r.pool.Get()
	defer conn.Close()

	// marshal SomeMessages
	r.m.SomeMessages = someMessages
	if data, err := github_com_json_iterator_go.Marshal(r.m.SomeMessages); err != nil {
		return err
	} else {
		// set SomeMessages field
		_, err = conn.Do("HSET", key, "SomeMessages", data)
		return err
	}

	return nil
}
