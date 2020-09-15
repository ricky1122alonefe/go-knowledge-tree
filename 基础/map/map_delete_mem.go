package main

import (
	"log"
	"runtime"
	"time"
)

var lastTotalFreed uint64
var intMap map[int]int
var cnt = 8192

//
func main() {
	log.Println("--------------------------------------------------------------")
	printMemStats()
	log.Println("--------------------------------------------------------------")
	initMap()
	runtime.GC()
	printMemStats()
	log.Println("--------------------------------------------------------------")
	log.Println(len(intMap))
	for i := 0; i < cnt; i++ {
		delete(intMap, i)
	}
	log.Println(len(intMap))

	runtime.GC()
	log.Println("--------------------------------------------------------------")
	printMemStats()
	log.Println("--------------------------------------------------------------")

	// intMap = nil
	runtime.GC()
	log.Println("--------------------------------------------------------------")
	printMemStats()
	log.Println("--------------------------------------------------------------")

	time.Sleep(time.Minute * 2)
	runtime.GC()
	log.Println("--------------------------------------------------------------")
	printMemStats()
	log.Println("--------------------------------------------------------------")

}

func initMap() {
	intMap = make(map[int]int, cnt)

	for i := 0; i < cnt; i++ {
		intMap[i] = i
	}
}

func printMemStats() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	log.Printf("Alloc = %v TotalAlloc = %v  Just Freed = %v Sys = %v NumGC = %v\n",
		m.Alloc/1024, m.TotalAlloc/1024, ((m.TotalAlloc-m.Alloc)-lastTotalFreed)/1024, m.Sys/1024, m.NumGC)

	lastTotalFreed = m.TotalAlloc - m.Alloc
}
