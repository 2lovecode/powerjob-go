package persistence

import (
	"github.com/2lovecode/powerjob-go/worker/common/constants"
	"github.com/2lovecode/powerjob-go/worker/core/processor"
)

type TaskPersistenceService interface {
	Init() error
	BatchSave(tasks []*TaskDO)
	UpdateTask(instanceId int, taskId string, updateEntity *TaskDO)
	UpdateTaskStatus(instanceId int, taskId string, status int, lastReportTime int, result string)
	UpdateLostTasks(instanceId int, addressList []string, retry bool)
	GetLastTask(instanceId int, subInstanceId int) *TaskDO
	GetAllUnFinishedTaskByAddress(instanceId int, address string) []*TaskDO
	GetTaskByStatus(instanceId int, status constants.TaskStatus, limit int) []*TaskDO
	GetTaskByQuery(instanceId int, customQuery string) []*TaskDO
	GetTaskStatusStatistics(instanceId int, subInstanceId int) map[constants.TaskStatus]int
	GetAllTaskResult(instanceId int, subInstanceId int) []*processor.TaskResult
	GetTask(instanceId int, taskId string) *TaskDO
	DeleteAllTasks(instanceId int) bool
	DeleteAllSubInstanceTasks(instanceId int, subInstanceId int) bool
	DeleteTasksByTaskIds(instanceId int, taskId []string) bool
}
