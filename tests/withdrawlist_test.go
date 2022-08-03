//go:build integration

package tests

import (
	"fmt"
	"github.com/alrund/yp-1-project/internal/application/app"
	"github.com/alrund/yp-1-project/internal/application/usecase"
	"github.com/alrund/yp-1-project/internal/infrastructure/handler"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
)

func SetupTestWithdrawList(a *app.App) {
	db := a.Storage.Connect()
	db.Exec("INSERT INTO users(id, login, password) VALUES ($1, $2, $3)",
		"9a110553-f962-4ab4-ba92-7c3fb7a107f3", "login", "password")
	db.Exec("INSERT INTO users(id, login, password) VALUES ($1, $2, $3)",
		"dbac0532-eaa4-44f8-8845-72be1b25a6ac", "login", "password")
	db.Exec(
		"INSERT INTO withdraws(id, order_number, user_id, sum, processed_at)"+
			" VALUES ($1, $2, $3, $4, $5)",
		uuid.NewString(), "333333", "dbac0532-eaa4-44f8-8845-72be1b25a6ac", 100.1, "2022-08-03T16:50:48Z")
	db.Exec(
		"INSERT INTO withdraws(id, order_number, user_id, sum, processed_at)"+
			" VALUES ($1, $2, $3, $4, $5)",
		uuid.NewString(), "444444", "dbac0532-eaa4-44f8-8845-72be1b25a6ac", 200.2, "2022-08-03T16:55:48Z")
}

func TearDownTestWithdrawList(a *app.App) {
	db := a.Storage.Connect()
	db.Exec("DELETE FROM users")
	db.Exec("DELETE FROM withdraws")
}

func (s *IntegrationTestSuite) TestWithdrawList() {
	SetupTestWithdrawList(s.app)
	defer func() {
		TearDownTestWithdrawList(s.app)
	}()

	tests := []struct {
		name    string
		request request
		want    want
	}{
		{
			name: "success",
			request: request{
				method:             http.MethodGet,
				target:             "/api/user/withdrawals",
				sessionCookieName:  s.app.Config.SessionCookieName,
				sessionCookieValue: "dbac0532-eaa4-44f8-8845-72be1b25a6ac",
				body:               "",
				contentType:        "application/json",
			},
			want: want{
				code: http.StatusOK,
				response: `[
    {
        "order": "333333",
        "sum": 100.1,
        "processed_at": "2022-08-03T16:50:48Z"
    },
    {
        "order": "444444",
        "sum": 200.2,
        "processed_at": "2022-08-03T16:55:48Z"
    }
]`,
				contentType: "application/json",
			},
		},
		{
			name: "fail with not authenticated",
			request: request{
				method:             http.MethodGet,
				target:             "/api/user/withdrawals",
				sessionCookieName:  s.app.Config.SessionCookieName,
				sessionCookieValue: "",
				body:               "",
				contentType:        "application/json",
			},
			want: want{
				code:        http.StatusUnauthorized,
				response:    fmt.Sprintf(`{"warning":"%s"}`, usecase.ErrNotAuthenticated),
				contentType: "application/json",
			},
		},
		{
			name: "fail with withdraw not found",
			request: request{
				method:             http.MethodGet,
				target:             "/api/user/withdrawals",
				sessionCookieName:  s.app.Config.SessionCookieName,
				sessionCookieValue: "9a110553-f962-4ab4-ba92-7c3fb7a107f3",
				body:               "",
				contentType:        "application/json",
			},
			want: want{
				code:        http.StatusNoContent,
				response:    fmt.Sprintf(`{"warning":"%s"}`, usecase.ErrWithdrawNotFound),
				contentType: "application/json",
			},
		},
	}
	t := s.T()
	for _, tt := range tests {
		s.Run(tt.name, func() {
			h := handler.WithdrawListHandler(s.app)

			w := s.MakeTestRequest(tt.request, h)
			res := w.Result()
			defer res.Body.Close()
			resBody, err := io.ReadAll(res.Body)
			if err != nil {
				s.logger.Fatal(err)
			}

			assert.Equal(t, tt.want.code, res.StatusCode)
			if tt.want.contentType == "application/json" {
				assert.JSONEqf(t, tt.want.response, string(resBody), w.Body.String())
			} else {
				assert.Equalf(t, tt.want.response, string(resBody), w.Body.String())
			}
			assert.Equal(t, tt.want.contentType, res.Header.Get("Content-Type"))
		})
	}
}
