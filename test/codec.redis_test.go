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

func TestJsonStringStorageTypeRedisController_Store(t *testing.T) {
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

	object := &JsonStringStorageType{}
	object.SomeString = "SomeString"
	object.SomeBool = true
	object.SomeInt32 = -125
	object.SomeUint32 = 255
	object.SomeInt64 = -255
	object.SomeUint64 = 255
	object.SomeFloat = -1.25
	object.SomeEnum = JsonStringStorageType_E1
	object.Timestamp = &timestamp.Timestamp{Seconds: 1, Nanos: 2}
	object.Timestamps = append(object.Timestamps, object.Timestamp, &timestamp.Timestamp{Seconds: 20, Nanos: 52})

	ctl := NewJsonStringStorageTypeRedisController(pool)

	ctl.JsonStringStorageType().SomeString = "SomeString"
	ctl.JsonStringStorageType().SomeBool = true
	ctl.JsonStringStorageType().SomeInt32 = -125
	ctl.JsonStringStorageType().SomeUint32 = 255
	ctl.JsonStringStorageType().SomeInt64 = -255
	ctl.JsonStringStorageType().SomeUint64 = 255
	ctl.JsonStringStorageType().SomeFloat = -1.25
	ctl.JsonStringStorageType().SomeEnum = JsonStringStorageType_E2
	ctl.JsonStringStorageType().Timestamp = &timestamp.Timestamp{Seconds: 3, Nanos: 4}
	ctl.JsonStringStorageType().JsonStringStorageType = object
	ctl.JsonStringStorageType().SomeMessages = append(ctl.JsonStringStorageType().SomeMessages, object)
	ctl.JsonStringStorageType().Timestamps = append(ctl.JsonStringStorageType().Timestamps, object.Timestamp, &timestamp.Timestamp{Seconds: 20, Nanos: 52})

	key := "TestJsonStringStorageTypeRedisController_Store"

	if err = ctl.Store(context.Background(), key); err != nil {
		t.Error(err)
	}

	s.CheckGet(t, key, "{\"some_string\":\"SomeString\",\"some_bool\":true,\"some_int32\":-125,\"some_uint32\":255,\"some_int64\":-255,\"some_uint64\":255,\"some_float\":-1.25,\"some_enum\":1,\"timestamp\":{\"seconds\":3,\"nanos\":4},\"JsonStringStorageType\":{\"some_string\":\"SomeString\",\"some_bool\":true,\"some_int32\":-125,\"some_uint32\":255,\"some_int64\":-255,\"some_uint64\":255,\"some_float\":-1.25,\"timestamp\":{\"seconds\":1,\"nanos\":2},\"timestamps\":[{\"seconds\":1,\"nanos\":2},{\"seconds\":20,\"nanos\":52}]},\"some_messages\":[{\"some_string\":\"SomeString\",\"some_bool\":true,\"some_int32\":-125,\"some_uint32\":255,\"some_int64\":-255,\"some_uint64\":255,\"some_float\":-1.25,\"timestamp\":{\"seconds\":1,\"nanos\":2},\"timestamps\":[{\"seconds\":1,\"nanos\":2},{\"seconds\":20,\"nanos\":52}]}],\"timestamps\":[{\"seconds\":1,\"nanos\":2},{\"seconds\":20,\"nanos\":52}]}")
}

func TestJsonStringStorageTypeRedisController_StoreWithTTL(t *testing.T) {
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

	object := &JsonStringStorageType{}
	object.SomeString = "SomeString"
	object.SomeBool = true
	object.SomeInt32 = -125
	object.SomeUint32 = 255
	object.SomeInt64 = -255
	object.SomeUint64 = 255
	object.SomeFloat = -1.25
	object.SomeEnum = JsonStringStorageType_E1
	object.Timestamp = &timestamp.Timestamp{Seconds: 1, Nanos: 2}
	object.Timestamps = append(object.Timestamps, object.Timestamp, &timestamp.Timestamp{Seconds: 20, Nanos: 52})

	ctl := NewJsonStringStorageTypeRedisController(pool)

	ctl.JsonStringStorageType().SomeString = "SomeString"
	ctl.JsonStringStorageType().SomeBool = true
	ctl.JsonStringStorageType().SomeInt32 = -125
	ctl.JsonStringStorageType().SomeUint32 = 255
	ctl.JsonStringStorageType().SomeInt64 = -255
	ctl.JsonStringStorageType().SomeUint64 = 255
	ctl.JsonStringStorageType().SomeFloat = -1.25
	ctl.JsonStringStorageType().SomeEnum = JsonStringStorageType_E2
	ctl.JsonStringStorageType().Timestamp = &timestamp.Timestamp{Seconds: 3, Nanos: 4}
	ctl.JsonStringStorageType().JsonStringStorageType = object
	ctl.JsonStringStorageType().SomeMessages = append(ctl.JsonStringStorageType().SomeMessages, object)
	ctl.JsonStringStorageType().Timestamps = append(ctl.JsonStringStorageType().Timestamps, object.Timestamp, &timestamp.Timestamp{Seconds: 20, Nanos: 52})

	key := "TestJsonStringStorageTypeRedisController_StoreWithTTL"

	if err = ctl.StoreWithTTL(context.Background(), key, 30); err != nil {
		t.Error(err)
	}

	s.CheckGet(t, key, "{\"some_string\":\"SomeString\",\"some_bool\":true,\"some_int32\":-125,\"some_uint32\":255,\"some_int64\":-255,\"some_uint64\":255,\"some_float\":-1.25,\"some_enum\":1,\"timestamp\":{\"seconds\":3,\"nanos\":4},\"JsonStringStorageType\":{\"some_string\":\"SomeString\",\"some_bool\":true,\"some_int32\":-125,\"some_uint32\":255,\"some_int64\":-255,\"some_uint64\":255,\"some_float\":-1.25,\"timestamp\":{\"seconds\":1,\"nanos\":2},\"timestamps\":[{\"seconds\":1,\"nanos\":2},{\"seconds\":20,\"nanos\":52}]},\"some_messages\":[{\"some_string\":\"SomeString\",\"some_bool\":true,\"some_int32\":-125,\"some_uint32\":255,\"some_int64\":-255,\"some_uint64\":255,\"some_float\":-1.25,\"timestamp\":{\"seconds\":1,\"nanos\":2},\"timestamps\":[{\"seconds\":1,\"nanos\":2},{\"seconds\":20,\"nanos\":52}]}],\"timestamps\":[{\"seconds\":1,\"nanos\":2},{\"seconds\":20,\"nanos\":52}]}")

	s.FastForward(29 * time.Second)

	if !s.Exists(key) {
		t.Fatalf("%s should have existed anymore", key)
	}

	s.FastForward(30 * time.Second)

	if s.Exists(key) {
		t.Fatalf("%s should not have existed anymore", key)
	}
}

func TestJsonStringStorageTypeRedisController_Load(t *testing.T) {
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

	object := &JsonStringStorageType{}
	object.SomeString = "SomeString"
	object.SomeBool = true
	object.SomeInt32 = -125
	object.SomeUint32 = 255
	object.SomeInt64 = -255
	object.SomeUint64 = 255
	object.SomeFloat = -1.25
	object.SomeEnum = JsonStringStorageType_E1
	object.Timestamp = &timestamp.Timestamp{Seconds: 1, Nanos: 2}
	object.Timestamps = append(object.Timestamps, object.Timestamp, &timestamp.Timestamp{Seconds: 20, Nanos: 52})

	ctl := NewJsonStringStorageTypeRedisController(pool)

	ctl.JsonStringStorageType().SomeString = "SomeString"
	ctl.JsonStringStorageType().SomeBool = true
	ctl.JsonStringStorageType().SomeInt32 = -125
	ctl.JsonStringStorageType().SomeUint32 = 255
	ctl.JsonStringStorageType().SomeInt64 = -255
	ctl.JsonStringStorageType().SomeUint64 = 255
	ctl.JsonStringStorageType().SomeFloat = -1.25
	ctl.JsonStringStorageType().SomeEnum = JsonStringStorageType_E2
	ctl.JsonStringStorageType().Timestamp = &timestamp.Timestamp{Seconds: 3, Nanos: 4}
	ctl.JsonStringStorageType().JsonStringStorageType = object
	ctl.JsonStringStorageType().SomeMessages = append(ctl.JsonStringStorageType().SomeMessages, object)
	ctl.JsonStringStorageType().Timestamps = append(ctl.JsonStringStorageType().Timestamps, object.Timestamp, &timestamp.Timestamp{Seconds: 20, Nanos: 52})

	key := "TestJsonStringStorageTypeRedisController_Load"

	if err = ctl.Store(context.Background(), key); err != nil {
		t.Error(err)
	}

	ctl2 := NewJsonStringStorageTypeRedisController(pool)

	if err = ctl2.Load(context.Background(), key); err != nil {
		t.Error(err)
	}

	if ctl.JsonStringStorageType().String() != ctl2.JsonStringStorageType().String() {
		t.Fatalf("load %v should be %v", ctl2, ctl)
	}
}

func TestJsonHashStorageTypeRedisController_Store(t *testing.T) {
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

	object := &JsonHashStorageType{}
	object.SomeString = "SomeString"
	object.SomeBool = true
	object.SomeInt32 = -125
	object.SomeUint32 = 255
	object.SomeInt64 = -255
	object.SomeUint64 = 255
	object.SomeFloat = -1.25
	object.SomeEnum = JsonHashStorageType_E3
	object.Timestamp = &timestamp.Timestamp{Seconds: 1, Nanos: 2}
	object.Timestamps = append(object.Timestamps, object.Timestamp, &timestamp.Timestamp{Seconds: 20, Nanos: 52})

	ctl := NewJsonHashStorageTypeRedisController(pool)

	ctl.JsonHashStorageType().SomeString = "SomeString"
	ctl.JsonHashStorageType().SomeBool = true
	ctl.JsonHashStorageType().SomeInt32 = -125
	ctl.JsonHashStorageType().SomeUint32 = 255
	ctl.JsonHashStorageType().SomeInt64 = -255
	ctl.JsonHashStorageType().SomeUint64 = 255
	ctl.JsonHashStorageType().SomeFloat = -1.25
	ctl.JsonHashStorageType().SomeEnum = JsonHashStorageType_E2
	ctl.JsonHashStorageType().Timestamp = &timestamp.Timestamp{Seconds: 3, Nanos: 4}
	ctl.JsonHashStorageType().JsonHashStorageType = object
	ctl.JsonHashStorageType().SomeMessages = append(ctl.JsonHashStorageType().SomeMessages, object)
	ctl.JsonHashStorageType().Timestamps = append(ctl.JsonHashStorageType().Timestamps, object.Timestamp, &timestamp.Timestamp{Seconds: 20, Nanos: 52})

	key := "TestJsonHashStorageTypeRedisController_Store"

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
	if str := s.HGet(key, "Timestamp"); bytes.Compare([]byte(str), []byte(`{"seconds":3,"nanos":4}`)) != 0 {
		t.Error("expect", "[]byte{8, 3, 16, 4}", "got", str)
	}
	if str := s.HGet(key, "JsonHashStorageType"); bytes.Compare([]byte(str), []byte(`{"some_string":"SomeString","some_bool":true,"some_int32":-125,"some_uint32":255,"some_int64":-255,"some_uint64":255,"some_float":-1.25,"some_enum":2,"timestamp":{"seconds":1,"nanos":2},"timestamps":[{"seconds":1,"nanos":2},{"seconds":20,"nanos":52}]}`)) != 0 {
		t.Error("expect", "[]byte{10,10,83,111,109,101,83...", "got", str)
	}
	if str := s.HGet(key, "SomeMessages"); str != `[{"some_string":"SomeString","some_bool":true,"some_int32":-125,"some_uint32":255,"some_int64":-255,"some_uint64":255,"some_float":-1.25,"some_enum":2,"timestamp":{"seconds":1,"nanos":2},"timestamps":[{"seconds":1,"nanos":2},{"seconds":20,"nanos":52}]}]` {
		t.Error("expect", `[{"some_string":"SomeString","some_bool":true,"some_int32":-125,"some_uint32":255,"some_int64":-255,"some_uint64":255,"some_float":-1.25,"some_enum":2,"timestamp":{"seconds":1,"nanos":2},"timestamps":[{"seconds":1,"nanos":2},{"seconds":20,"nanos":52}]}]`, "got", str)
	}
	if str := s.HGet(key, "Timestamps"); str != `[{"seconds":1,"nanos":2},{"seconds":20,"nanos":52}]` {
		t.Error("expect", `[{"seconds":1,"nanos":2},{"seconds":20,"nanos":52}]`, "got", str)
	}
}

func TestJsonHashStorageTypeRedisController_StoreWithTTL(t *testing.T) {
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

	object := &JsonHashStorageType{}
	object.SomeString = "SomeString"
	object.SomeBool = true
	object.SomeInt32 = -125
	object.SomeUint32 = 255
	object.SomeInt64 = -255
	object.SomeUint64 = 255
	object.SomeFloat = -1.25
	object.SomeEnum = JsonHashStorageType_E3
	object.Timestamp = &timestamp.Timestamp{Seconds: 1, Nanos: 2}
	object.Timestamps = append(object.Timestamps, object.Timestamp, &timestamp.Timestamp{Seconds: 20, Nanos: 52})

	ctl := NewJsonHashStorageTypeRedisController(pool)

	ctl.JsonHashStorageType().SomeString = "SomeString"
	ctl.JsonHashStorageType().SomeBool = true
	ctl.JsonHashStorageType().SomeInt32 = -125
	ctl.JsonHashStorageType().SomeUint32 = 255
	ctl.JsonHashStorageType().SomeInt64 = -255
	ctl.JsonHashStorageType().SomeUint64 = 255
	ctl.JsonHashStorageType().SomeFloat = -1.25
	ctl.JsonHashStorageType().SomeEnum = JsonHashStorageType_E2
	ctl.JsonHashStorageType().Timestamp = &timestamp.Timestamp{Seconds: 3, Nanos: 4}
	ctl.JsonHashStorageType().JsonHashStorageType = object
	ctl.JsonHashStorageType().SomeMessages = append(ctl.JsonHashStorageType().SomeMessages, object)
	ctl.JsonHashStorageType().Timestamps = append(ctl.JsonHashStorageType().Timestamps, object.Timestamp, &timestamp.Timestamp{Seconds: 20, Nanos: 52})

	key := "TestJsonHashStorageTypeRedisController_StoreWithTTL"

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
	if str := s.HGet(key, "Timestamp"); bytes.Compare([]byte(str), []byte(`{"seconds":3,"nanos":4}`)) != 0 {
		t.Error("expect", "[]byte{8, 3, 16, 4}", "got", str)
	}
	if str := s.HGet(key, "JsonHashStorageType"); bytes.Compare([]byte(str), []byte(`{"some_string":"SomeString","some_bool":true,"some_int32":-125,"some_uint32":255,"some_int64":-255,"some_uint64":255,"some_float":-1.25,"some_enum":2,"timestamp":{"seconds":1,"nanos":2},"timestamps":[{"seconds":1,"nanos":2},{"seconds":20,"nanos":52}]}`)) != 0 {
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

func TestJsonHashStorageTypeRedisController_Load(t *testing.T) {
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

	object := &JsonHashStorageType{}
	object.SomeString = "SomeString"
	object.SomeBool = true
	object.SomeInt32 = -125
	object.SomeUint32 = 255
	object.SomeInt64 = -255
	object.SomeUint64 = 255
	object.SomeFloat = -1.25
	object.SomeEnum = JsonHashStorageType_E3
	object.Timestamp = &timestamp.Timestamp{Seconds: 1, Nanos: 2}
	object.Timestamps = append(object.Timestamps, object.Timestamp, &timestamp.Timestamp{Seconds: 20, Nanos: 52})

	ctl := NewJsonHashStorageTypeRedisController(pool)

	ctl.JsonHashStorageType().SomeString = "SomeString"
	ctl.JsonHashStorageType().SomeBool = true
	ctl.JsonHashStorageType().SomeInt32 = -125
	ctl.JsonHashStorageType().SomeUint32 = 255
	ctl.JsonHashStorageType().SomeInt64 = -255
	ctl.JsonHashStorageType().SomeUint64 = 255
	ctl.JsonHashStorageType().SomeFloat = -1.25
	ctl.JsonHashStorageType().SomeEnum = JsonHashStorageType_E2
	ctl.JsonHashStorageType().Timestamp = &timestamp.Timestamp{Seconds: 3, Nanos: 4}
	ctl.JsonHashStorageType().JsonHashStorageType = object
	ctl.JsonHashStorageType().SomeMessages = append(ctl.JsonHashStorageType().SomeMessages, object)
	ctl.JsonHashStorageType().Timestamps = append(ctl.JsonHashStorageType().Timestamps, object.Timestamp, &timestamp.Timestamp{Seconds: 20, Nanos: 52})

	key := "TestJsonHashStorageTypeRedisController_Load"

	if err = ctl.Store(context.Background(), key); err != nil {
		t.Error(err)
	}

	ctl2 := NewJsonHashStorageTypeRedisController(pool)

	if err = ctl2.Load(context.Background(), key); err != nil {
		t.Error(err)
	}

	if ctl.JsonHashStorageType().String() != ctl2.JsonHashStorageType().String() {
		t.Fatalf("load %v should be %v", ctl2, ctl)
	}
}

func TestJsonHashStorageTypeRedisController_GetSomeString(t *testing.T) {
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

	key := "TestJsonHashStorageTypeRedisController_GetSomeString"
	value := "SomeString"

	ctl := NewJsonHashStorageTypeRedisController(pool)

	if err := ctl.SetSomeString(key, value); err != nil {
		t.Error(err)
	}

	if ctl.JsonHashStorageType().SomeString != value {
		t.Error(key, "value should be", value)
	}

	ctl2 := NewJsonHashStorageTypeRedisController(pool)

	if v, err := ctl2.GetSomeString(key); v != value || err != nil {
		t.Error(key, "value should be", value)
	}
}

func TestJsonHashStorageTypeRedisController_SetSomeString(t *testing.T) {
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

	key := "TestJsonHashStorageTypeRedisController_SetSomeString"
	value := "SomeString"

	ctl := NewJsonHashStorageTypeRedisController(pool)

	if err := ctl.SetSomeString(key, value); err != nil {
		t.Error(err)
	}

	if v := s.HGet(key, "SomeString"); v != value {
		t.Error(key, "value should be", value)
	}
}

func TestJsonHashStorageTypeRedisController_GetSomeBool(t *testing.T) {
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

	key := "TestJsonHashStorageTypeRedisController_GetSomeBool"
	var value bool

	value = true

	ctl := NewJsonHashStorageTypeRedisController(pool)

	if err := ctl.SetSomeBool(key, value); err != nil {
		t.Error(err)
	}

	if ctl.JsonHashStorageType().SomeBool != value {
		t.Error(key, "value should be", value)
	}

	ctl2 := NewJsonHashStorageTypeRedisController(pool)

	if v, err := ctl2.GetSomeBool(key); v != value || err != nil {
		t.Error(key, "value should be", value)
	}

	value = false

	if err := ctl.SetSomeBool(key, value); err != nil {
		t.Error(err)
	}

	if ctl.JsonHashStorageType().SomeBool != value {
		t.Error(key, "value should be", value)
	}

	if v, err := ctl2.GetSomeBool(key); v != value || err != nil {
		t.Error(key, "value should be", value)
	}
}

func TestJsonHashStorageTypeRedisController_SetSomeBool(t *testing.T) {
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

	key := "TestJsonHashStorageTypeRedisController_SetSomeBool"
	var value bool

	value = true

	ctl := NewJsonHashStorageTypeRedisController(pool)

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

func TestJsonHashStorageTypeRedisController_GetSomeInt32(t *testing.T) {
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

	key := "TestJsonHashStorageTypeRedisController_GetSomeInt32"
	value := int32(-125)

	ctl := NewJsonHashStorageTypeRedisController(pool)

	if err := ctl.SetSomeInt32(key, value); err != nil {
		t.Error(err)
	}

	if ctl.JsonHashStorageType().SomeInt32 != value {
		t.Error(key, "value should be", value)
	}

	ctl2 := NewJsonHashStorageTypeRedisController(pool)

	if v, err := ctl2.GetSomeInt32(key); v != value || err != nil {
		t.Error(key, "value should be", value)
	}
}

func TestJsonHashStorageTypeRedisController_SetSomeInt32(t *testing.T) {
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

	key := "TestJsonHashStorageTypeRedisController_SetSomeInt32"
	value := int32(-125)

	ctl := NewJsonHashStorageTypeRedisController(pool)

	if err := ctl.SetSomeInt32(key, value); err != nil {
		t.Error(err)
	}

	if v := s.HGet(key, "SomeInt32"); v != "-125" {
		t.Error(key, "value should be", value)
	}
}

func TestJsonHashStorageTypeRedisController_GetSomeUint32(t *testing.T) {
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

	key := "TestJsonHashStorageTypeRedisController_GetSomeUint32"
	value := uint32(125)

	ctl := NewJsonHashStorageTypeRedisController(pool)

	if err := ctl.SetSomeUint32(key, value); err != nil {
		t.Error(err)
	}

	if ctl.JsonHashStorageType().SomeUint32 != value {
		t.Error(key, "value should be", value)
	}

	ctl2 := NewJsonHashStorageTypeRedisController(pool)

	if v, err := ctl2.GetSomeUint32(key); v != value || err != nil {
		t.Error(key, "value should be", value)
	}
}

func TestJsonHashStorageTypeRedisController_SetSomeUint32(t *testing.T) {
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

	key := "TestJsonHashStorageTypeRedisController_SetSomeUint32"
	value := uint32(125)

	ctl := NewJsonHashStorageTypeRedisController(pool)

	if err := ctl.SetSomeUint32(key, value); err != nil {
		t.Error(err)
	}

	if v := s.HGet(key, "SomeUint32"); v != "125" {
		t.Error(key, "value should be", value)
	}
}

func TestJsonHashStorageTypeRedisController_GetSomeInt64(t *testing.T) {
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

	key := "TestJsonHashStorageTypeRedisController_GetSomeInt64"
	value := int64(-125)

	ctl := NewJsonHashStorageTypeRedisController(pool)

	if err := ctl.SetSomeInt64(key, value); err != nil {
		t.Error(err)
	}

	if ctl.JsonHashStorageType().SomeInt64 != value {
		t.Error(key, "value should be", value)
	}

	ctl2 := NewJsonHashStorageTypeRedisController(pool)

	if v, err := ctl2.GetSomeInt64(key); v != value || err != nil {
		t.Error(key, "value should be", value)
	}
}

func TestJsonHashStorageTypeRedisController_SetSomeInt64(t *testing.T) {
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

	key := "TestJsonHashStorageTypeRedisController_SetSomeInt64"
	value := int64(-125)

	ctl := NewJsonHashStorageTypeRedisController(pool)

	if err := ctl.SetSomeInt64(key, value); err != nil {
		t.Error(err)
	}

	if v := s.HGet(key, "SomeInt64"); v != "-125" {
		t.Error(key, "value should be", value)
	}
}

func TestJsonHashStorageTypeRedisController_GetSomeUint64(t *testing.T) {
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

	key := "TestJsonHashStorageTypeRedisController_GetSomeUint64"
	value := uint64(125)

	ctl := NewJsonHashStorageTypeRedisController(pool)

	if err := ctl.SetSomeUint64(key, value); err != nil {
		t.Error(err)
	}

	if ctl.JsonHashStorageType().SomeUint64 != value {
		t.Error(key, "value should be", value)
	}

	ctl2 := NewJsonHashStorageTypeRedisController(pool)

	if v, err := ctl2.GetSomeUint64(key); v != value || err != nil {
		t.Error(key, "value should be", value)
	}
}

func TestJsonHashStorageTypeRedisController_SetSomeUint64(t *testing.T) {
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

	key := "TestJsonHashStorageTypeRedisController_SetSomeUint64"
	value := uint64(125)

	ctl := NewJsonHashStorageTypeRedisController(pool)

	if err := ctl.SetSomeUint64(key, value); err != nil {
		t.Error(err)
	}

	if v := s.HGet(key, "SomeUint64"); v != "125" {
		t.Error(key, "value should be", value)
	}
}

func TestJsonHashStorageTypeRedisController_GetSomeFloat(t *testing.T) {
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

	key := "TestJsonHashStorageTypeRedisController_GetSomeFloat"
	value := float32(1.25000)

	ctl := NewJsonHashStorageTypeRedisController(pool)

	if err := ctl.SetSomeFloat(key, value); err != nil {
		t.Error(err)
	}

	if ctl.JsonHashStorageType().SomeFloat != value {
		t.Error(key, "value should be", value)
	}

	ctl2 := NewJsonHashStorageTypeRedisController(pool)

	if v, err := ctl2.GetSomeFloat(key); v != value || err != nil {
		t.Error(key, "value should be", value)
	}
}

func TestJsonHashStorageTypeRedisController_SetSomeFloat(t *testing.T) {
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

	key := "TestJsonHashStorageTypeRedisController_SetSomeFloat"
	value := float32(1.25000)

	ctl := NewJsonHashStorageTypeRedisController(pool)

	if err := ctl.SetSomeFloat(key, value); err != nil {
		t.Error(err)
	}

	if v := s.HGet(key, "SomeFloat"); v != "1.25" {
		t.Error(key, "value should be", value)
	}
}

func TestJsonHashStorageTypeRedisController_GetSomeEnum(t *testing.T) {
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

	key := "TestJsonHashStorageTypeRedisController_GetSomeEnum"
	value := JsonHashStorageType_E2

	ctl := NewJsonHashStorageTypeRedisController(pool)

	if err := ctl.SetSomeEnum(key, value); err != nil {
		t.Error(err)
	}

	if ctl.JsonHashStorageType().SomeEnum != value {
		t.Error(key, "value should be", value)
	}

	ctl2 := NewJsonHashStorageTypeRedisController(pool)

	if v, err := ctl2.GetSomeEnum(key); v != value || err != nil {
		t.Error(key, "value should be", value)
	}
}

func TestJsonHashStorageTypeRedisController_SetSomeEnum(t *testing.T) {
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

	key := "TestJsonHashStorageTypeRedisController_SetSomeEnum"
	value := JsonHashStorageType_E2

	ctl := NewJsonHashStorageTypeRedisController(pool)

	if err := ctl.SetSomeEnum(key, value); err != nil {
		t.Error(err)
	}

	if v := s.HGet(key, "SomeEnum"); v != "1" {
		t.Error(key, "value should be", value)
	}
}

func TestJsonHashStorageTypeRedisController_GetTimestamp(t *testing.T) {
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

	key := "TestJsonHashStorageTypeRedisController_GetTimestamp"
	value := &timestamp.Timestamp{Seconds: 1, Nanos: 4}

	ctl := NewJsonHashStorageTypeRedisController(pool)

	if err := ctl.SetTimestamp(key, value); err != nil {
		t.Error(err)
	}

	if ctl.JsonHashStorageType().Timestamp != value {
		t.Error(key, "value should be", value)
	}

	ctl2 := NewJsonHashStorageTypeRedisController(pool)

	if v, err := ctl2.GetTimestamp(key); (v.Seconds != 1 || v.Nanos != 4) || err != nil {
		t.Error(key, "is", v, "value should be", value)
	}
}

func TestJsonHashStorageTypeRedisController_SetTimestamp(t *testing.T) {
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

	key := "TestJsonHashStorageTypeRedisController_SetTimestamp"
	value := &timestamp.Timestamp{Seconds: 1, Nanos: 4}

	ctl := NewJsonHashStorageTypeRedisController(pool)

	if err := ctl.SetTimestamp(key, value); err != nil {
		t.Error(err)
	}

	if v := s.HGet(key, "Timestamp"); v != `{"seconds":1,"nanos":4}` {
		t.Error(key, "is", v, "value should be", value)

	}
}

func TestJsonHashStorageTypeRedisController_GetHashStorageType(t *testing.T) {
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

	key := "TestJsonHashStorageTypeRedisController_GetHashStorageType"
	object := &JsonHashStorageType{}
	object.SomeString = "SomeString"
	object.SomeBool = true
	object.SomeInt32 = -125
	object.SomeUint32 = 255
	object.SomeInt64 = -255
	object.SomeUint64 = 255
	object.SomeFloat = -1.25
	object.SomeEnum = JsonHashStorageType_E3
	object.Timestamp = &timestamp.Timestamp{Seconds: 1, Nanos: 2}
	object.Timestamps = append(object.Timestamps, object.Timestamp, &timestamp.Timestamp{Seconds: 20, Nanos: 52})

	ctl := NewJsonHashStorageTypeRedisController(pool)

	if err := ctl.SetJsonHashStorageTypeField(key, object); err != nil {
		t.Error(err)
	}

	if ctl.JsonHashStorageType().JsonHashStorageType.String() != object.String() {
		t.Error(key, "value should be", object)
	}

	ctl2 := NewJsonHashStorageTypeRedisController(pool)

	if v, err := ctl2.GetJsonHashStorageType(key); err != nil || v.String() != object.String() {
		t.Error(key, "is", v, "value should be", object)
	}
}

func TestJsonHashStorageTypeRedisController_SetHashStorageType(t *testing.T) {
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

	key := "TestJsonHashStorageTypeRedisController_SetHashStorageType"
	object := &JsonHashStorageType{}
	object.SomeString = "SomeString"
	object.SomeBool = true
	object.SomeInt32 = -125
	object.SomeUint32 = 255
	object.SomeInt64 = -255
	object.SomeUint64 = 255
	object.SomeFloat = -1.25
	object.SomeEnum = JsonHashStorageType_E3
	object.Timestamp = &timestamp.Timestamp{Seconds: 1, Nanos: 2}
	object.Timestamps = append(object.Timestamps, object.Timestamp, &timestamp.Timestamp{Seconds: 20, Nanos: 52})

	ctl := NewJsonHashStorageTypeRedisController(pool)

	if err := ctl.SetJsonHashStorageTypeField(key, object); err != nil {
		t.Error(err)
	}

	if v := s.HGet(key, "JsonHashStorageType"); v != `{"some_string":"SomeString","some_bool":true,"some_int32":-125,"some_uint32":255,"some_int64":-255,"some_uint64":255,"some_float":-1.25,"some_enum":2,"timestamp":{"seconds":1,"nanos":2},"timestamps":[{"seconds":1,"nanos":2},{"seconds":20,"nanos":52}]}` {
		t.Error(key, "is", v, "value should be", )
	}
}

func TestJsonHashStorageTypeRedisController_GetSomeMessages(t *testing.T) {
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

	key := "TestJsonHashStorageTypeRedisController_GetSomeMessages"
	object := &JsonHashStorageType{}
	object.SomeString = "SomeString"
	object.SomeBool = true
	object.SomeInt32 = -125
	object.SomeUint32 = 255
	object.SomeInt64 = -255
	object.SomeUint64 = 255
	object.SomeFloat = -1.25
	object.SomeEnum = JsonHashStorageType_E2
	object.Timestamp = &timestamp.Timestamp{Seconds: 1, Nanos: 2}
	object.Timestamps = append(object.Timestamps, object.Timestamp, &timestamp.Timestamp{Seconds: 20, Nanos: 52})

	var someMessages []*JsonHashStorageType
	someMessages = append(someMessages, object)

	ctl := NewJsonHashStorageTypeRedisController(pool)

	if err := ctl.SetSomeMessages(key, someMessages); err != nil {
		t.Error(err)
	}

	d1, _ := json.Marshal(&someMessages)
	d2, _ := json.Marshal(&ctl.JsonHashStorageType().SomeMessages)

	if bytes.Compare(d1, d2) != 0 {
		t.Error(key, "value should be", string(d1))
	}

	ctl2 := NewJsonHashStorageTypeRedisController(pool)

	if v, err := ctl2.GetSomeMessages(key); err != nil {
		t.Error(err)
	} else {
		d3, _ := json.Marshal(v)
		if bytes.Compare(d1, d3) != 0 {
			t.Error(key, "value should be", string(d2))
		}
	}
}

func TestJsonHashStorageTypeRedisController_SetSomeMessages(t *testing.T) {
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

	key := "TestJsonHashStorageTypeRedisController_SetSomeMessages"
	object := &JsonHashStorageType{}
	object.SomeString = "SomeString"
	object.SomeBool = true
	object.SomeInt32 = -125
	object.SomeUint32 = 255
	object.SomeInt64 = -255
	object.SomeUint64 = 255
	object.SomeFloat = -1.25
	object.SomeEnum = JsonHashStorageType_E2
	object.Timestamp = &timestamp.Timestamp{Seconds: 1, Nanos: 2}
	object.Timestamps = append(object.Timestamps, object.Timestamp, &timestamp.Timestamp{Seconds: 20, Nanos: 52})

	var someMessages []*JsonHashStorageType
	someMessages = append(someMessages, object)

	ctl := NewJsonHashStorageTypeRedisController(pool)

	if err := ctl.SetSomeMessages(key, someMessages); err != nil {
		t.Error(err)
	}

	d1, _ := json.Marshal(someMessages)

	if v := s.HGet(key, "SomeMessages"); v != string(d1) {
		t.Error(key, "is", v, "value should be", string(d1))
	}
}

func TestJsonHashStorageTypeRedisController_GetTimestamps(t *testing.T) {
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

	key := "TestJsonHashStorageTypeRedisController_GetTimestamps"

	var timestamps []*timestamp.Timestamp
	timestamps = append(timestamps, &timestamp.Timestamp{Seconds: 1, Nanos: 4}, &timestamp.Timestamp{Seconds: 2, Nanos: 3})

	ctl := NewJsonHashStorageTypeRedisController(pool)

	if err := ctl.SetTimestamps(key, timestamps); err != nil {
		t.Error(err)
	}

	d1, _ := json.Marshal(timestamps)
	d2, _ := json.Marshal(ctl.JsonHashStorageType().Timestamps)

	if bytes.Compare(d1, d2) != 0 {
		t.Error(key, "value should be", string(d1))
	}

	ctl2 := NewJsonHashStorageTypeRedisController(pool)

	if v, err := ctl2.GetTimestamps(key); err != nil {
		t.Error(err)
	} else {
		d3, _ := json.Marshal(v)
		if bytes.Compare(d1, d3) != 0 {
			t.Error(key, "value should be", string(d2))
		}
	}
}

func TestJsonHashStorageTypeRedisController_SetTimestamps(t *testing.T) {
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

	key := "TestJsonHashStorageTypeRedisController_SetTimestamp"
	var timestamps []*timestamp.Timestamp
	timestamps = append(timestamps, &timestamp.Timestamp{Seconds: 1, Nanos: 4}, &timestamp.Timestamp{Seconds: 2, Nanos: 3})

	ctl := NewJsonHashStorageTypeRedisController(pool)

	if err := ctl.SetTimestamps(key, timestamps); err != nil {
		t.Error(err)
	}

	d1, _ := json.Marshal(timestamps)
	d2, _ := json.Marshal(ctl.JsonHashStorageType().Timestamps)

	if bytes.Compare(d1, d2) != 0 {
		t.Error(key, "value should be", string(d1))
	}

	if v := s.HGet(key, "Timestamps"); v != string(d1) {
		t.Error(key, "is", v, "value should be", string(d1))
	}
}
