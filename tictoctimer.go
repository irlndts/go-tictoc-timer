package tictoctimer

import "time"

type TicTocTimer struct {
	duration time.Duration
	start    time.Time
}

func (self *TicTocTimer) Tic() {
	self.start = time.Now()
}

func (self *TicTocTimer) Toc() {
	self.duration = time.Now().Sub(self.start)
}

func (self *TicTocTimer) TicToc() time.Duration {
	return self.duration
}

//TODO
func (self *TicTocTimer) Average() {

}
