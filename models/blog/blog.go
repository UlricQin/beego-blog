package catalog

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/ulricqin/beego-blog/g"
	. "github.com/ulricqin/beego-blog/models"
)

func OneById(id int64) *Blog {
	if id == 0 {
		return nil
	}

	key := fmt.Sprintf("%d", id)
	val := g.BlogCacheGet(key)
	if val == nil {
		if p := OneByIdInDB(id); p != nil {
			g.BlogCachePut(key, *p)
			return p
		} else {
			return nil
		}
	}
	ret := val.(Blog)
	return &ret
}

func OneByIdInDB(id int64) *Blog {
	if id == 0 {
		return nil
	}

	o := Blog{Id: id}
	err := orm.NewOrm().Read(&o, "Id")
	if err != nil {
		return nil
	}
	return &o
}

func IdByIdent(ident string) int64 {
	if ident == "" {
		return 0
	}

	val := g.BlogCacheGet(ident)
	if val == nil {
		if p := OneByIdentInDB(ident); p != nil {
			g.BlogCachePut(ident, p.Id)
			return p.Id
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

func OneByIdent(ident string) *Blog {
	id := IdByIdent(ident)
	return OneById(id)
}

func OneByIdentInDB(ident string) *Blog {
	if ident == "" {
		return nil
	}

	c := Blog{Ident: ident}
	err := orm.NewOrm().Read(&c, "Ident")
	if err != nil {
		return nil
	}

	return &c
}

func IdsInDB(catalog_id int64) []int64 {
	var blogs []Blog
	Blogs().Filter("CatalogId", catalog_id).OrderBy("-Created").All(&blogs, "Id")
	size := len(blogs)
	if size == 0 {
		return []int64{}
	}

	ret := make([]int64, size)
	for i := 0; i < size; i++ {
		ret[i] = blogs[i].Id
	}

	return ret
}

func Ids(catalog_id int64) []int64 {
	key := fmt.Sprintf("article_ids_of_%d", catalog_id)
	val := g.BlogCacheGet(key)
	if val == nil {
		if ids := IdsInDB(catalog_id); len(ids) != 0 {
			g.BlogCachePut(key, ids)
			return ids
		} else {
			return []int64{}
		}
	}

	return val.([]int64)
}

//TODO: page and size
func All(catalog_id int64) []*Blog {
	ids := Ids(catalog_id)
	size := len(ids)
	if size == 0 {
		return []*Blog{}
	}

	ret := make([]*Blog, size)
	for i := 0; i < size; i++ {
		ret[i] = OneById(ids[i])
	}
	return ret
}

func Save(this *Blog) (int64, error) {
	if IdentExists(this.Ident) {
		return 0, fmt.Errorf("blog english identity exists")
	}
	num, err := orm.NewOrm().Insert(this)
	if err == nil {
		g.BlogCacheDel(fmt.Sprintf("article_ids_of_%d", this.CatalogId))
	}

	return num, err
}

func Update(this *Blog) error {
	if this.Id == 0 {
		return fmt.Errorf("primary key id not set")
	}
	_, err := orm.NewOrm().Update(this)
	if err == nil {
		g.BlogCacheDel(fmt.Sprintf("%d", this.Id))
	}
	return err
}

func Blogs() orm.QuerySeter {
	return orm.NewOrm().QueryTable(new(Blog))
}
