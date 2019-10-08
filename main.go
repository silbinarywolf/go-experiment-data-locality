package main

import "unsafe"

type flag uint32

// For a block of uint32, flagSize will be equal to 32
// For a block of uint64, flagSize will be equal to 64
const flagSize = uint32(unsafe.Sizeof(flag(0)) * 8)

type Entity struct {
	X, Y        float64
	W, H        float64
	SpriteIndex int
}

const entitySize = 1024

var (
	isDeadFlagList [entitySize / flagSize]flag
	entityList     [entitySize]Entity
)

func entityOnFireCount() int {
	count := 0
	for _, block := range isDeadFlagList {
		if block&(1<<1) != 0 {
			count++
		}
		if block&(1<<2) != 0 {
			count++
		}
		if block&(1<<3) != 0 {
			count++
		}
		if block&(1<<4) != 0 {
			count++
		}
		if block&(1<<5) != 0 {
			count++
		}
		if block&(1<<6) != 0 {
			count++
		}
		if block&(1<<7) != 0 {
			count++
		}
		if block&(1<<8) != 0 {
			count++
		}
		if block&(1<<9) != 0 {
			count++
		}
		if block&(1<<10) != 0 {
			count++
		}
		if block&(1<<11) != 0 {
			count++
		}
		if block&(1<<12) != 0 {
			count++
		}
		if block&(1<<13) != 0 {
			count++
		}
		if block&(1<<14) != 0 {
			count++
		}
		if block&(1<<15) != 0 {
			count++
		}
		if block&(1<<16) != 0 {
			count++
		}
		if block&(1<<17) != 0 {
			count++
		}
		if block&(1<<18) != 0 {
			count++
		}
		if block&(1<<19) != 0 {
			count++
		}
		if block&(1<<20) != 0 {
			count++
		}
		if block&(1<<21) != 0 {
			count++
		}
		if block&(1<<22) != 0 {
			count++
		}
		if block&(1<<23) != 0 {
			count++
		}
		if block&(1<<24) != 0 {
			count++
		}
		if block&(1<<25) != 0 {
			count++
		}
		if block&(1<<26) != 0 {
			count++
		}
		if block&(1<<27) != 0 {
			count++
		}
		if block&(1<<28) != 0 {
			count++
		}
		if block&(1<<29) != 0 {
			count++
		}
		if block&(1<<30) != 0 {
			count++
		}
		if block&(1<<31) != 0 {
			count++
		}
		/*for i := uint32(0); i < flagSize; i++ {
			if block&(1<<i) != 0 {
				count++
			}
		}*/
	}
	return count
}

func hasRecordAt(x, y float64) bool {
	for _, record := range entityList {
		if record.X == x && record.Y == y {
			return true
		}
	}
	return false
}
