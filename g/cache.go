package g

const (
	blogPrefix    = "b_"
	catalogPrefix = "c_"
)

func BlogCachePut(key string, val interface{}) error {
	return Cache.Put(blogPrefix+key, val, blogCacheExpire)
}

func CatalogCachePut(key string, val interface{}) error {
	return Cache.Put(catalogPrefix+key, val, catalogCacheExpire)
}

func BlogCacheGet(key string) interface{} {
	return Cache.Get(blogPrefix + key)
}

func CatalogCacheGet(key string) interface{} {
	return Cache.Get(catalogPrefix + key)
}

func CatalogCacheDel(key string) error {
	return Cache.Delete(catalogPrefix + key)
}

func BlogCacheDel(key string) error {
	return Cache.Delete(blogPrefix + key)
}


