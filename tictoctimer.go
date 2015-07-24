package tictoctimer

import "time"

type TicTocTimer struct {
	duration time.Duration
	start    time.Time
	repeats  int64
}

func (self *TicTocTimer) Tic() {
	self.start = time.Now()
}

func (self *TicTocTimer) Toc() {
	self.duration += time.Now().Sub(self.start)
	self.repeats++
}

func (self *TicTocTimer) TicToc() time.Duration {
	return self.duration
}

func (self *TicTocTimer) Average() time.Duration {
	return (self.duration / time.Duration(self.repeats))
}

func (self *TicTocTimer) Clear() {
	self.duration = 0
	self.repeats = 0
}
