package sql_convert

import (
	"database/sql/driver"
	"github.com/go-errors/errors"
	"github.com/pineal-niwan/sensor/json"
)

var (
	//不支持json序列化的类型
	ErrNotJsonSupportType = errors.New(`the data type is not support json`)
)

type Map map[string]interface{}

//json对应的map
type JsonMap struct {
	Map
}

//从数据库读取后组装
func (jsonMap *JsonMap) Scan(src interface{}) error {
	switch src := src.(type) {
	case string:
		if src == "" {
			jsonMap.Map = make(Map)
			return nil
		} else {
			return json.Unmarshal([]byte(src), &jsonMap.Map)
		}
	case []byte:
		if len(src) == 0 {
			jsonMap.Map = make(Map)
			return nil
		} else {
			return json.Unmarshal(src, &jsonMap.Map)
		}
	default:
		return ErrNotJsonSupportType
	}
}

//写回数据库时的组装
func (jsonMap *JsonMap) Value() (driver.Value, error) {
	jsonValue, err := json.Marshal(jsonMap.Map)
	if err != nil {
		return "", err
	}
	return string(jsonValue), nil
}

//Json marshal
func (jsonMap JsonMap) MarshalJSON() ([]byte, error) {
	return json.Marshal(jsonMap.Map)
}

//Json marshal
func (jsonMap *JsonMap) UnmarshalJSON(data []byte) error {
	err := json.Unmarshal(data, &jsonMap.Map)
	return err
}

//将interface转换为JsonMap
func Convert2JsonMap(v interface{}) (JsonMap, bool) {
	var j JsonMap
	ok := true
	switch v.(type) {
	case JsonMap:
		j = v.(JsonMap)
	case map[string]interface{}:
		j.Map = v.(map[string]interface{})
	case Map:
		j.Map = v.(Map)
	default:
		ok = false
	}
	return j, ok
}

//转换map中的key为JsonMap
func ConvertMapVal2JsonMap(m map[string]interface{}, k string) (JsonMap, bool) {
	x, ok := m[k]
	if !ok {
		return JsonMap{}, false
	}
	return Convert2JsonMap(x)
}
