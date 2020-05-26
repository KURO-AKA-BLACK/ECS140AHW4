package smash

import (
	"io"
	"bufio"
	"sync"
)

type word string

func Smash(r io.Reader, smasher func(word) uint32) map[uint32]uint {
	m := make(map[uint32]uint)
	var s []word
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		s = append(s, word(scanner.Text()))
	}
	var wg sync.WaitGroup
	var mutex sync.Mutex
	for i:= 0; i <len(s); i++ {
		wg.Add(1)
		mutex.Lock()
		go func(i int){
			_, ok := m[smasher(s[i])]
			if ok {
				m[smasher(s[i])] = m[smasher(s[i])] + 1
			} else {
				m[smasher(s[i])] = 1
			}
			mutex.Unlock()
			wg.Done()
		}(i)
	}
	wg.Wait()
	return m
}
