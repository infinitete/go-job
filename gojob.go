package gojob

type Callback func(...interface{}) interface{}
type OnSuccess Callback
type OnFail Callback

// 一个任务的数据结构
type Job struct {
	CreatedAt int64 // 创建时间
	StartedAt int64 // 开始时间

	// 失败的调度机制，第一次失败后会触发
	// 数组元素i是第i+1次失败以后的重启间隔
	Schedul []int64

	// 记录当前失败次数
	Current int64

	// 任务，是一个由调用者自行定义的回调函数
	Job Callback

	// 成功的回调
	OnSuccess OnSuccess

	// 失败的回调
	OnFail OnFail
}

// 开始一个任务
func (j *Job) Start() {}

// 任务成功
func (j *Job) Success() {}

// 任务单次失败
func (j *Job) Fail() {}

// 任务多次尝试后仍然失败
func (j *Job) FinnallyFail() {}

func GoJob() {
}
