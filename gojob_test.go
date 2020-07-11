package gojob

import "testing"
import "reflect"

func Test_Jobs(t *testing.T) {
	gojob := GoJob{
		jobs: make([]*Job, 0),
	}

	job1 := Job{
		Job: (func(args ...interface{}) interface{} {
			return args[0]
		}),
	}
	job2 := Job{}

	gojob.AddJob(&job1)
	gojob.AddJob(&job2)

	gojob.Abort(job1.id)
	if !reflect.DeepEqual(1, gojob.Size) {
		t.FailNow()
	}

	gojob.Abort(job2.id)
	if !reflect.DeepEqual(0, gojob.Size) {
		t.FailNow()
	}
}
