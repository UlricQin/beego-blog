package catalog

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/ulricqin/beego-blog/g"
	. "github.com/ulricqin/beego-blog/models"
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

func IdentExists(ident string) bool {
	id := IdByIdent(ident)
	return id > 0
}

func OneByIdent(ident string) *Catalog {
	id := IdByIdent(ident)
	return OneById(id)
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

func AllIdsInDB() []int64 {
	var catalogs []Catalog
	Catalogs().OrderBy("-DisplayOrder").All(&catalogs, "Id")
	size := len(catalogs)
	if size == 0 {
		return []int64{}
	}

	ret := make([]int64, size)
	for i := 0; i < size; i++ {
		ret[i] = catalogs[i].Id
	}

	return ret
}

func AllIds() []int64 {
	val := g.CatalogCacheGet("ids")
	if val == nil {
		if ids := AllIdsInDB(); len(ids) != 0 {
			g.CatalogCachePut("ids", ids)
			return ids
		} else {
			return []int64{}
		}
	}

	return val.([]int64)
}

func All() []*Catalog {
	ids := AllIds()
	size := len(ids)
	if size == 0 {
		return []*Catalog{}
	}

	ret := make([]*Catalog, size)
	for i := 0; i < size; i++ {
		ret[i] = OneById(ids[i])
	}
	return ret
}

func Save(this *Catalog) (int64, error) {
	if IdentExists(this.Ident) {
		return 0, fmt.Errorf("catalog english identity exists")
	}
	num, err := orm.NewOrm().Insert(this)
	if err == nil {
		g.CatalogCacheDel("ids")
	}

	return num, err
}

func Del(c *Catalog) error {
	num, err := orm.NewOrm().Delete(c)
	if err != nil {
		return err
	}

	if num > 0 {
		g.CatalogCacheDel("ids")
	}
	return nil
}

func Update(this *Catalog) error {
	if this.Id == 0 {
		return fmt.Errorf("primary key id not set")
	}
	_, err := orm.NewOrm().Update(this)
	if err == nil {
		g.CatalogCacheDel(fmt.Sprintf("%d", this.Id))
	}
	return err
}

func Catalogs() orm.QuerySeter {
	return orm.NewOrm().QueryTable(new(Catalog))
}
