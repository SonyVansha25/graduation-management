package graduation

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"log"
)

func GetAllStudent(ctx *fiber.Ctx) error {
	db, err := setupDb()
	if err != nil {
		log.Panic(err)
	}

	students := new([]Student)

	if err := db.Find(students).Error; err != nil {
		log.Panic(err)
	}

	return ctx.Render("pages/index", fiber.Map{
		"students": students,
	})
}

func DetailsStudent(ctx *fiber.Ctx) error {
	db, err := setupDb()
	if err != nil {
		log.Panic(err)
	}

	student := new(Student)

	if result := db.First(student, ctx.Params("id")); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return ctx.Status(fiber.StatusNotFound).Render("pages/error", fiber.Map{
				"code":    fiber.StatusNotFound,
				"message": fmt.Sprintf("Student with id %s not found", ctx.Params("id")),
			})
		}
		log.Panic(result.Error)
	}

	return ctx.Render("pages/details", fiber.Map{
		"student": student,
	})
}
