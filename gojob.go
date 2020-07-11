package gojob

import (
	"sync"

	"github.com/bwmarrin/snowflake"
)

var node *snowflake.Node

func init() {
	var err error
	node, err = snowflake.NewNode(1)
	if err != nil {
		panic(err)
	}
}

type GoJob struct {
	sync.Mutex
	jobs []*Job
}

func (g *GoJob) AddJob(job *Job) {
	g.Lock()
	job.setid(node.Generate().Base36())
	g.jobs = append(g.jobs, job)
	g.Unlock()
}

func (g *GoJob) Schedule() {
}

func (g *GoJob) Size() int {
	return len(g.jobs)
}

func (g *GoJob) Abort(id string) {
	size := len(g.jobs)
	if size == 0 {
		return
	}

	for index, job := range g.jobs {
		if job.id == id {
			job.Abort()
			g.jobs = g.jobs[0:index]
			next := index + 1
			if next == size {
				return
			}
			g.jobs = append(g.jobs, g.jobs[next:size]...)
		}
	}
}
