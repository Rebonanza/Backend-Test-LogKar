package main

import (
	"context"
	"fmt"
	"os"

	"github.com/local/be-test-logkar/internal/config"
	"github.com/local/be-test-logkar/internal/db"
	"github.com/local/be-test-logkar/internal/redis"
	"github.com/local/be-test-logkar/internal/server"
	"github.com/local/be-test-logkar/internal/user"

	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		fx.Provide(
			config.NewConfig,
			db.NewGormDB,
			redis.NewRedisClient,
			server.NewFiber,
			user.NewUserRepository,
			user.NewUserService,
			user.NewUserHandler,
			product.NewProductRepository,
			product.NewProductService,
			product.NewProductHandler,
			customer.NewCustomerRepository,
			customer.NewCustomerService,
			customer.NewCustomerHandler,
			transaction.NewTransactionRepository,
			transaction.NewTransactionService,
			transaction.NewTransactionHandler,
		),
		fx.Invoke(server.RegisterLifecycle, registerRoutes),
	)

	startCtx, cancel := context.WithTimeout(context.Background(), fx.DefaultTimeout)
	defer cancel()

	if err := app.Start(startCtx); err != nil {
		fmt.Fprintf(os.Stderr, "failed to start: %v\n", err)
		os.Exit(1)
	}

	<-app.Done()

	stopCtx, cancel := context.WithTimeout(context.Background(), fx.DefaultTimeout)
	defer cancel()
	_ = app.Stop(stopCtx)
}

func registerRoutes(app *fiber.App, uh *user.Handler, ph *product.Handler, ch *customer.Handler, th *transaction.Handler) {
	uh.RegisterRoutes(app)
	ph.RegisterRoutes(app)
	ch.RegisterRoutes(app)
	th.RegisterRoutes(app)
}
