package main

import (
	"testing"
)

var count int

var arr []int

func debugLog(b *testing.B, format string, args ...interface{}) {
	//b.Logf(format, args...)
}

func BenchmarkSlowCountEntitiesOnFireQuery10(b *testing.B) {
	// Setup data
	entityList[5].IsOnFire = true
	entityList[127].IsOnFire = true

	// Start
	var c int
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c = slowEntityOnFireCount()
	}
	count = c
	debugLog(b, "count: %d\n", count)
}

func BenchmarkCountEntitiesOnFire10(b *testing.B) {
	// Setup data
	for i := 0; i < len(entityList); i++ {
		clearEntityOnFire(i)
	}
	setEntityOnFire(5)
	setEntityOnFire(127)

	//b.Logf("Querying %d entities for flag\n", len(entityList))

	// Start
	var c int
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c = entityOnFireCount()
	}
	count = c // stop "c" from being optimized away
	debugLog(b, "count: %d\n", count)
}

func BenchmarkEntitiesSlowSmallList10(b *testing.B) {
	// Setup data
	entityList[5].IsOnFire = true
	entityList[127].IsOnFire = true

	// Start
	var a []int
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		a = slowEntityOnFireResult()
	}
	arr = a // stop "a" from being optimized away
	debugLog(b, "arr size: %d\n", len(arr))
}

func BenchmarkEntitiesSmallList10(b *testing.B) {
	// Setup data
	for i := 0; i < len(entityList); i++ {
		clearEntityOnFire(i)
	}
	setEntityOnFire(5)
	setEntityOnFire(127)

	// Start
	var a []int
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		a = entityOnFireResult()
	}
	arr = a // stop "a" from being optimized away
	debugLog(b, "arr size: %d\n", len(arr))
}

func BenchmarkEntitiesSmallListNoAlloc10(b *testing.B) {
	// Setup data
	for i := 0; i < len(entityList); i++ {
		clearEntityOnFire(i)
	}
	setEntityOnFire(5)
	setEntityOnFire(127)

	// Start
	var a []int
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		a = entityOnFireResultNoAlloc()
	}
	arr = a // stop "a" from being optimized away
	debugLog(b, "arr size: %d\n", len(arr))
}

func BenchmarkEntitiesListAllOnFire10(b *testing.B) {
	// Setup data
	for i := 0; i < len(entityList); i++ {
		setEntityOnFire(i)
	}

	// Start
	var a []int
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		a = entityOnFireResult()
	}
	arr = a // stop "a" from being optimized away
	debugLog(b, "arr size: %d\n", len(arr))
}

func BenchmarkEntitiesListAllOnFireNoAlloc10(b *testing.B) {
	// Setup data
	for i := 0; i < len(entityList); i++ {
		setEntityOnFire(i)
	}

	// Start
	var a []int
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		a = entityOnFireResultNoAlloc()
	}
	arr = a // stop "a" from being optimized away
	debugLog(b, "arr size: %d\n", len(arr))
}
