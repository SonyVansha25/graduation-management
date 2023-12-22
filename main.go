package main

import (
	"errors"
	"github.com/afifurrohman-id/kits-cloud-final-project/cmd/graduation"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/html/v2"
	"github.com/joho/godotenv"
	"io"
	"os"
	"path"
)

func init() {
	if os.Getenv("APP_ENV") != "production" {
		if err := godotenv.Load(path.Join("deployments", ".env")); err != nil {
			log.Panic(err)
		}
	}
}

func main() {
	engine := html.New(path.Join("web", "template"), ".go.html")

	loggerFile, err := os.OpenFile(os.Getenv("LOG_PATH"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0200)
	if err != nil {
		log.Panic(err)
	}
	defer loggerFile.Close()

	mw := io.MultiWriter(os.Stdout, loggerFile)

	app := fiber.New(fiber.Config{
		Views:              engine,
		AppName:            "Graduation Management",
		EnableIPValidation: true,
		CaseSensitive:      true,
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError

			// Retrieve the custom status code if it's a *fiber.Error
			fiberErr := new(fiber.Error)
			if errors.As(err, &fiberErr) {
				code = fiberErr.Code
			}

			log.Error(err)
			return ctx.Status(code).Render("pages/error", fiber.Map{
				"code":    code,
				"message": err.Error(),
			})
		},
	})

	app.Use(compress.New(compress.Config{Level: compress.LevelBestSpeed}), favicon.New(), logger.New(logger.Config{Output: mw,
		DisableColors: true}), recover.New())
	app.Static("/public", path.Join("web", "static"))

	app.Get("/monitor", monitor.New(monitor.Config{
		Title: "Graduation Management",
	}))

	app.Get("/", graduation.GetAllStudent)
	app.Get("/:id", graduation.DetailsStudent)
	app.Post("/", graduation.CreateStudent)
	app.Put("/:id", graduation.UpdateStudent)
	app.Delete("/", graduation.DeleteAllStudent)
	app.Delete("/:id", graduation.DeleteStudent)

	if err := app.Listen(":" + os.Getenv("PORT")); err != nil {
		log.Panic(err)
	}
}
