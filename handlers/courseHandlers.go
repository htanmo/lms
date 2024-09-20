package handlers

import (
	"context"
	"log"
	"time"

	"github.com/htanmo/lms/config"
	"github.com/htanmo/lms/models"

	"github.com/gofiber/fiber/v2"
)

func GetCourses(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rows, err := config.DBPool.Query(ctx, "SELECT id, title, description, instructor_id FROM courses")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get courses",
		})
	}
	defer rows.Close()

	var courses []models.Course
	for rows.Next() {
		var course models.Course
		err := rows.Scan(&course.ID, &course.Title, &course.Description, &course.InstructorID)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to parse courses",
			})
		}
		courses = append(courses, course)
	}

	return c.JSON(courses)
}

func GetCourseDetails(c *fiber.Ctx) error {
	id := c.Params("id")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var course models.Course
	err := config.DBPool.QueryRow(ctx, "SELECT id, title, description, instructor_id FROM courses WHERE id=$1", id).
		Scan(&course.ID, &course.Title, &course.Description, &course.InstructorID)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Course not found",
		})
	}

	return c.JSON(course)
}

func CreateCourse(c *fiber.Ctx) error {
	var course models.Course
	if err := c.BodyParser(&course); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid input",
		})
	}

	log.Println(course)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := config.DBPool.Exec(ctx, "INSERT INTO courses (title, description, instructor_id) VALUES ($1, $2, $3)",
		course.Title, course.Description, course.InstructorID)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create course",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Course created successfully",
	})
}

func UpdateCourse(c *fiber.Ctx) error {
	id := c.Params("id")
	var course models.Course
	if err := c.BodyParser(&course); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid input",
		})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := config.DBPool.Exec(ctx, "UPDATE courses SET title=$1, description=$2 WHERE id=$3",
		course.Title, course.Description, id)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update course",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Course updated successfully",
	})
}

func DeleteCourse(c *fiber.Ctx) error {
	id := c.Params("id")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := config.DBPool.Exec(ctx, "DELETE FROM courses WHERE id=$1", id)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete course",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Course deleted successfully",
	})
}
