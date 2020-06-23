package letter

import "sync"

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

type SafeCounter struct {
	m   FreqMap
	mux sync.Mutex
}

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

func ConcurrentFrequency(s string) FreqMap {

	counter := SafeCounter{m: make(map[rune]int)}

	for _, r := range s {
		counter.mux.Lock()
		counter.m[r]++
		counter.mux.Unlock()
	}

	return counter.m
}
