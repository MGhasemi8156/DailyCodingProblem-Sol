package main

import "fmt"

type LogManager struct {
	n    int
	last int
	logs []int
}

func newLogManager(n int) *LogManager {
	logManager := LogManager{n: n, last: 0, logs: make([]int, 2*n)}
	return &logManager
}

// has amortize complexity O(1)
func (logManager *LogManager) record(record_id int) {
	n := logManager.n
	if logManager.last == 2*n-1 {
		for i := 0; i < n-1; i++ {
			logManager.logs[i] = logManager.logs[i+n+1]
		}
		logManager.logs[n-1] = record_id
		logManager.last = n - 1
	} else {
		logManager.last += 1
		logManager.logs[logManager.last] = record_id
	}
}

func (logManager *LogManager) get_last(i int) int {
	return logManager.logs[logManager.last-i]
}

func main() {
	n := 2

	logManager := newLogManager(n)

	logManager.record(1111)
	logManager.record(2222)
	logManager.record(3333)
	logManager.record(4444)
	logManager.record(5555)

	fmt.Println(logManager.get_last(0))
	fmt.Println(logManager.get_last(1))

	fmt.Println("hello")
}
