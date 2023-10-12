package casbin

import (
	"errors"
	"fmt"
	gormadapter "github.com/casbin/gorm-adapter/v2"
	"github.com/hlf513/go-pkg/database/gorm/mysql"
)

const (
	PrefixUserID  = "u"
	PrefixRoleID  = "r"
	UserIDCtxName = "casbinUserId"
	SuperUserID   = 1
)

var (
	ErrUidEmpty = errors.New("uid not found")
	ErrCasbin   = errors.New("auth model exception")
)

var (
	defaultModel = `
	[request_definition] 
	r = sub, obj, act 
	
	[policy_definition] 
	p = sub, obj, act 
	
	[role_definition]
	g = _, _
	
	[policy_effect]
	e = some(where (p.eft == allow))
	
	[matchers]
	m = g(r.sub, p.sub) == true \
			&& keyMatch(r.obj, p.obj) == true \
			&& regexMatch(r.act, p.act) == true \
			|| r.sub == "root"`
)

func GormAdapter(dbConf mysql.Options) (*gormadapter.Adapter, error) {
	return gormadapter.NewAdapter(
		"mysql",
		fmt.Sprintf("%s:%s@tcp(%s)/%s", dbConf.Username, dbConf.Password, dbConf.Host, dbConf.Database),
		true,
	)
}
