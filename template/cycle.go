package main

import (
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

// Cycle is the component in charge of the life cycle of the application
// It is responsible for starting quickly your app and shutting it down gracefully

type Cycle interface {
	Setup() error
	Ignite() error
	Stop() error
}

type cycle struct {
	repositories []infrastructure.Repository
	applications []application.Application
	shutdown     chan os.Signal
}

func NewCycle() *cycle {
	return &cycle{
		shutdown:     make(chan os.Signal, 1),
		repositories: []infrastructure.Repository{},
		applications: []application.Application{},
	}
}

func (c *cycle) AddRepository(repository infrastructure.Repository) {
	c.repositories = append(c.repositories, repository)
}

func (c *cycle) AddApplications(applications []application.Application) {
	c.applications = append(c.applications, applications...)
}

func (c *cycle) AddRepositories(repository []infrastructure.Repository) {
	c.repositories = append(c.repositories, repository...)
}

func (c *cycle) AddApplication(application application.Application) {
	c.applications = append(c.applications, application)
}

func (c *cycle) Setup() error {
	if len(c.repositories) == 0 {
		slog.Info("No repository to setup")
		return nil
	}
	for _, repository := range c.repositories {
		if err := repository.Setup(); err != nil {
			return err
		}
	}
	return nil
}

// Ignite starts the application
func (c *cycle) Ignite() error {
	if len(c.applications) == 0 {
		slog.Info("No application to start")
		return nil
	}

	for _, application := range c.applications {
		go func(application application.Application) {
			if err := application.Ignite(); err != nil {
				slog.Error(err)
			}
		}(application)
	}

	signal.Notify(c.shutdown, os.Interrupt, syscall.SIGTERM)

	_ = <-c.shutdown

	_ = c.Stop()

	
	return nil
}

// Stop stops the application
func (c *cycle) Stop() error {
	slog.Info("Stopping applications with graceful shutdown")
	close(c.shutdown)
	for _, application := range c.applications {
		if err := application.Stop(); err != nil {
			return err
		}
	}
	return nil
}
