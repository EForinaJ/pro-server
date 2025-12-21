package utils_snow

import (
	"fmt"
	"sync"
	"time"
)

const (
	// 定义纪元，即初始时间戳，可以根据需要调整
	epoch = 1609459200000 // 例如: 2021-01-01 00:00:00 UTC 的毫秒数

	workerIdBits     = 5
	datacenterIdBits = 5
	maxWorkerId      = -1 ^ (-1 << workerIdBits)
	maxDatacenterId  = -1 ^ (-1 << datacenterIdBits)

	sequenceBits       = 12
	workerIdShift      = sequenceBits
	datacenterIdShift  = sequenceBits + workerIdBits
	timestampLeftShift = sequenceBits + workerIdBits + datacenterIdBits

	sequenceMask = -1 ^ (-1 << sequenceBits)
)

type SnowFlake struct {
	lastTimestamp int64
	sequence      int64
	datacenterId  int64
	workerId      int64
	lock          sync.Mutex
}

func NewSnowFlake(datacenterId, workerId int64) (*SnowFlake, error) {
	if datacenterId > maxDatacenterId || datacenterId < 0 {
		return nil, fmt.Errorf("datacenter Id can't be greater than %d or less than 0", maxDatacenterId)
	}
	if workerId > maxWorkerId || workerId < 0 {
		return nil, fmt.Errorf("worker Id can't be greater than %d or less than 0", maxWorkerId)
	}
	return &SnowFlake{
		lastTimestamp: -1,
		sequence:      0,
		datacenterId:  datacenterId,
		workerId:      workerId,
	}, nil
}

func (sf *SnowFlake) TillNextMillis(lastTimestamp int64) int64 {
	timestamp := sf.timeGen()
	for timestamp <= lastTimestamp {
		timestamp = sf.timeGen()
	}
	return timestamp
}

func (sf *SnowFlake) timeGen() int64 {
	return time.Now().UnixNano() / 1e6
}

func (sf *SnowFlake) Generate() int64 {
	sf.lock.Lock()
	defer sf.lock.Unlock()

	timestamp := sf.timeGen()

	if timestamp < sf.lastTimestamp {
		panic(fmt.Sprintf("Clock moved backwards. Refusing to generate id for %d milliseconds", sf.lastTimestamp-timestamp))
	}

	if sf.lastTimestamp == timestamp {
		sf.sequence = (sf.sequence + 1) & sequenceMask
		if sf.sequence == 0 {
			timestamp = sf.TillNextMillis(sf.lastTimestamp)
		}
	} else {
		sf.sequence = 0
	}

	sf.lastTimestamp = timestamp

	id := ((timestamp - epoch) << timestampLeftShift) |
		(sf.datacenterId << datacenterIdShift) |
		(sf.workerId << workerIdShift) |
		sf.sequence

	return id
}
