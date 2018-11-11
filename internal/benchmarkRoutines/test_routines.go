package benchmarkRoutines

import (
	"fmt"
	"sync"
	"testing"

	"github.com/xaionaro-go/atomicmap/errors"
	I "github.com/xaionaro-go/atomicmap/interfaces"
)

const (
	collisionCheckIterations = 1 << 20
)

type checkConsistencier interface {
	CheckConsistency() error
}

func expect(t *testing.T, m I.Map, key I.Key, expectedValue int) {
	value, err := m.Get(key)
	if err != nil {
		t.Errorf("Got an unexpected error: %v. key == %v; expectedValue == %v", err, key, expectedValue)
		return
	}
	if value != expectedValue {
		t.Errorf(`A wrong value "%v" (instead of %v)`, value, expectedValue)
	}
}

func DoTest(t *testing.T, factoryFunc mapFactoryFunc, keyHashFunc keyHashFunc) {
	m := factoryFunc(1024, keyHashFunc)

	if m.Count() != 0 && m.Count() != -1 { // "-1" means "unsupported"
		t.Errorf("m.Count() is not 0: %v", m.Count())
	}

	m.Set(1024*1024, 1)
	m.Set("a string", 2)

	expect(t, m, 1024*1024, 1)
	expect(t, m, "a string", 2)

	_, err := m.Get(3)
	if err != errors.NotFound {
		t.Errorf(`An expected "NotFound" error, but got: %v`, err)
	}

	if m.Count() != 2 && m.Count() != -1 { // "-1" means "unsupported"
		t.Errorf("m.Count() is not 2: %v", m.Count())
	}

	err = m.Unset(1024 * 1024)
	if err != nil {
		t.Errorf("Got an unexpected error: %v", err)
	}

	_, err = m.Get(1024 * 1024)
	if err != errors.NotFound {
		t.Errorf(`An expected "NotFound" error, but got: %v`, err)
	}

	if m.Count() != 1 && m.Count() != -1 { // "-1" means "unsupported"
		t.Errorf("m.Count() is not 1: %v", m.Count())
	}

	for i := 10; i < 1024*128; i++ {
		m.Set(i*6000, i)
	}
	err = m.Unset(60000)
	if err != nil {
		t.Errorf("Got an unexpected error: %v", err)
	}

	err = m.(checkConsistencier).CheckConsistency()
	if err != nil {
		t.Errorf("Got an unexpected error: %v", err)
		return
	}
	for i := 11; i < 1024*128; i++ {
		r, err := m.Get(i * 6000)
		if err != nil {
			t.Errorf("%v not found", i*6000)
			continue
		}
		if r.(int) != i {
			t.Errorf("%v != %v", r, i)
		}
	}

	err = m.(checkConsistencier).CheckConsistency()
	if err != nil {
		t.Errorf("Got an unexpected error: %v", err)
		return
	}

	for i := 11; i < 1024*128; i++ {
		err := m.Unset(i * 6000)
		if err != nil {
			t.Errorf("Cannot unset %v: %v", i*6000, err)
			continue
		}
	}

	err = m.(checkConsistencier).CheckConsistency()
	if err != nil {
		t.Errorf("Got an unexpected error: %v", err)
	}
}

func DoTestCollisions(t *testing.T, factoryFunc mapFactoryFunc, keyHashFunc keyHashFunc) {
	blockSize := uint64(16 * collisionCheckIterations)
	m := factoryFunc(blockSize, keyHashFunc)
	keys := generateKeys(collisionCheckIterations/2, "int")
	keys = append(keys, generateKeys(collisionCheckIterations/2, "string")...)

	collisions := 0
	for _, key := range keys {
		if m.(interface{ HasCollisionWithKey(I.Key) bool }).HasCollisionWithKey(key) {
			collisions++
		}
		m.Set(key, true)
	}

	fmt.Printf("Total collisions: %v/%v; bs%v (%.1f%%)\n", collisions, collisionCheckIterations, blockSize, float32(collisions)*100/float32(collisionCheckIterations))
}

func DoTestConcurrency(t *testing.T, factoryFunc mapFactoryFunc, keyHashFunc keyHashFunc) {
	blockSize := uint64(4)
	m := factoryFunc(blockSize, keyHashFunc)

	concurrency := 65536
	var wg sync.WaitGroup
	wg.Add(concurrency)
	for i := 0; i < concurrency; i++ {
		go func(i int) {
			defer wg.Done()
			err := m.Set(i, i)
			if err != nil {
				t.Errorf("Cannot m.Set(%v, %v): %v", i, i, err)
			}
			r, err := m.Get(i)
			if err != nil {
				t.Errorf("Cannot m.Get(%v): %v", i, err)
			} else {
				rInt, ok := r.(int)
				if i != rInt || !ok {
					t.Errorf("%v != %v", i, r)
				}
			}
			/*err = m.Unset(i)
			if err != nil {
				t.Errorf("Cannot m.Unset(%v): %v", i, err)
			}*/
			//fmt.Println("DoTestConcurrency", i)
		}(i)
	}

	wg.Wait()
}

func tryHashCollisions(keyHashFunc keyHashFunc, blockSize uint64, keys []interface{}) int {
	alreadyIsSet := map[uint64]bool{}

	collisions := 0
	for _, key := range keys {
		newHash := keyHashFunc(blockSize, key)
		if alreadyIsSet[newHash] {
			collisions++
		}
		alreadyIsSet[newHash] = true
	}

	return collisions
}

func DoTestHashCollisions(t *testing.T, keyHashFunc keyHashFunc, blockSize uint64, keyAmount uint64) {
	keys := generateKeys(keyAmount/2, "int")
	keys = append(keys, generateKeys(keyAmount/2, "string")...)

	collisions := tryHashCollisions(keyHashFunc, blockSize, keys)
	fmt.Printf("Total collisions on random keys: collisions %v, keyAmount %v and blockSize %v:\n\t%v/%v/%v (%.1f%%)\n", collisions, keyAmount, blockSize, collisions, keyAmount, blockSize, float32(collisions)*100/float32(keyAmount))

	keys = []interface{}{}
	for i := uint64(0); i < keyAmount; i++ {
		keys = append(keys, i*uint64(blockSize)*63)
	}

	collisions = tryHashCollisions(keyHashFunc, blockSize, keys)
	fmt.Printf("Total collisions on keys of pessimistic scenario (keys are multiple of blockSize): collisions %v, keyAmount %v and blockSize %v:\n\t%v/%v/%v (%.1f%%)\n", collisions, keyAmount, blockSize, collisions, keyAmount, blockSize, float32(collisions)*100/float32(keyAmount))

	keys = []interface{}{}
	for i := uint64(0); i < keyAmount; i++ {
		keys = append(keys, i)
	}

	collisions = tryHashCollisions(keyHashFunc, blockSize, keys)
	fmt.Printf("Total collisions on keys of pessimistic scenario (keys are consecutive): collisions %v, keyAmount %v and blockSize %v:\n\t%v/%v/%v (%.1f%%)\n", collisions, keyAmount, blockSize, collisions, keyAmount, blockSize, float32(collisions)*100/float32(keyAmount))
}
