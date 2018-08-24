package router

import (
	"github.com/duck8823/duci/application"
	"github.com/duck8823/duci/application/service/github"
	"github.com/duck8823/duci/application/service/log"
	"github.com/duck8823/duci/application/service/runner"
	"github.com/duck8823/duci/infrastructure/docker"
	"github.com/duck8823/duci/infrastructure/git"
	"github.com/duck8823/duci/presentation/controller"
	"github.com/go-chi/chi"
	"github.com/pkg/errors"
	"net/http"
)

func New() (http.Handler, error) {
	logStore, githubService, err := createCommonServices()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	dockerRunner, err := createRunner(logStore, githubService)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	webhooksCtrl := &controller.JobController{Runner: dockerRunner, GitHub: githubService}
	logCtrl := &controller.LogController{LogService: logStore}

	rtr := chi.NewRouter()
	rtr.Post("/", webhooksCtrl.ServeHTTP)
	rtr.Get("/logs/{uuid}", logCtrl.ServeHTTP)

	return rtr, nil
}

func createCommonServices() (log.StoreService, github.Service, error) {
	logStore, err := log.NewStoreService()
	if err != nil {
		return nil, nil, errors.WithStack(err)
	}
	githubService, err := github.NewWithEnv()
	if err != nil {
		return nil, nil, errors.WithStack(err)
	}

	return logStore, githubService, nil
}

func createRunner(logStore log.StoreService, github github.Service) (runner.Runner, error) {
	gitClient, err := git.New(application.Config.Server.SSHKeyPath)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	dockerClient, err := docker.New()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	dockerRunner := &runner.DockerRunner{
		Name:        application.Name,
		BaseWorkDir: application.Config.Server.WorkDir,
		Git:         gitClient,
		GitHub:      github,
		Docker:      dockerClient,
		LogStore:    logStore,
	}

	return dockerRunner, nil
}