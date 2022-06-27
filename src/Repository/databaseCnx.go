package database

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"
	"gopkg.in/yaml.v2"
)

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
func Connect() (*gorm.DB, error) {

	cfg, err := NewConfig("./config/config.prod.yml")
	if err != nil {
		log.Fatal(err)
	}
	dbDriver := cfg.Database.Type
	dbUser := cfg.Database.User
	dbPass := cfg.Database.Password
	dbProjectJson := cfg.Database.DbName
	host := cfg.Database.Host
	port := cfg.Database.Port
	db, err := gorm.Open(dbDriver, dbUser+":"+dbPass+"@tcp("+host+":"+port+")/"+dbProjectJson+"?parseTime=true")
	if err != nil {
		return nil, err
	}
	return db, nil
}
