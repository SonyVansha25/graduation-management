package graduation

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"time"
)

func CreateStudent(ctx *fiber.Ctx) error {
	db, err := setupDb()
	if err != nil {
		log.Panic(err)
	}

	student := new(Student)
	if err := ctx.BodyParser(student); err != nil {
		log.Panic(err)
	}

	student.Passed = verifyPassingGrade(student.FinalGrade)

	student.CreatedAt = formatDate(time.Now().UnixMilli())

	if err := db.Create(student).Error; err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	return ctx.Status(fiber.StatusCreated).JSON(student)
}

func UpdateStudent(ctx *fiber.Ctx) error {
	db, err := setupDb()
	if err != nil {
		log.Panic(err)
	}

	student := new(Student)
	if err := ctx.BodyParser(student); err != nil {
		log.Panic(err)
	}

	if err := db.Model(student).Where("id = ?", ctx.Params("id")).Updates(map[string]any{"name": student.Name,
		"final_grade": student.FinalGrade, "passed": verifyPassingGrade(student.FinalGrade), "updated_at": formatDate(time.Now().UnixMilli()), "school_cohort_of": student.SchoolCohortOf}).Error; err != nil {

		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	return ctx.Status(fiber.StatusOK).JSON(student)
}

func DeleteStudent(ctx *fiber.Ctx) error {
	db, err := setupDb()
	if err != nil {
		log.Panic(err)
	}

	student := new(Student)
	if err := db.Where("id = ?", ctx.Params("id")).Delete(student).Error; err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}

func DeleteAllStudent(ctx *fiber.Ctx) error {
	db, err := setupDb()
	if err != nil {
		log.Panic(err)
	}

	if err := db.Where("created_at IS NOT NULL").Delete(new(Student)).Error; err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}
