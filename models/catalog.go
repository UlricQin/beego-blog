package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/ulricqin/beego-blog/g"
)

func OneById(id int64) *Catalog {
	if id == 0 {
		return nil
	}

	key := fmt.Sprintf("%d", id)
	val := g.CatalogCacheGet(key)
	if val == nil {
		if cp := OneByIdInDB(id); cp != nil {
			g.CatalogCachePut(key, *cp)
			return cp
		} else {
			return nil
		}
	}
	ret := val.(Catalog)
	return &ret
}

func OneByIdInDB(id int64) *Catalog {
	if id == 0 {
		return nil
	}

	c := Catalog{Id: id}
	err := orm.NewOrm().Read(&c, "Id")
	if err != nil {
		return nil
	}

	return &c
}

func OneByIdentInDB(ident string) *Catalog {
	if ident == "" {
		return nil
	}

	c := Catalog{Ident: ident}
	err := orm.NewOrm().Read(&c, "Ident")
	if err != nil {
		return nil
	}

	return &c
}

func IdByIdent(ident string) int64 {
	if ident == "" {
		return 0
	}

	val := g.CatalogCacheGet(ident)
	if val == nil {
		if cp := OneByIdentInDB(ident); cp != nil {
			g.CatalogCachePut(ident, cp.Id)
			return cp.Id
		} else {
			return 0
		}
	}

	return val.(int64)
}

func OneByIdent(ident string) *Catalog {
	id := IdByIdent(ident)
	return OneById(id)
}

// all ids

// all

// save (clear ids cache)

// insert to db

// update (clear one cache)
