package link

import (
    "hanya-go/pkg/cache"
    "hanya-go/pkg/database"
    "hanya-go/pkg/helpers"
    "time"
)

func Get(idStr string) (link Link) {
    database.DB.Where("id", idStr).First(&link)
    return
}

func GetBy(field, value string) (link Link) {
    database.DB.Where("? = ?", field, value).First(&link)
    return
}

func All() (link []Link) {
    database.DB.Find(&link)
    return
}

func IsExist(field, value string) bool {
    var count int64
    database.DB.Model(Link{}).Where(" = ?", field, value).Count(&count)
    return count > 0
}

func AllCached() (links []Link) {
    // 设置缓存 key
    cacheKey := "links:all"

    // 设置过期时间
    expireTime := 120 * time.Minute

    // 取数据
    cache.GetObject(cacheKey, &links)

    if helpers.Empty(links) {
        links = All()

        if helpers.Empty(links) { // 查询依然没有
            return links
        } else {
            // 设置缓存
            cache.Set(cacheKey, links, expireTime)
        }
    }

    return
}
