package provider

import (
	"magician/model/dao"

	"github.com/google/wire"
)

var daoUserSet = wire.NewSet(wire.Struct(new(dao.User), "*"))

// DaoSet inject
var DaoSet = wire.NewSet(daoUserSet)
