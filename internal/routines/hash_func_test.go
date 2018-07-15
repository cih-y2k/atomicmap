package routines

import (
	"testing"

	benchmark "git.dx.center/trafficstars/testJob0/internal/benchmarkRoutines"
)

func TestHashCollisions_blockSize16_keyAmount16(t *testing.T) {
	benchmark.DoTestHashCollisions(t, HashFunc, 16, 16)
}
func TestHashCollisions_blockSize64_keyAmount16(t *testing.T) {
	benchmark.DoTestHashCollisions(t, HashFunc, 64, 16)
}
func TestHashCollisions_blockSize128_keyAmount16(t *testing.T) {
	benchmark.DoTestHashCollisions(t, HashFunc, 128, 16)
}
func TestHashCollisions_blockSize1024_keyAmount16(t *testing.T) {
	benchmark.DoTestHashCollisions(t, HashFunc, 1024, 16)
}

func TestHashCollisions_blockSize16_keyAmount64(t *testing.T) {
	benchmark.DoTestHashCollisions(t, HashFunc, 16, 64)
}
func TestHashCollisions_blockSize64_keyAmount64(t *testing.T) {
	benchmark.DoTestHashCollisions(t, HashFunc, 64, 64)
}
func TestHashCollisions_blockSize128_keyAmount64(t *testing.T) {
	benchmark.DoTestHashCollisions(t, HashFunc, 128, 64)
}
func TestHashCollisions_blockSize1024_keyAmount64(t *testing.T) {
	benchmark.DoTestHashCollisions(t, HashFunc, 1024, 64)
}

func TestHashCollisions_blockSize1024_keyAmount1024(t *testing.T) {
	benchmark.DoTestHashCollisions(t, HashFunc, 1024, 1024)
}

func BenchmarkHash_intKeyType_blockSize16(b *testing.B) {
	benchmark.DoBenchmarkHash(b, HashFunc, 16, "int")
}
func BenchmarkHash_stringKeyType_blockSize16(b *testing.B) {
	benchmark.DoBenchmarkHash(b, HashFunc, 16, "string")
}

func BenchmarkHash_intKeyType_blockSize64(b *testing.B) {
	benchmark.DoBenchmarkHash(b, HashFunc, 64, "int")
}
func BenchmarkHash_stringKeyType_blockSize64(b *testing.B) {
	benchmark.DoBenchmarkHash(b, HashFunc, 64, "string")
}

func BenchmarkHash_intKeyType_blockSize128(b *testing.B) {
	benchmark.DoBenchmarkHash(b, HashFunc, 128, "int")
}
func BenchmarkHash_stringKeyType_blockSize128(b *testing.B) {
	benchmark.DoBenchmarkHash(b, HashFunc, 128, "string")
}

func BenchmarkHash_intKeyType_blockSize1024(b *testing.B) {
	benchmark.DoBenchmarkHash(b, HashFunc, 1024, "int")
}
func BenchmarkHash_stringKeyType_blockSize1024(b *testing.B) {
	benchmark.DoBenchmarkHash(b, HashFunc, 1024, "string")
}
