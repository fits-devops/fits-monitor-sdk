package cache

import (
	"context"
	"fmt"
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/frame/g"
	"testing"
)

func TestBatch(t *testing.T) {
	t.Run("testMemory", testMemory)
	t.Run("testRedis", testRedis)
	t.Run("testMemoryWithTag", testMemoryWithTag)
	t.Run("testRedisWithTag", testRedisWithTag)
}

// 测试通过内存缓存普通键值
func testMemory(t *testing.T) {
	c := New("prefix")
	ctx := context.Background()
	//设置key
	c.Set(ctx, "person", g.Map{"name": "zhangsan", "age": 10}, 0)
	//查询key
	v := c.Get(ctx, "person")
	fmt.Println(v)
	//删除key
	c.Remove(ctx, "person")
}

// 测试通过Redis缓存普通键值
func testRedis(t *testing.T) {
	config := gredis.Config{
		Address: "127.0.0.1:6379",
		Db:      2,
	}
	ctx := context.Background()
	gredis.SetConfig(&config)
	c := NewRedis("fits")
	//设置key
	c.Set(ctx, "person", g.Map{"name": "honglm", "age": 22}, 0)
	//查询key
	v := c.Get(ctx, "person")
	fmt.Println(v)
	//删除key
	c.Remove(ctx, "person")
}

// 测试通过内存缓存带有标签的键值
func testMemoryWithTag(t *testing.T) {
	c := New("fits")
	ctx := context.Background()
	// tag can batch Management Cache

	c.Set(ctx, "person01", g.Map{"name": "zhangsan", "age": 10}, 0, "tag_person")
	c.Set(ctx, "family01", g.Map{"address": "Cai Yun street"}, 0, "tag_family")
	c.Set(ctx, "work01", g.Map{"unit": "qixun"}, 0, "tag_work")

	c.Set(ctx, "person02", g.Map{"name": "zhangsan", "age": 10}, 0, "tag_person")
	c.Set(ctx, "family02", g.Map{"address": "Cai Yun street"}, 0, "tag_family")
	c.Set(ctx, "work02", g.Map{"unit": "qixun"}, 0, "tag_work")

	p1 := c.Get(ctx, "person01")
	p2 := c.Get(ctx, "person02")
	fmt.Println(p1, p2)
	// 缓存标签在读取缓存数据时和直接缓存读取一样，差别只在删除时可以批量删除
	// 比如要删除 person01和person02两组对应的缓存
	// 不使用tag时
	c.Remove(ctx, "person01")
	c.Remove(ctx, "person02")
	//或
	c.Removes(ctx, []string{"person01", "person02"})
	// 使用缓存标签
	c.RemoveByTag(ctx, "tag_person") //直接就可以删除该标签下的缓存("person01","person02")
	// 甚至可以批量删除标签
	c.RemoveByTags(ctx, []string{"tag_work", "tag_family"}) // 同时删除多组标签下的数据
}

// 测试通过Redis缓存带有标签的键值
func testRedisWithTag(t *testing.T) {
	config := gredis.Config{
		Address: "127.0.0.1:6379",
		Db:      2,
	}
	ctx := context.Background()
	gredis.SetConfig(&config)
	c := NewRedis("fitsMonitor:")
	// tag can batch Management Cache

	c.Set(ctx, "person01", g.Map{"name": "zhangsan", "age": 10}, 0, "person")
	c.Set(ctx, "family01", g.Map{"address": "Cai Yun street"}, 0, "family")
	c.Set(ctx, "work01", g.Map{"unit": "qixun"}, 0, "work")

	c.Set(ctx, "person02", g.Map{"name": "zhangsan", "age": 10}, 0, "person")
	c.Set(ctx, "family02", g.Map{"address": "Cai Yun street"}, 0, "family")
	c.Set(ctx, "work02", g.Map{"unit": "qixun"}, 0, "work")

	p1 := c.Get(ctx, "person01")
	p2 := c.Get(ctx, "person02")
	fmt.Println(p1, p2)
	// 缓存标签在读取缓存数据时和直接缓存读取一样，差别只在删除时可以批量删除
	// 比如要删除 person01和person02两组对应的缓存
	// 不使用tag删除
	c.Remove(ctx, "person01")
	c.Remove(ctx, "person02")
	//或
	c.Removes(ctx, []string{"person01", "person02"})
	// 使用缓存标签
	c.RemoveByTag(ctx, "tag_person") //直接就可以删除该标签下的缓存("person01","person02")
	// 甚至可以批量删除标签
	c.RemoveByTags(ctx, []string{"tag_work", "tag_family"}) // 同时删除多组标签下的数据
}
