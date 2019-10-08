package main

import "testing"

func BenchmarkQuery10(b *testing.B) {
	b.Logf("Querying %d entities\n", len(entityList))
	for n := 0; n < b.N; n++ {
		hasRecordAt(10, 10)
	}
}

func BenchmarkBitQuery10(b *testing.B) {
	b.Logf("Querying %d entities for flag\n", len(entityList))
	for n := 0; n < b.N; n++ {
		b.Logf("Fire: %d", entityOnFireCount())
	}
}
