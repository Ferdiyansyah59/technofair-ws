package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func main() {
	InitConfig()
	// db := InitDatabase()
	InitDatabase()

	e := echo.New()

	e.GET("/weather", func(c echo.Context) error {

		lat := c.QueryParam("lat")
		long := c.QueryParam("lon")

		nLat, err := strconv.ParseFloat(lat, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, map[string]interface{}{
				"error" : "Latitude mintanya float",
			})
		}

		nLong, err := strconv.ParseFloat(long, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, map[string]interface{}{
				"error" : "Longtitude mintanya float",
			})
		}

		data, err := CurrentWeather(nLat, nLong, Config.Server.AppID)
		if err != nil {
			fmt.Println("Error get current weather ", err)
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"error" : "SERVER LU ERROR COK, COBA LAGI!",
			})
		}

		// model := &weatherModel{}

		// err = db.QueryRowContext(c.Request().Context(), "SELECT * FROM weter WHERE lat = ? AND long = ? ", nLat, nLong).Scan(&model.Long, &model.Long, &model.Temp)



		// _,err = db.ExecContext(c.Request().Context(), "INSERT INTO weter (lat, long, temp) VALUES(?,?,?)", nLat, nLong, data.Main.Temp,)

		// if err != nil {
		// 	return c.JSON(http.StatusInternalServerError, map[string]interface{}{
		// 		"error": "Something wrong bosq",
		// 	})
		// }

		return c.JSON(http.StatusOK, data)
		// return c.JSON(http.StatusOK, map[string]interface{}{
		// 	"message" : "Hello, World!",
		// })
	})

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", Config.Server.Port)))
}


type weatherModel struct {
	Lat float64
	Long float64
	Temp float64
}