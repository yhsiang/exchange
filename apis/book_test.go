package apis

import (
	"net/http"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/valyala/fastjson"
)

func mockDB(t *testing.T) gin.HandlerFunc {
	db, mock, err := sqlmock.New()

	mockSellRows := sqlmock.NewRows([]string{"price", "quantity"}).
		AddRow("2275", "20")

	mock.ExpectQuery("SELECT price, quantity FROM orders WHERE side = 'sell'").
		WithArgs("SBI").
		WillReturnRows(mockSellRows)

	mockBuyRows := sqlmock.NewRows([]string{"price", "quantity"}).
		AddRow("2272.5", "100").
		AddRow("2271", "80").
		AddRow("2271", "20")

	mock.ExpectQuery("SELECT price, quantity FROM orders WHERE side = 'buy'").
		WithArgs("SBI").
		WillReturnRows(mockBuyRows)

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	return func(c *gin.Context) {
		c.Set("DB", db)
		c.Next()
	}
}

func TestQueryOrderBook(t *testing.T) {

	gin.SetMode(gin.TestMode)
	router := gin.New()

	router.Use(mockDB(t))
	router.POST("/book", QueryOrderBook)

	w := performRequest(router, "POST", "/book", `{ "symbol": "SBI" }`)

	assert.Equal(t, http.StatusOK, w.Code)
	var parser fastjson.Parser
	v, err := parser.ParseBytes(w.Body.Bytes())
	assert.NoError(t, err)
	assert.Equal(t, 0, v.GetArray("sell")[0].GetInt("depth"))
	assert.Equal(t, "2275", string(v.GetArray("sell")[0].GetStringBytes("price")))
	assert.Equal(t, "20", string(v.GetArray("sell")[0].GetStringBytes("quantity")))

	assert.Equal(t, 0, v.GetArray("buy")[0].GetInt("depth"))
	assert.Equal(t, "2272.5", string(v.GetArray("buy")[0].GetStringBytes("price")))
	assert.Equal(t, "100", string(v.GetArray("buy")[0].GetStringBytes("quantity")))

	assert.Equal(t, 1, v.GetArray("buy")[1].GetInt("depth"))
	assert.Equal(t, "2271", string(v.GetArray("buy")[1].GetStringBytes("price")))
	assert.Equal(t, "100", string(v.GetArray("buy")[1].GetStringBytes("quantity")))
}
