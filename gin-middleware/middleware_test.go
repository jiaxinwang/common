package middleware

import (
	"os"
	"testing"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/sirupsen/logrus"
)

type Dog struct {
	ID        uint      `gorm:"primary_key" json:"id" lazy:"id" mapstructure:"id"`
	CreatedAt time.Time `json:"created_at" lazy:"created_at" mapstructure:"created_at"`
	UpdatedAt time.Time `json:"updated_at" lazy:"updated_at" mapstructure:"updated_at"`
	Name      string    `json:"name" lazy:"name" mapstructure:"id"`
}

type Profile struct {
	ID        uint      `gorm:"primary_key" lazy:"id" mapstructure:"id"`
	CreatedAt time.Time `json:"created_at" lazy:"created_at" mapstructure:"created_at"`
	UpdatedAt time.Time `json:"updated_at" lazy:"updated_at" mapstructure:"updated_at"`
	Breed     string    `json:"bread" lazy:"breed" mapstructure:"breed"`
	Age       uint      `json:"age" lazy:"age" mapstructure:"age"`
	DogID     uint      `json:"-" lazy:"dog_id" mapstructure:"dog_id"`
}

var gormDB *gorm.DB

func TestMain(m *testing.M) {
	setup()
	defer os.Exit(m.Run())
	defer func() {
		teardown()
	}()

}

func setup() {
	logrus.SetReportCaller(true)
	logrus.SetFormatter(&logrus.TextFormatter{DisableColors: true})
	var err error
	gormDB, err = gorm.Open("sqlite3", "./test.db")
	if err != nil {
		panic(err)
	}

	gormDB.AutoMigrate(&Dog{}, &Profile{})
	gormDB.Create(&Dog{Name: "gooddog"})
	gormDB.Create(&Dog{Name: "baddog"})

	gormDB.Create(&Profile{Breed: "Golden Retriever", Age: 3, DogID: 1})
	gormDB.Create(&Profile{Breed: "Husky", Age: 5, DogID: 2})
}

func teardown() {
	os.Remove("./test.db")
	gormDB.Close()
}
