package infrastructure

import (
	"fullExample/domain/procate"
	"fullExample/domain/product"
	"fullExample/infrastructure/repository"
	"github.com/farseer-go/data"
	"github.com/farseer-go/eventBus"
	"github.com/farseer-go/fs/container"
	"github.com/farseer-go/fs/modules"
	"github.com/farseer-go/queue"
	"github.com/farseer-go/redis"
)

type Module struct {
}

func (module Module) DependsModule() []modules.FarseerModule {
	// 使用到了redis模块、data(orm)模块、eventBus（事件总线）模块、queue（本地队列）模块
	return []modules.FarseerModule{redis.Module{}, data.Module{}, eventBus.Module{}, queue.Module{}}
}
func (module Module) PostInitialize() {
	// 初始化产品分类仓储
	repository.InitProCate()
	// 初始化产品仓储
	repository.InitProduct()
	// 初始化产品仓储
	repository.InitStock()

	// 为了方便演示，自动初始化产品分类
	procateRepository := container.Resolve[procate.Repository]()
	if procateRepository.Count() == 0 {
		procateRepository.Add(procate.DomainObject{Id: 1, Name: "苹果"})
		procateRepository.Add(procate.DomainObject{Id: 2, Name: "科沃斯"})
		procateRepository.Add(procate.DomainObject{Id: 3, Name: "游戏机"})
		procateRepository.Add(procate.DomainObject{Id: 4, Name: "电饭煲"})
		procateRepository.Add(procate.DomainObject{Id: 5, Name: "Bose"})
		procateRepository.Add(procate.DomainObject{Id: 6, Name: "戴森"})
	}

	// 为了方便演示，自动初始化10条产品记录
	productRepository := container.Resolve[product.Repository]()
	if productRepository.Count() == 0 {
		productRepository.Add(product.DomainObject{Cate: product.CateEO{Id: 1, Name: "苹果"}, Stock: 100, Price: 8999, Caption: "Apple iPhone 14 Pro Max", Desc: "Apple iPhone 14 Pro Max (A2896) 512GB 银色 支持移动联通电信5G 双卡双待手机", ImgSrc: "https://img10.360buyimg.com/n1/s450x450_jfs/t1/52631/5/21888/20074/63191b75E9234fd5f/1ce070cb00b6f896.jpg.avif"})
		productRepository.Add(product.DomainObject{Cate: product.CateEO{Id: 6, Name: "戴森"}, Stock: 100, Price: 4500, Caption: "戴森吸尘器V12", Desc: "戴森（DYSON）V12 Detect Slim Fluffy轻量手持吸尘器大吸力", ImgSrc: "https://img13.360buyimg.com/n1/jfs/t1/19688/36/16587/116628/6392d71cEaf3517f1/f8540270aa837711.jpg.avif"})
		productRepository.Add(product.DomainObject{Cate: product.CateEO{Id: 3, Name: "游戏机"}, Stock: 100, Price: 4300, Caption: "微软Xbox Series主机", Desc: "微软（Microsoft） 【国内保税仓】微软Xbox Series主机 次时代家用4K高清游戏主机 Xbox Series X日版", ImgSrc: "https://img11.360buyimg.com/n1/jfs/t1/148482/19/14661/44451/5fb37e27Ed4e29739/b803e7e6e730458f.jpg"})
		productRepository.Add(product.DomainObject{Cate: product.CateEO{Id: 2, Name: "科沃斯"}, Stock: 100, Price: 4600, Caption: "科沃斯扫地机器人 T10", Desc: "科沃斯扫地机器人 T10 OMNI扫拖一体机 吸拖洗烘一体拖地机器人洗地机擦地机 智能全自动集尘清洗", ImgSrc: "https://img14.360buyimg.com/n1/jfs/t1/221921/29/21654/66050/6343dbe6E65851991/26b01cc077765280.jpg.avif"})
		productRepository.Add(product.DomainObject{Cate: product.CateEO{Id: 4, Name: "电饭煲"}, Stock: 100, Price: 4999, Caption: "福库电饭煲", Desc: "韩国原装进口电饭锅3 IH加热高压无压可切换家用智能多功能双压电饭煲1-4人份", ImgSrc: "https://img12.360buyimg.com/n1/jfs/t1/220183/21/12895/369141/6214afa9Ed1efc3de/4c234fe4872bfe1b.jpg.avif"})
		productRepository.Add(product.DomainObject{Cate: product.CateEO{Id: 5, Name: "Bose"}, Stock: 100, Price: 43000, Caption: "Bose LifeStyle 650 家庭影院5", Desc: "Bose LifeStyle 650 博士电视音响 家庭影院5.1 回音壁 客厅 多功能 boss 黑色", ImgSrc: "https://img10.360buyimg.com/n1/jfs/t1/194207/27/22333/61466/6264f81aE57fd1558/fe3b5d2006e6f144.jpg.avif"})
		productRepository.Add(product.DomainObject{Cate: product.CateEO{Id: 5, Name: "Bose"}, Stock: 100, Price: 2999, Caption: "Bose 700 无线消噪耳机-银色", Desc: "Bose 700 无线消噪耳机-银色 手势触控蓝牙降噪耳机 主动降噪头戴式耳机 长久续航", ImgSrc: "https://img11.360buyimg.com/n1/s450x450_jfs/t1/9392/40/17588/329181/62a867caE1ac9df9c/076838c67d546f0f.jpg.avif"})
		productRepository.Add(product.DomainObject{Cate: product.CateEO{Id: 1, Name: "苹果"}, Stock: 100, Price: 14999, Caption: "Apple MacBook Pro 14英寸", Desc: "8核中央处理器 14核图形处理器 16G 256G 深空灰 笔记本", ImgSrc: "https://img14.360buyimg.com/n1/s450x450_jfs/t1/197505/24/13599/83205/616dc631E854d2563/e98b96e5044af9da.jpg.avif"})
		productRepository.Add(product.DomainObject{Cate: product.CateEO{Id: 1, Name: "苹果"}, Stock: 100, Price: 1400, Caption: "Apple苹果 Apple TV 7代 (2022款) ", Desc: "Apple苹果 Apple TV 7代 (2022款) 128GB WIFI+Ethernet版 A15仿生", ImgSrc: "https://img12.360buyimg.com/n1/jfs/t1/59455/13/22349/10684/6355f732E5940e6fe/edd0f8570d454be2.jpg"})
		productRepository.Add(product.DomainObject{Cate: product.CateEO{Id: 6, Name: "戴森"}, Stock: 100, Price: 2999, Caption: "戴森吹风机 HD08 长春花蓝礼盒款", Desc: "戴森(Dyson) 吹风机 Dyson Supersonic 电吹风 负离子 进口家用 礼物推荐 HD08 长春花蓝礼盒款", ImgSrc: "https://img14.360buyimg.com/n1/jfs/t1/126497/34/30895/26643/63490c3dEbc9de636/9cd45e32ca3b1e05.jpg"})
	}
}
