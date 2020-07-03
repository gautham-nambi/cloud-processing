package service

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"parallelSystems/user_gateway/cmd/db"
	"parallelSystems/user_gateway/cmd/service/resources"
)

func initPostgresService() {
	var configuration = getConfiguration()
	conn, err := sqlx.Connect("postgres", configuration.GetPostgresConnectionString())
	if err != nil {
		panic(err.Error())
	}
	db.SetConnection(conn)
	logger.Log("PostgreSQL connection established successfully.")
}

func ClosePostgresConnection() {
	conn := db.GetConnection()
	err := conn.Close()
	if err != nil {
		panic(err.Error())
	}
	logger.Log("Connection terminated successfully.")
}

func getConfiguration() *resources.PostgresConf {
	pwd, _ := os.Getwd()
	configuration, err := ioutil.ReadFile(pwd + "\\user_gateway\\cmd\\service\\resources\\configuration.yaml")
	if err != nil {
		panic(err.Error())
	}
	var conf = resources.Configuration{}
	err = yaml.Unmarshal(configuration, &conf)
	if err != nil {
		panic(err.Error())
	}
	return conf.GetPostgresConf()
}
