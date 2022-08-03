//go:build integration

package tests

import (
	"fmt"
	"github.com/alrund/yp-1-project/internal/application/app"
	"github.com/alrund/yp-1-project/internal/application/usecase"
	"github.com/alrund/yp-1-project/internal/infrastructure/handler"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
)

func SetupTestOrderList(a *app.App) {
	db := a.Storage.Connect()
	db.Exec("INSERT INTO users(id, login, password) VALUES ($1, $2, $3)",
		"9a110553-f962-4ab4-ba92-7c3fb7a107f3", "login", "password")
	db.Exec("INSERT INTO users(id, login, password) VALUES ($1, $2, $3)",
		"dbac0532-eaa4-44f8-8845-72be1b25a6ac", "login", "password")
	db.Exec(
		"INSERT INTO orders(number, user_id, status, accrual, uploaded_at, processed_at)"+
			" VALUES ($1, $2, $3, $4, $5, $6)",
		"333333", "dbac0532-eaa4-44f8-8845-72be1b25a6ac", 3, 100.1, "2022-08-03T16:50:48Z", "2022-08-03T16:50:48Z")
	db.Exec(
		"INSERT INTO orders(number, user_id, status, accrual, uploaded_at, processed_at)"+
			" VALUES ($1, $2, $3, $4, $5, $6)",
		"444444", "dbac0532-eaa4-44f8-8845-72be1b25a6ac", 3, 200.2, "2022-08-03T17:50:48Z", "2022-08-03T17:50:48Z")
}

func TearDownTestOrderList(a *app.App) {
	db := a.Storage.Connect()
	db.Exec("DELETE FROM users")
	db.Exec("DELETE FROM orders")
}

func (s *IntegrationTestSuite) TestOrderList() {
	SetupTestOrderList(s.app)
	defer func() {
		TearDownTestOrderList(s.app)
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
				target:             "/api/user/orders",
				sessionCookieName:  s.app.Config.SessionCookieName,
				sessionCookieValue: "dbac0532-eaa4-44f8-8845-72be1b25a6ac",
				body:               "",
				contentType:        "application/json",
			},
			want: want{
				code: http.StatusOK,
				response: `[
		   {
		       "accrual": 100.1,
		       "number": "333333",
		       "uploaded_at": "2022-08-03T16:50:48Z",
		       "processed_at": "2022-08-03T16:50:48Z",
		       "status": "PROCESSED"
		   },
		   {
		       "accrual": 200.2,
		       "number": "444444",
		       "uploaded_at": "2022-08-03T17:50:48Z",
		       "processed_at": "2022-08-03T17:50:48Z",
		       "status": "PROCESSED"
		   }
		]`,
				contentType: "application/json",
			},
		},
		{
			name: "fail with not authenticated",
			request: request{
				method:             http.MethodGet,
				target:             "/api/user/orders",
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
			name: "fail with order not found",
			request: request{
				method:             http.MethodGet,
				target:             "/api/user/orders",
				sessionCookieName:  s.app.Config.SessionCookieName,
				sessionCookieValue: "9a110553-f962-4ab4-ba92-7c3fb7a107f3",
				body:               "",
				contentType:        "application/json",
			},
			want: want{
				code:        http.StatusNoContent,
				response:    fmt.Sprintf(`{"warning":"%s"}`, usecase.ErrOrderNotFound),
				contentType: "application/json",
			},
		},
	}
	t := s.T()
	for _, tt := range tests {
		s.Run(tt.name, func() {
			h := handler.OrderListHandler(s.app)

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
