package test

type Glimit struct {
	n  int
	ch chan struct{}
}

//New initialization Glimit struct
func New(num int) *Glimit {
	return &Glimit{
		n:  num,
		ch: make(chan struct{}, num),
	}
}

//Run f in a new goroutine but with limit
func (g *Glimit) Run(f func()) {
	g.ch <- struct{}{}
	go func() {
		f()
		<-g.ch
	}()
}
