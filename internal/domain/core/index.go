package core

import (
	"birthday-bot/internal/adapters/logger"
	"birthday-bot/internal/adapters/mail"
	"birthday-bot/internal/adapters/repo"
	"birthday-bot/internal/adapters/repo/pg"
	"sync"
)

type St struct {
	lg   logger.Lite
	repo repo.Repo

	stopped   bool
	stoppedMu sync.RWMutex

	Mail mail.Client
	User *User
}

func New(
	lg logger.Lite,
	repo *pg.St,
	mail mail.Client,
) *St {
	c := &St{
		lg:   lg,
		repo: repo,
		Mail: mail,
	}
	c.User = NewUser(c)

	return c
}

func (c *St) IsStopped() bool {
	c.stoppedMu.RLock()
	defer c.stoppedMu.RUnlock()
	return c.stopped
}

func (c *St) StopAndWaitJobs() {
	c.stoppedMu.Lock()
	c.stopped = true
	c.stoppedMu.Unlock()

}
