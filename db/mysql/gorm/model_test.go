package gorm

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"gorm.io/plugin/soft_delete"
	"os"
	"testing"
	"time"
)

var (
	dbTest *gorm.DB
)

type Test struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
	IsDeleted soft_delete.DeletedAt `gorm:"softDelete:flag,DeletedAtField:DeletedAt"`
	Username  string
}

func TestMain(m *testing.M) {
	var err error
	dbTest, err = Connect(
		Host("host"),
		Username("user"),
		Password("pwd"),
		Database("dbname"),
	)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = dbTest.AutoMigrate(&Test{})
	if err != nil {
		fmt.Println(err)
		return
	}
	dbTest.Table("tests").Where("id > 0").Unscoped().Delete(&Test{})
	retCode := m.Run()
	os.Exit(retCode)
}

func TestModel_Create(t *testing.T) {
	var row = Test{
		Username: "username",
	}
	err := NewModel(dbTest).Create("tests", &row)
	assert.NoError(t, err)
	assert.Equal(t, true, row.ID > 0)
}

func TestModel_BatchInsert(t *testing.T) {
	var row = []Test{
		{Username: "user1"},
		{Username: "user2"},
	}
	m := NewModel(dbTest)
	err := m.BatchInsert("tests", &row, 10)
	assert.NoError(t, err)
	for _, item := range row {
		assert.Equal(t, true, item.ID > 0)
	}
}

func TestModel_FetchByWhere(t *testing.T) {
	dbTest.Create([]Test{
		{Username: "user3"},
		{Username: "user3"},
	})

	var result []Test
	m := NewModel(dbTest)
	err := m.FetchByWhere("tests", "username", map[string]interface{}{
		"username = ?": "user3",
	}, &result)
	assert.NoError(t, err)
	assert.Equal(t, 2, len(result))

	err2 := m.FetchByWhere("tests", "username", map[string]interface{}{
		"username = ?": "user3",
	}, &result, map[string]interface{}{
		OFFSET: 0,
		LIMIT:  1,
		ORDER:  "id desc",
		GROUP:  "username",
		HAVING: "username != ''",
	})
	assert.NoError(t, err2)
	assert.Equal(t, 1, len(result))

	var newResult Test
	err3 := m.FetchByWhere("tests", "username", map[string]interface{}{
		"username = ?": "user0",
	}, &newResult)
	assert.NoError(t, err3)
	assert.Equal(t, uint(0), newResult.ID)
}

func TestModel_DeleteByWhere(t *testing.T) {
	dbTest.Create([]Test{
		{Username: "user4"},
	})
	m := NewModel(dbTest)
	err := m.DeleteByWhere(&Test{}, map[string]interface{}{
		"username = ?": "user4",
	})
	assert.NoError(t, err)
	var result []Test
	dbTest.Model(Test{}).Where("username = ?", "user4").Find(&result)
	assert.Equal(t, 0, len(result))
}

func TestModel_UpdateByWhere(t *testing.T) {
	dbTest.Create(&Test{
		Username: "user51",
	})

	var newData = Test{
		Username: "user5",
	}
	m := NewModel(dbTest)
	err := m.UpdateByWhere("tests", map[string]interface{}{
		"username = ?": "user51",
	}, newData)
	assert.NoError(t, err)
	var updateData []Test
	dbTest.Model(updateData).Where("username = ?", "user5").Find(&updateData)
	assert.Equal(t, 1, len(updateData))
}

func TestModel_CountByWhere(t *testing.T) {
	dbTest.Create(&Test{
		Username: "user6",
	})
	m := NewModel(dbTest)
	c, err := m.CountByWhere("tests", map[string]interface{}{
		"username = ?": "user6",
	})
	assert.NoError(t, err)
	assert.Equal(t, int64(1), c)
}
