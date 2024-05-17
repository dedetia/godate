package handler

import (
	"github.com/dedetia/godate/internal/core/domain"
	"github.com/dedetia/godate/pkg/response"
	"github.com/dedetia/godate/shared/constant"
	"github.com/gofiber/fiber/v2"
	"os"
	"path/filepath"
)

func (h *Handler) Login(c *fiber.Ctx) error {
	var request domain.LoginRequest

	err := c.BodyParser(&request)
	if err != nil {
		return response.Error(c, fiber.ErrBadRequest, err.Error())
	}

	res, err := h.service.GetUserService().Login(c.Context(), &request)
	if err != nil {
		return response.AuthError(c, err)
	}

	return response.Success(c, fiber.StatusOK, res)
}

func (h *Handler) Signup(c *fiber.Ctx) error {
	var request domain.SignupRequest

	if err := c.BodyParser(&request); err != nil {
		return response.Error(c, fiber.ErrBadRequest, err.Error())
	}

	form, err := c.MultipartForm()
	if err != nil {
		return response.Error(c, fiber.ErrBadRequest, err.Error())
	}

	photoDir := os.Getenv("PHOTO_DIR")
	files := form.File["photos"]

	mFile := make(map[string]bool)
	for _, file := range files {
		if file.Size > constant.MaxFileSize {
			return response.Error(c, fiber.ErrBadRequest, "file size exceed maximum")
		}

		fileName := filepath.Base(file.Filename)

		if _, exists := mFile[fileName]; !exists {
			filePath := filepath.Join(photoDir, fileName)

			if err = c.SaveFile(file, filePath); err != nil {
				return response.Error(c, fiber.ErrInternalServerError, err.Error())
			}

			request.Photos = append(request.Photos, &domain.File{
				Name: fileName,
				Path: filePath,
			})

			mFile[fileName] = true
		}
	}

	res, err := h.service.GetUserService().Signup(c.Context(), &request)
	if err != nil {
		return response.AuthError(c, err)
	}

	return response.Success(c, fiber.StatusOK, res)
}

func (h *Handler) Profile(c *fiber.Ctx) error {
	var request domain.ProfileRequest

	err := c.QueryParser(&request)
	if err != nil {
		return response.Error(c, fiber.ErrBadRequest, err.Error())
	}

	if request.Page <= 0 {
		request.Page = 1
	}

	res, err := h.service.GetUserService().Profile(c.Context(), &request)
	if err != nil {
		return response.AuthError(c, err)
	}

	return response.Success(c, fiber.StatusOK, res)
}
