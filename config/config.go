package config

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"os"
	"strconv"
)

type PostgresConfig struct {
	Usr     string
	Pwd     string
	DbName  string
	ConnStr string
	Context *sql.DB
}

type OxygenConfig struct {
	Host  string
	Token string
	TMax  float64
	TMin  float64
}

var postgresConfig PostgresConfig
var oxygenConfig OxygenConfig

func InitConfig() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file – %s", err)
	}

	postgresConfig = PostgresConfig{
		Usr:    os.Getenv("POSTGRES_USR"),
		Pwd:    os.Getenv("POSTGRES_PWD"),
		DbName: os.Getenv("POSTGRES_DB_NAME"),
	}

	postgresConfig.ConnStr = fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", postgresConfig.Usr, postgresConfig.Pwd, postgresConfig.DbName)

	ctx, err := sql.Open("postgres", postgresConfig.ConnStr)
	if err != nil {
		log.Fatalf("Error connecting to DB. – %s", err)
	}

	postgresConfig.Context = ctx

	oxygenConfig = OxygenConfig{
		Host:  os.Getenv("OXYGEN_HOST"),
		Token: os.Getenv("OXYGEN_TOKEN"),
	}
	oxygenConfig.TMax, err = strconv.ParseFloat(os.Getenv("OXYGEN_TMAX"), 64)
	if err != nil {
		log.Fatalf("Error loading .env tmax – %s", err)
	}
	oxygenConfig.TMin, err = strconv.ParseFloat(os.Getenv("OXYGEN_TMIN"), 64)
	if err != nil {
		log.Fatalf("Error loading .env tmin – %s", err)
	}
}

func GetPostgresConfig() PostgresConfig {
	return postgresConfig
}

func GetOxygenConfig() OxygenConfig {
	return oxygenConfig
}
func Clear() {
	if err := postgresConfig.Context.Close(); err != nil {
		panic(err)
	}
}
