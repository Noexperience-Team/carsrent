package database

import (
	"fmt"
	"log"
	"os"

	model "github.com/Noexperience-Team/carsrent/src/Models"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"gopkg.in/yaml.v2"
)

var DB *gorm.DB

type DBConfig struct {
	Database struct {
		Type     string `yaml:"type"`
		Port     string `yaml:"port"`
		Host     string `yaml:"host"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		DbName   string `yaml:"dbName"`
	}
}

// get the data base configuration from the config file and parset to the database config struct
func NewConfig(configPath string) (*DBConfig, error) {
	// Create config structure
	config := &DBConfig{}

	// Open config file
	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Init new YAML decode
	d := yaml.NewDecoder(file)

	// Start YAML decoding from file
	if err := d.Decode(&config); err != nil {
		return nil, err
	}

	return config, nil
}

// func Connect to the data base it create a connection with the data base
// pass database variables into func in order to use it anywhere
func Connect(pathToConfig string) {

	cfg, err := NewConfig(pathToConfig)
	if err != nil {
		log.Fatal(err)
	}
	dbDriver := cfg.Database.Type
	dbUser := cfg.Database.User
	dbPass := cfg.Database.Password
	dbProjectJson := cfg.Database.DbName
	host := cfg.Database.Host
	port := cfg.Database.Port
	connection, err := gorm.Open(dbDriver, dbUser+":"+dbPass+"@tcp("+host+":"+port+")/"+dbProjectJson+"?parseTime=true")
	if err != nil {
		panic(err)
	}
	DB = connection
	connection.AutoMigrate(&model.Car{})
	//connection.Model(&model.Rent{}).AddForeignKey("car_id", "cars(id)", "CASCADE", "CASCADE")
	fmt.Println("Creating tabels ")

}
