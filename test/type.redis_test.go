package test

import (
	"testing"
	"github.com/alicebob/miniredis"
	"github.com/gomodule/redigo/redis"
	"time"
	"context"
	"github.com/golang/protobuf/ptypes/timestamp"
	"bytes"
	"encoding/json"
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

	object := &StringStorageType{}
	object.SomeString = "SomeString"
	object.SomeBool = true
	object.SomeInt32 = -125
	object.SomeUint32 = 255
	object.SomeInt64 = -255
	object.SomeUint64 = 255
	object.SomeFloat = -1.25
	object.SomeEnum = StringStorageType_E1
	object.Timestamp = &timestamp.Timestamp{Seconds: 1, Nanos: 2}
	object.Timestamps = append(object.Timestamps, object.Timestamp, &timestamp.Timestamp{Seconds: 20, Nanos: 52})

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
	ctl.StringStorageType().StringStorageType = object
	ctl.StringStorageType().SomeMessages = append(ctl.StringStorageType().SomeMessages, object)
	ctl.StringStorageType().Timestamps = append(ctl.StringStorageType().Timestamps, object.Timestamp, &timestamp.Timestamp{Seconds: 20, Nanos: 52})

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

	object := &StringStorageType{}
	object.SomeString = "SomeString"
	object.SomeBool = true
	object.SomeInt32 = -125
	object.SomeUint32 = 255
	object.SomeInt64 = -255
	object.SomeUint64 = 255
	object.SomeFloat = -1.25
	object.SomeEnum = StringStorageType_E1
	object.Timestamp = &timestamp.Timestamp{Seconds: 1, Nanos: 2}
	object.Timestamps = append(object.Timestamps, object.Timestamp, &timestamp.Timestamp{Seconds: 20, Nanos: 52})

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
	ctl.StringStorageType().StringStorageType = object
	ctl.StringStorageType().SomeMessages = append(ctl.StringStorageType().SomeMessages, object)
	ctl.StringStorageType().Timestamps = append(ctl.StringStorageType().Timestamps, object.Timestamp, &timestamp.Timestamp{Seconds: 20, Nanos: 52})

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

	object := &StringStorageType{}
	object.SomeString = "SomeString"
	object.SomeBool = true
	object.SomeInt32 = -125
	object.SomeUint32 = 255
	object.SomeInt64 = -255
	object.SomeUint64 = 255
	object.SomeFloat = -1.25
	object.SomeEnum = StringStorageType_E1
	object.Timestamp = &timestamp.Timestamp{Seconds: 1, Nanos: 2}
	object.Timestamps = append(object.Timestamps, object.Timestamp, &timestamp.Timestamp{Seconds: 20, Nanos: 52})

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
	ctl.StringStorageType().StringStorageType = object
	ctl.StringStorageType().SomeMessages = append(ctl.StringStorageType().SomeMessages, object)
	ctl.StringStorageType().Timestamps = append(ctl.StringStorageType().Timestamps, object.Timestamp, &timestamp.Timestamp{Seconds: 20, Nanos: 52})

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

	object := &HashStorageType{}
	object.SomeString = "SomeString"
	object.SomeBool = true
	object.SomeInt32 = -125
	object.SomeUint32 = 255
	object.SomeInt64 = -255
	object.SomeUint64 = 255
	object.SomeFloat = -1.25
	object.SomeEnum = HashStorageType_E3
	object.Timestamp = &timestamp.Timestamp{Seconds: 1, Nanos: 2}
	object.Timestamps = append(object.Timestamps, object.Timestamp, &timestamp.Timestamp{Seconds: 20, Nanos: 52})

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
	ctl.HashStorageType().HashStorageType = object
	ctl.HashStorageType().SomeMessages = append(ctl.HashStorageType().SomeMessages, object)
	ctl.HashStorageType().Timestamps = append(ctl.HashStorageType().Timestamps, object.Timestamp, &timestamp.Timestamp{Seconds: 20, Nanos: 52})

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
		t.Error("expect", "[]byte{10,10,83,111,109,101,83...", "got", str)
	}
	if str := s.HGet(key, "SomeMessages"); str != `[{"some_string":"SomeString","some_bool":true,"some_int32":-125,"some_uint32":255,"some_int64":-255,"some_uint64":255,"some_float":-1.25,"some_enum":2,"timestamp":{"seconds":1,"nanos":2},"timestamps":[{"seconds":1,"nanos":2},{"seconds":20,"nanos":52}]}]` {
		t.Error("expect", `[{"some_string":"SomeString","some_bool":true,"some_int32":-125,"some_uint32":255,"some_int64":-255,"some_uint64":255,"some_float":-1.25,"some_enum":2,"timestamp":{"seconds":1,"nanos":2},"timestamps":[{"seconds":1,"nanos":2},{"seconds":20,"nanos":52}]}]`, "got", str)
	}
	if str := s.HGet(key, "Timestamps"); str != `[{"seconds":1,"nanos":2},{"seconds":20,"nanos":52}]` {
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

	object := &HashStorageType{}
	object.SomeString = "SomeString"
	object.SomeBool = true
	object.SomeInt32 = -125
	object.SomeUint32 = 255
	object.SomeInt64 = -255
	object.SomeUint64 = 255
	object.SomeFloat = -1.25
	object.SomeEnum = HashStorageType_E3
	object.Timestamp = &timestamp.Timestamp{Seconds: 1, Nanos: 2}
	object.Timestamps = append(object.Timestamps, object.Timestamp, &timestamp.Timestamp{Seconds: 20, Nanos: 52})

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
	ctl.HashStorageType().HashStorageType = object
	ctl.HashStorageType().SomeMessages = append(ctl.HashStorageType().SomeMessages, object)
	ctl.HashStorageType().Timestamps = append(ctl.HashStorageType().Timestamps, object.Timestamp, &timestamp.Timestamp{Seconds: 20, Nanos: 52})

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
		t.Error("expect", "[]byte{10,10,83,111,109,101,83...", "got", str)
	}
	if str := s.HGet(key, "SomeMessages"); str != `[{"some_string":"SomeString","some_bool":true,"some_int32":-125,"some_uint32":255,"some_int64":-255,"some_uint64":255,"some_float":-1.25,"some_enum":2,"timestamp":{"seconds":1,"nanos":2},"timestamps":[{"seconds":1,"nanos":2},{"seconds":20,"nanos":52}]}]` {
		t.Error("expect", `[{"some_string":"SomeString","some_bool":true,"some_int32":-125,"some_uint32":255,"some_int64":-255,"some_uint64":255,"some_float":-1.25,"some_enum":2,"timestamp":{"seconds":1,"nanos":2},"timestamps":[{"seconds":1,"nanos":2},{"seconds":20,"nanos":52}]}]`, "got", str)
	}
	if str := s.HGet(key, "Timestamps"); str != `[{"seconds":1,"nanos":2},{"seconds":20,"nanos":52}]` {
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

	object := &HashStorageType{}
	object.SomeString = "SomeString"
	object.SomeBool = true
	object.SomeInt32 = -125
	object.SomeUint32 = 255
	object.SomeInt64 = -255
	object.SomeUint64 = 255
	object.SomeFloat = -1.25
	object.SomeEnum = HashStorageType_E3
	object.Timestamp = &timestamp.Timestamp{Seconds: 1, Nanos: 2}
	object.Timestamps = append(object.Timestamps, object.Timestamp, &timestamp.Timestamp{Seconds: 20, Nanos: 52})

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
	ctl.HashStorageType().HashStorageType = object
	ctl.HashStorageType().SomeMessages = append(ctl.HashStorageType().SomeMessages, object)
	ctl.HashStorageType().Timestamps = append(ctl.HashStorageType().Timestamps, object.Timestamp, &timestamp.Timestamp{Seconds: 20, Nanos: 52})

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

func TestHashStorageTypeRedisController_GetSomeString(t *testing.T) {
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

	key := "TestHashStorageTypeRedisController_GetSomeString"
	value := "SomeString"

	ctl := NewHashStorageTypeRedisController(pool)

	if err := ctl.SetSomeString(key, value); err != nil {
		t.Error(err)
	}

	if ctl.HashStorageType().SomeString != value {
		t.Error(key, "value should be", value)
	}

	ctl2 := NewHashStorageTypeRedisController(pool)

	if v, err := ctl2.GetSomeString(key); v != value || err != nil {
		t.Error(key, "value should be", value)
	}
}

func TestHashStorageTypeRedisController_SetSomeString(t *testing.T) {
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

	key := "TestHashStorageTypeRedisController_SetSomeString"
	value := "SomeString"

	ctl := NewHashStorageTypeRedisController(pool)

	if err := ctl.SetSomeString(key, value); err != nil {
		t.Error(err)
	}

	if v := s.HGet(key, "SomeString"); v != value {
		t.Error(key, "value should be", value)
	}
}

func TestHashStorageTypeRedisController_GetSomeBool(t *testing.T) {
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

	key := "TestHashStorageTypeRedisController_GetSomeBool"
	var value bool

	value = true

	ctl := NewHashStorageTypeRedisController(pool)

	if err := ctl.SetSomeBool(key, value); err != nil {
		t.Error(err)
	}

	if ctl.HashStorageType().SomeBool != value {
		t.Error(key, "value should be", value)
	}

	ctl2 := NewHashStorageTypeRedisController(pool)

	if v, err := ctl2.GetSomeBool(key); v != value || err != nil {
		t.Error(key, "value should be", value)
	}

	value = false

	if err := ctl.SetSomeBool(key, value); err != nil {
		t.Error(err)
	}

	if ctl.HashStorageType().SomeBool != value {
		t.Error(key, "value should be", value)
	}

	if v, err := ctl2.GetSomeBool(key); v != value || err != nil {
		t.Error(key, "value should be", value)
	}
}

func TestHashStorageTypeRedisController_SetSomeBool(t *testing.T) {
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

	key := "TestHashStorageTypeRedisController_SetSomeBool"
	var value bool

	value = true

	ctl := NewHashStorageTypeRedisController(pool)

	if err := ctl.SetSomeBool(key, value); err != nil {
		t.Error(err)
	}

	if v := s.HGet(key, "SomeBool"); v != "1" {
		t.Error(key, "value should be", value)
	}

	value = false

	if err := ctl.SetSomeBool(key, value); err != nil {
		t.Error(err)
	}

	if v := s.HGet(key, "SomeBool"); v != "0" {
		t.Error(key, "value should be", value)
	}
}

func TestHashStorageTypeRedisController_GetSomeInt32(t *testing.T) {
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

	key := "TestHashStorageTypeRedisController_GetSomeInt32"
	value := int32(-125)

	ctl := NewHashStorageTypeRedisController(pool)

	if err := ctl.SetSomeInt32(key, value); err != nil {
		t.Error(err)
	}

	if ctl.HashStorageType().SomeInt32 != value {
		t.Error(key, "value should be", value)
	}

	ctl2 := NewHashStorageTypeRedisController(pool)

	if v, err := ctl2.GetSomeInt32(key); v != value || err != nil {
		t.Error(key, "value should be", value)
	}
}

func TestHashStorageTypeRedisController_SetSomeInt32(t *testing.T) {
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

	key := "TestHashStorageTypeRedisController_SetSomeInt32"
	value := int32(-125)

	ctl := NewHashStorageTypeRedisController(pool)

	if err := ctl.SetSomeInt32(key, value); err != nil {
		t.Error(err)
	}

	if v := s.HGet(key, "SomeInt32"); v != "-125" {
		t.Error(key, "value should be", value)
	}
}

func TestHashStorageTypeRedisController_GetSomeUint32(t *testing.T) {
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

	key := "TestHashStorageTypeRedisController_GetSomeUint32"
	value := uint32(125)

	ctl := NewHashStorageTypeRedisController(pool)

	if err := ctl.SetSomeUint32(key, value); err != nil {
		t.Error(err)
	}

	if ctl.HashStorageType().SomeUint32 != value {
		t.Error(key, "value should be", value)
	}

	ctl2 := NewHashStorageTypeRedisController(pool)

	if v, err := ctl2.GetSomeUint32(key); v != value || err != nil {
		t.Error(key, "value should be", value)
	}
}

func TestHashStorageTypeRedisController_SetSomeUint32(t *testing.T) {
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

	key := "TestHashStorageTypeRedisController_SetSomeUint32"
	value := uint32(125)

	ctl := NewHashStorageTypeRedisController(pool)

	if err := ctl.SetSomeUint32(key, value); err != nil {
		t.Error(err)
	}

	if v := s.HGet(key, "SomeUint32"); v != "125" {
		t.Error(key, "value should be", value)
	}
}

func TestHashStorageTypeRedisController_GetSomeInt64(t *testing.T) {
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

	key := "TestHashStorageTypeRedisController_GetSomeInt64"
	value := int64(-125)

	ctl := NewHashStorageTypeRedisController(pool)

	if err := ctl.SetSomeInt64(key, value); err != nil {
		t.Error(err)
	}

	if ctl.HashStorageType().SomeInt64 != value {
		t.Error(key, "value should be", value)
	}

	ctl2 := NewHashStorageTypeRedisController(pool)

	if v, err := ctl2.GetSomeInt64(key); v != value || err != nil {
		t.Error(key, "value should be", value)
	}
}

func TestHashStorageTypeRedisController_SetSomeInt64(t *testing.T) {
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

	key := "TestHashStorageTypeRedisController_SetSomeInt64"
	value := int64(-125)

	ctl := NewHashStorageTypeRedisController(pool)

	if err := ctl.SetSomeInt64(key, value); err != nil {
		t.Error(err)
	}

	if v := s.HGet(key, "SomeInt64"); v != "-125" {
		t.Error(key, "value should be", value)
	}
}

func TestHashStorageTypeRedisController_GetSomeUint64(t *testing.T) {
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

	key := "TestHashStorageTypeRedisController_GetSomeUint64"
	value := uint64(125)

	ctl := NewHashStorageTypeRedisController(pool)

	if err := ctl.SetSomeUint64(key, value); err != nil {
		t.Error(err)
	}

	if ctl.HashStorageType().SomeUint64 != value {
		t.Error(key, "value should be", value)
	}

	ctl2 := NewHashStorageTypeRedisController(pool)

	if v, err := ctl2.GetSomeUint64(key); v != value || err != nil {
		t.Error(key, "value should be", value)
	}
}

func TestHashStorageTypeRedisController_SetSomeUint64(t *testing.T) {
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

	key := "TestHashStorageTypeRedisController_SetSomeUint64"
	value := uint64(125)

	ctl := NewHashStorageTypeRedisController(pool)

	if err := ctl.SetSomeUint64(key, value); err != nil {
		t.Error(err)
	}

	if v := s.HGet(key, "SomeUint64"); v != "125" {
		t.Error(key, "value should be", value)
	}
}

func TestHashStorageTypeRedisController_GetSomeFloat(t *testing.T) {
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

	key := "TestHashStorageTypeRedisController_GetSomeFloat"
	value := float32(1.25000)

	ctl := NewHashStorageTypeRedisController(pool)

	if err := ctl.SetSomeFloat(key, value); err != nil {
		t.Error(err)
	}

	if ctl.HashStorageType().SomeFloat != value {
		t.Error(key, "value should be", value)
	}

	ctl2 := NewHashStorageTypeRedisController(pool)

	if v, err := ctl2.GetSomeFloat(key); v != value || err != nil {
		t.Error(key, "value should be", value)
	}
}

func TestHashStorageTypeRedisController_SetSomeFloat(t *testing.T) {
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

	key := "TestHashStorageTypeRedisController_SetSomeFloat"
	value := float32(1.25000)

	ctl := NewHashStorageTypeRedisController(pool)

	if err := ctl.SetSomeFloat(key, value); err != nil {
		t.Error(err)
	}

	if v := s.HGet(key, "SomeFloat"); v != "1.25" {
		t.Error(key, "value should be", value)
	}
}

func TestHashStorageTypeRedisController_GetSomeEnum(t *testing.T) {
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

	key := "TestHashStorageTypeRedisController_GetSomeEnum"
	value := HashStorageType_E2

	ctl := NewHashStorageTypeRedisController(pool)

	if err := ctl.SetSomeEnum(key, value); err != nil {
		t.Error(err)
	}

	if ctl.HashStorageType().SomeEnum != value {
		t.Error(key, "value should be", value)
	}

	ctl2 := NewHashStorageTypeRedisController(pool)

	if v, err := ctl2.GetSomeEnum(key); v != value || err != nil {
		t.Error(key, "value should be", value)
	}
}

func TestHashStorageTypeRedisController_SetSomeEnum(t *testing.T) {
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

	key := "TestHashStorageTypeRedisController_SetSomeEnum"
	value := HashStorageType_E2

	ctl := NewHashStorageTypeRedisController(pool)

	if err := ctl.SetSomeEnum(key, value); err != nil {
		t.Error(err)
	}

	if v := s.HGet(key, "SomeEnum"); v != "1" {
		t.Error(key, "value should be", value)
	}
}

func TestHashStorageTypeRedisController_GetTimestamp(t *testing.T) {
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

	key := "TestHashStorageTypeRedisController_GetTimestamp"
	value := &timestamp.Timestamp{Seconds: 1, Nanos: 4}

	ctl := NewHashStorageTypeRedisController(pool)

	if err := ctl.SetTimestamp(key, value); err != nil {
		t.Error(err)
	}

	if ctl.HashStorageType().Timestamp != value {
		t.Error(key, "value should be", value)
	}

	ctl2 := NewHashStorageTypeRedisController(pool)

	if v, err := ctl2.GetTimestamp(key); (v.Seconds != 1 || v.Nanos != 4) || err != nil {
		t.Error(key, "is", v, "value should be", value)
	}
}

func TestHashStorageTypeRedisController_SetTimestamp(t *testing.T) {
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

	key := "TestHashStorageTypeRedisController_SetTimestamp"
	value := &timestamp.Timestamp{Seconds: 1, Nanos: 4}

	ctl := NewHashStorageTypeRedisController(pool)

	if err := ctl.SetTimestamp(key, value); err != nil {
		t.Error(err)
	}

	if v := s.HGet(key, "Timestamp"); v != "\b\x01\x10\x04" {
		t.Error(key, "is", v, "value should be", value)

	}
}

func TestHashStorageTypeRedisController_GetHashStorageType(t *testing.T) {
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

	key := "TestHashStorageTypeRedisController_GetHashStorageType"
	object := &HashStorageType{}
	object.SomeString = "SomeString"
	object.SomeBool = true
	object.SomeInt32 = -125
	object.SomeUint32 = 255
	object.SomeInt64 = -255
	object.SomeUint64 = 255
	object.SomeFloat = -1.25
	object.SomeEnum = HashStorageType_E3
	object.Timestamp = &timestamp.Timestamp{Seconds: 1, Nanos: 2}
	object.Timestamps = append(object.Timestamps, object.Timestamp, &timestamp.Timestamp{Seconds: 20, Nanos: 52})

	ctl := NewHashStorageTypeRedisController(pool)

	if err := ctl.SetHashStorageTypeField(key, object); err != nil {
		t.Error(err)
	}

	if ctl.HashStorageType().HashStorageType.String() != object.String() {
		t.Error(key, "value should be", object)
	}

	ctl2 := NewHashStorageTypeRedisController(pool)

	if v, err := ctl2.GetHashStorageType(key); err != nil || v.String() != object.String() {
		t.Error(key, "is", v, "value should be", object)
	}
}

func TestHashStorageTypeRedisController_SetHashStorageType(t *testing.T) {
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

	key := "TestHashStorageTypeRedisController_SetHashStorageType"
	object := &HashStorageType{}
	object.SomeString = "SomeString"
	object.SomeBool = true
	object.SomeInt32 = -125
	object.SomeUint32 = 255
	object.SomeInt64 = -255
	object.SomeUint64 = 255
	object.SomeFloat = -1.25
	object.SomeEnum = HashStorageType_E3
	object.Timestamp = &timestamp.Timestamp{Seconds: 1, Nanos: 2}
	object.Timestamps = append(object.Timestamps, object.Timestamp, &timestamp.Timestamp{Seconds: 20, Nanos: 52})

	ctl := NewHashStorageTypeRedisController(pool)

	if err := ctl.SetHashStorageTypeField(key, object); err != nil {
		t.Error(err)
	}

	if v := s.HGet(key, "HashStorageType"); v != "\n\nSomeString\x10\x01\x18\x83\xff\xff\xff\xff\xff\xff\xff\xff\x01 \xff\x01(\x81\xfe\xff\xff\xff\xff\xff\xff\xff\x010\xff\x01=\x00\x00\xa0\xbf@\x02J\x04\b\x01\x10\x02b\x04\b\x01\x10\x02b\x04\b\x14\x104" {
		t.Error(key, "is", v, "value should be", )
	}
}

func TestHashStorageTypeRedisController_GetSomeMessages(t *testing.T) {
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

	key := "TestHashStorageTypeRedisController_GetSomeMessages"
	object := &HashStorageType{}
	object.SomeString = "SomeString"
	object.SomeBool = true
	object.SomeInt32 = -125
	object.SomeUint32 = 255
	object.SomeInt64 = -255
	object.SomeUint64 = 255
	object.SomeFloat = -1.25
	object.SomeEnum = HashStorageType_E2
	object.Timestamp = &timestamp.Timestamp{Seconds: 1, Nanos: 2}
	object.Timestamps = append(object.Timestamps, object.Timestamp, &timestamp.Timestamp{Seconds: 20, Nanos: 52})

	var someMessages []*HashStorageType
	someMessages = append(someMessages, object)

	ctl := NewHashStorageTypeRedisController(pool)

	if err := ctl.SetSomeMessages(key, someMessages); err != nil {
		t.Error(err)
	}

	d1, _ := json.Marshal(&someMessages)
	d2, _ := json.Marshal(&ctl.HashStorageType().SomeMessages)

	if bytes.Compare(d1, d2) != 0 {
		t.Error(key, "value should be", string(d1))
	}

	ctl2 := NewHashStorageTypeRedisController(pool)

	if v, err := ctl2.GetSomeMessages(key); err != nil {
		t.Error(err)
	} else {
		d3, _ := json.Marshal(v)
		if bytes.Compare(d1, d3) != 0 {
			t.Error(key, "value should be", string(d2))
		}
	}
}

func TestHashStorageTypeRedisController_SetSomeMessages(t *testing.T) {
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

	key := "TestHashStorageTypeRedisController_SetSomeMessages"
	object := &HashStorageType{}
	object.SomeString = "SomeString"
	object.SomeBool = true
	object.SomeInt32 = -125
	object.SomeUint32 = 255
	object.SomeInt64 = -255
	object.SomeUint64 = 255
	object.SomeFloat = -1.25
	object.SomeEnum = HashStorageType_E2
	object.Timestamp = &timestamp.Timestamp{Seconds: 1, Nanos: 2}
	object.Timestamps = append(object.Timestamps, object.Timestamp, &timestamp.Timestamp{Seconds: 20, Nanos: 52})

	var someMessages []*HashStorageType
	someMessages = append(someMessages, object)

	ctl := NewHashStorageTypeRedisController(pool)

	if err := ctl.SetSomeMessages(key, someMessages); err != nil {
		t.Error(err)
	}

	d1, _ := json.Marshal(someMessages)

	if v := s.HGet(key, "SomeMessages"); v != string(d1) {
		t.Error(key, "is", v, "value should be", string(d1))
	}
}

func TestHashStorageTypeRedisController_GetTimestamps(t *testing.T) {
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

	key := "TestHashStorageTypeRedisController_GetTimestamps"

	var timestamps []*timestamp.Timestamp
	timestamps = append(timestamps, &timestamp.Timestamp{Seconds: 1, Nanos: 4}, &timestamp.Timestamp{Seconds: 2, Nanos: 3})

	ctl := NewHashStorageTypeRedisController(pool)

	if err := ctl.SetTimestamps(key, timestamps); err != nil {
		t.Error(err)
	}

	d1, _ := json.Marshal(timestamps)
	d2, _ := json.Marshal(ctl.HashStorageType().Timestamps)

	if bytes.Compare(d1, d2) != 0 {
		t.Error(key, "value should be", string(d1))
	}

	ctl2 := NewHashStorageTypeRedisController(pool)

	if v, err := ctl2.GetTimestamps(key); err != nil {
		t.Error(err)
	} else {
		d3, _ := json.Marshal(v)
		if bytes.Compare(d1, d3) != 0 {
			t.Error(key, "value should be", string(d2))
		}
	}
}

func TestHashStorageTypeRedisController_SetTimestamps(t *testing.T) {
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

	key := "TestHashStorageTypeRedisController_SetTimestamp"
	var timestamps []*timestamp.Timestamp
	timestamps = append(timestamps, &timestamp.Timestamp{Seconds: 1, Nanos: 4}, &timestamp.Timestamp{Seconds: 2, Nanos: 3})

	ctl := NewHashStorageTypeRedisController(pool)

	if err := ctl.SetTimestamps(key, timestamps); err != nil {
		t.Error(err)
	}

	d1, _ := json.Marshal(timestamps)
	d2, _ := json.Marshal(ctl.HashStorageType().Timestamps)

	if bytes.Compare(d1, d2) != 0 {
		t.Error(key, "value should be", string(d1))
	}

	if v := s.HGet(key, "Timestamps"); v != string(d1) {
		t.Error(key, "is", v, "value should be", string(d1))
	}
}
