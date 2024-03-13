package persistence

// TaskDO
/**
 * 为了简化 DAO 层，一张表实现两种功能
 * 对于 TaskTracker，task_info 存储了当前 JobInstance 所有的子任务信息
 * 对于普通的 Worker，task_info 存储了当前无法处理的任务信息
 *
 */
type TaskDO struct {
	/**
	 * 层次命名法，可以表示 Map 后的父子关系，如 0.1.2 代表 rootTask map 的第一个 task map 的第二个 task
	 */
	taskId string
	/**
	 * 任务实例 ID
	 */
	instanceId int

	/**
	 * 秒级任务专用
	 * 对于普通任务而言 等于 instanceId
	 * 对于秒级（固定频率）任务 自增长
	 */
	subInstanceId int
	/**
	 * 任务名称
	 */
	taskName string
	/**
	 *  任务对象（序列化后的二进制数据）
	 */
	taskContent []byte
	/**
	 * 对于TaskTracker为workerAddress（派发地址），对于普通Worker为TaskTrackerAddress（汇报地址），所有地址都是 IP:Port
	 */
	address string
	/**
	 * 任务状态，0～10代表 JobTracker 使用，11～20代表普通Worker使用
	 */
	status int
	/**
	 * 执行结果
	 */
	result string
	/**
	 * 失败次数
	 */
	failedCnt int
	/**
	 * 创建时间
	 */
	createdTime int
	/**
	 * 最后修改时间
	 */
	lastModifiedTime int
	/**
	 * ProcessorTracker 最后上报时间
	 */
	lastReportTime int
}

func (td *TaskDO) FetchUpdateSQL() string {
	if td.address != "" {

	}

	return ""
}
