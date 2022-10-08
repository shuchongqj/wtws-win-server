package service

import (
	"sync"
	"time"
	"wtws-server/common"
	common_struct "wtws-server/common/common-struct"
	"wtws-server/conf"
	"wtws-server/dto"
	wtws_mysql "wtws-server/models/wtws-mysql"
	request_entity "wtws-server/service-struct/request-entity"
)

func GetAllGoodsList() common_struct.ResponseStruct {

	list := wtws_mysql.GetAllGoodsList()
	return common.ResponseStatus(0, "", dto.GoodsList{list, len(list)})

}

func GetGoodsList(data *request_entity.GoodsList) common_struct.ResponseStruct {

	if goodsList, count, err := wtws_mysql.GetGoodsList(data.PageNum, data.PageSize, data.CategoryID, data.GoodsName, data.GoodsNo); err != nil {
		return common.ResponseStatus(0, "", dto.GoodsList{
			List:  []wtws_mysql.GoodListItem{},
			Count: 0,
		})
	} else {
		return common.ResponseStatus(0, "", dto.GoodsList{
			List:  goodsList,
			Count: count,
		})
	}

}

func AddGoods(data *request_entity.AddGoods) common_struct.ResponseStruct {

	if id, err := wtws_mysql.AddGGoods(&wtws_mysql.GGoods{
		Name:          data.GoodsName,
		CategoryId:    data.CategoryID,
		GoodNo:        data.GoodNo,
		Specification: data.Specification,
		BagWeight:     data.BagWeight,
		DeductWeight:  data.DeductWeight,
		ExtraWeight:   data.ExtraWeight,
		IsDelete:      conf.UN_DELETE,
		InsertTime:    time.Now(),
		UpdateTime:    time.Now(),
	}); err != nil || id <= 0 {
		return common.ResponseStatus(-13, "", nil)
	}

	return common.ResponseStatus(12, "", nil)
}

func DeleteGoods(data *request_entity.DeleteGoods) common_struct.ResponseStruct {

	wg := &sync.WaitGroup{}
	wg.Add(len(data.GoodsIDs))

	deleteErrs := []error{}

	for _, id := range data.GoodsIDs {
		go func(id int) {
			if err := wtws_mysql.UpdateByGoodsId(&wtws_mysql.GGoods{
				Id:         id,
				IsDelete:   conf.IS_DELETE,
				UpdateTime: time.Now(),
			}, []string{"IsDelete", "UpdateTime"}); err != nil {
				deleteErrs = append(deleteErrs, err)
			}
			wg.Done()
		}(id)
	}

	wg.Wait()

	if len(deleteErrs) > 0 {
		return common.ResponseStatus(-12, "", nil)
	}

	return common.ResponseStatus(16, "", nil)
}

func UpdateGoods(data *request_entity.UpdateGoods) common_struct.ResponseStruct {

	if err := wtws_mysql.UpdateByGoodsId(&wtws_mysql.GGoods{
		Id:         data.GoodsID,
		Name:       data.GoodsName,
		UpdateTime: time.Now(),
	}, []string{"Name", "ContactName", "Tel", "Address", "Type", "UpdateTime"}); err != nil {
		return common.ResponseStatus(-15, "", nil)
	}

	return common.ResponseStatus(14, "", nil)
}
