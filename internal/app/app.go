package app

import (
	"net/http"

	"github.com/null0routed/mariner/internal/config"
)

type App struct {
	Config Config

	Router *http.Server
}

func New(ctx context.Context, cfg Config) (*App, error) {
	db, err := OpenDatabase(cfg.Database)
	if err != nil {
		return nil, err
	}

	_ = db

	mux := http.NewServeMux()
	handler := &http.Server{
		Addr: ":8080",
		Handler: mux,
		ReadTimeout: 120,
		WriteTimeout: 120,
		IdelTimeout: 300,
	}

	return &App{
		Config: cfg,
		Router: handler
	}, nil
}

func (a *App) Run(ctx context.Context) error {
	a.Router.ListenAndServe()
}