package day10

type StopCb = func() bool
type CycleCb = func(i int)

type CpuClock struct {
	StopCb    StopCb
	Callbacks []CycleCb
}

func NewCpuClock(stopCb StopCb) *CpuClock {
	return &CpuClock{
		StopCb:    stopCb,
		Callbacks: []CycleCb{},
	}
}

func (cc *CpuClock) RegisterCb(cb CycleCb) {
	cc.Callbacks = append(cc.Callbacks, cb)
}

func (cc *CpuClock) Start() {
	// We start at one for easier cycle number handling.
	for i := 1; !cc.StopCb(); i += 1 {
		for _, cb := range cc.Callbacks {
			cb(i)
		}
	}
}
