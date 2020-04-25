package sonyflake

import "testing"

func TestPKG_GetID(t *testing.T) {
	value, err := GetID()
	println(value)
	print(err)
}

/*
1 ns = 10 -9 s
// Benchmark 名字 - CPU            循环次数             平均每次执行时间
BenchmarkPKG_GetID-12    	   31036	     38987 ns/op
*/
func BenchmarkPKG_GetID(b *testing.B) {
	for cnt := 0; cnt < b.N; cnt++ {
		GetID()
	}
}

/*
pkg: box2api/pkg/utils/sonyflake
BenchmarkLoopsParallel-12    	   30918	     39063 ns/op
*/
// 测试并发效率
func BenchmarkLoopsPKG_GetID(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			GetID()
		}
	})
}

// 原文链接：https://blog.csdn.net/cchd0001/article/details/48181239
