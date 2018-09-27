# protoc-gen-redis

Generate redis load and store function for protobuffer message.
use redis string save the message proto data.

## using the following extensions:

* message options

    - enabled

        enable generate redis load and store function

        default: false

    - ttl

        enable generate  store function with expire ttl

        default: true

## Installing and using

> go get -u github.com/galaxyobe/protoc-gen-redis

> protoc $GOPATH/src -I . --go_out=. --redis_out=. *.proto


## example

See [test.proto](github.com/galaxyobe/protoc-gen-redis/tree/master/test/test.proto)
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

