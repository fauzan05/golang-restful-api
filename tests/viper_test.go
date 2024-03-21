package tests

import (
	"fmt"
	"golang-restful-api/helper"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestViperJSON(t *testing.T) {
	config := viper.New()
	config.SetConfigName("config")
	config.SetConfigType("json")
	config.AddConfigPath("..")

	// membaca config
	err := config.ReadInConfig()
	assert.Nil(t, err)

	assert.Equal(t, "golang-restful-api", config.GetString("app.name"))
	assert.Equal(t, "localhost", config.GetString("database.host"))
	assert.Equal(t, "root", config.GetString("database.username"))
}

func TestDatabaseConfigProd(t *testing.T) {
	configDatabaseProd := helper.GetDatabaseConfigProd()
	fmt.Println(configDatabaseProd)
}

func TestENVFile(t *testing.T) {
	config := viper.New()
	config.SetConfigFile("config.env")
	config.AddConfigPath(".")

	err := config.ReadInConfig()
	assert.Nil(t, err)
}