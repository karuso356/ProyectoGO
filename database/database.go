package database

import (
	"context"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/karuso356/ProjectGoLE/configuracion"
)

func Conec(ctx context.Context, c *configuracion.Configuracion) (*sqlx.DB, error) {
	connectionString := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?parseTime=true", c.BaseDatos.User, c.BaseDatos.Password, c.BaseDatos.Host, c.BaseDatos.Port, c.BaseDatos.Nombre,
	)

	return sqlx.ConnectContext(ctx, "mysql", connectionString)

}
