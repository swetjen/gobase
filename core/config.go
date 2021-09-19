package core

import (
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"reflect"
)
import "log"

type Config struct {
	Environment string `env:"ENVIRONMENT" default:"dev"`
	AdminEmail  string `env:"ADMIN_EMAIL" default:"spence@rig.hteo.us"`
	ProjectName string `env:"PROJECT_NAME" default:"My Go Project"`
	ServerName  string `env:"SERVER_NAME" default:"dev1"`
	ServerPort  string `env:"SERVER_PORT" default:":3000"`
	ApiV1Prefix string `env:"API_V1_PREFIX" default:"/api/v1"`

	FirstSuperUser         string `env:"FIRST_SUPER_USER" default:"spence@rig.hteo.us"`
	FirstSuperUserPassword string `env:"FIRST_SUPER_USER_PASSWORD" default:"righteous"`

	MongoUri string `env:"MONGO_URI"`
	Db       string `env:"DB" default:"test"`
}

// NewConfig return a new Config, first with environmental variables and then with defaults
// if an env variable isn't found.  Panics on required values if missing.
func NewConfig() Config {
	config := Config{}

	err := godotenv.Load()
	if err != nil {
		log.Println("No .ENV file in project root")
	} else {
		log.Printf("LOADED .ENV")
	}

	dict := make(map[string]string)

	fields := reflect.VisibleFields(reflect.TypeOf(struct{ Config }{}))
	for count, field := range fields {
		if count != 0 {
			v := os.Getenv(field.Tag.Get("env"))
			if v != "" {
				// Set the Environmental value
				dict[field.Name] = v
			} else {
				// Set the Default value
				dict[field.Name] = field.Tag.Get("default")
			}
		}
	}

	jsonbody, err := json.Marshal(dict)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	if err := json.Unmarshal(jsonbody, &config); err != nil {
		fmt.Println(err)
		panic(err)
	}

	// Required Values
	if config.Db == "" {
		log.Fatalln("Error:  Config.Db=='' but expected a name")
	}

	if config.MongoUri == "" {
		log.Fatalln("Error:  Config.MongoUri=='' but expected resource URI.")
	}

	log.Println("==== SETTINGS =============================================")
	log.Printf("%#v\n", config)
	log.Println("===========================================================")

	return config
}

var Settings = NewConfig()
