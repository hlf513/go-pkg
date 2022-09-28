package database

type Tabler interface {
	TableName() string
}

type Query interface {
	Create(row Tabler) error
	BatchInsert(newRows interface{}, opts ...QueryOption) error
	FetchByWhere(rows interface{}, opts ...QueryOption) error
	DeleteByWhere(opts ...QueryOption) error
	UpdateByWhere(newRow interface{}, opts ...QueryOption) error
	CountByWhere(count *int64, opts ...QueryOption) error
}
