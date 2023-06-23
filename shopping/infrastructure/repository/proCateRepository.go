package repository

import (
	"github.com/farseer-go/cache"
	"github.com/farseer-go/collections"
	"github.com/farseer-go/data"
	"github.com/farseer-go/fs/container"
	"github.com/farseer-go/fs/parse"
	"github.com/farseer-go/mapper"
	"github.com/farseer-go/redis"
	"shopping/domain/procate"
	"shopping/infrastructure/repository/context"
	"shopping/infrastructure/repository/model"
)

// InitProCate 注册商品分类仓储 ioc procate.Repository
func InitProCate() {
	// 这是一个抽象的缓存集合（可以选择Redis、本地内存作为存储源，这里使用的是Redis作为数据源）
	// 这里为了让Redis能自动同步数据库的数据而使用。
	// key：是全局唯一的，同时也是IOC的别名，用于注入时，区分哪个缓存集合
	procateCache := redis.SetProfiles[procate.DomainObject]("procate", "Id", "default")
	// 如果单个item不存在，是否需要重新读整个List数据源（这里开启）
	procateCache.EnableItemNullToLoadAll()
	// 当单个item不存在于Redis时，则要到数据库中获取（如果你的应用不需要，这里就不需要定义）
	procateCache.SetItemSource(func(cacheId any) (procate.DomainObject, bool) {
		// 万能的基础类型转换，将any转成int（也可以转成其它任意基础类型）
		cateId := parse.Convert(cacheId, 0)
		// 数据库中读取
		po := context.MysqlContextIns.ProCate.Where("Id", cateId).ToEntity()
		do := mapper.Single[procate.DomainObject](&po)
		// bool 表示数据是否存在
		return do, do.Id > 0
	})
	// 当集合不存于Redis时，则到数据库中获取（如果你的应用不需要，这里就不需要定义）
	procateCache.SetListSource(func() collections.List[procate.DomainObject] {
		// 数据库中读取
		lst := context.MysqlContextIns.ProCate.ToList()
		return mapper.ToList[procate.DomainObject](lst)
	})

	container.Register(func() procate.Repository {
		return &ProCateRepository{}
	})
}

type ProCateRepository struct { // IRepository 通用的仓储接口
	data.IRepository[procate.DomainObject]
	Cache cache.ICacheManage[procate.DomainObject] `inject:"procate"` // 通过容器注入，获得实例
}

func (p *ProCateRepository) ToEntity(cateId any) procate.DomainObject {
	item, _ := p.Cache.GetItem(cateId)
	return item
}

func (p *ProCateRepository) ToList() collections.List[procate.DomainObject] {
	return p.Cache.Get()
}

func (p *ProCateRepository) Count() int64 {
	return int64(p.Cache.Count())
}

func (p *ProCateRepository) Add(product procate.DomainObject) {
	po := mapper.Single[model.ProCatePO](&product)
	_ = context.MysqlContextIns.ProCate.Insert(&po)
	p.Cache.SaveItem(product)
}
