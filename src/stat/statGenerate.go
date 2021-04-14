package stat

import (
	"time"
)

import (
	"github.com/mackerelio/go-osstat/memory"
	"github.com/mackerelio/go-osstat/cpu"
)

type SystemStat struct {
	MEM_used_total uint64
	MEM_used_used uint64
	MEM_used_cached uint64
	MEM_used_free uint64

	CPU_user_percent float64
	CPU_sys_percent float64
	CPU_idle_percent float64
}

func GetServerSystemStats() SystemStat {
	s := SystemStat{0,0,0,0,0.0,0.0,0.0};

	before, errBc := cpu.Get()
	if errBc == nil {
		time.Sleep(time.Duration(1) * time.Second)
		after, errAc := cpu.Get()
		if errAc == nil {
			total := float64(after.Total - before.Total)
			s.CPU_user_percent = float64(after.User-before.User)/total*100;
			s.CPU_sys_percent = float64(after.System-before.System)/total*100;
			s.CPU_idle_percent = float64(after.Idle-before.Idle)/total*100;
		}
	}

	memory, errM := memory.Get()
	if errM == nil {
		s.MEM_used_total = memory.Total;
		s.MEM_used_used = memory.Used;
		s.MEM_used_cached = memory.Cached;
		s.MEM_used_free = memory.Free;
	}

	return s
}