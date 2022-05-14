package category

import (
    "hanya-go/pkg/database"
)

func Get(idStr string) (category Category) {
    database.DB.Where("id", idStr).First(&category)
    return
}

func GetBy(field, value string) (category Category) {
    database.DB.Where("? = ?", field, value).First(&category)
    return
}

func All() (category []Category) {
    database.DB.Find(&category)
    return
}

func IsExist(field, value string) bool {
    var count int64
    database.DB.Model(Category{}).Where(" = ?", field, value).Count(&count)
    return count > 0
}