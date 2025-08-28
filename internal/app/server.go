package app

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	httpSwagger "github.com/swaggo/http-swagger"

	"order-service/internal/app/logger"
	"order-service/internal/controllers/handler"
	"order-service/internal/lib/logger/sl"
	"order-service/internal/service/order"
)

const (
	readTimeout  time.Duration = 5 * time.Second
	writeTimeout time.Duration = 15 * time.Second
	idleTimeout  time.Duration = 60 * time.Second
)

const Addr = ":8081"

func setupServer(service *order.OrderService) *http.Server {
	r := chi.NewRouter()
	handler := handler.New(service)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*", "http://*", "null"},
		AllowedMethods: []string{"GET"},
		AllowedHeaders: []string{"Accept", "Content-Type"},
		MaxAge:         300,
	}))

	r.Route("/order", func(r chi.Router) {
		r.Get("/{order_uid}", handler.GetOrder)
	})
	r.Get("/swagger/*", httpSwagger.WrapHandler)
	return &http.Server{
		Addr:         Addr,
		Handler:      r,
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
		IdleTimeout:  idleTimeout,
	}
}

func (a *Application) runServer() {
	logger.Logger.Info("HTTP listening on %s", a.srv.Addr)
	if err := a.srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		logger.Logger.Error("listen: ", sl.Err(err))
	}
}

func (a *Application) shutdownServer() error {
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	return a.srv.Shutdown(shutdownCtx)
}
