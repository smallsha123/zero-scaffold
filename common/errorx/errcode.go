package errorx

var message = make(map[uint32]string)

const (
	OK                = 200  // 成功
	ERR_DEFAULT       = 100  // 错误
	ERR_EMPTY         = 400  // 未查到数据
	ERR_PARAMS        = 499  // 参数错误
	ERR_SHOP_EXIST    = 1000 // 店铺已存在
	ERR_SKU_EMPTY     = 1001 // sku信息不能为空
	ERR_PRODUCT_FOUND = 1002 // 商品不存在
)

func init() {
	message[OK] = "请求成功"
	message[ERR_DEFAULT] = "系统未知错误"
	message[ERR_EMPTY] = "未查到数据"
	message[ERR_PARAMS] = "参数错误"
	message[ERR_SHOP_EXIST] = "店铺已存在"
	message[ERR_SKU_EMPTY] = "sku信息不能为空"
	message[ERR_PRODUCT_FOUND] = "商品不存在"
}
