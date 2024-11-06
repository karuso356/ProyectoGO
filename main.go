package main

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/karuso356/ProjectGoLE/configuracion"
	"github.com/karuso356/ProjectGoLE/database"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		fx.Provide(
			context.Background,
			configuracion.Inicio,
			database.Conec,
		),
		fx.Invoke(
			func(db *sqlx.DB) {
				_, err := db.Query("SELECT * FROM USER")
				if err != nil {
					panic(err)
				}
			},
		),
	)

	app.Run()

}
