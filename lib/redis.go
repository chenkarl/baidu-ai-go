package lib

import (
	"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/redis"
)

// Redis ...
// connect local redis
func Redis() (bm cache.Cache, err error) {
	bm, err = cache.NewCache("redis", `{"conn": "127.0.0.1:6379","dbNum":"0"}`)
	return
}
