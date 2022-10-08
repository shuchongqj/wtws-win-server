package common

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	uuid "github.com/satori/go.uuid"
	"math/rand"
	"reflect"
	"regexp"
	"strconv"
	"time"
	common_struct "wtws-server/common/common-struct"
)

func UUID() string {
	uuid4 := uuid.NewV4()
	uuids := uuid4.String()
	//通过函数进行替换
	re3, _ := regexp.Compile("-")
	//把匹配的所有字符a替换成b
	rep2 := re3.ReplaceAllString(uuids, "")
	// fmt.Println(rep2)
	return rep2
}

// 求交集
func Intersect(slice1, slice2 []string) []string {
	m := make(map[string]int)
	nn := make([]string, 0)
	for _, v := range slice1 {
		m[v]++
	}

	for _, v := range slice2 {
		times, _ := m[v]
		if times == 1 {
			nn = append(nn, v)
		}
	}
	return nn
}

// 求差集 slice1-并集
func Difference(slice1, slice2 []string) []string {
	m := make(map[string]int)
	nn := make([]string, 0)
	inter := Intersect(slice1, slice2)
	for _, v := range inter {
		m[v]++
	}

	for _, value := range slice1 {
		times, _ := m[value]
		if times == 0 {
			nn = append(nn, value)
		}
	}
	return nn
}

// 求并集
func Union(slice1, slice2 []string) []string {
	m := make(map[string]int)
	for _, v := range slice1 {
		m[v]++
	}

	for _, v := range slice2 {
		times, _ := m[v]
		if times == 0 {
			slice1 = append(slice1, v)
		}
	}
	return slice1
}

// value type interace 有数据的结构体 binding type interface 要修改的结构体
func StructAssign(binding interface{}, value interface{}) {
	bVal := reflect.ValueOf(binding).Elem() //获取reflect.Type类型
	vVal := reflect.ValueOf(value).Elem()   //获取reflect.Type类型
	vTypeOfT := vVal.Type()
	for i := 0; i < vVal.NumField(); i++ {
		// 在要修改的结构体中查询有数据结构体中相同属性的字段，有则修改其值
		name := vTypeOfT.Field(i).Name
		if ok := bVal.FieldByName(name).IsValid(); ok {
			bVal.FieldByName(name).Set(reflect.ValueOf(vVal.Field(i).Interface()))
		}
	}
}

// GetTyrePatternMapString 获取花纹map 数组string
func GetTyrePatternMapString(patternNum int8, initPatterns []float32) (string, error) {
	var patternMapBytes []byte
	switch patternNum {
	case 2:
		patternMapBytes, _ = json.Marshal(common_struct.PatternMap1{
			Pattern1: initPatterns[0],
			Pattern2: initPatterns[1],
		})

	case 3:
		patternMapBytes, _ = json.Marshal(common_struct.PatternMap2{
			Pattern1: initPatterns[0],
			Pattern2: initPatterns[1],
			Pattern3: initPatterns[2],
		})
	case 4:
		patternMapBytes, _ = json.Marshal(common_struct.PatternMap3{
			Pattern1: initPatterns[0],
			Pattern2: initPatterns[1],
			Pattern3: initPatterns[2],
			Pattern4: initPatterns[3],
		})
	case 5:
		patternMapBytes, _ = json.Marshal(common_struct.PatternMap4{
			Pattern1: initPatterns[0],
			Pattern2: initPatterns[1],
			Pattern3: initPatterns[2],
			Pattern4: initPatterns[3],
			Pattern5: initPatterns[4],
		})
	case 6:
		patternMapBytes, _ = json.Marshal(common_struct.PatternMap5{
			Pattern1: initPatterns[0],
			Pattern2: initPatterns[1],
			Pattern3: initPatterns[2],
			Pattern4: initPatterns[3],
			Pattern5: initPatterns[4],
			Pattern6: initPatterns[5],
		})
	default:
		return "", errors.New("花纹个数有误")
	}
	return string(patternMapBytes), nil
}

func ExportMonthTyreExcel(excelHeader []string, dates []common_struct.TyreMonthReportItem, sheetName string) (buffer *bytes.Buffer) {
	wf := excelize.NewFile()
	wf.SetSheetName("Sheet1", sheetName)

	headerMap := []interface{}{}
	for _, item := range excelHeader {
		var itemInterface interface{}
		itemInterface = item
		headerMap = append(headerMap, itemInterface)
	}

	wf.SetSheetRow(sheetName, "A1", &headerMap)
	for i, v := range dates {
		wf.SetSheetRow(sheetName, "A"+strconv.Itoa(i+2), &[]interface{}{
			v.VehicleNumber, v.WheelPosition, v.FirstTestPattern, v.FirstTestTime, v.LastTestPattern, v.FirstTestTime})
	}

	buffer, _ = wf.WriteToBuffer()

	return buffer

}

// ExcelDateToDate excel中的时间转成time对象
func ExcelDateToDate(excelDate string) (excelTime time.Time, err error) {
	excelTime = time.Date(1899, time.December, 30, 0, 0, 0, 0, time.UTC)
	var days int
	if days, err = strconv.Atoi(excelDate); err == nil {
		return excelTime.Add(time.Second * time.Duration(days*86400)), nil
	}
	return excelTime, err

}

var DefaultLetters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
var NumberLetters = []rune("0123456789")

func GenerateRandNumber(n int, allowedChars []rune) string {
	var letters []rune

	if len(allowedChars) == 0 {
		letters = DefaultLetters
	} else {
		letters = allowedChars
	}

	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return string(b)
}

func GenerateCode(len int8) string {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	bytes := make([]byte, len)
	for i := int8(0); i < len; i++ {
		c := r.Intn(2)
		if c == 0 {
			b := r.Intn(10) + 48
			bytes[i] = byte(b)
		} else {
			b := r.Intn(26) + 65
			bytes[i] = byte(b)
		}
	}
	return string(bytes)
}

func GenerateOrderNo(prefix string) string {
	nowDate := time.Now()
	nowDateTimeStr := nowDate.Format("20060102150405")
	nowDateTimeStr = fmt.Sprintf("%s%d", nowDateTimeStr, nowDate.Nanosecond()/1e6)

	return fmt.Sprintf("%s%s%s", prefix, nowDateTimeStr, GenerateCode(3))

}
