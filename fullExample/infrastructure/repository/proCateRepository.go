package repository

import (
	"fullExample/domain/procate"
	"fullExample/infrastructure/repository/model"
	"github.com/farseer-go/cache"
	"github.com/farseer-go/collections"
	"github.com/farseer-go/data"
	"github.com/farseer-go/fs/container"
	"github.com/farseer-go/fs/parse"
	"github.com/farseer-go/mapper"
	"github.com/farseer-go/redis"
)

// InitProCate 注册产品分类仓储 ioc procate.Repository
func InitProCate() {
	// 初始化数据库上下文
	context := data.NewContext[ProCateRepository]("default", true)
	// 这是一个抽象的缓存集合（可以选择Redis、本地内存作为存储源，这里使用的是Redis作为数据源）
	// 这里为了让Redis能自动同步数据库的数据而使用。
	// key：是全局唯一的，同时也是IOC的别名，用于注入时，区分哪个缓存集合
	context.Cache = redis.SetProfiles[procate.DomainObject]("procate", "Id", 0, "default")
	// 如果单个item不存在，是否需要重新读整个List数据源（这里开启）
	context.Cache.EnableItemNullToLoadAll()
	// 当单个item不存在于Redis时，则要到数据库中获取（如果你的应用不需要，这里就不需要定义）
	context.Cache.SetItemSource(func(cacheId any) (procate.DomainObject, bool) {
		// 万能的基础类型转换，将any转成int（也可以转成其它任意基础类型）
		cateId := parse.Convert(cacheId, 0)
		// 数据库中读取
		po := context.ProCate.Where("Id", cateId).ToEntity()
		do := mapper.Single[procate.DomainObject](&po)
		// bool 表示数据是否存在
		return do, do.Id > 0
	})
	// 当集合不存于Redis时，则到数据库中获取（如果你的应用不需要，这里就不需要定义）
	context.Cache.SetListSource(func() collections.List[procate.DomainObject] {
		// 数据库中读取
		lst := context.ProCate.ToList()
		return mapper.ToList[procate.DomainObject](lst)
	})

	// 通过已存在的实例来注册到IOC
	container.RegisterInstance[procate.Repository](context)
}

type ProCateRepository struct {
	// 定义数据库表映射TableSet
	ProCate data.TableSet[model.ProCatePO] `data:"name=farseer_go_procate"`
	Cache   cache.ICacheManage[procate.DomainObject]
}

func (p *ProCateRepository) ToEntity(cateId int) procate.DomainObject {
	item, _ := p.Cache.GetItem(cateId)
	return item
}

func (p *ProCateRepository) ToList() collections.List[procate.DomainObject] {
	return p.Cache.Get()
}

func (p *ProCateRepository) Count() int {
	return p.Cache.Count()
}

func (p *ProCateRepository) Add(product procate.DomainObject) {
	po := mapper.Single[model.ProCatePO](&product)
	_ = p.ProCate.Insert(&po)
}
