package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// R1: 审计/状态机并发读写必须无数据竞争。

func TestOpsAuditConcurrentAddForNoRace(t *testing.T) {
	a := newOpsAudit()
	start := make(chan struct{})
	var wg sync.WaitGroup
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			<-start
			for j := 0; j < 50; j++ {
				a.Add(fmt.Sprintf("rec-%d", n%2), "created", "tester")
			}
		}(i)
	}
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			<-start
			for j := 0; j < 50; j++ {
				_ = a.For("rec-0")
			}
		}()
	}
	close(start)
	wg.Wait()
}

func TestOpsAuditConcurrentAddSinceNoRace(t *testing.T) {
	a := newOpsAudit()
	start := make(chan struct{})
	var wg sync.WaitGroup
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			<-start
			for j := 0; j < 50; j++ {
				a.Add("rec-x", "status_changed", "tester")
			}
		}()
	}
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			<-start
			for j := 0; j < 50; j++ {
				_ = a.Since(time.Time{})
			}
		}()
	}
	close(start)
	wg.Wait()
}

func TestOpsAuditConcurrentAddCountNoRace(t *testing.T) {
	a := newOpsAudit()
	start := make(chan struct{})
	var wg sync.WaitGroup
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			<-start
			for j := 0; j < 50; j++ {
				a.Add("rec-c", "created", "tester")
			}
		}()
	}
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			<-start
			for j := 0; j < 50; j++ {
				_ = a.Count()
			}
		}()
	}
	close(start)
	wg.Wait()
}

func TestOpsStateConcurrentMoveResetNoRace(t *testing.T) {
	m := newOpsStateMachine()
	start := make(chan struct{})
	var wg sync.WaitGroup
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			<-start
			for j := 0; j < 50; j++ {
				_ = m.Move(OpsStatusQueued, OpsStatusActive, "t")
			}
		}()
	}
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			<-start
			for j := 0; j < 50; j++ {
				m.Reset()
			}
		}()
	}
	close(start)
	wg.Wait()
}

func TestOpsStateConcurrentMoveOnlyNoRace(t *testing.T) {
	m := newOpsStateMachine()
	start := make(chan struct{})
	var wg sync.WaitGroup
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			<-start
			for j := 0; j < 50; j++ {
				_ = m.Move(OpsStatusQueued, OpsStatusActive, "t")
			}
		}()
	}
	close(start)
	wg.Wait()
}

func TestOpsStateConcurrentMoveHistoryNoRace(t *testing.T) {
	m := newOpsStateMachine()
	start := make(chan struct{})
	var wg sync.WaitGroup
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			<-start
			for j := 0; j < 50; j++ {
				_ = m.Move(OpsStatusQueued, OpsStatusActive, "t")
				_ = m.Move(OpsStatusActive, OpsStatusPaused, "t")
				_ = m.Move(OpsStatusPaused, OpsStatusClosed, "t")
			}
		}()
	}
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			<-start
			for j := 0; j < 50; j++ {
				_ = m.History()
			}
		}()
	}
	close(start)
	wg.Wait()
}

func TestOpsStateConcurrentMoveLastNoRace(t *testing.T) {
	m := newOpsStateMachine()
	start := make(chan struct{})
	var wg sync.WaitGroup
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			<-start
			for j := 0; j < 50; j++ {
				_ = m.Move(OpsStatusQueued, OpsStatusActive, "t")
			}
		}()
	}
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			<-start
			for j := 0; j < 50; j++ {
				_, _ = m.Last()
			}
		}()
	}
	close(start)
	wg.Wait()
}
