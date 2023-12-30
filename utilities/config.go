package utilities

import (
	"context"

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

func EnableNoSqlDatabaseConfiguration() (*mongo.Client, error) {
	// Viper essential config
	viper.SetConfigName("env")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./conf")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.Get("local.mongo") != nil {
		mongoConf := viper.Get("local.mongo").(map[string]interface{})
		uri := "mongo+srv://" + mongoConf["user"].(string) + ":" + mongoConf["password"].(string) + "@" + mongoConf["host"].(string) + mongoConf["database"].(string)
		client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
		if err != nil {
			panic(err)
		}
		return client, nil
	}
	return nil, err
}
