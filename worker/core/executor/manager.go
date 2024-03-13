package executor

import "github.com/2lovecode/powerjob-go/worker/common"

type Manager struct {
	coreExecutor                       interface{}
	lightweightTaskStatusCheckExecutor interface{}
	lightweightTaskExecutorService     interface{}
}

func NewManager(config *common.Config) *Manager {
	return &Manager{}
}

func (m *Manager) Shutdown() {

}
