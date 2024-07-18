package main

import (
	"os"
	"wan-api-verify-user/AppConfig/Config"
	"wan-api-verify-user/Controller"
	"wan-api-verify-user/Data"
	"wan-api-verify-user/Service"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	// * Load environemnt file and start connection to database
	env := ".env"
	if len(os.Args) > 1 {
		env = os.Args[1]
	}

	Config.SetEnvironment(env)
	config, err := Config.LoadConfig()
	if err != nil {
		e.Logger.Fatal("Error loading .env file")
	}

	db_data, _ := Config.Connect(config)

	dataLayer := Data.NewKolDataLayer(db_data)
	service := Service.NewKOLService(dataLayer)
	Controller.NewKOLController(e, service)

	e.Logger.Info(e.Start(":8081"))
}
