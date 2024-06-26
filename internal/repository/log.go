package repository

import (
	"GoWAFer/internal/model"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type LogRepository struct {
	db *gorm.DB
}

func NewLogRepository(db *gorm.DB) *LogRepository {
	return &LogRepository{db: db}
}

// Create 新增日志记录
func (r *LogRepository) Create(l *model.Log) error {
	return r.db.Create(l).Error
}

// FindLog 查询指定天数和小时数的日志记录
func (r *LogRepository) FindLog(days, hours int) []model.Log {
	// 初始化开始日期为当前时间
	startTime := time.Now()

	// 如果提供了小时数并且大于0，则优先使用小时计算开始时间
	if hours > 0 {
		startTime = startTime.Add(time.Duration(-hours) * time.Hour)
	} else if days > 0 {
		// 否则，如果提供了天数并且大于0，使用天数计算开始时间
		startTime = startTime.AddDate(0, 0, -days)
	}

	var logs []model.Log

	r.db.Where("CreatedAt >= ?", startTime).Find(&logs)

	return logs
}

// FindPaginated 分页查询被拦截的流量日志
func (r *LogRepository) FindPaginated(page, pageSize int, query map[string]interface{}, order string) ([]*model.Log, int) {
	var logs []*model.Log
	var count int64
	sql := r.db.Model(&model.Log{}).Order("CreatedAt " + order).
		Where("BlockBy IS NOT NULL AND BlockBy <> ''")
	for k, v := range query {
		sql = sql.Where(k+" LIKE ?", fmt.Sprintf("%%%s%%", v))
	}
	sql.Count(&count).Limit(pageSize).Offset((page - 1) * pageSize).Find(&logs)
	return logs, int(count)
}
