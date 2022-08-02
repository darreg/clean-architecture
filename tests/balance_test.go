//go:build integration

package tests

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/alrund/yp-1-project/internal/application/app"
	"github.com/alrund/yp-1-project/internal/infrastructure/handler"
	"github.com/stretchr/testify/assert"
)

func (s *IntegrationTestSuite) TestBalance() {
	type request struct {
		method             string
		target             string
		sessionCookieName  string
		sessionCookieValue string
		body               string
		contentType        string
	}
	type want struct {
		code        int
		response    string
		contentType string
	}
	tests := []struct {
		name    string
		request request
		want    want
	}{
		{
			name: "success",
			request: request{
				method:             http.MethodGet,
				target:             "/api/user/balance",
				sessionCookieValue: "XXX-YYY-ZZZ",
				body:               "",
				contentType:        "application/json; charset=utf-8",
			},
			want: want{
				code:        http.StatusOK,
				response:    `[]`,
				contentType: "application/json; charset=utf-8",
			},
		},
	}
	t := s.T()
	for _, tt := range tests {
		s.Run(tt.name, func() {
			request := httptest.NewRequest(tt.request.method, tt.request.target, strings.NewReader(tt.request.body))
			request.WithContext(context.WithValue(
				request.Context(),
				app.SessionContextKey(tt.request.sessionCookieName),
				tt.request.sessionCookieValue,
			))
			request.Header.Set("Content-type", tt.request.contentType)
			w := httptest.NewRecorder()
			h := handler.BalanceHandler(s.app)
			h.ServeHTTP(w, request)
			res := w.Result()

			assert.Equal(t, tt.want.code, res.StatusCode)

			defer res.Body.Close()
			resBody, err := io.ReadAll(res.Body)
			if err != nil {
				t.Fatal(err)
			}

			assert.Equalf(t, tt.want.response, string(resBody), w.Body.String())
			assert.Equal(t, tt.want.contentType, res.Header.Get("Content-Type"))
		})
	}
}
