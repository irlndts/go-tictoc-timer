package tictoctimer

import "time"

//structure for timer
type TicTocTimer struct {
	duration time.Duration
	start    time.Time
	repeats  int64
}

//constructor with default values for timer
func Init() *TicTocTimer {
	return &TicTocTimer{duration: 0, start: time.Now(), repeats: 0}
}

//start timer
func (self *TicTocTimer) Tic() {
	self.start = time.Now()
}

//pause timer
func (self *TicTocTimer) Toc() {
	self.duration += time.Now().Sub(self.start)
	self.repeats++
}

//total time of timer
func (self *TicTocTimer) TicToc() time.Duration {
	return self.duration
}

//average time for timer
func (self *TicTocTimer) Average() time.Duration {
	return (self.duration / time.Duration(self.repeats))
}

//clear value
func (self *TicTocTimer) Clear() {
	self.duration = 0
	self.repeats = 0
}
