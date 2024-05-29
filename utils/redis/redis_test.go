package redis

import (
	"github.com/orzlinuxcn/gotool/consts"
	"testing"
	"time"
)

func TestRedis(t *testing.T) {
	client := NewRedis()
	te := map[string]interface{}{
		"a": "1",
		"b": "2",
	}
	client.HMSet("bb", te)
	//client.HS
	//get, err := client.HGetAll("bb").Result()
	//println(err)
	//println(get)

	eval := client.Eval(consts.LuaRedisLimiter, []string{consts.RedisLimiterDefaultName},
		1, time.Now().UnixNano(), int64(time.Second), "wait1")
	println(eval.Val())
	eval = client.Eval(consts.LuaRedisLimiter, []string{consts.RedisLimiterDefaultName},
		1, time.Now().UnixNano(), int64(time.Second), "wait1")
	println(eval.Result())
	eval = client.Eval(consts.LuaRedisLimiter, []string{consts.RedisLimiterDefaultName},
		1, time.Now().UnixNano(), int64(time.Second), "wait1")
	println(eval.Val())
	eval = client.Eval(consts.LuaRedisLimiter, []string{consts.RedisLimiterDefaultName},
		1, time.Now().UnixNano(), int64(time.Second), "wait1")

	i, err := eval.Int64()
	println(err)
	println(i)
}
