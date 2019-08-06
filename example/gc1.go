package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func makeBuffer() []byte {
	return make([]byte, rand.Intn(5000000)+5000000)
}

func main() {
	pool := make([][]byte, 20)

	buffer := make(chan []byte, 5)

	var m runtime.MemStats
	makes := 0
	for {
		var b []byte
		select {
		//如果能从channel取值，就取值赋予b
		case b = <-buffer:
			//如果不能取值，channel为空，则直接生成一个随机
		default:
			makes += 1
			b = makeBuffer()
		}
		//生成一个基于pool长度为基数生成的随机数
		i := rand.Intn(len(pool))
		//如果pool切片中随机一个[]byte不为空
		if pool[i] != nil {
			select {
			//如果buffer channel把pool channel中i的【】byte读出来，则设置pool[i]为空
			case buffer <- pool[i]:
				pool[i] = nil
			default:
			}
		}
		//pool[i]赋值为b，这个b是从buffer中取出的，当buffer channel为空读取出来的是0为false，则用上面定义的makebuffer函数生成一个[]bytes
		pool[i] = b

		time.Sleep(time.Second)

		bytes := 0
		//启动一个for循环，次数为pool的长度
		for i := 0; i < len(pool); i++ {
			//如果pool[循环次数]不为空，则把pool[i]的长度和bytes相加，计算pool的长度
			if pool[i] != nil {
				bytes += len(pool[i])
			}
		}

		runtime.ReadMemStats(&m)
		fmt.Printf("%d,%d,%d,%d,%d,%d\n", m.HeapSys, bytes, m.HeapAlloc,
			m.HeapIdle, m.HeapReleased, makes)
	}
}