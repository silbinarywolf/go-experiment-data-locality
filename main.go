package main

import "unsafe"

type flag uint32

// For a block of uint32, flagSize will be equal to 32
// For a block of uint64, flagSize will be equal to 64
const flagSize = flag(unsafe.Sizeof(flag(0)) * 8)

// Entity is a rough mock-up of a typical entity object
// in my game engine
type Entity struct {
	X, Y              float64
	W, H              float64
	IsOnFire          bool
	SpriteIndex       int
	ImageIndex        int
	Direction         int
	ObjectIndex       int
	Depth             int
	ImageAngleRadians float64
	TimerA            int
	TimerB            int
	TimerC            int
	TimerD            int
	_                 bloatStructSize
	_                 bloatStructSize
	_                 bloatStructSize
	_                 bloatStructSize
	_                 bloatStructSize
}

// bloatStructSize purposefully bloats the size of the struct
// so that there are more cache misses and thus, iterating over
// game objects is slower
type bloatStructSize struct {
	_ int
	_ int
	_ int
	_ int
	_ int
	_ int
	_ int
	_ int
	_ int
	_ int
	_ int
	_ int
}

const entitySize = 1024

var (
	isOnFireFlagList [entitySize / flagSize]flag
	entityList       [entitySize]Entity
	temporaryMemory  [entitySize]int
)

func B2i(b bool) int {
	if b {
		return 1
	}
	return 0
}

func clearEntityOnFire(index int) {
	ind := flag(index)
	slotIndex := ind / flagSize
	bitIndex := ind % flagSize
	isOnFireFlagList[slotIndex] &^= (1 << bitIndex)
}

func setEntityOnFire(index int) {
	ind := flag(index)
	slotIndex := ind / flagSize
	bitIndex := ind % flagSize
	isOnFireFlagList[slotIndex] |= (1 << bitIndex)
}

func slowEntityOnFireCount() int {
	count := 0
	for _, record := range entityList {
		count += B2i(record.IsOnFire)
	}
	return count
}

func entityOnFireCount() int {
	count := 0
	for _, block := range isOnFireFlagList {
		for i := flag(0); i < flagSize; i++ {
			count += B2i(block&(1<<i) != 0)
		}
	}
	return count
}

func slowEntityOnFireResult() []int {
	var arr []int
	for i, _ := range entityList {
		arr = append(arr, i)
	}
	return arr
}

func entityOnFireResult() []int {
	var arr []int
	for slotIndex, block := range isOnFireFlagList {
		for bitIndex := flag(0); bitIndex < flagSize; bitIndex++ {
			if block&(1<<bitIndex) != 0 {
				arr = append(arr, int(slotIndex)+int(bitIndex))
			}
		}
	}
	return arr
}

func entityOnFireResultNoAlloc() []int {
	// NOTE(Jake): 2019-10-08
	// reuse the same int slice so that we avoid
	// allocations.
	arr := temporaryMemory[:0]

	for slotIndex, block := range isOnFireFlagList {
		for bitIndex := flag(0); bitIndex < flagSize; bitIndex++ {
			if block&(1<<bitIndex) != 0 {
				arr = append(arr, int(slotIndex)+int(bitIndex))
			}
		}
	}
	return arr
}
