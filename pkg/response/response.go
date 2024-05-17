package response

import (
	"errors"
	"github.com/dedetia/godate/pkg/response/custerr"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type JSONResponse struct {
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message,omitempty"`
	Data       interface{} `json:"data,omitempty"`
}

func Success(ctx *fiber.Ctx, code int, data ...interface{}) error {
	hte := JSONResponse{
		StatusCode: code,
		Message:    http.StatusText(code),
	}

	if len(data) > 0 {
		hte.Data = data[0]
	}

	ctx.Status(code)
	return ctx.JSON(hte)
}

func Error(ctx *fiber.Ctx, err *fiber.Error, message ...string) error {
	res := JSONResponse{
		StatusCode: err.Code,
		Message:    err.Message,
	}

	if len(message) > 0 {
		res.Message = message[0]
	}

	ctx.Status(err.Code)
	return ctx.JSON(res)
}

func AuthError(ctx *fiber.Ctx, err error) error {
	var he *custerr.HttpError
	ok := errors.As(err, &he)
	if ok {
		res := JSONResponse{
			StatusCode: he.Code,
			Message:    he.Error(),
		}
		ctx.Status(he.Code)
		return ctx.JSON(res)
	}

	res := JSONResponse{
		StatusCode: http.StatusInternalServerError,
		Message:    err.Error(),
	}

	ctx.Status(http.StatusInternalServerError)
	return ctx.JSON(res)
}
