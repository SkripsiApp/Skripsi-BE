package main

import (
	"fmt"
	"skripsi/app/config"
	"skripsi/app/database"
	"skripsi/app/migration"
	"skripsi/app/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	md "skripsi/utils/middleware"
)

func main() {
	e := echo.New()

	e.HTTPErrorHandler = md.ErrorMiddleware

	cfg := config.InitConfig()
	mysql := database.InitDBMysql(cfg)
	migration.InitMigrationMysql(mysql)

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())

	routes.InitRoutes(e, mysql)
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `[${time_rfc3339}] ${status} ${method} ${host}${path} ${latency_human}` + "\n",
	}))

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", cfg.SERVERPORT)))
}
