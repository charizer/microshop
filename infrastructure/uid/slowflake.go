package uid

import (
	"errors"
	"sync"
	"time"
)

const (
	BaseTs    = 1530687318000 //2018/7/4 14:55:18
	SeqLen    = 12
	WorkerLen = 10

	WorkerMask = 1<<WorkerLen - 1
	SeqMask    = 1<<SeqLen - 1

	TsShift   = SeqLen + WorkerLen
	SeqShift  = WorkerLen
	MaxWorkId = 1<<WorkerLen - 1
	MaxSeq    = 1<<SeqLen - 1
)

type UidGenerator struct {
	idStart int64
	idEnd   int64
	idCur   int64
	lastTs  int64
	seq     int64
	lock    sync.Mutex
}

var Idg *UidGenerator

func init(){
	Idg, _                       = NewDefaultIdGenerator(0)
}

func NewDefaultIdGenerator(instanceID int) (*UidGenerator, error) {
	start := 0
	end := 0

	if instanceID != -1 {
		start = instanceID
		end = instanceID
	}

	idg := new(UidGenerator)
	if start > MaxWorkId || start < 0 || end > MaxWorkId || end < 0 || start > end {
		panic(errors.New("wrong para"))
	}

	idg.idStart = int64(start)
	idg.idEnd = int64(end)
	idg.idCur = int64(start)

	idg.lastTs = -1
	idg.seq = 0
	return idg, nil
}

func NewIdGenerator(start, end int64) (*UidGenerator, error) {
	idg := new(UidGenerator)
	if start > MaxWorkId || start < 0 || end > MaxWorkId || end < 0 || start > end {
		panic(errors.New("wrong para"))
	}

	idg.idStart = start
	idg.idEnd = end
	idg.idCur = start

	idg.lastTs = -1
	idg.seq = 0
	return idg, nil
}

func (idg *UidGenerator) getMs() int64 {
	return time.Now().UnixNano() / 1000000
}

func (idg *UidGenerator) Next() (id int64, err error) {
	idg.lock.Lock()
	defer idg.lock.Unlock()

	now := idg.getMs()
	if now < idg.lastTs { //incase time schew
		time.Sleep(time.Duration(idg.lastTs-now) * time.Millisecond)
		now = idg.getMs()
	}

	if now == idg.lastTs {
		idg.seq = (idg.seq + 1) % SeqMask
		if idg.seq == 0 {
			if idg.idCur < idg.idEnd {
				idg.idCur = idg.idCur + 1
			} else {
				cur := idg.lastTs * 1000000
				for {
					now = time.Now().UnixNano() //wait next slot
					if now > cur {
						now = now / 1000000
						break
					}
				}
				idg.seq = 0
				idg.idCur = idg.idStart
			}
		}
	} else {
		idg.seq = 0
		idg.idCur = idg.idStart
	}

	idg.lastTs = now
	id = (now-BaseTs)<<TsShift | idg.seq<<SeqShift | idg.idCur
	return id, nil
}
