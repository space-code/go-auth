// Copyright (c) 2024 space-code
//
// Permission is hereby granted, free of charge, to any person obtaining a copy of
// this software and associated documentation files (the "Software"), to deal in
// the Software without restriction, including without limitation the rights to
// use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
// the Software, and to permit persons to whom the Software is furnished to do so,
// subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
// FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
// COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
// IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
// CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package endpoints

import (
	"context"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/mehdihadeli/go-mediatr"
	commandsv1 "github.com/space-code/go-auth/internal/features/registering_user/v1/commands"
	dtosv1 "github.com/space-code/go-auth/internal/features/registering_user/v1/dtos"
	"github.com/space-code/go-auth/internal/middlewares"
)

// MapRoute defines the route for user-related endpoints in the Echo framework.
func MapRoute(validator *validator.Validate, echo *echo.Echo, ctx context.Context) {
	group := echo.Group("api/v1/users")
	group.POST("", createUser(validator, ctx), middlewares.ValidateBearerToken())
}

// createUser handles the HTTP POST request to register a new user.
func createUser(validator *validator.Validate, ctx context.Context) echo.HandlerFunc {
	return func(c echo.Context) error {
		request := dtosv1.RegisterUserRequestDto{}

		if err := c.Bind(request); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err)
		}

		command := commandsv1.NewRegisterUser(request.FirstName, request.LastName, request.UserName, request.Email, request.Password)

		if err := validator.StructCtx(ctx, command); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err)
		}

		result, err := mediatr.Send[*commandsv1.RegisterUser, *dtosv1.RegisterUserResponseDto](ctx, command)

		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err)
		}

		return c.JSON(http.StatusCreated, result)
	}
}
