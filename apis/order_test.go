package apis

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/valyala/fastjson"
)

func performRequest(r http.Handler, method, path, body string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, strings.NewReader(body))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func mockDatabase(t *testing.T) gin.HandlerFunc {
	db, mock, err := sqlmock.New()

	mock.ExpectExec("^INSERT INTO orders").WillReturnResult(sqlmock.NewResult(1, 1))

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	return func(c *gin.Context) {
		c.Set("DB", db)
		c.Next()
	}
}

func TestIndex(t *testing.T) {
	body := `It works`
	gin.SetMode(gin.TestMode)
	router := gin.New()

	router.GET("/", Index)

	w := performRequest(router, "GET", "/", "")

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, body, w.Body.String())
}

func TestCreateOrder(t *testing.T) {

	gin.SetMode(gin.TestMode)
	router := gin.New()

	router.Use(mockDatabase(t))
	router.POST("/order", CreateOrder)

	w := performRequest(router, "POST", "/order", `{ "symbol": "SBI", "side": "sell", "price": "2275.5", "quantity":"100" }`)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.True(t, fastjson.Exists(w.Body.Bytes(), "order_id"))
}
