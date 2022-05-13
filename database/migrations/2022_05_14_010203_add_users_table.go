package migrations

import (
    "database/sql"
    "hanya-go/app/models"
    "hanya-go/pkg/migrate"

    "gorm.io/gorm"
)

func init() {

    type User struct {
        UserID uint `gorm:"column:user_id;primaryKey;autoIncrement;type:int(10)" json:"user_id,omitempty"`

        Nickname string `gorm:"column:nickname;type:varchar(128)" json:"nickname,omitempty"`
        Avatar   string `gorm:"column:avatar;type:varchar(256)" json:"avatar,omitempty"`
        Openid   string `gorm:"column:openid;uniqueIndex;type:varchar(128)" json:"openid,omitempty"`
        Email    string `gorm:"uniqueIndex;type:varchar(128)" json:"-"`
        Phone    string `gorm:"uniqueIndex;type:char(11)" json:"-"`
        Password string `gorm:"type:varchar(256)" json:"-"`

        models.Datetime
    }

    up := func(migrator gorm.Migrator, DB *sql.DB) {
        _ = migrator.AutoMigrate(&User{})
    }

    down := func(migrator gorm.Migrator, DB *sql.DB) {
        _ = migrator.DropTable(&User{})
    }

    migrate.Add("2022_05_14_010203_add_users_table", up, down)
}