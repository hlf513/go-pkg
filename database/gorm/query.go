package gorm

import (
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type Query interface {
	GetDB() *gorm.DB
	Create(newRow schema.Tabler) error
	BatchInsert(table schema.Tabler, newRows interface{}, opts ...QueryOption) error
	FindByWhere(rows interface{}, opts ...QueryOption) error
	ScanByWhere(table schema.Tabler, row interface{}, opts ...QueryOption) error // 用于自定义结构体
	DeleteByWhere(table schema.Tabler, opts ...QueryOption) error
	UpdateByWhere(newRow interface{}, opts ...QueryOption) error
	CountByWhere(table schema.Tabler, count *int64, opts ...QueryOption) error
}

func NewQuery(db *gorm.DB) Query {
	return &query{
		db:   db,
		opts: newQueryOptions(),
	}
}

type query struct {
	db   *gorm.DB
	opts QueryOptions
}

func (m *query) GetDB() *gorm.DB {
	return m.db
}

func (m *query) Create(newRow schema.Tabler) error {
	return m.db.Create(newRow).Error
}

func (m *query) BatchInsert(table schema.Tabler, rows interface{}, opts ...QueryOption) error {
	opt := newQueryOptions(append(opts, Table(table))...)
	return m.db.Model(m.opts.Table).CreateInBatches(rows, opt.BatchInsertSize).Error
}

func (m *query) FindByWhere(rows interface{}, opts ...QueryOption) error {
	return m.buildCondition(opts...).Select(m.opts.Fields).Find(rows).Error
}

func (m *query) ScanByWhere(table schema.Tabler, row interface{}, opts ...QueryOption) error {
	return m.buildCondition(append(opts, Table(table))...).Model(m.opts.Table).Select(m.opts.Fields).Scan(row).Error
}

func (m *query) DeleteByWhere(table schema.Tabler, opts ...QueryOption) error {
	q := m.buildCondition(append(opts, Table(table))...)
	if m.opts.HardDelete {
		return q.Unscoped().Delete(&m.opts.Table).Error
	} else {
		return q.Delete(&m.opts.Table).Error
	}
}

func (m *query) UpdateByWhere(newRow interface{}, opts ...QueryOption) error {
	return m.buildCondition(opts...).Updates(newRow).Error
}

func (m *query) CountByWhere(table schema.Tabler, count *int64, opts ...QueryOption) error {
	return m.buildCondition(append(opts, Table(table))...).Model(m.opts.Table).Count(count).Error
}

func (m *query) buildCondition(opts ...QueryOption) *gorm.DB {
	m.opts = newQueryOptions(opts...)

	q := m.db

	if m.opts.Where[0] != "" {
		q = q.Where(m.opts.Where[0], m.opts.Where[1])
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
