package validata

import (
	"encoding/json"
	"errors"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/validation"
)

// CommonValidate 通用参数验证
func CommonValidate[T any](bytes []byte) (data T, err error) {
	if err = json.Unmarshal(bytes, &data); err != nil {
		logs.Error(err.Error())
		return data, err
	}

	valid := validation.Validation{}
	b, validErr := valid.Valid(&data)
	if validErr != nil {
		logs.Error("参数验证失败。", validErr.Error())
		return data, errors.New("参数验证失败")
	}
	if !b {
		for _, err := range valid.Errors {
			logs.Error(err.Key, err.Message)
		}
	}
	return data, nil
}
