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

package repositories

import (
	"context"

	"github.com/pkg/errors"
	"github.com/space-code/go-auth/config"
	"github.com/space-code/go-auth/internal/data/contracts"
	"github.com/space-code/go-auth/pkg/model"
	"gorm.io/gorm"
)

// PostgresUserRepository implements the UserRepository interface for PostgreSQL using GORM.
// It provides methods to interact with the user data in a PostgreSQL database.
type PostgresUserRepository struct {
	cfg  *config.Config
	gorm *gorm.DB
}

// NewPostgresUserRepository creates a new instance of PostgresUserRepository.
// It initializes the repository with the provided configuration and GORM database instance.
// It returns a UserRepository interface implementation.
func NewPostgresUserRepository(cfg *config.Config, gorm *gorm.DB) contracts.UserRepository {
	return PostgresUserRepository{cfg: cfg, gorm: gorm}
}

// RegisterUser inserts a new user into the PostgreSQL database.
func (p PostgresUserRepository) RegisterUser(ctx context.Context, user *model.User) (*model.User, error) {
	if err := p.gorm.Create(&user).Error; err != nil {
		return nil, errors.Wrap(err, "error in the inserting user into the database")
	}
	return user, nil
}
