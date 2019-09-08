package test

import (
	"roundrobin"
	"testing"
)

func BenchmarkNRR_Next(b *testing.B) {
	b.ReportAllocs()

	initBenchmark()

	rr := new(roundrobin.NRR)
	rr.InitRR(benchmarkNodes)

	b.StartTimer()

	for i := 0; i < b.N; i++ {
		rr.Next()
	}
}

func BenchmarkRW_Next(b *testing.B) {
	b.ReportAllocs()

	initBenchmark()

	rr := new(roundrobin.RW)
	rr.InitRR(benchmarkNodes)

	b.StartTimer()

	for i := 0; i < b.N; i++ {
		rr.Next()
	}
}

func BenchmarkWRR_Next(b *testing.B) {
	b.ReportAllocs()

	initBenchmark()

	rr := new(roundrobin.WRR)
	rr.InitRR(benchmarkNodes)

	b.StartTimer()

	for i := 0; i < b.N; i++ {
		rr.Next()
	}
}

func BenchmarkSW_Next(b *testing.B) {
	b.ReportAllocs()

	initBenchmark()

	rr := new(roundrobin.SW)
	rr.InitRR(benchmarkNodes)

	b.StartTimer()

	for i := 0; i < b.N; i++ {
		rr.Next()
	}
}
