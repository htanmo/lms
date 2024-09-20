package handlers

import (
	"context"
	"log"
	"time"

	"github.com/htanmo/lms/config"

	"github.com/gofiber/fiber/v2"
)

func EnrollCourse(c *fiber.Ctx) error {
	userRole, ok := c.Locals("user").(string)
	if !ok || userRole != "student" {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "Unauthorized",
		})
	}

	courseID := c.Params("id")
	studentID, ok := c.Locals("user_id").(float64)
	log.Println("student id", studentID)

	if !ok {
		c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized",
		})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var enrollmentCount int
	err := config.DBPool.QueryRow(ctx, "SELECT COUNT(*) FROM enrollments WHERE course_id=$1 AND student_id=$2", courseID, studentID).Scan(&enrollmentCount)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to check enrollment",
		})
	}
	if enrollmentCount > 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Already enrolled in the course",
		})
	}

	_, err = config.DBPool.Exec(ctx, "INSERT INTO enrollments (course_id, student_id, progress) VALUES ($1, $2, 0)",
		courseID, studentID)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to enroll in course",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Enrolled successfully",
	})
}

func TrackProgress(c *fiber.Ctx) error {
	userRole, ok := c.Locals("user").(string)
	if !ok || userRole != "student" {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "Unauthorized",
		})
	}

	courseID := c.Params("id")
	studentID := c.Locals("user_id").(float64)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var progress int
	err := config.DBPool.QueryRow(ctx, "SELECT progress FROM enrollments WHERE course_id=$1 AND student_id=$2", courseID, studentID).
		Scan(&progress)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Enrollment not found",
		})
	}

	return c.JSON(fiber.Map{
		"course_id": courseID,
		"progress":  progress,
	})
}
