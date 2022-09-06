package main

import "testing"

func benchmark(b *testing.B, f func(int, string) string) {
	var str = randomString(10)
	for i := 0; i < b.N; i++ {
		f(10000, str)
	}
}

func BenchmarkPlusConcat(b *testing.B)    { benchmark(b, concatPlus) }
func BenchmarkSprintfConcat(b *testing.B) { benchmark(b, concatSprintf) }
func BenchmarkBuilderConcat(b *testing.B) { benchmark(b, concatBuilder) }
func BenchmarkBufferConcat(b *testing.B)  { benchmark(b, concatBuffer) }
func BenchmarkByteConcat(b *testing.B)    { benchmark(b, concatByte) }
func BenchmarkPreByteConcat(b *testing.B) { benchmark(b, concatPreByte) }
