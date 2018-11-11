// This file had been automatically generated by utility "github.com/xaionaro-go/atomicmap/internal/benchmarkCodeGen"

package atomicmap

import (
	"testing"

	"github.com/xaionaro-go/atomicmap/hash"
	benchmark "github.com/xaionaro-go/atomicmap/internal/benchmarkRoutines"
)

func TestMap(t *testing.T) {
	benchmark.DoTest(t, NewWithArgs, hash.KeyHashFunc)
}

func Benchmark_atomicmap_Set_intKeyType_blockSize16_keyAmount16_falseThreadSafety(b *testing.B) {
	benchmark.DoBenchmarkOfSet(b, NewWithArgs, hash.KeyHashFunc, 16, 16, "int")
}

func Benchmark_atomicmap_Set_stringKeyType_blockSize16_keyAmount16_falseThreadSafety(b *testing.B) {
	benchmark.DoBenchmarkOfSet(b, NewWithArgs, hash.KeyHashFunc, 16, 16, "string")
}

func Benchmark_atomicmap_Set_intKeyType_blockSize16_keyAmount512_falseThreadSafety(b *testing.B) {
	benchmark.DoBenchmarkOfSet(b, NewWithArgs, hash.KeyHashFunc, 16, 512, "int")
}

func Benchmark_atomicmap_Set_stringKeyType_blockSize16_keyAmount512_falseThreadSafety(b *testing.B) {
	benchmark.DoBenchmarkOfSet(b, NewWithArgs, hash.KeyHashFunc, 16, 512, "string")
}

func Benchmark_atomicmap_Set_intKeyType_blockSize16_keyAmount65536_falseThreadSafety(b *testing.B) {
	benchmark.DoBenchmarkOfSet(b, NewWithArgs, hash.KeyHashFunc, 16, 65536, "int")
}

func Benchmark_atomicmap_Set_stringKeyType_blockSize16_keyAmount65536_falseThreadSafety(b *testing.B) {
	benchmark.DoBenchmarkOfSet(b, NewWithArgs, hash.KeyHashFunc, 16, 65536, "string")
}

func Benchmark_atomicmap_Set_intKeyType_blockSize16_keyAmount1048576_falseThreadSafety(b *testing.B) {
	benchmark.DoBenchmarkOfSet(b, NewWithArgs, hash.KeyHashFunc, 16, 1048576, "int")
}

func Benchmark_atomicmap_Set_stringKeyType_blockSize16_keyAmount1048576_falseThreadSafety(b *testing.B) {
	benchmark.DoBenchmarkOfSet(b, NewWithArgs, hash.KeyHashFunc, 16, 1048576, "string")
}

func Benchmark_atomicmap_Set_intKeyType_blockSize64_keyAmount16_falseThreadSafety(b *testing.B) {
	benchmark.DoBenchmarkOfSet(b, NewWithArgs, hash.KeyHashFunc, 64, 16, "int")
}

func Benchmark_atomicmap_Set_stringKeyType_blockSize64_keyAmount16_falseThreadSafety(b *testing.B) {
	benchmark.DoBenchmarkOfSet(b, NewWithArgs, hash.KeyHashFunc, 64, 16, "string")
}

func Benchmark_atomicmap_Set_intKeyType_blockSize64_keyAmount512_falseThreadSafety(b *testing.B) {
	benchmark.DoBenchmarkOfSet(b, NewWithArgs, hash.KeyHashFunc, 64, 512, "int")
}

func Benchmark_atomicmap_Set_stringKeyType_blockSize64_keyAmount512_falseThreadSafety(b *testing.B) {
	benchmark.DoBenchmarkOfSet(b, NewWithArgs, hash.KeyHashFunc, 64, 512, "string")
}

func Benchmark_atomicmap_Set_intKeyType_blockSize64_keyAmount65536_falseThreadSafety(b *testing.B) {
	benchmark.DoBenchmarkOfSet(b, NewWithArgs, hash.KeyHashFunc, 64, 65536, "int")
}

func Benchmark_atomicmap_Set_stringKeyType_blockSize64_keyAmount65536_falseThreadSafety(b *testing.B) {
	benchmark.DoBenchmarkOfSet(b, NewWithArgs, hash.KeyHashFunc, 64, 65536, "string")
}

func Benchmark_atomicmap_Set_intKeyType_blockSize64_keyAmount1048576_falseThreadSafety(b *testing.B) {
	benchmark.DoBenchmarkOfSet(b, NewWithArgs, hash.KeyHashFunc, 64, 1048576, "int")
}

func Benchmark_atomicmap_Set_stringKeyType_blockSize64_keyAmount1048576_falseThreadSafety(b *testing.B) {
	benchmark.DoBenchmarkOfSet(b, NewWithArgs, hash.KeyHashFunc, 64, 1048576, "string")
}

func Benchmark_atomicmap_Set_intKeyType_blockSize128_keyAmount16_falseThreadSafety(b *testing.B) {
	benchmark.DoBenchmarkOfSet(b, NewWithArgs, hash.KeyHashFunc, 128, 16, "int")
}

func Benchmark_atomicmap_Set_stringKeyType_blockSize128_keyAmount16_falseThreadSafety(b *testing.B) {
	benchmark.DoBenchmarkOfSet(b, NewWithArgs, hash.KeyHashFunc, 128, 16, "string")
}

func Benchmark_atomicmap_Set_intKeyType_blockSize128_keyAmount512_falseThreadSafety(b *testing.B) {
	benchmark.DoBenchmarkOfSet(b, NewWithArgs, hash.KeyHashFunc, 128, 512, "int")
}

func Benchmark_atomicmap_Set_stringKeyType_blockSize128_keyAmount512_falseThreadSafety(b *testing.B) {
	benchmark.DoBenchmarkOfSet(b, NewWithArgs, hash.KeyHashFunc, 128, 512, "string")
}

func Benchmark_atomicmap_Set_intKeyType_blockSize128_keyAmount65536_falseThreadSafety(b *testing.B) {
	benchmark.DoBenchmarkOfSet(b, NewWithArgs, hash.KeyHashFunc, 128, 65536, "int")
}

func Benchmark_atomicmap_Set_stringKeyType_blockSize128_keyAmount65536_falseThreadSafety(b *testing.B) {
	benchmark.DoBenchmarkOfSet(b, NewWithArgs, hash.KeyHashFunc, 128, 65536, "string")
}

func Benchmark_atomicmap_Set_intKeyType_blockSize128_keyAmount1048576_falseThreadSafety(b *testing.B) {
	benchmark.DoBenchmarkOfSet(b, NewWithArgs, hash.KeyHashFunc, 128, 1048576, "int")
}

func Benchmark_atomicmap_Set_stringKeyType_blockSize128_keyAmount1048576_falseThreadSafety(b *testing.B) {
	benchmark.DoBenchmarkOfSet(b, NewWithArgs, hash.KeyHashFunc, 128, 1048576, "string")
}

func Benchmark_atomicmap_Set_intKeyType_blockSize1024_keyAmount16_falseThreadSafety(b *testing.B) {
	benchmark.DoBenchmarkOfSet(b, NewWithArgs, hash.KeyHashFunc, 1024, 16, "int")
}

func Benchmark_atomicmap_Set_stringKeyType_blockSize1024_keyAmount16_falseThreadSafety(b *testing.B) {
	benchmark.DoBenchmarkOfSet(b, NewWithArgs, hash.KeyHashFunc, 1024, 16, "string")
}

func Benchmark_atomicmap_Set_intKeyType_blockSize1024_keyAmount512_falseThreadSafety(b *testing.B) {
	benchmark.DoBenchmarkOfSet(b, NewWithArgs, hash.KeyHashFunc, 1024, 512, "int")
}

func Benchmark_atomicmap_Set_stringKeyType_blockSize1024_keyAmount512_falseThreadSafety(b *testing.B) {
	benchmark.DoBenchmarkOfSet(b, NewWithArgs, hash.KeyHashFunc, 1024, 512, "string")
}

func Benchmark_atomicmap_Set_intKeyType_blockSize1024_keyAmount65536_falseThreadSafety(b *testing.B) {
	benchmark.DoBenchmarkOfSet(b, NewWithArgs, hash.KeyHashFunc, 1024, 65536, "int")
}

func Benchmark_atomicmap_Set_stringKeyType_blockSize1024_keyAmount65536_falseThreadSafety(b *testing.B) {
	benchmark.DoBenchmarkOfSet(b, NewWithArgs, hash.KeyHashFunc, 1024, 65536, "string")
}

func Benchmark_atomicmap_Set_intKeyType_blockSize1024_keyAmount1048576_falseThreadSafety(b *testing.B) {
	benchmark.DoBenchmarkOfSet(b, NewWithArgs, hash.KeyHashFunc, 1024, 1048576, "int")
}

func Benchmark_atomicmap_Set_stringKeyType_blockSize1024_keyAmount1048576_falseThreadSafety(b *testing.B) {
	benchmark.DoBenchmarkOfSet(b, NewWithArgs, hash.KeyHashFunc, 1024, 1048576, "string")
}

func Benchmark_atomicmap_Set_intKeyType_blockSize65536_keyAmount512_falseThreadSafety(b *testing.B) {
	benchmark.DoBenchmarkOfSet(b, NewWithArgs, hash.KeyHashFunc, 65536, 512, "int")
}

func Benchmark_atomicmap_Set_stringKeyType_blockSize65536_keyAmount512_falseThreadSafety(b *testing.B) {
	benchmark.DoBenchmarkOfSet(b, NewWithArgs, hash.KeyHashFunc, 65536, 512, "string")
}

func Benchmark_atomicmap_Set_intKeyType_blockSize65536_keyAmount65536_falseThreadSafety(b *testing.B) {
	benchmark.DoBenchmarkOfSet(b, NewWithArgs, hash.KeyHashFunc, 65536, 65536, "int")
}

func Benchmark_atomicmap_Set_stringKeyType_blockSize65536_keyAmount65536_falseThreadSafety(b *testing.B) {
	benchmark.DoBenchmarkOfSet(b, NewWithArgs, hash.KeyHashFunc, 65536, 65536, "string")
}

func Benchmark_atomicmap_Set_intKeyType_blockSize65536_keyAmount1048576_falseThreadSafety(b *testing.B) {
	benchmark.DoBenchmarkOfSet(b, NewWithArgs, hash.KeyHashFunc, 65536, 1048576, "int")
}

func Benchmark_atomicmap_Set_stringKeyType_blockSize65536_keyAmount1048576_falseThreadSafety(b *testing.B) {
	benchmark.DoBenchmarkOfSet(b, NewWithArgs, hash.KeyHashFunc, 65536, 1048576, "string")
}

func Benchmark_atomicmap_Set_intKeyType_blockSize4194304_keyAmount65536_falseThreadSafety(b *testing.B) {
	benchmark.DoBenchmarkOfSet(b, NewWithArgs, hash.KeyHashFunc, 4194304, 65536, "int")
}

func Benchmark_atomicmap_Set_stringKeyType_blockSize4194304_keyAmount65536_falseThreadSafety(b *testing.B) {
	benchmark.DoBenchmarkOfSet(b, NewWithArgs, hash.KeyHashFunc, 4194304, 65536, "string")
}

func Benchmark_atomicmap_Set_intKeyType_blockSize4194304_keyAmount1048576_falseThreadSafety(b *testing.B) {
	benchmark.DoBenchmarkOfSet(b, NewWithArgs, hash.KeyHashFunc, 4194304, 1048576, "int")
}

func Benchmark_atomicmap_Set_stringKeyType_blockSize4194304_keyAmount1048576_falseThreadSafety(b *testing.B) {
	benchmark.DoBenchmarkOfSet(b, NewWithArgs, hash.KeyHashFunc, 4194304, 1048576, "string")
}

func Benchmark_atomicmap_Set_intKeyType_blockSize16777216_keyAmount65536_falseThreadSafety(b *testing.B) {
	benchmark.DoBenchmarkOfSet(b, NewWithArgs, hash.KeyHashFunc, 16777216, 65536, "int")
}

func Benchmark_atomicmap_Set_stringKeyType_blockSize16777216_keyAmount65536_falseThreadSafety(b *testing.B) {
	benchmark.DoBenchmarkOfSet(b, NewWithArgs, hash.KeyHashFunc, 16777216, 65536, "string")
}

func Benchmark_atomicmap_Set_intKeyType_blockSize16777216_keyAmount1048576_falseThreadSafety(b *testing.B) {
	benchmark.DoBenchmarkOfSet(b, NewWithArgs, hash.KeyHashFunc, 16777216, 1048576, "int")
}

func Benchmark_atomicmap_Set_stringKeyType_blockSize16777216_keyAmount1048576_falseThreadSafety(b *testing.B) {
	benchmark.DoBenchmarkOfSet(b, NewWithArgs, hash.KeyHashFunc, 16777216, 1048576, "string")
}

func Benchmark_atomicmap_Get_intKeyType_blockSize16_keyAmount16_falseThreadSafety(b *testing.B) {
	benchmark.DoBenchmarkOfGet(b, NewWithArgs, hash.KeyHashFunc, 16, 16, "int")
}

func Benchmark_atomicmap_Get_stringKeyType_blockSize16_keyAmount16_falseThreadSafety(b *testing.B) {
	benchmark.DoBenchmarkOfGet(b, NewWithArgs, hash.KeyHashFunc, 16, 16, "string")
}

func Benchmark_atomicmap_Get_intKeyType_blockSize16_keyAmount512_falseThreadSafety(b *testing.B) {
	benchmark.DoBenchmarkOfGet(b, NewWithArgs, hash.KeyHashFunc, 16, 512, "int")
}

func Benchmark_atomicmap_Get_stringKeyType_blockSize16_keyAmount512_falseThreadSafety(b *testing.B) {
	benchmark.DoBenchmarkOfGet(b, NewWithArgs, hash.KeyHashFunc, 16, 512, "string")
}

func Benchmark_atomicmap_Get_intKeyType_blockSize16_keyAmount65536_falseThreadSafety(b *testing.B) {
	benchmark.DoBenchmarkOfGet(b, NewWithArgs, hash.KeyHashFunc, 16, 65536, "int")
}

func Benchmark_atomicmap_Get_stringKeyType_blockSize16_keyAmount65536_falseThreadSafety(b *testing.B) {
	benchmark.DoBenchmarkOfGet(b, NewWithArgs, hash.KeyHashFunc, 16, 65536, "string")
}

func Benchmark_atomicmap_Get_intKeyType_blockSize16_keyAmount1048576_falseThreadSafety(b *testing.B) {
	benchmark.DoBenchmarkOfGet(b, NewWithArgs, hash.KeyHashFunc, 16, 1048576, "int")
}

func Benchmark_atomicmap_Get_stringKeyType_blockSize16_keyAmount1048576_falseThreadSafety(b *testing.B) {
	benchmark.DoBenchmarkOfGet(b, NewWithArgs, hash.KeyHashFunc, 16, 1048576, "string")
}

func Benchmark_atomicmap_Get_intKeyType_blockSize64_keyAmount16_falseThreadSafety(b *testing.B) {
	benchmark.DoBenchmarkOfGet(b, NewWithArgs, hash.KeyHashFunc, 64, 16, "int")
}

func Benchmark_atomicmap_Get_stringKeyType_blockSize64_keyAmount16_falseThreadSafety(b *testing.B) {
	benchmark.DoBenchmarkOfGet(b, NewWithArgs, hash.KeyHashFunc, 64, 16, "string")
}

func Benchmark_atomicmap_Get_intKeyType_blockSize64_keyAmount512_falseThreadSafety(b *testing.B) {
	benchmark.DoBenchmarkOfGet(b, NewWithArgs, hash.KeyHashFunc, 64, 512, "int")
}

func Benchmark_atomicmap_Get_stringKeyType_blockSize64_keyAmount512_falseThreadSafety(b *testing.B) {
	benchmark.DoBenchmarkOfGet(b, NewWithArgs, hash.KeyHashFunc, 64, 512, "string")
}

func Benchmark_atomicmap_Get_intKeyType_blockSize64_keyAmount65536_falseThreadSafety(b *testing.B) {
	benchmark.DoBenchmarkOfGet(b, NewWithArgs, hash.KeyHashFunc, 64, 65536, "int")
}

func Benchmark_atomicmap_Get_stringKeyType_blockSize64_keyAmount65536_falseThreadSafety(b *testing.B) {
	benchmark.DoBenchmarkOfGet(b, NewWithArgs, hash.KeyHashFunc, 64, 65536, "string")
}

func Benchmark_atomicmap_Get_intKeyType_blockSize64_keyAmount1048576_falseThreadSafety(b *testing.B) {
	benchmark.DoBenchmarkOfGet(b, NewWithArgs, hash.KeyHashFunc, 64, 1048576, "int")
}

func Benchmark_atomicmap_Get_stringKeyType_blockSize64_keyAmount1048576_falseThreadSafety(b *testing.B) {
	benchmark.DoBenchmarkOfGet(b, NewWithArgs, hash.KeyHashFunc, 64, 1048576, "string")
}

func Benchmark_atomicmap_Get_intKeyType_blockSize128_keyAmount16_falseThreadSafety(b *testing.B) {
	benchmark.DoBenchmarkOfGet(b, NewWithArgs, hash.KeyHashFunc, 128, 16, "int")
}

func Benchmark_atomicmap_Get_stringKeyType_blockSize128_keyAmount16_falseThreadSafety(b *testing.B) {
	benchmark.DoBenchmarkOfGet(b, NewWithArgs, hash.KeyHashFunc, 128, 16, "string")
}

func Benchmark_atomicmap_Get_intKeyType_blockSize128_keyAmount512_falseThreadSafety(b *testing.B) {
	benchmark.DoBenchmarkOfGet(b, NewWithArgs, hash.KeyHashFunc, 128, 512, "int")
}

func Benchmark_atomicmap_Get_stringKeyType_blockSize128_keyAmount512_falseThreadSafety(b *testing.B) {
	benchmark.DoBenchmarkOfGet(b, NewWithArgs, hash.KeyHashFunc, 128, 512, "string")
}

func Benchmark_atomicmap_Get_intKeyType_blockSize128_keyAmount65536_falseThreadSafety(b *testing.B) {
	benchmark.DoBenchmarkOfGet(b, NewWithArgs, hash.KeyHashFunc, 128, 65536, "int")
}

func Benchmark_atomicmap_Get_stringKeyType_blockSize128_keyAmount65536_falseThreadSafety(b *testing.B) {
	benchmark.DoBenchmarkOfGet(b, NewWithArgs, hash.KeyHashFunc, 128, 65536, "string")
}

func Benchmark_atomicmap_Get_intKeyType_blockSize128_keyAmount1048576_falseThreadSafety(b *testing.B) {
	benchmark.DoBenchmarkOfGet(b, NewWithArgs, hash.KeyHashFunc, 128, 1048576, "int")
}

func Benchmark_atomicmap_Get_stringKeyType_blockSize128_keyAmount1048576_falseThreadSafety(b *testing.B) {
	benchmark.DoBenchmarkOfGet(b, NewWithArgs, hash.KeyHashFunc, 128, 1048576, "string")
}

func Benchmark_atomicmap_Get_intKeyType_blockSize1024_keyAmount16_falseThreadSafety(b *testing.B) {
	benchmark.DoBenchmarkOfGet(b, NewWithArgs, hash.KeyHashFunc, 1024, 16, "int")
}

func Benchmark_atomicmap_Get_stringKeyType_blockSize1024_keyAmount16_falseThreadSafety(b *testing.B) {
	benchmark.DoBenchmarkOfGet(b, NewWithArgs, hash.KeyHashFunc, 1024, 16, "string")
}

func Benchmark_atomicmap_Get_intKeyType_blockSize1024_keyAmount512_falseThreadSafety(b *testing.B) {
	benchmark.DoBenchmarkOfGet(b, NewWithArgs, hash.KeyHashFunc, 1024, 512, "int")
}

func Benchmark_atomicmap_Get_stringKeyType_blockSize1024_keyAmount512_falseThreadSafety(b *testing.B) {
	benchmark.DoBenchmarkOfGet(b, NewWithArgs, hash.KeyHashFunc, 1024, 512, "string")
}

func Benchmark_atomicmap_Get_intKeyType_blockSize1024_keyAmount65536_falseThreadSafety(b *testing.B) {
	benchmark.DoBenchmarkOfGet(b, NewWithArgs, hash.KeyHashFunc, 1024, 65536, "int")
}

func Benchmark_atomicmap_Get_stringKeyType_blockSize1024_keyAmount65536_falseThreadSafety(b *testing.B) {
	benchmark.DoBenchmarkOfGet(b, NewWithArgs, hash.KeyHashFunc, 1024, 65536, "string")
}

func Benchmark_atomicmap_Get_intKeyType_blockSize1024_keyAmount1048576_falseThreadSafety(b *testing.B) {
	benchmark.DoBenchmarkOfGet(b, NewWithArgs, hash.KeyHashFunc, 1024, 1048576, "int")
}

func Benchmark_atomicmap_Get_stringKeyType_blockSize1024_keyAmount1048576_falseThreadSafety(b *testing.B) {
	benchmark.DoBenchmarkOfGet(b, NewWithArgs, hash.KeyHashFunc, 1024, 1048576, "string")
}

func Benchmark_atomicmap_Get_intKeyType_blockSize65536_keyAmount512_falseThreadSafety(b *testing.B) {
	benchmark.DoBenchmarkOfGet(b, NewWithArgs, hash.KeyHashFunc, 65536, 512, "int")
}

func Benchmark_atomicmap_Get_stringKeyType_blockSize65536_keyAmount512_falseThreadSafety(b *testing.B) {
	benchmark.DoBenchmarkOfGet(b, NewWithArgs, hash.KeyHashFunc, 65536, 512, "string")
}

func Benchmark_atomicmap_Get_intKeyType_blockSize65536_keyAmount65536_falseThreadSafety(b *testing.B) {
	benchmark.DoBenchmarkOfGet(b, NewWithArgs, hash.KeyHashFunc, 65536, 65536, "int")
}

func Benchmark_atomicmap_Get_stringKeyType_blockSize65536_keyAmount65536_falseThreadSafety(b *testing.B) {
	benchmark.DoBenchmarkOfGet(b, NewWithArgs, hash.KeyHashFunc, 65536, 65536, "string")
}

func Benchmark_atomicmap_Get_intKeyType_blockSize65536_keyAmount1048576_falseThreadSafety(b *testing.B) {
	benchmark.DoBenchmarkOfGet(b, NewWithArgs, hash.KeyHashFunc, 65536, 1048576, "int")
}

func Benchmark_atomicmap_Get_stringKeyType_blockSize65536_keyAmount1048576_falseThreadSafety(b *testing.B) {
	benchmark.DoBenchmarkOfGet(b, NewWithArgs, hash.KeyHashFunc, 65536, 1048576, "string")
}

func Benchmark_atomicmap_Get_intKeyType_blockSize4194304_keyAmount65536_falseThreadSafety(b *testing.B) {
	benchmark.DoBenchmarkOfGet(b, NewWithArgs, hash.KeyHashFunc, 4194304, 65536, "int")
}

func Benchmark_atomicmap_Get_stringKeyType_blockSize4194304_keyAmount65536_falseThreadSafety(b *testing.B) {
	benchmark.DoBenchmarkOfGet(b, NewWithArgs, hash.KeyHashFunc, 4194304, 65536, "string")
}

func Benchmark_atomicmap_Get_intKeyType_blockSize4194304_keyAmount1048576_falseThreadSafety(b *testing.B) {
	benchmark.DoBenchmarkOfGet(b, NewWithArgs, hash.KeyHashFunc, 4194304, 1048576, "int")
}

func Benchmark_atomicmap_Get_stringKeyType_blockSize4194304_keyAmount1048576_falseThreadSafety(b *testing.B) {
	benchmark.DoBenchmarkOfGet(b, NewWithArgs, hash.KeyHashFunc, 4194304, 1048576, "string")
}

func Benchmark_atomicmap_Get_intKeyType_blockSize16777216_keyAmount65536_falseThreadSafety(b *testing.B) {
	benchmark.DoBenchmarkOfGet(b, NewWithArgs, hash.KeyHashFunc, 16777216, 65536, "int")
}

func Benchmark_atomicmap_Get_stringKeyType_blockSize16777216_keyAmount65536_falseThreadSafety(b *testing.B) {
	benchmark.DoBenchmarkOfGet(b, NewWithArgs, hash.KeyHashFunc, 16777216, 65536, "string")
}

func Benchmark_atomicmap_Get_intKeyType_blockSize16777216_keyAmount1048576_falseThreadSafety(b *testing.B) {
	benchmark.DoBenchmarkOfGet(b, NewWithArgs, hash.KeyHashFunc, 16777216, 1048576, "int")
}

func Benchmark_atomicmap_Get_stringKeyType_blockSize16777216_keyAmount1048576_falseThreadSafety(b *testing.B) {
	benchmark.DoBenchmarkOfGet(b, NewWithArgs, hash.KeyHashFunc, 16777216, 1048576, "string")
}

func Benchmark_atomicmap_Unset_intKeyType_blockSize16_keyAmount16_falseThreadSafety(b *testing.B) {
	benchmark.DoBenchmarkOfUnset(b, NewWithArgs, hash.KeyHashFunc, 16, 16, "int")
}

func Benchmark_atomicmap_Unset_stringKeyType_blockSize16_keyAmount16_falseThreadSafety(b *testing.B) {
	benchmark.DoBenchmarkOfUnset(b, NewWithArgs, hash.KeyHashFunc, 16, 16, "string")
}

func Benchmark_atomicmap_Unset_intKeyType_blockSize16_keyAmount512_falseThreadSafety(b *testing.B) {
	benchmark.DoBenchmarkOfUnset(b, NewWithArgs, hash.KeyHashFunc, 16, 512, "int")
}

func Benchmark_atomicmap_Unset_stringKeyType_blockSize16_keyAmount512_falseThreadSafety(b *testing.B) {
	benchmark.DoBenchmarkOfUnset(b, NewWithArgs, hash.KeyHashFunc, 16, 512, "string")
}

func Benchmark_atomicmap_Unset_intKeyType_blockSize16_keyAmount65536_falseThreadSafety(b *testing.B) {
	benchmark.DoBenchmarkOfUnset(b, NewWithArgs, hash.KeyHashFunc, 16, 65536, "int")
}

func Benchmark_atomicmap_Unset_stringKeyType_blockSize16_keyAmount65536_falseThreadSafety(b *testing.B) {
	benchmark.DoBenchmarkOfUnset(b, NewWithArgs, hash.KeyHashFunc, 16, 65536, "string")
}

func Benchmark_atomicmap_Unset_intKeyType_blockSize16_keyAmount1048576_falseThreadSafety(b *testing.B) {
	benchmark.DoBenchmarkOfUnset(b, NewWithArgs, hash.KeyHashFunc, 16, 1048576, "int")
}

func Benchmark_atomicmap_Unset_stringKeyType_blockSize16_keyAmount1048576_falseThreadSafety(b *testing.B) {
	benchmark.DoBenchmarkOfUnset(b, NewWithArgs, hash.KeyHashFunc, 16, 1048576, "string")
}

func Benchmark_atomicmap_Unset_intKeyType_blockSize64_keyAmount16_falseThreadSafety(b *testing.B) {
	benchmark.DoBenchmarkOfUnset(b, NewWithArgs, hash.KeyHashFunc, 64, 16, "int")
}

func Benchmark_atomicmap_Unset_stringKeyType_blockSize64_keyAmount16_falseThreadSafety(b *testing.B) {
	benchmark.DoBenchmarkOfUnset(b, NewWithArgs, hash.KeyHashFunc, 64, 16, "string")
}

func Benchmark_atomicmap_Unset_intKeyType_blockSize64_keyAmount512_falseThreadSafety(b *testing.B) {
	benchmark.DoBenchmarkOfUnset(b, NewWithArgs, hash.KeyHashFunc, 64, 512, "int")
}

func Benchmark_atomicmap_Unset_stringKeyType_blockSize64_keyAmount512_falseThreadSafety(b *testing.B) {
	benchmark.DoBenchmarkOfUnset(b, NewWithArgs, hash.KeyHashFunc, 64, 512, "string")
}

func Benchmark_atomicmap_Unset_intKeyType_blockSize64_keyAmount65536_falseThreadSafety(b *testing.B) {
	benchmark.DoBenchmarkOfUnset(b, NewWithArgs, hash.KeyHashFunc, 64, 65536, "int")
}

func Benchmark_atomicmap_Unset_stringKeyType_blockSize64_keyAmount65536_falseThreadSafety(b *testing.B) {
	benchmark.DoBenchmarkOfUnset(b, NewWithArgs, hash.KeyHashFunc, 64, 65536, "string")
}

func Benchmark_atomicmap_Unset_intKeyType_blockSize64_keyAmount1048576_falseThreadSafety(b *testing.B) {
	benchmark.DoBenchmarkOfUnset(b, NewWithArgs, hash.KeyHashFunc, 64, 1048576, "int")
}

func Benchmark_atomicmap_Unset_stringKeyType_blockSize64_keyAmount1048576_falseThreadSafety(b *testing.B) {
	benchmark.DoBenchmarkOfUnset(b, NewWithArgs, hash.KeyHashFunc, 64, 1048576, "string")
}

func Benchmark_atomicmap_Unset_intKeyType_blockSize128_keyAmount16_falseThreadSafety(b *testing.B) {
	benchmark.DoBenchmarkOfUnset(b, NewWithArgs, hash.KeyHashFunc, 128, 16, "int")
}

func Benchmark_atomicmap_Unset_stringKeyType_blockSize128_keyAmount16_falseThreadSafety(b *testing.B) {
	benchmark.DoBenchmarkOfUnset(b, NewWithArgs, hash.KeyHashFunc, 128, 16, "string")
}

func Benchmark_atomicmap_Unset_intKeyType_blockSize128_keyAmount512_falseThreadSafety(b *testing.B) {
	benchmark.DoBenchmarkOfUnset(b, NewWithArgs, hash.KeyHashFunc, 128, 512, "int")
}

func Benchmark_atomicmap_Unset_stringKeyType_blockSize128_keyAmount512_falseThreadSafety(b *testing.B) {
	benchmark.DoBenchmarkOfUnset(b, NewWithArgs, hash.KeyHashFunc, 128, 512, "string")
}

func Benchmark_atomicmap_Unset_intKeyType_blockSize128_keyAmount65536_falseThreadSafety(b *testing.B) {
	benchmark.DoBenchmarkOfUnset(b, NewWithArgs, hash.KeyHashFunc, 128, 65536, "int")
}

func Benchmark_atomicmap_Unset_stringKeyType_blockSize128_keyAmount65536_falseThreadSafety(b *testing.B) {
	benchmark.DoBenchmarkOfUnset(b, NewWithArgs, hash.KeyHashFunc, 128, 65536, "string")
}

func Benchmark_atomicmap_Unset_intKeyType_blockSize128_keyAmount1048576_falseThreadSafety(b *testing.B) {
	benchmark.DoBenchmarkOfUnset(b, NewWithArgs, hash.KeyHashFunc, 128, 1048576, "int")
}

func Benchmark_atomicmap_Unset_stringKeyType_blockSize128_keyAmount1048576_falseThreadSafety(b *testing.B) {
	benchmark.DoBenchmarkOfUnset(b, NewWithArgs, hash.KeyHashFunc, 128, 1048576, "string")
}

func Benchmark_atomicmap_Unset_intKeyType_blockSize1024_keyAmount16_falseThreadSafety(b *testing.B) {
	benchmark.DoBenchmarkOfUnset(b, NewWithArgs, hash.KeyHashFunc, 1024, 16, "int")
}

func Benchmark_atomicmap_Unset_stringKeyType_blockSize1024_keyAmount16_falseThreadSafety(b *testing.B) {
	benchmark.DoBenchmarkOfUnset(b, NewWithArgs, hash.KeyHashFunc, 1024, 16, "string")
}

func Benchmark_atomicmap_Unset_intKeyType_blockSize1024_keyAmount512_falseThreadSafety(b *testing.B) {
	benchmark.DoBenchmarkOfUnset(b, NewWithArgs, hash.KeyHashFunc, 1024, 512, "int")
}

func Benchmark_atomicmap_Unset_stringKeyType_blockSize1024_keyAmount512_falseThreadSafety(b *testing.B) {
	benchmark.DoBenchmarkOfUnset(b, NewWithArgs, hash.KeyHashFunc, 1024, 512, "string")
}

func Benchmark_atomicmap_Unset_intKeyType_blockSize1024_keyAmount65536_falseThreadSafety(b *testing.B) {
	benchmark.DoBenchmarkOfUnset(b, NewWithArgs, hash.KeyHashFunc, 1024, 65536, "int")
}

func Benchmark_atomicmap_Unset_stringKeyType_blockSize1024_keyAmount65536_falseThreadSafety(b *testing.B) {
	benchmark.DoBenchmarkOfUnset(b, NewWithArgs, hash.KeyHashFunc, 1024, 65536, "string")
}

func Benchmark_atomicmap_Unset_intKeyType_blockSize1024_keyAmount1048576_falseThreadSafety(b *testing.B) {
	benchmark.DoBenchmarkOfUnset(b, NewWithArgs, hash.KeyHashFunc, 1024, 1048576, "int")
}

func Benchmark_atomicmap_Unset_stringKeyType_blockSize1024_keyAmount1048576_falseThreadSafety(b *testing.B) {
	benchmark.DoBenchmarkOfUnset(b, NewWithArgs, hash.KeyHashFunc, 1024, 1048576, "string")
}

func Benchmark_atomicmap_Unset_intKeyType_blockSize65536_keyAmount512_falseThreadSafety(b *testing.B) {
	benchmark.DoBenchmarkOfUnset(b, NewWithArgs, hash.KeyHashFunc, 65536, 512, "int")
}

func Benchmark_atomicmap_Unset_stringKeyType_blockSize65536_keyAmount512_falseThreadSafety(b *testing.B) {
	benchmark.DoBenchmarkOfUnset(b, NewWithArgs, hash.KeyHashFunc, 65536, 512, "string")
}

func Benchmark_atomicmap_Unset_intKeyType_blockSize65536_keyAmount65536_falseThreadSafety(b *testing.B) {
	benchmark.DoBenchmarkOfUnset(b, NewWithArgs, hash.KeyHashFunc, 65536, 65536, "int")
}

func Benchmark_atomicmap_Unset_stringKeyType_blockSize65536_keyAmount65536_falseThreadSafety(b *testing.B) {
	benchmark.DoBenchmarkOfUnset(b, NewWithArgs, hash.KeyHashFunc, 65536, 65536, "string")
}

func Benchmark_atomicmap_Unset_intKeyType_blockSize65536_keyAmount1048576_falseThreadSafety(b *testing.B) {
	benchmark.DoBenchmarkOfUnset(b, NewWithArgs, hash.KeyHashFunc, 65536, 1048576, "int")
}

func Benchmark_atomicmap_Unset_stringKeyType_blockSize65536_keyAmount1048576_falseThreadSafety(b *testing.B) {
	benchmark.DoBenchmarkOfUnset(b, NewWithArgs, hash.KeyHashFunc, 65536, 1048576, "string")
}

func Benchmark_atomicmap_Unset_intKeyType_blockSize4194304_keyAmount65536_falseThreadSafety(b *testing.B) {
	benchmark.DoBenchmarkOfUnset(b, NewWithArgs, hash.KeyHashFunc, 4194304, 65536, "int")
}

func Benchmark_atomicmap_Unset_stringKeyType_blockSize4194304_keyAmount65536_falseThreadSafety(b *testing.B) {
	benchmark.DoBenchmarkOfUnset(b, NewWithArgs, hash.KeyHashFunc, 4194304, 65536, "string")
}

func Benchmark_atomicmap_Unset_intKeyType_blockSize4194304_keyAmount1048576_falseThreadSafety(b *testing.B) {
	benchmark.DoBenchmarkOfUnset(b, NewWithArgs, hash.KeyHashFunc, 4194304, 1048576, "int")
}

func Benchmark_atomicmap_Unset_stringKeyType_blockSize4194304_keyAmount1048576_falseThreadSafety(b *testing.B) {
	benchmark.DoBenchmarkOfUnset(b, NewWithArgs, hash.KeyHashFunc, 4194304, 1048576, "string")
}

func Benchmark_atomicmap_Unset_intKeyType_blockSize16777216_keyAmount65536_falseThreadSafety(b *testing.B) {
	benchmark.DoBenchmarkOfUnset(b, NewWithArgs, hash.KeyHashFunc, 16777216, 65536, "int")
}

func Benchmark_atomicmap_Unset_stringKeyType_blockSize16777216_keyAmount65536_falseThreadSafety(b *testing.B) {
	benchmark.DoBenchmarkOfUnset(b, NewWithArgs, hash.KeyHashFunc, 16777216, 65536, "string")
}

func Benchmark_atomicmap_Unset_intKeyType_blockSize16777216_keyAmount1048576_falseThreadSafety(b *testing.B) {
	benchmark.DoBenchmarkOfUnset(b, NewWithArgs, hash.KeyHashFunc, 16777216, 1048576, "int")
}

func Benchmark_atomicmap_Unset_stringKeyType_blockSize16777216_keyAmount1048576_falseThreadSafety(b *testing.B) {
	benchmark.DoBenchmarkOfUnset(b, NewWithArgs, hash.KeyHashFunc, 16777216, 1048576, "string")
}
