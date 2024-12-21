package internal

import (
	_ "github.com/go-sql-driver/mysql" //导入mysql驱动
	"github.com/gone-io/gone"
	"github.com/gone-io/gone/goner"
)

//go:generate gonectr generate -m=../cmd/server
func Load(loader gone.Cemetery) error {
	_ = goner.XormPriest(loader)
	return goner.GinPriest(loader)
}
