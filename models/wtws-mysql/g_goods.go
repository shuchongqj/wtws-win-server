package wtws_mysql

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego/logs"
	"reflect"
	"strings"
	"sync"
	"time"
	"wtws-server/conf"

	"github.com/astaxie/beego/orm"
)

type GGoods struct {
	Id            int       `orm:"column(goods_id);auto" description:"id" json:"goodsID"`
	CategoryId    int       `orm:"column(category_id)" description:"分类id（销售和采购）" json:"categoryID"`
	GoodNo        string    `orm:"column(goods_no);size(64);null" description:"产品编号" json:"goodsNo"`
	Name          string    `orm:"column(name);size(32)" description:"产品名称" json:"name"`
	Specification float64   `orm:"column(specification)" description:"产品规格" json:"specification"`
	BagWeight     float64   `orm:"column(bag_weight)" description:"编织袋重" json:"bagWeight"`
	DeductWeight  float64   `orm:"column(deduct_weight)" description:"每吨扣Kg数（吨）" json:"deductWeight"`
	ExtraWeight   float64   `orm:"column(extra_weight)" description:"允许额外可配发重量	" json:"extraWeight"`
	IsDelete      int8      `orm:"column(is_delete)" description:"是否已删除  1-未删除  2-已删除" json:"isDelete"`
	UpdateTime    time.Time `orm:"column(update_time);type(datetime)" description:"记录更新时间" json:"updateTime"`
	InsertTime    time.Time `orm:"column(insert_time);type(datetime)" description:"记录创建时间" json:"insertTime"`
}

type GoodListItem struct {
	GGoods
	CategoryName string `json:"categoryName"`
}

func (t *GGoods) TableName() string {
	return "g_goods"
}

func init() {
	orm.RegisterModel(new(GGoods))
}

// AddGGoods insert a new GGoods into database and returns
// last inserted Id on success.
func AddGGoods(m *GGoods) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetGGoodsById retrieves GGoods by Id. Returns error if
// Id doesn't exist
func GetGGoodsById(id int) (v *GGoods, err error) {
	o := orm.NewOrm()
	v = &GGoods{Id: id, IsDelete: conf.UN_DELETE}
	if err = o.Read(v, "Id", "IsDelete"); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllGGoods retrieves all GGoods matches certain condition. Returns empty list if
// no records exist
func GetAllGGoods(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(GGoods))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, v == "true" || v == "1")
		} else {
			qs = qs.Filter(k, v)
		}
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []GGoods
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateGGoods updates GGoods by Id and returns error if
// the record to be updated doesn't exist
func UpdateGGoodsById(m *GGoods) (err error) {
	o := orm.NewOrm()
	v := GGoods{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteGGoods deletes GGoods by Id and returns error if
// the record to be deleted doesn't exist
func DeleteGGoods(id int) (err error) {
	o := orm.NewOrm()
	v := GGoods{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&GGoods{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

func GetGoodsList(pageNum, pageSize, categoryID int, goodsName, goodsNo string) (list []GoodListItem, count int, err error) {
	o := orm.NewOrm()
	qd, _ := orm.NewQueryBuilder("mysql")
	qd.Select("gg.*,gc.name as category_name,gc.category_id as category_id").
		From("g_goods as gg").
		LeftJoin("g_category as gc").
		On("gg.category_id = gc.category_id").
		And("gc.is_delete = 1").
		Where("gg.is_delete = 1")

	if categoryID > 0 {
		qd.And(fmt.Sprintf("gg.category_id = %d", categoryID))
	}

	if len(goodsName) > 0 {
		qd.And("gg.name LIKE '%" + goodsName + "%'")
	}

	if len(goodsNo) > 0 {
		qd.And("gg.goods_no LIKE '%" + goodsNo + "%'")
	}

	qd.OrderBy("gg.goods_id desc")

	countSql := qd.String()
	qd.Limit(pageSize).Offset((pageNum - 1) * pageSize)
	sql := qd.String()

	wg := &sync.WaitGroup{}
	wg.Add(2)

	list = []GoodListItem{}
	var countInt64 int64 = 0
	var getListErr, getCountErr error

	go func() {
		if _, getListErr = o.Raw(sql).QueryRows(&list); getListErr != nil {
			logs.Error("[mysql]  查询货物失败，失败信息:", getListErr.Error())
		}
		wg.Done()
	}()

	go func() {
		if countInt64, getCountErr = o.Raw(countSql).QueryRows(&[]GoodListItem{}); getCountErr != nil {
			logs.Error("[mysql]  查询货物失败，失败信息:", getCountErr.Error())
		}
		wg.Done()
	}()

	wg.Wait()

	if getListErr != nil {
		return []GoodListItem{}, 0, getListErr
	}

	if getCountErr != nil {
		return []GoodListItem{}, 0, getListErr
	}

	return list, int(countInt64), nil
}

func UpdateByGoodsId(u *GGoods, cols []string) (err error) {
	o := orm.NewOrm()
	v := GGoods{Id: u.Id}

	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(u, cols...); err == nil {
			logs.Info("[mysql]  Number of g_goods records updated in database:", num)
		}
	}
	return err
}

func GetAllGoodsList() []GoodListItem {
	o := orm.NewOrm()
	qd, _ := orm.NewQueryBuilder("mysql")
	qd.Select("*").From("g_goods").Where("is_delete = 1")
	sql := qd.String()
	list := []GoodListItem{}
	if _, err := o.Raw(sql).QueryRows(&list); err != nil {
		logs.Error("[mysql]  查询所有的货品失败，失败信息：", err.Error())
		return []GoodListItem{}
	}

	return list
}

func GetGoodsAndCategoryByGoodsID(goodsID int) (*GoodListItem, error) {
	o := orm.NewOrm()
	qd, _ := orm.NewQueryBuilder("mysql")

	qd.Select("gg.*,gc.name as category_name").
		From("g_goods as gg").
		LeftJoin("g_category as gc").
		On("gg.category_id = gc.category_id").
		Where("gg.is_delete = 1").And("gg.goods_id = ?")
	sql := qd.String()
	data := []GoodListItem{}
	if _, err := o.Raw(sql, goodsID).QueryRows(&data); err != nil {
		logs.Error("[mysql]  根据货品ID查询货品失败，失败信息：", err.Error())
		return nil, err
	} else if len(data) == 0 {
		logs.Error("[mysql]  根据货品ID查询货品失败，未查询到货品信息，货品ID：", goodsID)
		return nil, errors.New(" 根据货品ID查询货品失败，未查询到货品信息")
	}

	return &data[0], nil

}
