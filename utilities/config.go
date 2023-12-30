package utilities

import (
	"context"
	"log"
	"time"

	"github.com/beego/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	maxIdle = 5
	maxConn = 30
)

func EnableSQLDatabasesConfiguration() {
	// Viper essential config
	viper.SetConfigName("env")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./conf")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	if viper.Get("local"+".mysql") != nil {
		orm.RegisterDriver("mysql", orm.DRMySQL)
		mysql := viper.Get("local" + ".mysql").(map[string]interface{})
		mysqlConf := mysql["user"].(string) + ":" + mysql["password"].(string) + "@tcp(" + mysql["host"].(string) + ")/" + mysql["database"].(string) + "?charset=utf8"
		if err := orm.RegisterDataBase("default", "mysql", mysqlConf, maxIdle, maxConn); err != nil {
			return
		}
		orm.Debug = true
	}
}

func GetDBInstance() (*mongo.Client, error) {
	uri := "mongodb://192.168.29.186:27017/products"
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))

	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return client, nil
}
