package good

type ProductDetail struct {
    Title       string   // 商品标题
	Price       float64  // 价格
    Description string   // 商品描述
    Specs       string   // 规格参数 
    Origin      string   // 原产地
    StorageTips string   // 存储建议
    ShelfLife   int      // 保质期（天数）
	pic_urls    []string // 图片URL列表

}

