package internal

import (
	_ "github.com/go-sql-driver/mysql" //导入mysql驱动
	"github.com/gone-io/gone/v2"
	"github.com/gone-io/goner"
	"github.com/gone-io/goner/xorm"
)

//go:generate gonectr generate -m=../cmd/server
func Load(loader gone.Loader) error {
	err := xorm.Load(loader)
	if err != nil {
		return gone.ToError(err)
	}
	return goner.GinLoad(loader)
}

func TestLoader(loader gone.Loader) error {
	return xorm.Load(loader)
}
