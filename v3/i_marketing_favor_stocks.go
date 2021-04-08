package v3

import "strconv"

// 条件查询批次列表API https://pay.weixin.qq.com/wiki/doc/apiv3/wxpay/marketing/convention/chapter3_4.shtml
// 通过此接口可查询多个批次的信息，包括批次的配置信息以及批次概况数据。

// 条件查询批次列表
func MarketingFavorStocks(cfg *Config, offset, limit int, stock_creator_mchid string, create_start_time, create_end_time string, status string) (o *MarketingFavorStocksRes, err error) {
	err = Call(cfg, "/v3/marketing/favor/stocks"+
		"?offset="+strconv.Itoa(offset)+
		"&limit="+strconv.Itoa(offset)+
		"&stock_creator_mchid="+stock_creator_mchid+
		"&create_start_time="+create_start_time+
		"&create_end_time="+create_end_time+
		"&status="+status,
		"GET",
		nil, &o)
	return o, err
}

type MarketingFavorStocksRes struct {
	Total_count int64
	Offset      int64
	Limit       int64
	Data        []*MarketingFavorStocksData
}

type MarketingFavorStocksData struct {
	Stock_id            string
	Stock_creator_mchid string
	Stock_name          string
	Status              string
	Create_time         string
	Description         string
	//stock_use_rule
	Available_begin_time string
	Available_end_time   string
	Distributed_coupons  int64
	No_cash              bool
	Singleitem           bool
	Stock_type           string
}
