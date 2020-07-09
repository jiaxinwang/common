package middleware

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Dog struct {
	gorm.Model
	Name string
}

type Profile struct {
	gorm.Model
	Age   int
	DogID uint
}

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open("sqlite3", "./before_test.db")
	if err != nil {
		panic(err)
	}
	defer func() {
		db.Close()
		os.Remove("./before_test.db")
	}()

	db.AutoMigrate(&Dog{}, &Profile{})
	db.Create(&Dog{Name: "meat"})
	db.Create(&Profile{Age: 3, DogID: 1})
}

func TestBeforeQuery(t *testing.T) {
	router := beforeRoute()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/dog", nil)
	router.ServeHTTP(w, req)

	// type args struct {
	// 	c *gin.Context
	// }
	// tests := []struct {
	// 	name string
	// 	args args
	// }{
	// 	// TODO: Add test cases.
	// }
	// for _, tt := range tests {
	// 	t.Run(tt.name, func(t *testing.T) {
	// 		BeforeQuery(tt.args.c)
	// 	})
	// }
}

func beforeRoute() *gin.Engine {
	r := gin.Default()
	r.Use(Trace)
	r.Use(LazyResponse)
	r.GET("/dog", func(c *gin.Context) {
		c.Set("ret", map[string]interface{}{"data": nil})
		return
	})

	return r
}
