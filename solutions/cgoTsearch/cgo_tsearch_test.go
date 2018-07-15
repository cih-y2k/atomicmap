
// This file had been automatically generated by utility "git.dx.center/trafficstars/testJob0/internal/benchmarkCodeGen"

package cgoTsearch

import (
	"testing"

	benchmark "git.dx.center/trafficstars/testJob0/internal/benchmarkRoutines"
)

func TestMap(t *testing.T) {
	benchmark.DoTest(t, NewHashMap)
}

func Benchmark_cgoTsearch_Set_intKeyType_blockSize1048576_keyAmount65536(b *testing.B) {
	benchmark.DoBenchmarkOfSet(b, NewHashMap, 1048576, 65536, "int")
}

func Benchmark_cgoTsearch_Set_intKeyType_blockSize1048576_keyAmount1048576(b *testing.B) {
	benchmark.DoBenchmarkOfSet(b, NewHashMap, 1048576, 1048576, "int")
}

func Benchmark_cgoTsearch_ReSet_intKeyType_blockSize1048576_keyAmount65536(b *testing.B) {
	benchmark.DoBenchmarkOfReSet(b, NewHashMap, 1048576, 65536, "int")
}

func Benchmark_cgoTsearch_ReSet_intKeyType_blockSize1048576_keyAmount1048576(b *testing.B) {
	benchmark.DoBenchmarkOfReSet(b, NewHashMap, 1048576, 1048576, "int")
}

func Benchmark_cgoTsearch_Get_intKeyType_blockSize1048576_keyAmount65536(b *testing.B) {
	benchmark.DoBenchmarkOfGet(b, NewHashMap, 1048576, 65536, "int")
}

func Benchmark_cgoTsearch_Get_intKeyType_blockSize1048576_keyAmount1048576(b *testing.B) {
	benchmark.DoBenchmarkOfGet(b, NewHashMap, 1048576, 1048576, "int")
}

func Benchmark_cgoTsearch_GetMiss_intKeyType_blockSize1048576_keyAmount65536(b *testing.B) {
	benchmark.DoBenchmarkOfGetMiss(b, NewHashMap, 1048576, 65536, "int")
}

func Benchmark_cgoTsearch_GetMiss_intKeyType_blockSize1048576_keyAmount1048576(b *testing.B) {
	benchmark.DoBenchmarkOfGetMiss(b, NewHashMap, 1048576, 1048576, "int")
}

func Benchmark_cgoTsearch_Unset_intKeyType_blockSize1048576_keyAmount65536(b *testing.B) {
	benchmark.DoBenchmarkOfUnset(b, NewHashMap, 1048576, 65536, "int")
}

func Benchmark_cgoTsearch_Unset_intKeyType_blockSize1048576_keyAmount1048576(b *testing.B) {
	benchmark.DoBenchmarkOfUnset(b, NewHashMap, 1048576, 1048576, "int")
}

func Benchmark_cgoTsearch_UnsetMiss_intKeyType_blockSize1048576_keyAmount65536(b *testing.B) {
	benchmark.DoBenchmarkOfUnsetMiss(b, NewHashMap, 1048576, 65536, "int")
}

func Benchmark_cgoTsearch_UnsetMiss_intKeyType_blockSize1048576_keyAmount1048576(b *testing.B) {
	benchmark.DoBenchmarkOfUnsetMiss(b, NewHashMap, 1048576, 1048576, "int")
}
