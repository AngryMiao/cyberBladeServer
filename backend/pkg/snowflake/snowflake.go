package snowflake

import (
	"errors"
	"net"
	"os"
	"strconv"
	"sync"
	"time"
)

const (
	WorkerIDKey = "WORKER_ID"
	WorkerIPKey = "WORKER_IP"
)
const (
	workerBits  uint8 = 16
	numberBits  uint8 = 5
	workerMax   int64 = -1 ^ (-1 << workerBits)
	numberMax   int64 = -1 ^ (-1 << numberBits)
	timeShift   uint8 = workerBits + numberBits + 1
	workerShift uint8 = numberBits
	startTime   int64 = 1420041600000 // 如果在程序跑了一段时间修改了epoch这个值 可能会导致生成相同的ID
)

var (
	One       sync.Once
	SnowFlake *Worker
)

type Worker struct {
	mu        sync.Mutex
	timestamp int64
	workerId  int64
	number    int64
}

func getWorkerIdByIP(address string) int64 {
	b := net.ParseIP(address).To4()

	return int64(b[3]) | int64(b[2])<<8
}

func GetWorker() *Worker {
	One.Do(func() {
		var workerID int64
		var err error

		if idValue := os.Getenv(WorkerIDKey); idValue != "" {
			workerID, err = strconv.ParseInt(idValue, 10, 64)
			if err != nil {
				workerID = 1
			}
		} else if ipValue := os.Getenv(WorkerIPKey); ipValue != "" {
			workerID = getWorkerIdByIP(ipValue)
		} else {
			workerID = 1
		}

		initWorker(workerID)
	})

	return SnowFlake

}

func initWorker(workerId int64) {
	worker, _ := NewWorker(workerId)
	SnowFlake = worker

}

func NewWorker(workerId int64) (*Worker, error) {
	if workerId < 0 || workerId > workerMax {
		return nil, errors.New("worker ID excess of quantity")
	}

	// 生成一个新节点
	return &Worker{
		timestamp: 0,
		workerId:  workerId,
		number:    0,
	}, nil
}

func (w *Worker) GetID() int64 {
	w.mu.Lock()
	defer w.mu.Unlock()

	now := time.Now().UnixNano() / 1e6
	if w.timestamp == now {
		w.number++
		if w.number > numberMax {
			for now <= w.timestamp {
				now = time.Now().UnixNano() / 1e6
			}
		}
	} else {
		w.number = 0
		w.timestamp = now
	}
	return (now-startTime)<<timeShift | (w.workerId << workerShift) | (w.number)
}

func (w *Worker) GetStringID() string {
	return strconv.FormatInt(w.GetID(), 10)
}
