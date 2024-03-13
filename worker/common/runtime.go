package common

import (
	"github.com/2lovecode/powerjob-go/common/model"
	"github.com/2lovecode/powerjob-go/remote/framework/transporter"
	"github.com/2lovecode/powerjob-go/worker/background"
	"github.com/2lovecode/powerjob-go/worker/background/discovery"
	"github.com/2lovecode/powerjob-go/worker/core/executor"
	"github.com/2lovecode/powerjob-go/worker/persistence"
	"github.com/2lovecode/powerjob-go/worker/processor"
)

type Runtime struct {
	// App 基础信息
	appInfo *model.AppInfo
	// 当前执行器地址
	address string
	// 用户配置
	config *Config
	// 通讯器
	transporter transporter.Transporter
	// 处理器加载器
	processorLoader        processor.Loader
	executorManager        *executor.Manager
	omsLogHandler          *background.OmsLogHandler
	serverDiscoverService  discovery.ServerDiscoveryService
	taskPersistenceService persistence.TaskPersistenceService
}

func (r *Runtime) GetAppId() int {
	return r.appInfo.GetAppId()
}

func NewRuntime() *Runtime {
	return &Runtime{}
}

func (r *Runtime) SetConfig(cnf *Config) {
	r.config = cnf
}

func (r *Runtime) GetExecutorManager() *executor.Manager {
	return r.executorManager
}
