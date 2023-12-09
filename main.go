package main

import (
	"fmt"
	"log"
	"massage/databases"
	"massage/handlers"
	"massage/logs"
	"massage/repositories"
	"massage/services"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

func main() {
	initConfig()
	db, err := databases.CreateDB()
	if err != nil {
		panic(err)
	}
	databases.AutoMigrate(db)
	userRepositoryDB := repositories.NewUserRespositoryDB(db)
	userService := services.NewUserService(userRepositoryDB)
	userHandler := handlers.NewUserHandler(userService)
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-Max",
	}))

	app.Use(logger.New(logger.Config{
		CustomTags: map[string]logger.LogFunc{
			"port": func(output logger.Buffer, c *fiber.Ctx, data *logger.Data, extraParam string) (int, error) {
				return output.WriteString(viper.GetString("app.port"))
			},
			"msg": func(output logger.Buffer, c *fiber.Ctx, data *logger.Data, extraParam string) (int, error) {
				message := fmt.Sprintf("Response body: %s", c.Response().Body())
				return output.WriteString(message)
			},
		},
		Format:     "[${time}] Status: ${status} - Medthod: ${method} API: ${path} Port:${port}\n Message: ${msg}\n\n",
		TimeFormat: "2 Jan 2006 15:04:05",
		TimeZone:   "Asia/Bangkok",
	}))
	app.Post("/api/register", userHandler.Registers)
	app.Post("/api/login", userHandler.Login)
	authorized := app.Group("/api", handlers.JWTAuthen())
	authorized.Get("/users", userHandler.GetAllUsers)

	logs.Info("Product service started at port " + viper.GetString("APP_PORT"))
	app.Listen(fmt.Sprintf(":%v", viper.GetInt("APP_PORT")))
}
func initConfig() {
	err := godotenv.Load("configs/config.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	viper.SetConfigName("config")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

}
