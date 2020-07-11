package gojob

type Callback func(args ...interface{}) interface{}
type OnSuccess func(args ...interface{}) interface{}
type OnFail func(args ...interface{}) interface{}

// 一个任务的数据结构
type Job struct {
	id        string
	CreatedAt int64 // 创建时间
	StartedAt int64 // 开始时间

	// 失败的调度机制，第一次失败后会触发
	// 数组元素i是第i+1次失败以后的重启间隔
	Schedul []int64

	// 记录当前失败次数
	Current int64

	// 任务，是一个由调用者自行定义的回调函数
	Job interface{}

	// 成功的回调
	OnSuccess interface{}

	// 失败的回调
	OnFail interface{}
}

func (j *Job) setid(id string) {
	j.id = id
}

func (j *Job) GetID() string {
	return j.id
}

// 开始一个任务
func (j *Job) Start() {}

// 取消一个任务
func (j *Job) Abort() {}

// 任务成功
func (j *Job) Success() {}

// 任务单次失败
func (j *Job) Fail() {}

// 任务多次尝试后仍然失败
func (j *Job) FinnallyFail() {}
