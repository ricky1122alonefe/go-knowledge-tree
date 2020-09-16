package main

import (
	"log"
	"runtime"
	"fmt"
	"reflect"
)

var intMapMap map[int]*A
type A struct {
	T int
}

var cnt = 1024
var lastTotalFreed uint64 // size of last memory has been freed

func main() {
	// 1
	log.Println("--------------------------------------------------------------")
	printMemStats()

	// 2
	log.Println("--------------------------------------------------------------")
	initMapMap()
	runtime.GC()
	printMemStats()

	// 3
	log.Println("--------------------------------------------------------------")
	// fillMapMap()
	runtime.GC()
	printMemStats()

	// 4
	log.Println("--------------------------------------------------------------")
	log.Println(len(intMapMap))
	for i := 0; i < cnt; i++ {
		delete(intMapMap, i)
	}
	log.Println(len(intMapMap))
	runtime.GC()
	printMemStats()

	// 5
	log.Println("--------------------------------------------------------------")
	intMapMap = nil
	runtime.GC()
	printMemStats()
}

func initMapMap() {
	intMapMap = make(map[int]*A, cnt)
	for i:=0;i<cnt;i++{
		a:=new(A)
		intMapMap[i] = a
	}

}


func printMemStats() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	log.Printf("Alloc = %v TotalAlloc = %v  Just Freed = %v Sys = %v NumGC = %v\n",
		m.Alloc/1024, m.TotalAlloc/1024, ((m.TotalAlloc-m.Alloc)-lastTotalFreed)/1024, m.Sys/1024, m.NumGC)

	lastTotalFreed = m.TotalAlloc - m.Alloc
}
