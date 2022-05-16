package link

import (
    "hanya-go/pkg/database"
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