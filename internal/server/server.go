package server

import (
	"Short/internal/config"
	"Short/internal/db"
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/sirupsen/logrus"
)

type HTTPService struct {
	server *http.Server
	config *config.AppConfig
	logger *logrus.Logger
	router *chi.Mux
	db     *db.DB
}

func New(ctx context.Context, cfg *config.AppConfig) *HTTPService {
	appLogger := logrus.New()
	appRouter := chi.NewRouter()
	database := db.New(cfg)
	appServer := &http.Server{
		Addr:    cfg.Server.Bind,
		Handler: appRouter,
	}

	server := &HTTPService{
		server: appServer,
		config: cfg,
		logger: appLogger,
		router: appRouter,
		db:     database,
	}

	if err := server.configureServer(); err != nil {
		logrus.Error(err.Error())
		ctx.Done()

		return nil // ???
	}

	return server
}

func (s *HTTPService) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.Logger.Level)
	if err != nil {
		return err
	}

	s.logger.SetLevel(level)

	return nil
}

func (s *HTTPService) configureRouter() {
	// A good base middleware stack
	s.router.Use(middleware.RequestID)
	s.router.Use(middleware.RealIP)
	s.router.Use(middleware.Logger)
	s.router.Use(middleware.Recoverer)

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	s.router.Use(middleware.Timeout(60 * time.Second))

	s.router.Route("/", func(r chi.Router) {
		r.With(s.dataValidationMiddleware).Post("/create", s.HandleCreate())
	})
}

func (s *HTTPService) configureServer() error {
	if err := s.configureLogger(); err != nil {
		return err
	}

	s.configureRouter()

	return nil
}

func (s *HTTPService) Shutdown(shutdownCtx context.Context) error {
	return s.server.Shutdown(shutdownCtx)
}

func (s *HTTPService) Run(ctx context.Context, cancel context.CancelFunc) {
	go func() {
		s.logger.Info(fmt.Sprintf("Start httpserver on %s!", s.config.Server.Bind))

		err := s.server.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			s.logger.Error("error serve s httpserver", logrus.Fields{"error": err})
			cancel()
		}
	}()

	<-ctx.Done()

	s.logger.Info(fmt.Sprintf("Shutdown http httpserver on %s!", s.config.Server.Bind))

	err := s.server.Shutdown(context.Background())
	if err != nil {
		s.logger.Error("error shutdown http httpserver", logrus.Fields{"error": err})
	}
}
