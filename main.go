package main

import (
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis"
)

var rdb *redis.Client

func init() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     YourRedisIpPort,   // use default Addr
		Password: YourRedisPassword, // no password set
		DB:       0,                 // use default DB
	})

	_, err := rdb.Ping().Result()
	if err != nil {
		log.Fatal("连接redis服务器失败!")
	}
}

func main() {

	key := "hyperLogLog"
	switch key {
	case "string":
		stringDemo()
	case "list":
		listDemo()
	case "hash":
		hashDemo()
	case "set":
		setDemo()
	case "zset":
		zsetDemo()
	case "geo":
		geoDemo()
	case "hyperLogLog":
		hyperLogLogDemo()
	default:

	}
}

func stringDemo() {

	// 清空数据库
	rdb.FlushAll()

	// set
	statusCmd := rdb.Set("strKey", "hello", 0)
	val := statusCmd.Val()          // 命令返回值
	result, _ := statusCmd.Result() // 命令执行结果
	name := statusCmd.Name()        // 命令名称
	str := statusCmd.String()       // 执行的命令以及返回的结果
	fmt.Printf("val: %s \t result: %v \t name: %s \t str: %s\n", val, result, name, str)

	rdb.Set("key1", "v111111111", 0)
	rdb.Set("key2", "v211111", 0)
	rdb.Set("key3", "v3133333", 0)
	rdb.Set("key4", "v4445566", 0)

	// append, 返回追加后字符串的长度
	intCmd := rdb.Append("strKey", " world")
	fmt.Println("append: ", intCmd.Val())

	// BitCount
	bitCount := redis.BitCount{Start: 2, End: 4}
	intCmd = rdb.BitCount("strKey", &bitCount)
	fmt.Println("BitCount: ", intCmd.Val())

	// Get
	stringCmd := rdb.Get("strKey")
	fmt.Println("Get: ", stringCmd.Val())

	// BitOpAnd
	intCmd = rdb.BitOpAnd("key1", "key2", "key3", "key4")
	fmt.Println("BitOpAnd: ", intCmd.Val())

	// BitOpOr
	intCmd = rdb.BitOpOr("key1", "key2", "key3", "key4")
	fmt.Println("BitOpOr: ", intCmd.Val())

	// BitOpXor
	intCmd = rdb.BitOpXor("key1", "key2", "key3", "key4")
	fmt.Println("BitOpXor: ", intCmd.Val())

	// BitOpNot
	intCmd = rdb.BitOpNot("key2", "key2")
	fmt.Println("BitOpNot: ", intCmd.Val())

	// BitPos
	intCmd = rdb.BitPos("strKey", 1, 0, 3)
	fmt.Println("BitPos: ", intCmd.Val())

	// SetBit
	rdb.SetBit("key1", 20, 0)

	rdb.Set("intKey", 1, 0)
	// Incr
	intCmd = rdb.Incr("intKey")
	fmt.Println("Incr: ", intCmd.Val())

	// IncrBy
	intCmd = rdb.IncrBy("intKey", 5)
	fmt.Println("IncrBy: ", intCmd.Val())

	// IncrByFloat
	floatCmd := rdb.IncrByFloat("intKey", 2.2)
	fmt.Println("IncrByFloat: ", floatCmd.Val())

	// MSet 批量设置
	rdb.MSet("k1", "v1", "k2", "v2")

	// MGet 批量获取
	sliceCmd := rdb.MGet("k1", "k2")
	vals := sliceCmd.Val()
	fmt.Print("MGet: \t")
	for _, val := range vals {
		fmt.Print(val, "\t")
	}
	fmt.Println()

	// SetNX
	boolCmd := rdb.SetNX("kx", "vx", 0)
	fmt.Println("SetNX: ", boolCmd.Val())
	boolCmd = rdb.SetNX("k1", "v1", 0)
	fmt.Println("SetNX: ", boolCmd.Val())

	// MSetNX 必须全部都不存在才会返回true
	boolCmd = rdb.MSetNX("k4", "v4", "k3", "v3")
	fmt.Println("MSetNX: ", boolCmd.Val())

	// StrLen
	intCmd = rdb.StrLen("strKey")
	fmt.Println("StrLen: ", intCmd.Val())

	// SetRange
	// "hello world"
	intCmd = rdb.SetRange("strKey", 6, "re")
	fmt.Println("SetRange: ", intCmd.Val(), "\t", rdb.Get("strKey").Val())
}

func listDemo() {

	// 清空数据库
	rdb.FlushAll()

	// LPush
	statusCmd := rdb.LPush("listKey", "v1", "v2", "v3")
	fmt.Println("LPush statusCmd: ", statusCmd)
	fmt.Println("LPush: ", rdb.LRange("listKey", 0, -1).Val())

	// LRange
	stringSliceCmd := rdb.LRange("listKey", 0, -1)
	fmt.Println("LRange: ", stringSliceCmd.Val())

	// RPush
	statusCmd = rdb.RPush("listKey", "v7", "v8", "v9")
	fmt.Println("RPush statusCmd: ", statusCmd)

	fmt.Println("RPush: ", rdb.LRange("listKey", 0, -1).Val())

	// LLen()
	intCmd := rdb.LLen("listKey")
	fmt.Println("LLen(): ", intCmd.Val())

	// BLPop
	rdb.BLPop(1*time.Second, "listKey")
	fmt.Println("BLPop: ", rdb.LRange("listKey", 0, -1).Val())

	// BRPop
	rdb.BRPop(1*time.Second, "listKey")
	fmt.Println("BRPop: ", rdb.LRange("listKey", 0, -1).Val())

	// LIndex
	stringCmd := rdb.LIndex("listKey", 1)
	fmt.Println("LIndex: ", stringCmd.Val())

	// LInsert
	rdb.LInsert("listKey", "Before", "v7", "v6")
	fmt.Println("LInsert Before: ", rdb.LRange("listKey", 0, -1).Val())

	rdb.LInsert("listKey", "After", "v1", "v0")
	fmt.Println("LInsert After: ", rdb.LRange("listKey", 0, -1).Val())

	// goredis 封装了LInsert
	rdb.LInsertAfter("listKey", "v8", "v9")
	rdb.LInsertBefore("listKey", "v2", "v3")
	fmt.Println("goredis 封装了LInsert: ", rdb.LRange("listKey", 0, -1).Val())

	// LPop
	stringCmd = rdb.LPop("listKey")
	fmt.Println("LPop: ", stringCmd.Val())

	// RPop
	stringCmd = rdb.RPop("listKey")
	fmt.Println("LPop: ", stringCmd.Val())

	fmt.Println("listKey: ", rdb.LRange("listKey", 0, -1).Val())

	// RPopLPush
	stringCmd = rdb.RPopLPush("listKey", "destKey")
	fmt.Println(stringCmd.Val())
	fmt.Println("RPopLPush listKey: ", rdb.LRange("listKey", 0, -1).Val())
	fmt.Println("RPopLPush destKey: ", rdb.LRange("destKey", 0, -1).Val())

	// LPushX
	intCmd = rdb.LPushX("listKey", "v3")
	fmt.Println("LPushX exist: ", intCmd.Val())
	intCmd = rdb.LPushX("listKey_no_exist", "v3")
	fmt.Println("LPushX not exist: ", intCmd.Val())

	fmt.Println("listKey", rdb.LRange("listKey", 0, -1).Val())

	// LSet
	rdb.LSet("listKey", 0, "lset")
	fmt.Println("LSet", rdb.LRange("listKey", 0, -1).Val())

	// LRem
	// 	count > 0: 从头往尾移除值为 value 的元素。
	// count < 0: 从尾往头移除值为 value 的元素。
	// count = 0: 移除所有值为 value 的元素
	rdb.LRem("listKey", 0, "lset")
	fmt.Println("LRem: ", rdb.LRange("listKey", 0, -1).Val())

	// LTrim
	rdb.LTrim("listKey", 0, 3)
	fmt.Println("LTrim: ", rdb.LRange("listKey", 0, -1).Val())
}
func hashDemo() {

	// 清空数据库
	rdb.FlushAll()

	// HGetAll
	stringStringMapCmd := rdb.HGetAll("hashKey")
	fmt.Printf("HGetAll: %#v \n", stringStringMapCmd.Val())

	// HSet
	boolCmd := rdb.HSet("hashKey", "hk1", "hv1")
	fmt.Println("BoolCmd: ", boolCmd.Val())
	rdb.HSet("hashKey", "hk2", "hv2")
	rdb.HSet("hashKey", "hk3", "hv3")
	rdb.HSet("hashKey", "hk4", "hv4")
	fmt.Printf("redis状态: %#v \n", rdb.HGetAll("hashKey").Val())

	// HSetNX
	// 存在设置失败
	boolCmd = rdb.HSetNX("hashKey", "hk1", "hv1_new")
	fmt.Println("BoolCmd: ", boolCmd.Val())
	// 不存在设置成功
	boolCmd = rdb.HSetNX("hashKey", "hk1_new", "hv1_new")
	fmt.Println("BoolCmd: ", boolCmd.Val())
	fmt.Printf("redis状态: %#v \n", rdb.HGetAll("hashKey").Val())

	// HGet
	stringCmd := rdb.HGet("hashKey", "hk1")
	fmt.Println("HGet: ", stringCmd.Val())

	// HMSet
	statusCmd := rdb.HMSet("hashKey", map[string]interface{}{"hk8": "hv8", "hk9": "hv8", "hk10": "hv8", "intKey": 25})
	fmt.Println("HMSet: ", statusCmd.Val())
	fmt.Printf("redis状态: %#v \n", rdb.HGetAll("hashKey").Val())

	// HMGet
	sliceCmd := rdb.HMGet("hashKey", "hk1", "hk8", "intKey")
	slice := sliceCmd.Val()
	fmt.Print("HMGet: ")
	for _, e := range slice {
		fmt.Print(e, "\t")
	}
	fmt.Println()

	// HKeys
	stringSliceCmd := rdb.HKeys("hashKey")
	fmt.Print("HKeys: ")
	for _, s := range stringSliceCmd.Val() {
		fmt.Print(s, "\t")
	}
	fmt.Println()

	// HVals
	stringSliceCmd = rdb.HVals("hashKey")
	fmt.Print("HVals: ")
	for _, s := range stringSliceCmd.Val() {
		fmt.Print(s, "\t")
	}
	fmt.Println()

	// HLen
	intCmd := rdb.HLen("hashKey")
	fmt.Println("HLen: ", intCmd.Val())

	// Exists
	boolCmd = rdb.HExists("hashKey", "hk1")
	fmt.Println("HExists: ", boolCmd.Val())

	// HDel
	intCmd = rdb.HDel("hashKey", "hk1", "hk3", "hk5")
	fmt.Println("HDel: ", intCmd.Val())
	fmt.Printf("redis状态: %#v \n", rdb.HGetAll("hashKey").Val())

	// HScan
	// TODO count好像没有生效
	scanCmd := rdb.HScan("hashKey", 0, "hk*", 10)
	keys, cursor := scanCmd.Val()
	fmt.Println("HScan: ", keys, "===", cursor)

	// HIncrBy
	intCmd = rdb.HIncrBy("hashKey", "intKey", 10)
	fmt.Println("HIncrBy:  ", intCmd.Val())
	fmt.Printf("redis状态: %#v \n", rdb.HGetAll("hashKey").Val())

	// HIncrByFloat
	floatCmd := rdb.HIncrByFloat("hashKey", "intKey", 10.2)
	fmt.Println("HIncrByFloat:  ", floatCmd.Val())
	fmt.Printf("redis状态: %#v \n", rdb.HGetAll("hashKey").Val())
}

func setDemo() {
	// 清空数据库
	rdb.FlushAll()

	// SMembers
	stringSliceCmd := rdb.SMembers("setKey")
	fmt.Print("SMembers: ")
	for s := range stringSliceCmd.Val() {
		fmt.Print(s, "\t")
	}
	fmt.Println()

	// SAdd
	intCmd := rdb.SAdd("setKey", "s1", "s2", "s2", "s3", "s4", "s5", "s6")
	fmt.Println("SAdd: ", intCmd.Val())
	fmt.Printf("redis状态: %#v \n", rdb.SMembers("setKey").Val())

	// SCard
	intCmd = rdb.SCard("setKey")
	fmt.Println("SCard: ", intCmd.Val())

	// SIsMember
	boolCmd := rdb.SIsMember("setKey", "s1")
	fmt.Print("SIsMember: ", boolCmd.Val(), "\t")
	boolCmd = rdb.SIsMember("setKey", "s0")
	fmt.Println(boolCmd.Val())

	// SPop
	stringCmd := rdb.SPop("setKey")
	fmt.Println("SPop: ", stringCmd.Val())
	fmt.Printf("redis状态: %#v \n", rdb.SMembers("setKey").Val())

	// go-redis封装的
	stringSliceCmd = rdb.SPopN("setKey", 2)
	fmt.Printf("SPopN: %#v \n", stringSliceCmd.Val())
	fmt.Printf("redis状态: %#v \n", rdb.SMembers("setKey").Val())

	// SRandMember
	stringCmd = rdb.SRandMember("setKey")
	fmt.Println("SRandMember: ", stringCmd.Val())

	// go-redis封装的
	stringSliceCmd = rdb.SRandMemberN("setKey", 3)
	fmt.Printf("SPopN: %#v \n", stringSliceCmd.Val())

	// SRem
	intCmd = rdb.SRem("setKey", "s1", "s4", "s6")
	fmt.Println("SRem: ", intCmd.Val())
	fmt.Printf("redis状态: %#v \n", rdb.SMembers("setKey").Val())

	// SScan
	scanCmd := rdb.SScan("setKey", 0, "s*", 10)
	iter := scanCmd.Iterator()
	fmt.Print("SScan:\t")
	for iter.Next() {
		fmt.Print(iter.Val(), "\t")
	}
	fmt.Println()

	// 为多个set之间的操作创建点数据
	rdb.SAdd("set1", "v1", "v2", "v3", "v4")
	rdb.SAdd("set2", "v2", "v3", "v4", "v5")

	// SDiffStore: 指定key保存结果. 返回结果的元素个数
	intCmd = rdb.SDiffStore("storeDiff", "set1", "set2")
	fmt.Printf("SDiffStore: %d, storeDiff: %#v, set1: %#v, set2: %#v \n", intCmd.Val(), rdb.SMembers("storeDiff").Val(), rdb.SMembers("set1").Val(), rdb.SMembers("set2").Val())

	// SDiff: 返回差后的结果
	stringSliceCmd = rdb.SDiff("set1", "set2")
	fmt.Printf("SDiff: %#v,  set1: %#v,  set2: %#v \n", stringSliceCmd.Val(), rdb.SMembers("set1").Val(), rdb.SMembers("set2").Val())

	// SInterStore
	intCmd = rdb.SInterStore("storeInter", "set1", "set2")
	fmt.Printf("SInterStore: %d, storeInter: %#v, set1: %#v, set2: %#v \n", intCmd.Val(), rdb.SMembers("storeInter").Val(), rdb.SMembers("set1").Val(), rdb.SMembers("set2").Val())

	// SInter
	stringSliceCmd = rdb.SInter("set1", "set2")
	fmt.Printf("SInter: %#v,  set1: %#v,  set2: %#v \n", stringSliceCmd.Val(), rdb.SMembers("set1").Val(), rdb.SMembers("set2").Val())

	// SUnionStore
	intCmd = rdb.SUnionStore("storeUnion", "set1", "set2")
	fmt.Printf("SUnionStore: %d, storeUnion: %#v, set1: %#v, set2: %#v \n", intCmd.Val(), rdb.SMembers("storeInter").Val(), rdb.SMembers("set1").Val(), rdb.SMembers("set2").Val())

	// SUnion
	stringSliceCmd = rdb.SUnion("set1", "set2")
	fmt.Printf("SUnion: %#v,  set1: %#v,  set2: %#v \n", stringSliceCmd.Val(), rdb.SMembers("set1").Val(), rdb.SMembers("set2").Val())

	// SMove
	// SMove 一个不存在于src中的key
	boolCmd = rdb.SMove("set1", "set2", "v0")
	fmt.Printf("SMove 一个不存在于src中的key: %v \n", boolCmd.Val())

	// SMove 一个存在于src中的key
	boolCmd = rdb.SMove("set1", "set2", "v1")
	fmt.Printf("SMove 一个存在于src中的key: %v \n", boolCmd.Val())
	fmt.Printf("set1: %#v,  set2: %#v \n", rdb.SMembers("set1").Val(), rdb.SMembers("set2").Val())

}

func zsetDemo() {

	// 清空数据库
	rdb.FlushAll()

	// ZAdd
	intCmd := rdb.ZAdd("zsetKey", redis.Z{Member: "k0", Score: 1})
	fmt.Printf("%s: %d \n", intCmd.Name(), intCmd.Val())

	rdb.ZAdd("zsetKey", redis.Z{Member: "k1", Score: 1})
	rdb.ZAdd("zsetKey", redis.Z{Member: "k4", Score: 4})
	rdb.ZAdd("zsetKey", redis.Z{Member: "k5", Score: 4})
	rdb.ZAdd("zsetKey", redis.Z{Member: "k2", Score: 2})
	rdb.ZAdd("zsetKey", redis.Z{Member: "ak3", Score: 3})
	rdb.ZAdd("zsetKey", redis.Z{Member: "bk3", Score: 3})
	rdb.ZAdd("zsetKey", redis.Z{Member: "kk3", Score: 3})
	rdb.ZAdd("zsetKey", redis.Z{Member: "ka3", Score: 3})

	// ZIncrBy
	floatCmd := rdb.ZIncrBy("zsetKey", 0.2, "k5")
	fmt.Printf("%s: %f \n", floatCmd.Name(), floatCmd.Val())

	// go-redis封装的
	floatCmd = rdb.ZIncr("zsetKey", redis.Z{Member: "k5", Score: 0.5})
	fmt.Printf("%s: %f \n", floatCmd.Name(), floatCmd.Val())

	floatCmd = rdb.ZScore("zsetKey", "k5")
	fmt.Printf("%s: %f \n", floatCmd.Name(), floatCmd.Val())

	// ZCard
	intCmd = rdb.ZCard("zsetKey")
	fmt.Printf("%s: %d \n", intCmd.Name(), intCmd.Val())

	// ZCount
	intCmd = rdb.ZCount("zsetKey", "(2", "3")
	fmt.Printf("%s: %d \n", intCmd.Name(), intCmd.Val())

	// BZPopMax
	zWithKeyCmd := rdb.BZPopMax(1*time.Second, "zsetKey")
	fmt.Printf("%s: %#v \n", zWithKeyCmd.Name(), zWithKeyCmd.Val())

	// BZPopMin
	zWithKeyCmd = rdb.BZPopMin(1*time.Second, "zsetKey")
	fmt.Printf("%s: %#v \n", zWithKeyCmd.Name(), zWithKeyCmd.Val())

	// 对应的有非阻塞的
	// rdb.ZPopMax("zsetKey")
	// rdb.ZPopMin("zsetKey")

	// ZRange
	stringSliceCmd := rdb.ZRange("zsetKey", 0, -1)
	fmt.Printf("%s: %#v \n", stringSliceCmd.Name(), stringSliceCmd.Val())

	// ZRevRange
	stringSliceCmd = rdb.ZRevRange("zsetKey", 0, -1)
	fmt.Printf("%s: %#v \n", stringSliceCmd.Name(), stringSliceCmd.Val())

	// ZRangeWithScores: withScores这是将score也返回了
	zSliceCmd := rdb.ZRangeWithScores("zsetKey", 0, 1)
	fmt.Printf("%s: %#v \n", zSliceCmd.Name(), zSliceCmd.Val())

	// ZRevRangeWithScores
	zSliceCmd = rdb.ZRevRangeWithScores("zsetKey", 0, 1)
	fmt.Printf("%s: %#v \n", zSliceCmd.Name(), zSliceCmd.Val())

	// ZRangeByScore
	stringSliceCmd = rdb.ZRangeByScore("zsetKey", redis.ZRangeBy{Min: "(2", Max: "(4", Offset: 0, Count: 10})
	fmt.Printf("%s: %#v \n", stringSliceCmd.Name(), stringSliceCmd.Val())

	// ZRevRangeByScore
	stringSliceCmd = rdb.ZRevRangeByScore("zsetKey", redis.ZRangeBy{Min: "(2", Max: "(4", Offset: 0, Count: 10})
	fmt.Printf("%s: %#v \n", stringSliceCmd.Name(), stringSliceCmd.Val())

	// ZRangeByScoreWithScores
	zSliceCmd = rdb.ZRangeByScoreWithScores("zsetKey", redis.ZRangeBy{Min: "(2", Max: "(4", Offset: 0, Count: 10})
	fmt.Printf("%s: %#v \n", zSliceCmd.Name(), zSliceCmd.Val())

	// ZRangeByScoreWithScores
	zSliceCmd = rdb.ZRevRangeByScoreWithScores("zsetKey", redis.ZRangeBy{Min: "(2", Max: "(4", Offset: 0, Count: 10})
	fmt.Printf("%s: %#v \n", zSliceCmd.Name(), zSliceCmd.Val())

	// 添加相同score的zset， 测试lex相关命令
	rdb.ZAdd("lexKey", redis.Z{Member: "ak3", Score: 3})
	rdb.ZAdd("lexKey", redis.Z{Member: "bk3", Score: 3})
	rdb.ZAdd("lexKey", redis.Z{Member: "ka3", Score: 3})
	rdb.ZAdd("lexKey", redis.Z{Member: "kk3", Score: 3})

	// ZRangeByLex 只有当所有score是相同时， 才是确定结果
	stringSliceCmd = rdb.ZRangeByLex("lexKey", redis.ZRangeBy{Min: "(ak3", Max: "(kk3", Offset: 0, Count: 10})
	fmt.Printf("%s: %#v \n", stringSliceCmd.Name(), stringSliceCmd.Val())

	// ZRevRangeByLex
	stringSliceCmd = rdb.ZRevRangeByLex("lexKey", redis.ZRangeBy{Min: "(ak3", Max: "(kk3", Offset: 0, Count: 10})
	fmt.Printf("%s: %#v \n", stringSliceCmd.Name(), stringSliceCmd.Val())

	// ZLexCount
	intCmd = rdb.ZLexCount("lexKey", "(ak3", "(kk3")
	fmt.Printf("%s: %#v \n", intCmd.Name(), intCmd.Val())

	// ZRank: 从0开始的
	intCmd = rdb.ZRank("zsetKey", "k2")
	fmt.Printf("%s: %#v \n", intCmd.Name(), intCmd.Val())

	// ZRevRank
	intCmd = rdb.ZRevRank("zsetKey", "k2")
	fmt.Printf("%s: %#v \n", intCmd.Name(), intCmd.Val())

	// ZRem
	intCmd = rdb.ZRem("zsetKey", "k2")
	fmt.Printf("%s: %#v \n", intCmd.Name(), intCmd.Val())
	fmt.Printf("zsetKey 状态: %#v \n", rdb.ZRange("zsetKey", 0, -1).Val())

	// ZRemRangeByScore
	intCmd = rdb.ZRemRangeByScore("zsetKey", "-inf", "(2")
	fmt.Printf("%s: %#v \n", intCmd.Name(), intCmd.Val())
	fmt.Printf("zsetKey 状态: %#v \n", rdb.ZRange("zsetKey", 0, -1).Val())

	// ZRemRangeByRank: start, stop 都包含
	intCmd = rdb.ZRemRangeByRank("zsetKey", 0, 1)
	fmt.Printf("%s: %#v \n", intCmd.Name(), intCmd.Val())
	fmt.Printf("zsetKey 状态: %#v \n", rdb.ZRange("zsetKey", 0, -1).Val())

	// ZRemRangeByLex
	intCmd = rdb.ZRemRangeByLex("lexKey", "(ak3", "(kk3")
	fmt.Printf("%s: %#v \n", intCmd.Name(), intCmd.Val())
	fmt.Printf("lexKey 状态: %#v \n", rdb.ZRange("lexKey", 0, -1).Val())

	// 创建三个zset用于ZInterStore、ZUnionStore
	rdb.ZAdd("zs1", redis.Z{Score: 1, Member: "k1"},
		redis.Z{Score: 2, Member: "k2"},
		redis.Z{Score: 3, Member: "k3"})

	rdb.ZAdd("zs2", redis.Z{Score: 1, Member: "k1"},
		redis.Z{Score: 2, Member: "k2"},
		redis.Z{Score: 3, Member: "k4"})

	rdb.ZAdd("zs3", redis.Z{Score: 1, Member: "k1"},
		redis.Z{Score: 2, Member: "k4"},
		redis.Z{Score: 3, Member: "k5"})

	rdb.ZInterStore("InterKey", redis.ZStore{Weights: []float64{1.0, 1.0, 1.0}, Aggregate: "SUM"}, "zs1", "zs2", "zs3")
	fmt.Printf("InterKey 状态: %#v \n", rdb.ZRange("InterKey", 0, -1).Val())
	rdb.ZUnionStore("UnionKey", redis.ZStore{Weights: []float64{1.0, 1.0, 1.0}, Aggregate: "SUM"}, "zs1", "zs2", "zs3")
	fmt.Printf("UnionKey 状态: %#v \n", rdb.ZRange("UnionKey", 0, -1).Val())
}

func geoDemo() {
	// 清空数据库
	rdb.FlushAll()

	// GeoAdd
	intCmd := rdb.GeoAdd("geoKey", &redis.GeoLocation{Name: "Palermo", Longitude: 13.361389, Latitude: 38.115556}, &redis.GeoLocation{Name: "Catania", Longitude: 15.087269, Latitude: 37.502669})
	fmt.Printf("%s: %#v \n", intCmd.Name(), intCmd.Val())

	// GeoHash
	stringSliceCmd := rdb.GeoHash("geoKey", "Palermo", "Catania")
	fmt.Printf("%s: %#v \n", stringSliceCmd.Name(), stringSliceCmd.Val())

	// GeoPos: 返回指定位置的经纬度
	geoPosCmd := rdb.GeoPos("geoKey", "Palermo", "Catania", "nonexist")
	fmt.Printf("%s: %#v \n", geoPosCmd.Name(), geoPosCmd.Val())

	for i, geoPos := range geoPosCmd.Val() {
		if geoPos == nil {
			continue
		}
		fmt.Printf("第%d个: (%f, %f) \n", i, geoPos.Latitude, geoPos.Longitude)
	}

	// GeoDist
	floatCmd := rdb.GeoDist("geoKey", "Palermo", "Catania", "km")
	fmt.Printf("%s: %#v \n", floatCmd.Name(), floatCmd.Val())

	// GeoRadius
	geoLocationCmd := rdb.GeoRadius("geoKey", 15, 37, &redis.GeoRadiusQuery{Radius: 200, Unit: "km", WithGeoHash: true, WithCoord: true, WithDist: true})
	fmt.Printf("%s: %#v \n", geoLocationCmd.Name(), geoLocationCmd.Val())

	// GeoRadiusByMember
	geoLocationCmd = rdb.GeoRadiusByMember("geoKey", "Palermo", &redis.GeoRadiusQuery{Radius: 170, Unit: "km", WithGeoHash: true, WithCoord: true, WithDist: true})
	fmt.Printf("%s: %#v \n", geoLocationCmd.Name(), geoLocationCmd.Val())
}

func hyperLogLogDemo() {
	// 清空数据库
	rdb.FlushAll()

	// PFAdd
	intCmd := rdb.PFAdd("hll", "a", "b", "c", "d", "e", "f", "g")
	fmt.Printf("%s: %#v \n", intCmd.Name(), intCmd.Val())

	// PFCount
	intCmd = rdb.PFCount("hll")
	fmt.Printf("%s: %#v \n", intCmd.Name(), intCmd.Val())

	rdb.PFAdd("hll1", "foo", "bar", "zap", "a")
	rdb.PFAdd("hll2", "a", "b", "c", "foo")
	// PFMerge
	statusCmd := rdb.PFMerge("hll3", "hll1", "hll2")
	fmt.Printf("%s: %#v. hll3: %#v \n", statusCmd.Name(), statusCmd.Val(), rdb.PFCount("hll3").Val())
}
