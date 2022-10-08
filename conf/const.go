package conf

const (
	CTX_CONTEXT_USER   = "userInfo"
	IS_DELETE          = 2
	UN_DELETE          = 1
	IS_VALID           = 1
	UN_VALID           = 2
	DEFAULT_STATION_ID = 1
	DEFAULT_PAGE_SIZE  = 20
	START_PAGE_NUM     = 1
)

const (
	USER_MANAGER_DEFAULT_TYPE = 1
	USER_DRIVER_DEFAULT_TYPE  = 2
	USER_DRIVER_TITLE         = "司机"
	USER_DRIVER_ROLE_ID       = 2
	USER_ADMIN_ROLE_ID        = 1
	DEFAULT_PASS_WORD         = "e978e272e335e0e5da73fd70b1a9be51"
	DEFAULT_USER_STATUS       = 1
)

const (
	ORDER_TYPE_PURCHASE    = 1
	ORDER_TYPE_SALE        = 2
	ORDER_TYPE_SENT_DIRECT = 3

	ORDER_STATUS_WAIT      = 1
	ORDER_STATUS_PASS      = 2
	ORDER_STATUS_REJECT    = 3
	ORDER_STATUS_FAILURE   = 4
	ORDER_STATUS_FINISH    = 5
	ORDER_STATUS_HAS_TRUCK = 6

	TRUCK_ORDER_STATUS_WAIT    = 1
	TRUCK_ORDER_STATUS_PASS    = 2
	TRUCK_ORDER_STATUS_REJECT  = 3
	TRUCK_ORDER_STATUS_FAILURE = 4
	TRUCK_ORDER_STATUS_FINISH  = 5

	//WEIGH_ORDER_STATUS_CREATED           = 1
	WEIGH_ORDER_STATUS_PROCESS           = 1
	WEIGH_ORDER_STATUS_FAILURE           = 2
	WEIGH_ORDER_STATUS_WAREHOUSE_CONFIRM = 3
	WEIGH_ORDER_STATUS_WAITE_FINISH      = 4
	WEIGH_ORDER_STATUS_FINISH            = 5

	ORDER_GOODS_UNIT = "吨"

	ORDER_PURCHASE_TYPE    = 1
	ORDER_SALE_TYPE        = 2
	ORDER_SENT_DIRECT_TYPE = 3
)

const (
	ANALYSIS_DATE_TYPE_DAY   = "day"
	ANALYSIS_DATE_TYPE_MONTH = "month"
	ANALYSIS_DATE_TYPE_YEAR  = "year"

	ANALYSIS_ORDER_TYPE_PURCHASE    = "采购单"
	ANALYSIS_ORDER_TYPE_SALE        = "销售单"
	ANALYSIS_ORDER_TYPE_SENT_DIRECT = "直发/倒短单"
	ANALYSIS_ORDER_TYPE_TRUCK       = "派车单"
	ANALYSIS_ORDER_TYPE_WEIGH       = "过磅单"
)

var ANALYSIS_ORDER_TYPE_MAP = map[string]int{
	"采购单":    1,
	"销售单":    2,
	"直发/倒短单": 3,
}

var GENDER_ENUM = map[int]string{
	1: "男",
	2: "女",
	3: "未知",
}

var ORDER_TYPE_MAP = map[int]string{
	1: "采购单",
	2: "销售单",
	3: "直发/倒短",
}

var TRUCK_ORDER_TYPE_MAP = map[int]string{
	1: "采购单",
	2: "销售单",
	3: "直发/倒短",
}

var TRUCK_ORDER_IS_LIMIT_LOAD = map[int]string{
	1: "不限重",
	2: "限重",
}

var PAYMENT_METHOD_MAP = map[int]string{
	1: "供方结算",
	2: "客户结算",
	3: "无需结算",
}

// 1-等待审核 2-审核通过 3-审核驳回 4-失效 5-结束
var ORDER_STATUS_MAP = map[int]string{
	1: "等待审核",
	2: "审核通过",
	3: "审核驳回",
	4: "失败",
	5: "结束",
	6: "已派车",
}

// 1-等待审核 2-审核通过 3-审核驳回 4-失效 5-结束
var TRUCK_ORDER_STATUS_MAP = map[int]string{
	1: "等待审核",
	2: "审核通过",
	3: "审核驳回",
	4: "失败",
	5: "结束",
}

var WEIGH_ORDER_STATUS_MAP = map[int]string{
	1: "运输中",
	2: "已失效",
	3: "仓库确认",
	4: "待完成",
	5: "已完成",
}

var DEVICE_TRUCK_ORDER_STATUS_MAP = map[int]string{
	1: "wait",
	2: "pass",
	3: "reject",
	4: "invalid",
}
