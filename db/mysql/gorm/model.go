package gorm

import (
	"gorm.io/gorm"
)

const (
	GROUP  = "group"
	ORDER  = "order"
	JOIN   = "join"
	HAVING = "having"
	OFFSET = "offset"
	LIMIT  = "limit"
)

type Model interface {
	GetDB() *gorm.DB
	Create(tableName string, data interface{}) error
	BatchInsert(tableName string, data interface{}, size int) error
	FetchByWhere(tableName, fields string, where map[string]interface{}, data interface{}, others ...map[string]interface{}) error
	DeleteByWhere(table interface{}, where map[string]interface{}, others ...map[string]interface{}) error
	UpdateByWhere(tableName string, where map[string]interface{}, set interface{}, others ...map[string]interface{}) error
	CountByWhere(tableName string, where map[string]interface{}, others ...map[string]interface{}) (int64, error)
}


func NewModel(db *gorm.DB) Model {
	return &model{
		db:   db,
	}
}

type model struct {
	db   *gorm.DB
}

func (m *model) GetDB() *gorm.DB {
	return m.db
}

func (m *model) Create(tableName string, data interface{}) error {
	return m.db.Table(tableName).Create(data).Error
}

func (m *model) BatchInsert(tableName string, data interface{}, size int) error {
	return m.db.Table(tableName).CreateInBatches(data, size).Error
}

func (m *model) FetchByWhere(tableName, fields string, where map[string]interface{}, data interface{}, others ...map[string]interface{}) error {
	return m.prepare(where, others...).Table(tableName).Select(fields).Find(data).Error
}

func (m *model) DeleteByWhere(table interface{}, where map[string]interface{}, others ...map[string]interface{}) error {
	return m.prepare(where, others...).Delete(table).Error
}

func (m *model) UpdateByWhere(tableName string, where map[string]interface{}, set interface{}, others ...map[string]interface{}) error {
	return m.prepare(where, others...).Table(tableName).Updates(set).Error
}

func (m *model) CountByWhere(tableName string, where map[string]interface{}, others ...map[string]interface{}) (int64, error) {
	var c int64
	if err := m.prepare(where, others...).Table(tableName).Count(&c).Error; err != nil {
		return c, err
	}
	return c, nil
}

func (m *model) prepare(where map[string]interface{}, others ...map[string]interface{}) *gorm.DB {
	q := m.db
	for k, v := range where {
		if v != nil {
			q = q.Where(k, v)
		} else {
			q = q.Where(k)
		}
	}

	if others != nil {
		for _, other := range others {
			if g, ok := other[GROUP]; ok {
				q = q.Group(g.(string))
			}

			if o, ok := other[ORDER]; ok {
				q = q.Order(o)
			}

			if o, ok := other[JOIN]; ok {
				for _, j := range o.([]string) {
					q = q.Joins(j)
				}
			}

			if h, ok := other[HAVING]; ok {
				q = q.Having(h)
			}

			if o, ok := other[OFFSET]; ok {
				q = q.Offset(o.(int))
			}
			if l, ok := other[LIMIT]; ok {
				q = q.Limit(l.(int))
			}
		}
	}
	return q
}
