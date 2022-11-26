package utils

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

func HandleSingleFIle(ctx *fiber.Ctx) error {

	// Handle File
	file, errFile := ctx.FormFile("file")

	if errFile != nil {
		log.Println("Error file =", errFile)
	}

	var filename *string

	if file != nil {
		filename = &file.Filename

		errSaveFile := ctx.SaveFile(file, fmt.Sprintf("./public/images/%s", *filename))

		if errSaveFile != nil {
			log.Println("Failed to store into public/images directory.")
		}

	} else {
		log.Println("Nothing file to be upload.")

	}

	if filename != nil {
		ctx.Locals("filename", *filename)

	} else {
		ctx.Locals("filename", nil)
	}

	return ctx.Next()

}
