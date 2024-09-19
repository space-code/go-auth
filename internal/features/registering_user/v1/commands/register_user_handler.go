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

package commands

import (
	"context"

	"github.com/space-code/go-auth/internal/data/contracts"
	"github.com/space-code/go-auth/internal/features/registering_user/v1/dtos"
	"github.com/space-code/go-auth/internal/pkg/utils"
	"github.com/space-code/go-auth/pkg/model"
)

// RegisterUserHandler handles the registration of a new user.
type RegisterUserHandler struct {
	userRepository contracts.UserRepository
	ctx            context.Context
}

// NewRegisterUserHandler creates a new instance of RegisterUserHandler.
// It returns a pointer to the RegisterUserHandler.
func NewRegisterUserHandler(userRepository contracts.UserRepository, ctx context.Context) *RegisterUserHandler {
	return &RegisterUserHandler{userRepository: userRepository, ctx: ctx}
}

// Handle processes the registration command for a new user.
func (c *RegisterUserHandler) Handle(ctx context.Context, command *RegisterUser) (*dtos.RegisterUserResponseDto, error) {
	pass, err := utils.HashPassword(command.Password)

	if err != nil {
		return nil, err
	}

	product := &model.User{
		Email:     command.Email,
		Password:  pass,
		UserName:  command.UserName,
		LastName:  command.LastName,
		FirstName: command.FirstName,
		CreatedAt: command.CreatedAt,
	}

	registeredUser, err := c.userRepository.RegisterUser(ctx, product)

	if err != nil {
		return nil, err
	}

	response := &dtos.RegisterUserResponseDto{
		UserID:    registeredUser.UserID,
		FirstName: registeredUser.FirstName,
		LastName:  registeredUser.LastName,
		UserName:  registeredUser.UserName,
		Email:     registeredUser.Email,
		Password:  registeredUser.Password,
	}

	return response, nil
}
