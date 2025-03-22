package dal

import (
	"eCommerce/app/frontend/biz/dal/mysql"
	"eCommerce/app/frontend/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
