package trace

import (
	"bytes"
	"github.com/google/uuid"
	"runtime"
	"strconv"
	"sync"
)

var traceMap = map[uint64]string{}
var lock = sync.Mutex{}

func InitTraceId() (uint64, string) {
	lock.Lock()
	defer lock.Unlock()
	gid := getGid()
	traceId := getDefaultTraceId()
	traceMap[getGid()] = traceId
	return gid, traceId
}
func InitWithTraceId(traceId string) (uint64, string) {
	lock.Lock()
	defer lock.Unlock()
	gid := getGid()
	traceMap[getGid()] = traceId
	return gid, traceId
}

func getGid() (gid uint64) {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, err := strconv.ParseUint(string(b), 10, 64)
	if err != nil {
		panic(err)
	}
	return n
}

func GetTraceId() (uint64, string) {
	gid := getGid()
	traceId, b := traceMap[getGid()]
	if !b {
		traceId = getDefaultTraceId()
		return InitTraceId()
	}
	return gid, traceId
}

func getDefaultTraceId() string {
	return uuid.New().String()
}
