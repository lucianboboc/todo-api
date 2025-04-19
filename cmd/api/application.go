package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/lucianboboc/todo-api/config"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type Application struct {
	mux    *http.ServeMux
	config *config.Config
	logger *slog.Logger
}

func newApplication(mux *http.ServeMux, config *config.Config, logger *slog.Logger) *Application {
	return &Application{
		mux:    mux,
		config: config,
		logger: logger,
	}
}

func (a *Application) createServer() http.Server {
	return http.Server{
		Addr:         fmt.Sprintf(":%d", a.config.Port),
		Handler:      a.mux,
		IdleTimeout:  time.Minute,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 20,
		ErrorLog:     slog.NewLogLogger(a.logger.Handler(), slog.LevelError),
	}
}

func (a *Application) Start() {
	s := a.createServer()

	go func() {
		a.logger.Info("Starting server", slog.String("addr", s.Addr))
		err := s.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			a.logger.Error("Server error", slog.String("error", err.Error()))
			os.Exit(1)
		}
	}()

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer cancel()
	<-ctx.Done()
	a.logger.Info("Received shutdown signal")

	shutdownContext, shutdownCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer shutdownCancel()

	if err := s.Shutdown(shutdownContext); err != nil {
		a.logger.Error("Server shutdown error", slog.String("error", err.Error()))
	} else {
		a.logger.Info("Server shutdown gracefully")
	}
}
