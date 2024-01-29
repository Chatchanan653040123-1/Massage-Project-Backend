package main

import (
	"fmt"
	"massage/databases"
	"massage/handlers"
	"massage/logs"
	"massage/repositories"
	"massage/services"
	"strings"
	"time"

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
		logs.Error(err)
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
	//authentication
	app.Post("/register", userHandler.Registers)
	app.Post("/login", userHandler.Login)
	//when entity is authenticated
	authorized := app.Group("/authorized", handlers.JWTAuthen())
	//normal user
	normalUser := authorized.Group("/user", userHandler.UserPermissionLevel1())
	normalUser.Get("/account", userHandler.GetMyAccount)
	normalUser.Put("/update", userHandler.UpdateMyAccount)
	normalUser.Post("/create_group", userHandler.CreateGroup)
	//admin
	normalAdmin := authorized.Group("/permission_level1", userHandler.AdminPermissionLevel1())
	normalAdmin.Get("/get/:uuid", userHandler.GetUser)
	normalAdmin.Get("/getall", userHandler.GetAllUsers)
	//super admin
	superAdmin := authorized.Group("/permission_level2", userHandler.AdminPermissionLevel2())
	superAdmin.Delete("/delete/:uuid", userHandler.DeleteAccount)
	superAdmin.Put("/update/:uuid", userHandler.UpdateAccount)

	logData := fmt.Sprintf("Server is running....\nService: %v\n Ip: %v\n Port: %v\n Date: %v %v %v\n Time: %v:%v:%v", viper.GetString("APP_NAME"), viper.GetString("DB_HOST"), viper.GetInt("APP_PORT"), time.Now().Day(), time.Now().Month(), time.Now().Year(), time.Now().Hour(), time.Now().Minute(), time.Now().Second())
	logs.Info(logData)
	app.Listen(fmt.Sprintf(":%v", viper.GetInt("APP_PORT")))
}
func initConfig() {
	err := godotenv.Load("configs/.env")
	if err != nil {
		logs.Error("Error loading env file")
	}

	viper.SetConfigName("config")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

}
