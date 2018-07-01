package main

import (
	"flag"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"ywopt.com/carrier.v1/server"
	"golang-serialization-benchmark/src/config"
)

const (
	Conf         = "f"
	DefaultConf = "/src/config/setting.json"
	ConfDes     = "config path"
)

type (
	Config struct {
		DbUri   string `json:"db_uri"`
		DbType  string `json:"db_type"`
		Address string `json:"address"`
	}
)

func main() {
	confPath := flag.String(Conf, DefaultConf, ConfDes)
	flag.Parse()
	c := Config{}
	fmt.Println(*confPath)
	config.Load(*confPath, &c)

	db, err := gorm.Open(c.DbType, c.DbUri)
	//defer db.Close()
	if err != nil {
		fmt.Println(err)
		fmt.Println("connect error")
		return
	}
	conf := &server.Config{
		Addr:         c.Address,
		Db:           *db,
		DbUri:        c.DbUri,
	}
	s := server.NewServer(conf)
	s.Start()
}
