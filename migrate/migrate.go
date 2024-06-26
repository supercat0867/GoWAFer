package migrate

import (
	"GoWAFer/internal/model"
	"GoWAFer/pkg/hash_handler"
	"fmt"
	"gorm.io/gorm"
	"log"
)

func migrate(db *gorm.DB) {
	err := db.AutoMigrate(&model.Admin{})
	if err != nil {
		panic(fmt.Sprintf("管理员表创建失败：%v", err))
	}
	log.Println("管理员表迁移成功")

	err = db.AutoMigrate(&model.Log{})
	if err != nil {
		panic(fmt.Sprintf("流量日志表创建失败：%v", err))
	}
	log.Println("流量日志表迁移成功")
}

func insert(db *gorm.DB) {
	// 哈希加盐加密
	defaultPassword, _ := hash_handler.EncryptPassword("123456")
	defaultUser := model.Admin{Model: gorm.Model{ID: 1}, Username: "admin", Password: defaultPassword, Nickname: "admin"}
	err := db.FirstOrCreate(&defaultUser, model.Admin{Model: gorm.Model{ID: 1}}).Error
	if err != nil {
		panic(fmt.Sprintf("默认管理员创建失败：%v", err))
	}

	//sqlDefault := []model.SqlInjectionRules{
	//	{Rule: "(?i)(union)(.*)(select)"},
	//	{Rule: "(?i)select(.*)from"},
	//	{Rule: "(?i)insert into"},
	//	{Rule: "(?i)delete from"},
	//	{Rule: "(?i)drop table"},
	//	{Rule: "(?i)update(.*)set"},
	//	{Rule: "--"},
	//	{Rule: "(\\b|\\')(OR|or|oR|Or)('|\\b)\\s*('\\d+'|'\\d+'--\\s*|'\\d+'(\\s+)(--)?|\\d+)(\\s+)(=|like)(\\s+)(\\b|\\')\\d+('|\\b)"},
	//	{Rule: "/\\*.*\\*/"},
	//	{Rule: ";"},
	//}
	//for _, sql := range sqlDefault {
	//	db.FirstOrCreate(&sql, model.SqlInjectionRules{Rule: sql.Rule})
	//}
	//
	//xssDefault := []model.XssDetectRules{
	//	{Rule: "<script[^>]*>.*?</script>"},
	//	{Rule: "on\\w+=\"[^\"]*\""},
	//	{Rule: "on\\w+='[^']*'"},
	//	{Rule: "javascript:[^\\s]*"},
	//	{Rule: "<iframe[^>]*>.*?</iframe>"},
	//	{Rule: "<embed[^>]*>.*?</embed>"},
	//	{Rule: "<object[^>]*>.*?</object>"},
	//	{Rule: "srcdoc=\"[^\"]*\""},
	//	{Rule: "<img[^>]*src=\"[^\"]*javascript:[^\"]*\"[^>]*>"},
	//}
	//for _, xss := range xssDefault {
	//	db.FirstOrCreate(&xss, model.XssDetectRules{Rule: xss.Rule})
	//}
}

func AutoMigrateAndInsertData(db *gorm.DB) {
	migrate(db)
	insert(db)
}
