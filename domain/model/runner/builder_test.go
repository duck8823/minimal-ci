package runner_test

import (
	"context"
	"github.com/duck8823/duci/domain/model/docker"
	"github.com/duck8823/duci/domain/model/job"
	"github.com/duck8823/duci/domain/model/runner"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"reflect"
	"testing"
)

func TestDefaultDockerRunnerBuilder(t *testing.T) {
	// given
	opts := []cmp.Option{
		cmp.AllowUnexported(runner.Builder{}),
		cmp.Transformer("LogFunc", func(f runner.LogFunc) uintptr {
			return reflect.ValueOf(f).Pointer()
		}),
		cmpopts.IgnoreInterfaces(struct{ docker.Docker }{}),
	}

	// and
	want := &runner.Builder{}
	defer want.SetLogFunc(runner.NothingToDo)()

	// when
	got := runner.DefaultDockerRunnerBuilder()

	// then
	if !cmp.Equal(want, got, opts...) {
		t.Errorf("must be equal, but: %+v", cmp.Diff(want, got, opts...))
	}
}

func TestBuilder_LogFunc(t *testing.T) {
	// given
	opts := []cmp.Option{
		cmp.AllowUnexported(runner.Builder{}),
		cmp.Transformer("LogFunc", func(f runner.LogFunc) uintptr {
			return reflect.ValueOf(f).Pointer()
		}),
		cmpopts.IgnoreInterfaces(struct{ docker.Docker }{}),
	}

	// and
	var wantFunc runner.LogFunc = func(context.Context, job.Log) {}

	// and
	want := &runner.Builder{}
	defer want.SetLogFunc(wantFunc)()

	// and
	sut := &runner.Builder{}

	// when
	got := sut.LogFunc(wantFunc)

	// then
	if !cmp.Equal(want, got, opts...) {
		t.Errorf("must be equal, but: %+v", cmp.Diff(want, got, opts...))
	}

	// and
	gotFunc := sut.GetLogFunc()
	if !cmp.Equal(wantFunc, gotFunc, opts...) {
		t.Errorf("must be equal, but: %+v", cmp.Diff(wantFunc, gotFunc, opts...))
	}

}

func TestBuilder_Build(t *testing.T) {
	// given
	opts := []cmp.Option{
		cmp.AllowUnexported(runner.DockerRunnerImpl{}),
		cmp.Transformer("LogFunc", func(f runner.LogFunc) uintptr {
			return reflect.ValueOf(f).Pointer()
		}),
	}

	// and
	want := &runner.DockerRunnerImpl{}
	defer want.SetLogFunc(func(context.Context, job.Log) {})()

	// and
	sut := &runner.Builder{}

	// when
	got := sut.LogFunc(want.GetLogFunc()).Build()

	// then
	if !cmp.Equal(want, got, opts...) {
		t.Errorf("must be equal, but: %+v", cmp.Diff(want, got, opts...))
	}
}
