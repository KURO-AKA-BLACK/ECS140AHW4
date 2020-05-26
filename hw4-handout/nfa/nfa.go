package nfa
import (
	"sync"
)
// A state in the NFA is labeled by a single integer.
type state uint

// TransitionFunction tells us, given a current state and some symbol, which
// other states the NFA can move to.
//
// Deterministic automata have only one possible destination state,
// but we're working with non-deterministic automata.
type TransitionFunction func(st state, act rune) []state
func flatten(r []state, f []state) []state {
	for i:=0; i < len(f); i++ {
		r = append(r, f[i])
	}
	return r
}
func Reachable(
	// `transitions` tells us what our NFA looks like
	transitions TransitionFunction,
	// `start` and `final` tell us where to start, and where we want to end up
	start, final state,
	// `input` is a (possible empty) list of symbols to apply.
	input []rune,
) bool {
	if input == nil {
		if start != final {
			return false
		}
		return true
	}
	var s []state
	var new_s []state
	var mutex sync.Mutex
	s = append(s, start)
	for index:=0; index < len(input); index ++{
		length := len(s)
		var wg sync.WaitGroup
		for i:=0; i < length; i++ {
			wg.Add(1)
			mutex.Lock()
			go func(){
				head := s[0]
				s = s[1:]
				new_s = flatten(new_s, transitions(head,input[index]))
				mutex.Unlock()
				wg.Done()
			}()
		}
		wg.Wait()
		s = new_s
		new_s = new_s[len(s):]
	}
	for j:=0; j < len(s); j++ {
		if s[j] == final {
			return true
		}
	}
	return false
}
