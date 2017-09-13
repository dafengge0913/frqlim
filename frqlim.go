package frqlim

import (
	"container/list"
	"time"
)

type Limiter struct {
	q      *list.List
	size   int
	period time.Duration
}

func New(size int, period time.Duration) *Limiter {
	return &Limiter{
		q:      list.New(),
		size:   size,
		period: period,
	}
}

func (lim *Limiter) Do() bool {
	now := time.Now()
	if lim.q.Len() >= lim.size {
		for lim.q.Len() > 0 {
			t := lim.q.Front().Value.(time.Time)
			if now.Sub(t) > lim.period {
				lim.q.Remove(lim.q.Front())
			} else {
				break
			}
		}
	}
	if lim.q.Len() >= lim.size {
		return false
	}
	lim.q.PushBack(now)
	return true
}
