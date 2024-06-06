package luhn

import (
	"math/rand"
	"sync"
	"time"
)

func New() *Generator {
	return &Generator{
		seed:  rand.New(rand.NewSource(time.Now().UnixNano())),
		mutex: sync.Mutex{},
	}
}

func Calc(str string) int {
	value := make([]int, len(str))
	for i := 0; i < len(str); i++ {
		value[i] = int(str[i] - '0')
	}
	return calc(value)
}

type Generator struct {
	seed  *rand.Rand
	mutex sync.Mutex
}

func (gen *Generator) Generate(length int, prefix ...string) string {
	gen.mutex.Lock()
	defer gen.mutex.Unlock()
	var iin string
	if len(prefix) > 0 {
		iin = prefix[0]
	}
	value := make([]int, length)
	for i := 0; i < len(iin); i++ {
		value[i] = int(iin[i] - '0')
	}
	for i := len(iin); i < length-1; i++ {
		value[i] = gen.seed.Intn(9)
	}
	value[length-1] = calc(clone(value))
	result := make([]byte, length)
	for i := 0; i < length; i++ {
		result[i] = byte(value[i] + '0')
	}
	return string(result)
}

func clone(value []int) []int {
	buf := make([]int, len(value))
	for i, v := range value {
		buf[i] = v
	}
	return buf
}

func calc(value []int) int {
	sum := 0
	for i := 0; i < len(value); i++ {
		if i%2 == 0 {
			value[i] *= 2
			if value[i] > 9 {
				value[i] -= 9
			}
		}
		sum += value[i]
	}
	return (10 - sum%10) % 10
}

func Verify(ccn string) bool {
	value := make([]int, len(ccn))
	for i := 0; i < len(ccn); i++ {
		value[i] = int(ccn[i] - '0')
	}
	return calc(value) == 0
}
