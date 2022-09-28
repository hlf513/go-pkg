package gorm

import (
	"errors"
	"github.com/hlf513/go-pkg/database"
	"gorm.io/gorm"
)

func NewQuery(db *gorm.DB) database.Query {
	return &query{
		db:   db,
		opts: database.NewQueryOptions(),
	}
}

type query struct {
	db   *gorm.DB
	opts database.QueryOptions
}

func (m *query) Create(row database.Tabler) error {
	return m.db.Create(row).Error
}

func (m *query) BatchInsert(rows interface{}, opts ...database.QueryOption) error {
	opt := database.NewQueryOptions(opts...)
	if opt.Table == nil || opt.BatchInsertSize <= 0 {
		return errors.New("table,bathInsertSize required")
	}
	return m.db.Model(opt.Table).CreateInBatches(rows, opt.BatchInsertSize).Error
}

func (m *query) FetchByWhere(rows interface{}, opts ...database.QueryOption) error {
	return m.buildCondition(opts...).Select(m.opts.Fields).Find(rows).Error
}

func (m *query) DeleteByWhere(opts ...database.QueryOption) error {
	q := m.buildCondition(opts...)
	if m.opts.Table == nil {
		return errors.New("table required")
	}
	return q.Delete(&m.opts.Table).Error
}

func (m *query) UpdateByWhere(newRow interface{}, opts ...database.QueryOption) error {
	q := m.buildCondition(opts...)
	return q.Updates(newRow).Error
}

func (m *query) CountByWhere(count *int64, opts ...database.QueryOption) error {
	q := m.buildCondition(opts...)
	if m.opts.Table == nil {
		return errors.New("table required")
	}
	return q.Model(m.opts.Table).Count(count).Error
}

func (m *query) GetDB() *gorm.DB {
	return m.db
}

func (m *query) buildCondition(opts ...database.QueryOption) *gorm.DB {
	m.opts = database.NewQueryOptions(opts...)

	q := m.db

	if m.opts.Fields != "" {
		q = q.Select(m.opts.Fields)
	}
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
