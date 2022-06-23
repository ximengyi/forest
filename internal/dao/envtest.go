package dao

import (
	"forest/pkg/db"
	"gorm.io/gorm"
)

type Envtest struct {
	Name string `json:"name" gorm:"column:name" description:"名称"`
	Text string `json:"text" gorm:"column:text" description:"text"`
	Base
}

type ApiEnvtest struct {
	Id   int    `json:"id" gorm:"primary_key" description:"自增主键"`
	Name string `json:"name" gorm:"column:name" description:"名称"`
	Text string `json:"text" gorm:"column:text" description:"text"`
}

func (t *Envtest) TableName() string {
	return "envtest"
}

func (t *Envtest) Dbtest() ([]ApiEnvtest, error) {
	var envTestList []ApiEnvtest
	db, _ := db.GetGormPool("default")
	err := db.Model(&Envtest{}).Find(&envTestList).Error
	if err != nil {
		return nil, err
	}

	return envTestList, nil
}

func (t *Envtest) PageList(page int, pageSize int, conditions interface{}, args ...interface{}) ([]Envtest, int64, error) {
	var results []Envtest
	var count int64
	offset := (page - 1) * pageSize
	query, _ := db.GetGormPool("default")
	err := query.Model(&Envtest{}).Limit(pageSize).Offset(offset).Order("id desc").Find(&results).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0, err
	}
	errCount := query.Model(&Envtest{}).Where(conditions, args...).Count(&count).Error
	if errCount != nil {
		return nil, 0, err
	}
	return results, count, nil
}
