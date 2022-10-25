package gorm

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/plugin/soft_delete"
	"os"
	"testing"
	"time"
)

var db *gorm.DB

type Test struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
	IsDeleted soft_delete.DeletedAt `gorm:"softDelete:flag,DeletedAtField:DeletedAt"`
	Username  string
}

func (t Test) TableName() string {
	return "tests_202209"
}

func TestMain(m *testing.M) {
	var err error
	db, err = gorm.Open(mysql.Open("user:pwd@tcp(1.1.1.2:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	err = db.AutoMigrate(&Test{})
	if err != nil {
		fmt.Println(err)
		return
	}
	db.Model(Test{}).Where("id > 0").Unscoped().Delete(&Test{})
	retCode := m.Run()
	os.Exit(retCode)
}

func TestModel_Create(t *testing.T) {
	var row = Test{
		Username: "username",
	}
	err := NewQuery(db).Create(&row)
	assert.NoError(t, err)
	assert.Equal(t, true, row.ID > 0)
}

func TestModel_BatchInsert(t *testing.T) {
	var rows = []Test{
		{Username: "user1"},
		{Username: "user2"},
	}
	err := NewQuery(db).BatchInsert(Test{}, rows, BatchInsertSize(1))
	assert.NoError(t, err)
	for _, item := range rows {
		assert.Equal(t, true, item.ID > 0)
	}
}

func TestModel_FindByWhere(t *testing.T) {
	db.Create([]Test{
		{Username: "user3"},
		{Username: "user3"},
	})

	var result []Test
	m := NewQuery(db)
	err := m.FindByWhere(&result, Where("username = ?", "user3"))
	assert.NoError(t, err)
	assert.Equal(t, 2, len(result))

	err2 := m.FindByWhere(
		&result,
		Where("username = ?", "user3"),
		Limit(1),
		Offset(0),
		Order("id desc"),
		Group("username"),
		Having("username != ''"),
	)
	assert.NoError(t, err2)
	assert.Equal(t, 1, len(result))

	var newResult Test
	err3 := m.ScanByWhere(Test{}, &newResult,
		Where("username = ?", "user3"),
	)
	assert.NoError(t, err3)
	assert.Equal(t, true, newResult.ID > 0)

}

func TestModel_DeleteByWhere(t *testing.T) {
	db.Create([]Test{
		{Username: "user4"},
	})
	m := NewQuery(db)
	err := m.DeleteByWhere(
		Test{},
		Where("username = ?", "user4"),
	)
	assert.NoError(t, err)
	var result []Test
	db.Model(Test{}).Where("username = ?", "user4").Find(&result)
	assert.Equal(t, 0, len(result))

	db.Model(Test{}).Unscoped().Where("username = ?", "user4").Find(&result)
	assert.Equal(t, 1, len(result))

	err = m.DeleteByWhere(Test{}, Where("username = ?", "user4"), HardDelete())
	assert.NoError(t, err)

	db.Model(Test{}).Unscoped().Where("username = ?", "user4").Find(&result)
	assert.Equal(t, 0, len(result))
}

func TestModel_UpdateByWhere(t *testing.T) {
	db.Create(&Test{
		Username: "user51",
	})

	var newData = Test{
		Username: "user5",
	}
	m := NewQuery(db)
	err := m.UpdateByWhere(
		newData,
		Where("username = ?", "user51"),
	)
	assert.NoError(t, err)
	var updateData []Test
	db.Model(updateData).Where("username = ?", "user5").Find(&updateData)
	assert.Equal(t, 1, len(updateData))
}

func TestModel_CountByWhere(t *testing.T) {
	db.Create(&Test{
		Username: "user6",
	})
	m := NewQuery(db)
	var c int64
	err := m.CountByWhere(
		Test{},
		&c,
		Where("username = ?", "user6"),
	)
	assert.NoError(t, err)
	assert.Equal(t, int64(1), c)
}
