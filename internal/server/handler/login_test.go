package handler

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/UnLess24/coin/client/internal/database"
	"github.com/gin-gonic/gin"
)

func TestLogin(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("email isn't valid", func(t *testing.T) {
		db := database.NewFake()

		rdr := bytes.NewReader([]byte(`{"email":"test","password":"password"}`))
		req, err := http.NewRequest("POST", "/login", rdr)
		if err != nil {
			t.Fatal(err)
		}

		res := httptest.NewRecorder()
		_, r := gin.CreateTestContext(res)
		r.POST("/login", Login(db, []byte{}))
		r.ServeHTTP(res, req)

		expect := fmt.Sprintf(`{"errorMessage":"%s"}`, database.ErrEmailOrPasswordIsIncorrect)
		if res.Body.String() != expect {
			t.Fatalf("expected %v, got %v", expect, res.Body.String())
		}

		if res.Code != http.StatusBadRequest {
			t.Fatalf("expected %v, got %v", http.StatusBadRequest, res.Code)
		}
	})

	t.Run("password isn't set", func(t *testing.T) {
		db := database.NewFake()

		rdr := bytes.NewReader([]byte(`{"email":"test","password":""}`))
		req, err := http.NewRequest("POST", "/login", rdr)
		if err != nil {
			t.Fatal(err)
		}

		res := httptest.NewRecorder()
		_, r := gin.CreateTestContext(res)
		r.POST("/login", Login(db, []byte{}))
		r.ServeHTTP(res, req)

		expect := fmt.Sprintf(`{"errorMessage":"%s"}`, database.ErrEmailOrPasswordIsIncorrect)
		if res.Body.String() != expect {
			t.Fatalf("expected %v, got %v", expect, res.Body.String())
		}

		if res.Code != http.StatusBadRequest {
			t.Fatalf("expected %v, got %v", http.StatusBadRequest, res.Code)
		}
	})

	t.Run("email isn't set", func(t *testing.T) {
		db := database.NewFake()

		rdr := bytes.NewReader([]byte(`{"email":"","password":"test"}`))
		req, err := http.NewRequest("POST", "/login", rdr)
		if err != nil {
			t.Fatal(err)
		}

		res := httptest.NewRecorder()
		_, r := gin.CreateTestContext(res)
		r.POST("/login", Login(db, []byte{}))
		r.ServeHTTP(res, req)

		expect := fmt.Sprintf(`{"errorMessage":"%s"}`, database.ErrEmailOrPasswordIsIncorrect)
		if res.Body.String() != expect {
			t.Fatalf("expected %v, got %v", expect, res.Body.String())
		}

		if res.Code != http.StatusBadRequest {
			t.Fatalf("expected %v, got %v", http.StatusBadRequest, res.Code)
		}
	})

	t.Run("not return JWT", func(t *testing.T) {
		db := database.NewFake()

		rdr := bytes.NewReader([]byte(`{"email":"test@test.ru","password":"test"}`))

		res := httptest.NewRecorder()
		_, r := gin.CreateTestContext(res)

		req, err := http.NewRequest("POST", "/register", rdr)
		if err != nil {
			t.Fatal(err)
		}
		r.POST("/register", Register(db))
		r.ServeHTTP(res, req)

		rdr = bytes.NewReader([]byte(`{"email":"test@test.ru","password":"test"}`))
		req, err = http.NewRequest("POST", "/login", rdr)
		if err != nil {
			t.Fatal(err)
		}

		res = httptest.NewRecorder()
		r.POST("/login", Login(db, nil))
		r.ServeHTTP(res, req)

		expect := fmt.Sprintf(`{"errorMessage":"%s"}`, errBadCredentials)
		if res.Body.String() != expect {
			t.Fatalf("expected %v, got %v", expect, res.Body.String())
		}

		if res.Code != http.StatusInternalServerError {
			t.Fatalf("expected %v, got %v", http.StatusInternalServerError, res.Code)
		}
	})

	t.Run("empty reader data", func(t *testing.T) {
		db := database.NewFake()

		rdr := bytes.NewReader([]byte{})
		req, err := http.NewRequest("POST", "/login", rdr)
		if err != nil {
			t.Fatal(err)
		}

		res := httptest.NewRecorder()
		_, r := gin.CreateTestContext(res)
		r.POST("/login", Login(db, []byte{}))
		r.ServeHTTP(res, req)

		expect := fmt.Sprintf(`{"errorMessage":"%s"}`, errBadCredentials)
		if res.Body.String() != expect {
			t.Fatalf("expected %v, got %v", expect, res.Body.String())
		}

		if res.Code != http.StatusBadRequest {
			t.Fatalf("expected %v, got %v", http.StatusBadRequest, res.Code)
		}
	})

	t.Run("success", func(t *testing.T) {
		db := database.NewFake()

		rdr := bytes.NewReader([]byte(`{"email":"test@test.ru","password":"test"}`))

		res := httptest.NewRecorder()
		_, r := gin.CreateTestContext(res)

		req, err := http.NewRequest("POST", "/register", rdr)
		if err != nil {
			t.Fatal(err)
		}
		r.POST("/register", Register(db))
		r.ServeHTTP(res, req)

		rdr = bytes.NewReader([]byte(`{"email":"test@test.ru","password":"test"}`))
		req, err = http.NewRequest("POST", "/login", rdr)
		if err != nil {
			t.Fatal(err)
		}

		res = httptest.NewRecorder()
		r.POST("/login", Login(db, []byte("test")))
		r.ServeHTTP(res, req)

		if res.Code != http.StatusOK {
			t.Fatalf("expected %v, got %v", http.StatusOK, res.Code)
		}
	})

	t.Run("context is canceled", func(t *testing.T) {
		db := database.NewFake()

		rdr := bytes.NewReader([]byte(`{"email":"test@test.ru","password":"test"}`))

		res := httptest.NewRecorder()
		_, r := gin.CreateTestContext(res)

		req, err := http.NewRequest("POST", "/register", rdr)
		if err != nil {
			t.Fatal(err)
		}
		r.POST("/register", Register(db))
		r.ServeHTTP(res, req)

		rdr = bytes.NewReader([]byte(`{"email":"test@test.ru","password":"test"}`))
		ctx, cancel := context.WithCancel(context.Background())
		cancel()
		req, err = http.NewRequestWithContext(ctx, "POST", "/login", rdr)
		if err != nil {
			t.Fatal(err)
		}

		res = httptest.NewRecorder()
		r.POST("/login", Login(db, []byte("test")))
		r.ServeHTTP(res, req)

		expect := fmt.Sprintf(`{"errorMessage":"%s"}`, database.ErrContextIsCanceled)
		if res.Body.String() != expect {
			t.Fatalf("expected %v, got %v", expect, res.Body.String())
		}

		if res.Code != http.StatusBadRequest {
			t.Fatalf("expected %v, got %v", http.StatusBadRequest, res.Code)
		}
	})
}
