package query

import (
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type Query interface {
	GetDB() *gorm.DB
	SetDB(db *gorm.DB)
	Transaction(func(tx *gorm.DB) error) error
	Save(row schema.Tabler) error
	Create(newRow schema.Tabler) error
	BatchInsert(table schema.Tabler, newRows any, opts ...Option) error
	FindOneByWhere(rows any, opts ...Option) error
	FindByWhere(rows any, opts ...Option) error
	ScanByWhere(table schema.Tabler, row any, opts ...Option) error // 用于自定义结构体
	DeleteByWhere(table schema.Tabler, opts ...Option) error
	UpdateByWhere(newRow any, opts ...Option) error
	CountByWhere(table schema.Tabler, count *int64, opts ...Option) error
	SearchAll(table schema.Tabler, total *int64, list any, page, limit int, opt ...Option) error
}

func NewQuery(db *gorm.DB) Query {
	return &query{
		db:   db,
		opts: newOptions(),
	}
}

type query struct {
	db   *gorm.DB
	opts Options
}

func (m *query) Save(row schema.Tabler) error {
	return m.db.Save(row).Error
}

func (m *query) Transaction(f func(tx *gorm.DB) error) error {
	return m.db.Transaction(f)
}

func (m *query) SetDB(db *gorm.DB) {
	m.db = db
}

func (m *query) FindOneByWhere(rows any, opts ...Option) error {
	opts = append(opts, Limit(1))
	return m.FindByWhere(rows, opts...)
}

func (m *query) SearchAll(table schema.Tabler, total *int64, list any, page, limit int, opts ...Option) error {
	q := m.buildCondition(append(opts, Table(table))...)
	if err := q.Count(total).Error; err != nil {
		return err
	}
	if *total == 0 {
		return nil
	}
	offset := (page - 1) * limit
	if offset < 0 {
		offset = 0
	}
	return q.Select(m.opts.Select).Offset(offset).Limit(limit).Scan(list).Error
}

func (m *query) GetDB() *gorm.DB {
	return m.db
}

func (m *query) Create(newRow schema.Tabler) error {
	return m.db.Create(newRow).Error
}

func (m *query) BatchInsert(table schema.Tabler, rows any, opts ...Option) error {
	opt := newOptions(append(opts, Table(table))...)
	return m.db.CreateInBatches(rows, opt.BatchInsertSize).Error
}

func (m *query) FindByWhere(rows any, opts ...Option) error {
	return m.buildCondition(opts...).Select(m.opts.Select).Find(rows).Error
}

func (m *query) ScanByWhere(table schema.Tabler, row any, opts ...Option) error {
	return m.buildCondition(append(opts, Table(table))...).Select(m.opts.Select).Scan(row).Error
}

func (m *query) DeleteByWhere(table schema.Tabler, opts ...Option) error {
	q := m.buildCondition(append(opts, Table(table))...)
	if m.opts.HardDelete {
		return q.Unscoped().Delete(&m.opts.Table).Error
	} else {
		return q.Delete(&m.opts.Table).Error
	}
}

func (m *query) UpdateByWhere(newRow any, opts ...Option) error {
	return m.buildCondition(opts...).Updates(newRow).Error
}

func (m *query) CountByWhere(table schema.Tabler, count *int64, opts ...Option) error {
	return m.buildCondition(append(opts, Table(table))...).Count(count).Error
}

func (m *query) buildCondition(opts ...Option) *gorm.DB {
	m.opts = newOptions(opts...)

	q := m.db

	if m.opts.Table != nil {
		q = q.Model(m.opts.Table)
	}

	for _, w := range m.opts.Where {
		q = q.Where(w[0], w[1:]...)
	}

	if m.opts.Group != "" {
		q = q.Group(m.opts.Group)
	}

	if m.opts.Order != "" {
		q = q.Order(m.opts.Order)
	}

	if m.opts.Join != "" {
		q = q.Joins(m.opts.Join)
	}

	if m.opts.Having != "" {
		q = q.Having(m.opts.Having)
	}

	if m.opts.Limit > 0 {
		q = q.Limit(m.opts.Limit)
	}

	if m.opts.Offset > 0 {
		q = q.Offset(m.opts.Offset)
	}

	return q
}
