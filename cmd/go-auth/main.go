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

package main

import (
	"github.com/go-playground/validator"
	"github.com/space-code/go-auth/config"
	"github.com/space-code/go-auth/internal/configurations"
	"github.com/space-code/go-auth/internal/data/repositories"
	echoserver "github.com/space-code/go-auth/internal/pkg/echo/server"
	gormpgsql "github.com/space-code/go-auth/internal/pkg/gorm_pgsql"
	"github.com/space-code/go-auth/internal/pkg/http"
	"github.com/space-code/go-auth/server"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Options(
			fx.Provide(
				config.InitConfig,
				http.NewContext,
				echoserver.NewEchoServer,
				validator.New,
				gormpgsql.NewGorm,
				repositories.NewPostgresUserRepository,
			),
			fx.Invoke(server.RunServer),
			fx.Invoke(configurations.ConfigMiddlewares),
			fx.Invoke(configurations.ConfigEndpoints),
			fx.Invoke(configurations.ConfigUsersMediator),
		),
	).Run()
}
