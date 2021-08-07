package mysql

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"microshop/infrastructure/config"
	"microshop/infrastructure/logger"
	"strings"
)

var (
	DB *gorm.DB
	cfg = config.GetConfig()
	log = logger.GetLogger()
)

func StartUp() {
	user := cfg.MysqlUserName
	pwd := cfg.MysqlPassWord
	server := cfg.MysqlUrl
	port := cfg.MysqlPort
	dbName := cfg.DBName
	network := "tcp"
	url := fmt.Sprintf("%s:%s@%s(%s:%d)/%s?charset=utf8mb4", user, pwd, network, server, port, dbName)
	log.Infof("mysql connect url is:%s", url)
	var err error
	DB, err = gorm.Open("mysql", url)
	if err != nil {
		panic(err)
	}
	DB.DB().SetMaxIdleConns(2)
	DB.DB().SetMaxOpenConns(20)
	DB.SingularTable(true)
}

func SQLBuilder(where map[string]interface{}) (whereSql string,
	values []interface{}, err error) {
	for key, value := range where {
		conditionKey := strings.Split(key, " ")
		if len(conditionKey) > 2 {
			return "", nil, fmt.Errorf("" +
				"map构建的条件格式不对，类似于'age >'")
		}
		if whereSql != "" {
			whereSql += " AND "
		}
		switch len(conditionKey) {
		case 1:
			whereSql += fmt.Sprint(conditionKey[0], " = ?")
			values = append(values, value)
			break
		case 2:
			field := conditionKey[0]
			switch conditionKey[1] {
			case "=":
				whereSql += fmt.Sprint(field, " = ?")
				values = append(values, value)
				break
			case ">":
				whereSql += fmt.Sprint(field, " > ?")
				values = append(values, value)
				break
			case ">=":
				whereSql += fmt.Sprint(field, " >= ?")
				values = append(values, value)
				break
			case "<":
				whereSql += fmt.Sprint(field, " < ?")
				values = append(values, value)
				break
			case "<=":
				whereSql += fmt.Sprint(field, " <= ?")
				values = append(values, value)
				break
			case "in":
				whereSql += fmt.Sprint(field, " in (?)")
				values = append(values, value)
				break
			case "like":
				whereSql += fmt.Sprint(field, " like ?")
				values = append(values, value)
				break
			case "<>":
				whereSql += fmt.Sprint(field, " != ?")
				values = append(values, value)
				break
			case "!=":
				whereSql += fmt.Sprint(field, " != ?")
				values = append(values, value)
				break
			}
			break
		}
	}
	return
}
