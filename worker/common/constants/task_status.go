package constants

type TaskStatus int

//WAITING_DISPATCH(1, "等待调度器调度", "dispatching"),
//DISPATCH_SUCCESS_WORKER_UNCHECK(2, "调度成功（但不保证worker收到）", "unreceived"),
//WORKER_RECEIVED(3, "worker接收成功，但未开始执行", "received"),
//WORKER_PROCESSING(4, "worker正在执行", "running"),
//WORKER_PROCESS_FAILED(5, "worker执行失败", "failed"),
//WORKER_PROCESS_SUCCESS(6, "worker执行成功", "succeed")

const (
	WaitingDispatch              TaskStatus = 1 // 等待调度器调度
	DispatchSuccessWorkerUncheck TaskStatus = 2 // 调度成功（但不保证worker收到）
	WorkerReceived               TaskStatus = 3 // worker接收成功，但未开始执行
	WorkerProcessing             TaskStatus = 4 // worker正在执行
	WorkerProcessFailed          TaskStatus = 5 // worker执行失败
	WorkerProcessSuccess         TaskStatus = 6 // worker执行成功
)

func GetTaskStatusDesc(status TaskStatus) string {
	switch status {
	case WaitingDispatch:
		return "等待调度器调度"
	case DispatchSuccessWorkerUncheck:
		return "调度成功（但不保证worker收到）"
	case WorkerReceived:
		return "worker接收成功，但未开始执行"
	case WorkerProcessing:
		return "worker正在执行"
	case WorkerProcessFailed:
		return "worker执行失败"
	case WorkerProcessSuccess:
		return "worker执行成功"
	}
	return ""
}

func GetTaskStatusSimplyDesc(status TaskStatus) string {
	switch status {
	case WaitingDispatch:
		return "dispatching"
	case DispatchSuccessWorkerUncheck:
		return "unreceived"
	case WorkerReceived:
		return "received"
	case WorkerProcessing:
		return "running"
	case WorkerProcessFailed:
		return "failed"
	case WorkerProcessSuccess:
		return "succeed"
	}
	return ""
}
