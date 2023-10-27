package server

import (
	"fmt"
	"madyasantosa/ruangkegiatan/config"
	reservationHandlerPkg "madyasantosa/ruangkegiatan/features/reservation/handler"
	resevationRepositoryPkg "madyasantosa/ruangkegiatan/features/reservation/repository"
	reservationServicePkg "madyasantosa/ruangkegiatan/features/reservation/service"
	roomHandlerPkg "madyasantosa/ruangkegiatan/features/rooms/handler"
	roomRepositoryPkg "madyasantosa/ruangkegiatan/features/rooms/repository"
	roomServicePkg "madyasantosa/ruangkegiatan/features/rooms/service"
	userHandlerPkg "madyasantosa/ruangkegiatan/features/users/handler"
	userRepositoryPkg "madyasantosa/ruangkegiatan/features/users/repository"
	userServicePkg "madyasantosa/ruangkegiatan/features/users/service"
	"madyasantosa/ruangkegiatan/pkg/database"
	"madyasantosa/ruangkegiatan/pkg/s3"
	"madyasantosa/ruangkegiatan/routes"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitServer() error {
	config := *config.InitConfig()

	validate := validator.New()

	db := database.InitDB(config)
	database.Migrate(db)

	s3.NewUploader(config)

	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	userRepository := userRepositoryPkg.NewUserRepository(db)
	userService := userServicePkg.NewUserService(userRepository, validate)
	userHandler := userHandlerPkg.NewUserHandler(userService)

	roomRespository := roomRepositoryPkg.NewRoomRespository(db)
	roomService := roomServicePkg.NewRoomService(roomRespository, validate)
	roomHandler := roomHandlerPkg.NewRoomHandler(roomService)

	reservationRepository := resevationRepositoryPkg.NewReservationRepository(db)
	reservationService := reservationServicePkg.NewReservationService(reservationRepository, validate)
	reservationHandler := reservationHandlerPkg.NewReservationHanlder(reservationService)

	e.GET("/", func(c echo.Context) error {
		return c.String(200, "PORT: "+config.AppPort)
	})

	routes.UserRoutes(e, userHandler)
	routes.RoomRoutes(e, roomHandler)
	routes.ReservationRoutes(e, reservationHandler)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", config.AppPort)))

	return nil
}
