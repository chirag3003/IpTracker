package routes

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/oschwald/geoip2-golang"
	"log"
	"net"
)

var App *fiber.App

type homeRouteResponse struct {
	Name string `json:"name"`
}

func Init(a *fiber.App) {
	App = a
	App.Get("/", func(ctx *fiber.Ctx) error {
		db, err := geoip2.Open("db.mmdb")
		if err != nil {
			log.Println(err)
		}
		defer func(db *geoip2.Reader) {
			err := db.Close()
			if err != nil {
				log.Println(err)
			}
		}(db)
		// If you are using strings that may be invalid, check that ip is not nil
		ip := net.ParseIP(ctx.IPs()[0])

		record, err := db.ASN(ip)
		if err != nil {
			log.Println(err)
		}
		log.Println(record)
		res, err := json.Marshal(record)
		if err != nil {
			fmt.Println(err)
			return ctx.SendStatus(500)
		}
		return ctx.SendString(string(res))
	})
}
