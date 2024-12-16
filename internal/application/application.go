package application

import (
	"fmt"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humachi"
	"github.com/danielgtaylor/huma/v2/humacli"
	"github.com/go-chi/chi/v5"
	v1 "github.com/mendelgusmao/device-manager/internal/domain/devices/interfaces/rest/v1"
	"github.com/mendelgusmao/device-manager/internal/domain/devices/repository"
	"github.com/mendelgusmao/device-manager/internal/domain/devices/services"
	"github.com/mendelgusmao/device-manager/internal/infrastructure/database"
)

type Application struct {
	config Configuration
}

func NewApplication(config Configuration) *Application {
	return &Application{
		config: config,
	}
}

func (a *Application) Run() {
	cli := humacli.New(func(hooks humacli.Hooks, options *Configuration) {
		router := chi.NewMux()
		api := humachi.New(
			router,
			huma.DefaultConfig("Device Management API", "1.0.0"),
		)

		db, err := database.NewSQLiteDatabase(options.DSN)

		if err != nil {
			panic(fmt.Sprintf("%s (%s)", err, options.DSN))
		}

		r := repository.NewDeviceRepository(db)

		if err := r.Setup(); err != nil {
			panic(err)
		}

		s := services.NewDeviceService(r)
		h := v1.NewDeviceHandlers(s)
		h.RegisterRoutes(api)

		hooks.OnStart(func() {
			fmt.Printf("running device manager api revision %s", revision)

			if err := http.ListenAndServe(options.ServerAddress, router); err != nil {
				panic(err)
			}
		})
	})

	cli.Run()
}
