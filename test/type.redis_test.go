package test

import (
	"testing"
	"github.com/alicebob/miniredis"
	"github.com/gomodule/redigo/redis"
	"time"
	"context"
	"github.com/golang/protobuf/ptypes/timestamp"
	"bytes"
)

func TestStringStorageTypeRedisController_Store(t *testing.T) {
	s, err := miniredis.Run()
	if err != nil {
		t.Error(err)
	}
	defer s.Close()

	pool := &redis.Pool{
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", s.Addr())
			if err != nil {
				t.Error(err)
				return nil, err
			}
			return c, nil
		},
	}

	stringStorageType := &StringStorageType{}
	stringStorageType.SomeString = "SomeString"
	stringStorageType.SomeBool = true
	stringStorageType.SomeInt32 = -125
	stringStorageType.SomeUint32 = 255
	stringStorageType.SomeInt64 = -255
	stringStorageType.SomeUint64 = 255
	stringStorageType.SomeFloat = -1.25
	stringStorageType.SomeEnum = StringStorageType_E1
	stringStorageType.Timestamp = &timestamp.Timestamp{Seconds: 1, Nanos: 2}
	stringStorageType.Timestamps = append(stringStorageType.Timestamps, stringStorageType.Timestamp, &timestamp.Timestamp{Seconds: 20, Nanos: 52})

	ctl := NewStringStorageTypeRedisController(pool)

	ctl.StringStorageType().SomeString = "SomeString"
	ctl.StringStorageType().SomeBool = true
	ctl.StringStorageType().SomeInt32 = -125
	ctl.StringStorageType().SomeUint32 = 255
	ctl.StringStorageType().SomeInt64 = -255
	ctl.StringStorageType().SomeUint64 = 255
	ctl.StringStorageType().SomeFloat = -1.25
	ctl.StringStorageType().SomeEnum = StringStorageType_E2
	ctl.StringStorageType().Timestamp = &timestamp.Timestamp{Seconds: 3, Nanos: 4}
	ctl.StringStorageType().StringStorageType = stringStorageType
	ctl.StringStorageType().SomeMessages = append(ctl.StringStorageType().SomeMessages, stringStorageType)
	ctl.StringStorageType().Timestamps = append(ctl.StringStorageType().Timestamps, stringStorageType.Timestamp, &timestamp.Timestamp{Seconds: 20, Nanos: 52})

	key := "TestStringStorageTypeRedisController_Store"

	if err = ctl.Store(context.Background(), key); err != nil {
		t.Error(err)
	}

	s.CheckGet(t, key, "\n\nSomeString\x10\x01\x18\x83\xff\xff\xff\xff\xff\xff\xff\xff\x01 \xff\x01(\x81\xfe\xff\xff\xff\xff\xff\xff\xff\x010\xff\x01=\x00\x00\xa0\xbf@\x01J\x04\b\x03\x10\x04RA\n\nSomeString\x10\x01\x18\x83\xff\xff\xff\xff\xff\xff\xff\xff\x01 \xff\x01(\x81\xfe\xff\xff\xff\xff\xff\xff\xff\x010\xff\x01=\x00\x00\xa0\xbfJ\x04\b\x01\x10\x02b\x04\b\x01\x10\x02b\x04\b\x14\x104ZA\n\nSomeString\x10\x01\x18\x83\xff\xff\xff\xff\xff\xff\xff\xff\x01 \xff\x01(\x81\xfe\xff\xff\xff\xff\xff\xff\xff\x010\xff\x01=\x00\x00\xa0\xbfJ\x04\b\x01\x10\x02b\x04\b\x01\x10\x02b\x04\b\x14\x104b\x04\b\x01\x10\x02b\x04\b\x14\x104")
}

func TestStringStorageTypeRedisController_StoreWithTTL(t *testing.T) {
	s, err := miniredis.Run()
	if err != nil {
		t.Error(err)
	}
	defer s.Close()

	pool := &redis.Pool{
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", s.Addr())
			if err != nil {
				t.Error(err)
				return nil, err
			}
			return c, nil
		},
	}

	stringStorageType := &StringStorageType{}
	stringStorageType.SomeString = "SomeString"
	stringStorageType.SomeBool = true
	stringStorageType.SomeInt32 = -125
	stringStorageType.SomeUint32 = 255
	stringStorageType.SomeInt64 = -255
	stringStorageType.SomeUint64 = 255
	stringStorageType.SomeFloat = -1.25
	stringStorageType.SomeEnum = StringStorageType_E1
	stringStorageType.Timestamp = &timestamp.Timestamp{Seconds: 1, Nanos: 2}
	stringStorageType.Timestamps = append(stringStorageType.Timestamps, stringStorageType.Timestamp, &timestamp.Timestamp{Seconds: 20, Nanos: 52})

	ctl := NewStringStorageTypeRedisController(pool)

	ctl.StringStorageType().SomeString = "SomeString"
	ctl.StringStorageType().SomeBool = true
	ctl.StringStorageType().SomeInt32 = -125
	ctl.StringStorageType().SomeUint32 = 255
	ctl.StringStorageType().SomeInt64 = -255
	ctl.StringStorageType().SomeUint64 = 255
	ctl.StringStorageType().SomeFloat = -1.25
	ctl.StringStorageType().SomeEnum = StringStorageType_E2
	ctl.StringStorageType().Timestamp = &timestamp.Timestamp{Seconds: 3, Nanos: 4}
	ctl.StringStorageType().StringStorageType = stringStorageType
	ctl.StringStorageType().SomeMessages = append(ctl.StringStorageType().SomeMessages, stringStorageType)
	ctl.StringStorageType().Timestamps = append(ctl.StringStorageType().Timestamps, stringStorageType.Timestamp, &timestamp.Timestamp{Seconds: 20, Nanos: 52})

	key := "TestStringStorageTypeRedisController_StoreWithTTL"

	if err = ctl.StoreWithTTL(context.Background(), key, 30); err != nil {
		t.Error(err)
	}

	s.CheckGet(t, key, "\n\nSomeString\x10\x01\x18\x83\xff\xff\xff\xff\xff\xff\xff\xff\x01 \xff\x01(\x81\xfe\xff\xff\xff\xff\xff\xff\xff\x010\xff\x01=\x00\x00\xa0\xbf@\x01J\x04\b\x03\x10\x04RA\n\nSomeString\x10\x01\x18\x83\xff\xff\xff\xff\xff\xff\xff\xff\x01 \xff\x01(\x81\xfe\xff\xff\xff\xff\xff\xff\xff\x010\xff\x01=\x00\x00\xa0\xbfJ\x04\b\x01\x10\x02b\x04\b\x01\x10\x02b\x04\b\x14\x104ZA\n\nSomeString\x10\x01\x18\x83\xff\xff\xff\xff\xff\xff\xff\xff\x01 \xff\x01(\x81\xfe\xff\xff\xff\xff\xff\xff\xff\x010\xff\x01=\x00\x00\xa0\xbfJ\x04\b\x01\x10\x02b\x04\b\x01\x10\x02b\x04\b\x14\x104b\x04\b\x01\x10\x02b\x04\b\x14\x104")

	s.FastForward(29 * time.Second)

	if !s.Exists(key) {
		t.Fatalf("%s should have existed anymore", key)
	}

	s.FastForward(30 * time.Second)

	if s.Exists(key) {
		t.Fatalf("%s should not have existed anymore", key)
	}
}

func TestStringStorageTypeRedisController_Load(t *testing.T) {
	s, err := miniredis.Run()
	if err != nil {
		t.Error(err)
	}
	defer s.Close()

	pool := &redis.Pool{
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", s.Addr())
			if err != nil {
				t.Error(err)
				return nil, err
			}
			return c, nil
		},
	}

	stringStorageType := &StringStorageType{}
	stringStorageType.SomeString = "SomeString"
	stringStorageType.SomeBool = true
	stringStorageType.SomeInt32 = -125
	stringStorageType.SomeUint32 = 255
	stringStorageType.SomeInt64 = -255
	stringStorageType.SomeUint64 = 255
	stringStorageType.SomeFloat = -1.25
	stringStorageType.SomeEnum = StringStorageType_E1
	stringStorageType.Timestamp = &timestamp.Timestamp{Seconds: 1, Nanos: 2}
	stringStorageType.Timestamps = append(stringStorageType.Timestamps, stringStorageType.Timestamp, &timestamp.Timestamp{Seconds: 20, Nanos: 52})

	ctl := NewStringStorageTypeRedisController(pool)

	ctl.StringStorageType().SomeString = "SomeString"
	ctl.StringStorageType().SomeBool = true
	ctl.StringStorageType().SomeInt32 = -125
	ctl.StringStorageType().SomeUint32 = 255
	ctl.StringStorageType().SomeInt64 = -255
	ctl.StringStorageType().SomeUint64 = 255
	ctl.StringStorageType().SomeFloat = -1.25
	ctl.StringStorageType().SomeEnum = StringStorageType_E2
	ctl.StringStorageType().Timestamp = &timestamp.Timestamp{Seconds: 3, Nanos: 4}
	ctl.StringStorageType().StringStorageType = stringStorageType
	ctl.StringStorageType().SomeMessages = append(ctl.StringStorageType().SomeMessages, stringStorageType)
	ctl.StringStorageType().Timestamps = append(ctl.StringStorageType().Timestamps, stringStorageType.Timestamp, &timestamp.Timestamp{Seconds: 20, Nanos: 52})

	key := "TestStringStorageTypeRedisController_Load"

	if err = ctl.Store(context.Background(), key); err != nil {
		t.Error(err)
	}

	ctl2 := NewStringStorageTypeRedisController(pool)

	if err = ctl2.Load(context.Background(), key); err != nil {
		t.Error(err)
	}

	if ctl.StringStorageType().String() != ctl2.StringStorageType().String() {
		t.Fatalf("load %v should be %v", ctl2, ctl)
	}
}

func TestHashStorageTypeRedisController_Store(t *testing.T) {
	s, err := miniredis.Run()
	if err != nil {
		t.Error(err)
	}
	defer s.Close()

	pool := &redis.Pool{
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", s.Addr())
			if err != nil {
				t.Error(err)
				return nil, err
			}
			return c, nil
		},
	}

	stringStorageType := &HashStorageType{}
	stringStorageType.SomeString = "SomeString"
	stringStorageType.SomeBool = true
	stringStorageType.SomeInt32 = -125
	stringStorageType.SomeUint32 = 255
	stringStorageType.SomeInt64 = -255
	stringStorageType.SomeUint64 = 255
	stringStorageType.SomeFloat = -1.25
	stringStorageType.SomeEnum = HashStorageType_E3
	stringStorageType.Timestamp = &timestamp.Timestamp{Seconds: 1, Nanos: 2}
	stringStorageType.Timestamps = append(stringStorageType.Timestamps, stringStorageType.Timestamp, &timestamp.Timestamp{Seconds: 20, Nanos: 52})

	ctl := NewHashStorageTypeRedisController(pool)

	ctl.HashStorageType().SomeString = "SomeString"
	ctl.HashStorageType().SomeBool = true
	ctl.HashStorageType().SomeInt32 = -125
	ctl.HashStorageType().SomeUint32 = 255
	ctl.HashStorageType().SomeInt64 = -255
	ctl.HashStorageType().SomeUint64 = 255
	ctl.HashStorageType().SomeFloat = -1.25
	ctl.HashStorageType().SomeEnum = HashStorageType_E2
	ctl.HashStorageType().Timestamp = &timestamp.Timestamp{Seconds: 3, Nanos: 4}
	ctl.HashStorageType().HashStorageType = stringStorageType
	ctl.HashStorageType().SomeMessages = append(ctl.HashStorageType().SomeMessages, stringStorageType)
	ctl.HashStorageType().Timestamps = append(ctl.HashStorageType().Timestamps, stringStorageType.Timestamp, &timestamp.Timestamp{Seconds: 20, Nanos: 52})

	key := "TestHashStorageTypeRedisController_Store"

	if err = ctl.Store(context.Background(), key); err != nil {
		t.Fatal(err)
	}

	if str := s.HGet(key, "SomeString"); str != "SomeString" {
		t.Error("expect", "SomeString", "got", str)
	}
	if str := s.HGet(key, "SomeBool"); str != "1" {
		t.Error("expect", "1", "got", str)
	}
	if str := s.HGet(key, "SomeInt32"); str != "-125" {
		t.Error("expect", "-125", "got", str)
	}
	if str := s.HGet(key, "SomeUint32"); str != "255" {
		t.Error("expect", "255", "got", str)
	}
	if str := s.HGet(key, "SomeInt64"); str != "-255" {
		t.Error("expect", "-255", "got", str)
	}
	if str := s.HGet(key, "SomeUint64"); str != "255" {
		t.Error("expect", "255", "got", str)
	}
	if str := s.HGet(key, "SomeFloat"); str != "-1.25" {
		t.Error("expect", "-1.25", "got", str)
	}
	if str := s.HGet(key, "SomeEnum"); str != "1" {
		t.Error("expect", "1", "got", str)
	}
	if str := s.HGet(key, "Timestamp"); bytes.Compare([]byte(str), []byte{8, 3, 16, 4}) != 0 {
		t.Error("expect", "[]byte{8, 3, 16, 4}", "got", str)
	}
	if str := s.HGet(key, "HashStorageType"); bytes.Compare([]byte(str), []byte{10, 10, 83, 111, 109, 101, 83, 116, 114, 105, 110, 103, 16, 1, 24, 131, 255, 255, 255, 255, 255, 255, 255, 255, 1, 32, 255, 1, 40, 129, 254, 255, 255, 255, 255, 255, 255, 255, 1, 48, 255, 1, 61, 0, 0, 160, 191, 64, 2, 74, 4, 8, 1, 16, 2, 98, 4, 8, 1, 16, 2, 98, 4, 8, 20, 16, 52}) != 0 {
		t.Log([]byte(str))
		t.Error("expect", "[]byte{10,10,83,111,109,101,83...", "got", str)
	}
	if str := s.HGet(key, "SomeMessages"); str != `[{"some_string":"SomeString","some_bool":true,"some_int32":-125,"some_uint32":255,"some_int64":-255,"some_uint64":255,"some_float":-1.25,"some_enum":2,"timestamp":{"seconds":1,"nanos":2},"timestamps":[{"seconds":1,"nanos":2},{"seconds":20,"nanos":52}]}]` {
		t.Log(str)
		t.Error("expect", `[{"some_string":"SomeString","some_bool":true,"some_int32":-125,"some_uint32":255,"some_int64":-255,"some_uint64":255,"some_float":-1.25,"some_enum":2,"timestamp":{"seconds":1,"nanos":2},"timestamps":[{"seconds":1,"nanos":2},{"seconds":20,"nanos":52}]}]`, "got", str)
	}
	if str := s.HGet(key, "Timestamps"); str != `[{"seconds":1,"nanos":2},{"seconds":20,"nanos":52}]` {
		t.Log(str)
		t.Error("expect", `[{"seconds":1,"nanos":2},{"seconds":20,"nanos":52}]`, "got", str)
	}
}

func TestHashStorageTypeRedisController_StoreWithTTL(t *testing.T) {
	s, err := miniredis.Run()
	if err != nil {
		t.Error(err)
	}
	defer s.Close()

	pool := &redis.Pool{
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", s.Addr())
			if err != nil {
				t.Error(err)
				return nil, err
			}
			return c, nil
		},
	}

	stringStorageType := &HashStorageType{}
	stringStorageType.SomeString = "SomeString"
	stringStorageType.SomeBool = true
	stringStorageType.SomeInt32 = -125
	stringStorageType.SomeUint32 = 255
	stringStorageType.SomeInt64 = -255
	stringStorageType.SomeUint64 = 255
	stringStorageType.SomeFloat = -1.25
	stringStorageType.SomeEnum = HashStorageType_E3
	stringStorageType.Timestamp = &timestamp.Timestamp{Seconds: 1, Nanos: 2}
	stringStorageType.Timestamps = append(stringStorageType.Timestamps, stringStorageType.Timestamp, &timestamp.Timestamp{Seconds: 20, Nanos: 52})

	ctl := NewHashStorageTypeRedisController(pool)

	ctl.HashStorageType().SomeString = "SomeString"
	ctl.HashStorageType().SomeBool = true
	ctl.HashStorageType().SomeInt32 = -125
	ctl.HashStorageType().SomeUint32 = 255
	ctl.HashStorageType().SomeInt64 = -255
	ctl.HashStorageType().SomeUint64 = 255
	ctl.HashStorageType().SomeFloat = -1.25
	ctl.HashStorageType().SomeEnum = HashStorageType_E2
	ctl.HashStorageType().Timestamp = &timestamp.Timestamp{Seconds: 3, Nanos: 4}
	ctl.HashStorageType().HashStorageType = stringStorageType
	ctl.HashStorageType().SomeMessages = append(ctl.HashStorageType().SomeMessages, stringStorageType)
	ctl.HashStorageType().Timestamps = append(ctl.HashStorageType().Timestamps, stringStorageType.Timestamp, &timestamp.Timestamp{Seconds: 20, Nanos: 52})

	key := "TestHashStorageTypeRedisController_StoreWithTTL"

	if err = ctl.StoreWithTTL(context.Background(), key, 30); err != nil {
		t.Error(err)
	}

	if str := s.HGet(key, "SomeString"); str != "SomeString" {
		t.Error("expect", "SomeString", "got", str)
	}
	if str := s.HGet(key, "SomeBool"); str != "1" {
		t.Error("expect", "1", "got", str)
	}
	if str := s.HGet(key, "SomeInt32"); str != "-125" {
		t.Error("expect", "-125", "got", str)
	}
	if str := s.HGet(key, "SomeUint32"); str != "255" {
		t.Error("expect", "255", "got", str)
	}
	if str := s.HGet(key, "SomeInt64"); str != "-255" {
		t.Error("expect", "-255", "got", str)
	}
	if str := s.HGet(key, "SomeUint64"); str != "255" {
		t.Error("expect", "255", "got", str)
	}
	if str := s.HGet(key, "SomeFloat"); str != "-1.25" {
		t.Error("expect", "-1.25", "got", str)
	}
	if str := s.HGet(key, "SomeEnum"); str != "1" {
		t.Error("expect", "1", "got", str)
	}
	if str := s.HGet(key, "Timestamp"); bytes.Compare([]byte(str), []byte{8, 3, 16, 4}) != 0 {
		t.Error("expect", "[]byte{8, 3, 16, 4}", "got", str)
	}
	if str := s.HGet(key, "HashStorageType"); bytes.Compare([]byte(str), []byte{10, 10, 83, 111, 109, 101, 83, 116, 114, 105, 110, 103, 16, 1, 24, 131, 255, 255, 255, 255, 255, 255, 255, 255, 1, 32, 255, 1, 40, 129, 254, 255, 255, 255, 255, 255, 255, 255, 1, 48, 255, 1, 61, 0, 0, 160, 191, 64, 2, 74, 4, 8, 1, 16, 2, 98, 4, 8, 1, 16, 2, 98, 4, 8, 20, 16, 52}) != 0 {
		t.Log([]byte(str))
		t.Error("expect", "[]byte{10,10,83,111,109,101,83...", "got", str)
	}
	if str := s.HGet(key, "SomeMessages"); str != `[{"some_string":"SomeString","some_bool":true,"some_int32":-125,"some_uint32":255,"some_int64":-255,"some_uint64":255,"some_float":-1.25,"some_enum":2,"timestamp":{"seconds":1,"nanos":2},"timestamps":[{"seconds":1,"nanos":2},{"seconds":20,"nanos":52}]}]` {
		t.Log(str)
		t.Error("expect", `[{"some_string":"SomeString","some_bool":true,"some_int32":-125,"some_uint32":255,"some_int64":-255,"some_uint64":255,"some_float":-1.25,"some_enum":2,"timestamp":{"seconds":1,"nanos":2},"timestamps":[{"seconds":1,"nanos":2},{"seconds":20,"nanos":52}]}]`, "got", str)
	}
	if str := s.HGet(key, "Timestamps"); str != `[{"seconds":1,"nanos":2},{"seconds":20,"nanos":52}]` {
		t.Log(str)
		t.Error("expect", `[{"seconds":1,"nanos":2},{"seconds":20,"nanos":52}]`, "got", str)
	}
	s.FastForward(29 * time.Second)

	if !s.Exists(key) {
		t.Fatalf("%s should have existed anymore", key)
	}

	s.FastForward(30 * time.Second)

	if s.Exists(key) {
		t.Fatalf("%s should not have existed anymore", key)
	}
}

func TestHashStorageTypeRedisController_Load(t *testing.T) {
	s, err := miniredis.Run()
	if err != nil {
		t.Error(err)
	}
	defer s.Close()

	pool := &redis.Pool{
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", s.Addr())
			if err != nil {
				t.Error(err)
				return nil, err
			}
			return c, nil
		},
	}

	stringStorageType := &HashStorageType{}
	stringStorageType.SomeString = "SomeString"
	stringStorageType.SomeBool = true
	stringStorageType.SomeInt32 = -125
	stringStorageType.SomeUint32 = 255
	stringStorageType.SomeInt64 = -255
	stringStorageType.SomeUint64 = 255
	stringStorageType.SomeFloat = -1.25
	stringStorageType.SomeEnum = HashStorageType_E3
	stringStorageType.Timestamp = &timestamp.Timestamp{Seconds: 1, Nanos: 2}
	stringStorageType.Timestamps = append(stringStorageType.Timestamps, stringStorageType.Timestamp, &timestamp.Timestamp{Seconds: 20, Nanos: 52})

	ctl := NewHashStorageTypeRedisController(pool)

	ctl.HashStorageType().SomeString = "SomeString"
	ctl.HashStorageType().SomeBool = true
	ctl.HashStorageType().SomeInt32 = -125
	ctl.HashStorageType().SomeUint32 = 255
	ctl.HashStorageType().SomeInt64 = -255
	ctl.HashStorageType().SomeUint64 = 255
	ctl.HashStorageType().SomeFloat = -1.25
	ctl.HashStorageType().SomeEnum = HashStorageType_E2
	ctl.HashStorageType().Timestamp = &timestamp.Timestamp{Seconds: 3, Nanos: 4}
	ctl.HashStorageType().HashStorageType = stringStorageType
	ctl.HashStorageType().SomeMessages = append(ctl.HashStorageType().SomeMessages, stringStorageType)
	ctl.HashStorageType().Timestamps = append(ctl.HashStorageType().Timestamps, stringStorageType.Timestamp, &timestamp.Timestamp{Seconds: 20, Nanos: 52})

	key := "TestHashStorageTypeRedisController_Load"

	if err = ctl.Store(context.Background(), key); err != nil {
		t.Error(err)
	}

	ctl2 := NewHashStorageTypeRedisController(pool)

	if err = ctl2.Load(context.Background(), key); err != nil {
		t.Error(err)
	}

	if ctl.HashStorageType().String() != ctl2.HashStorageType().String() {
		t.Fatalf("load %v should be %v", ctl2, ctl)
	}
}
