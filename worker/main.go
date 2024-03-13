package worker

import (
	"github.com/2lovecode/powerjob-go/remote/framework/engine"
	"github.com/2lovecode/powerjob-go/worker/common"
	"github.com/2lovecode/powerjob-go/worker/processor"
)

type PowerJobWorker struct {
	remoteEngine engine.Remote
	runtime      *common.Runtime
}

func NewPowerJobWorker(cnf *common.Config) *PowerJobWorker {
	worker := &PowerJobWorker{
		remoteEngine: engine.NewPowerJobRemoteEngine(),
		runtime:      common.NewRuntime(),
	}
	worker.runtime.SetConfig(cnf)
	return worker
}

func (pjw *PowerJobWorker) buildProcessorLoader(runtime *common.Runtime) processor.Loader {
	return nil
}

func (pjw *PowerJobWorker) Init() error {
	return nil
}

func (pjw *PowerJobWorker) Destroy() error {
	if pjw.runtime != nil && pjw.runtime.GetExecutorManager() != nil {
		pjw.runtime.GetExecutorManager().Shutdown()
	}
	if pjw.remoteEngine != nil {
		_ = pjw.remoteEngine.Close()
	}

	return nil
}
