package controllers

import (
	"github.com/andritroops/go-latihan/config"
	"github.com/andritroops/go-latihan/models/entity"
	"github.com/andritroops/go-latihan/models/request"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func CategoryStore(ctx *fiber.Ctx) error {

	category := new(request.CategoryCreateRequest)

	if err := ctx.BodyParser(category); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})

	}

	var errors []*request.ErrorResponse
	validate := validator.New()
	errValidate := validate.Struct(category)

	if errValidate != nil {

		for _, err := range errValidate.(validator.ValidationErrors) {
			var element request.ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Field()
			errors = append(errors, &element)
		}

		return ctx.Status(fiber.StatusBadRequest).JSON(errors)
	}

	// var filenameString string

	filenames := ctx.Locals("filenames")

	if filenames == nil {

		return ctx.Status(422).JSON(fiber.Map{
			"message": "File is required",
		})
	} else {
		// filenameString = fmt.Sprintf("%v", filenames)

		filenamesData := filenames.([]string)

		for _, filename := range filenamesData {

			newCategory := entity.Category{
				Name: category.Name,
				File: filename,
			}

			errCreateCategory := config.DB.Create(&newCategory).Error

			if errCreateCategory != nil {

				return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"message": errCreateCategory,
				})
			}
		}
	}

	// log.Println(filenameString)

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success",
	})

}
