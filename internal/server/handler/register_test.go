package handler

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/UnLess24/coin/client/internal/database"
	"github.com/gin-gonic/gin"
)

func TestRegister(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("invalid request", func(t *testing.T) {
		db := database.NewFake()

		res := httptest.NewRecorder()

		_, r := gin.CreateTestContext(res)
		r.POST("/register", Register(db))
		r.ServeHTTP(res, httptest.NewRequest("POST", "/register", nil))

		expected := fmt.Sprintf(`{"errorMessage":"%s"}`, invalidRequest)
		if res.Body.String() != expected {
			t.Errorf("expected %s, got %s", expected, res.Body.String())
		}

		if res.Code != http.StatusBadRequest {
			t.Errorf("expected status code 400, got %d", res.Code)
		}
	})

	t.Run("invalid email", func(t *testing.T) {
		db := database.NewFake()

		res := httptest.NewRecorder()

		_, r := gin.CreateTestContext(res)
		r.POST("/register", Register(db))

		rdr := bytes.NewReader([]byte(`{"email":"", "password":"password"}`))
		r.ServeHTTP(res, httptest.NewRequest("POST", "/register", rdr))

		expected := fmt.Sprintf(`{"errorMessage":"%s"}`, invalidEmail)
		if res.Body.String() != expected {
			t.Errorf("expected %s, got %s", expected, res.Body.String())
		}

		if res.Code != http.StatusBadRequest {
			t.Errorf("expected status code 400, got %d", res.Code)
		}
	})

	t.Run("invalid email 2", func(t *testing.T) {
		db := database.NewFake()

		res := httptest.NewRecorder()

		_, r := gin.CreateTestContext(res)
		r.POST("/register", Register(db))

		rdr := bytes.NewReader([]byte(`{"email":"test", "password":"password"}`))
		r.ServeHTTP(res, httptest.NewRequest("POST", "/register", rdr))

		expected := fmt.Sprintf(`{"errorMessage":"%s"}`, invalidEmail)
		if res.Body.String() != expected {
			t.Errorf("expected %s, got %s", expected, res.Body.String())
		}

		if res.Code != http.StatusBadRequest {
			t.Errorf("expected status code 400, got %d", res.Code)
		}
	})

	t.Run("success", func(t *testing.T) {
		db := database.NewFake()

		res := httptest.NewRecorder()

		_, r := gin.CreateTestContext(res)
		r.POST("/register", Register(db))

		rdr := bytes.NewReader([]byte(`{"email":"test@test.ru", "password":"password"}`))
		r.ServeHTTP(res, httptest.NewRequest("POST", "/register", rdr))

		expected := ""
		if res.Body.String() != expected {
			t.Errorf("expected %s, got %s", expected, res.Body.String())
		}

		if res.Code != http.StatusCreated {
			t.Errorf("expected status code 201, got %d", res.Code)
		}
	})

	t.Run("user already exists", func(t *testing.T) {
		db := database.NewFake()

		res := httptest.NewRecorder()

		_, r := gin.CreateTestContext(res)
		r.POST("/register", Register(db))

		rdr := bytes.NewReader([]byte(`{"email":"test@test.ru", "password":"password"}`))
		r.ServeHTTP(res, httptest.NewRequest("POST", "/register", rdr))

		res = httptest.NewRecorder()
		rdr = bytes.NewReader([]byte(`{"email":"test@test.ru", "password":"password"}`))
		r.ServeHTTP(res, httptest.NewRequest("POST", "/register", rdr))

		expected := fmt.Sprintf(`{"errorMessage":"%s"}`, database.ErrUserAlreadyExists)
		if res.Body.String() != expected {
			t.Errorf("expected %s, got %s", expected, res.Body.String())
		}

		if res.Code != http.StatusBadRequest {
			t.Errorf("expected status code 400, got %d", res.Code)
		}
	})
}
