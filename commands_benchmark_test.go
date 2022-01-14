/*
 * @Author: feige
 * @Date: 2022/1/14 4:21 PM
 */
package redis

import (
	"testing"
)

var client = NewClient(&Options{
	Addr: "127.0.0.1:6379",
})

func BenchmarkCmdable_Set(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		client.Set("hello", "world", 0)
	}
}

func BenchmarkCmdable_CustomSet(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		cmd := client.SetString("hello", "world")
		status, err := cmd.Result()
		if err != nil {
			panic(err)
		}
		if status != "OK" {
			panic(status)
		}
		cmd.Return()
	}
}

func TestCmdable_Set(t *testing.T) {
	StartPProf()
	for i := 0; i < 10000000; i++ {
		client.SetString("hello", "world")
	}
}
