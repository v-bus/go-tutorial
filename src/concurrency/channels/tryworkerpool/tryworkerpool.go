package main

import (
	"concurrency/channels/tryworkerpool/workerpool"
)

func main() {
	workerpool.RunWorkerPool(100, 1)
}
