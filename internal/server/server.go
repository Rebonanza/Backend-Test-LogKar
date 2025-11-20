package server

import (
	"context"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/local/be-test-logkar/internal/config"
	"go.uber.org/fx"
)

func NewFiber(cfg *config.Config) (*fiber.App, error) {
	app := fiber.New(fiber.Config{
		Prefork:       false,
		CaseSensitive: true,
	})
	return app, nil
}


func RegisterLifecycle(lc fx.Lifecycle, app *fiber.App, cfg *config.Config) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				addr := fmt.Sprintf(":%s", cfg.AppPort)
				_ = app.Listen(addr)
			}()

			select {
			case <-time.After(200 * time.Millisecond):
			case <-ctx.Done():
				return ctx.Err()
			}
			return nil
		},
		OnStop: func(ctx context.Context) error {
			shutdownCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
			defer cancel()
			return app.Shutdown()
			_ = shutdownCtx
		},
	})
}
