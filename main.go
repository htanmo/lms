package main

import (
	"log"

	"github.com/htanmo/lms/config"
	"github.com/htanmo/lms/handlers"
	"github.com/htanmo/lms/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	config.ConnectDB()

	app := fiber.New()

	app.Use(logger.New())

	api := app.Group("/api")

	// v1
	v1 := api.Group("/v1")

	// Public routes
	public := v1.Group("/public")
	public.Post("/register", handlers.Register)
	public.Post("/login", handlers.Login)
	public.Get("/courses", handlers.GetCourses)
	public.Get("/courses/:id", handlers.GetCourseDetails)

	// Student-only routes
	student := v1.Group("/student", middleware.JWTMiddleware("student"))
	student.Post("/courses/:id/enroll", handlers.EnrollCourse)
	student.Get("/courses/:id/progress", handlers.TrackProgress)

	// Protected routes (Admin and Instructor only)
	adminInstructor := v1.Group("/admin", middleware.JWTMiddleware("admin", "instructor"))
	adminInstructor.Post("/courses", handlers.CreateCourse)
	adminInstructor.Put("/courses/:id", handlers.UpdateCourse)
	adminInstructor.Delete("/courses/:id", handlers.DeleteCourse)

	log.Fatal(app.Listen(":8000"))
}
