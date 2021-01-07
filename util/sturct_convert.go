package util

import "reflect"

// Convert 结构体转化，有相同字段的值会被赋值
func Convert(source interface{}, target interface{}) {

	m := GetMap(source)

	keys := reflect.TypeOf(target).Elem()
	values := reflect.ValueOf(target).Elem()

	for index := 0; index < keys.NumField(); index++ {
		sourceValue := m[keys.Field(index).Name]
		if sourceValue.IsValid() {
			values.Field(index).Set(sourceValue)
		}
	}
	Log().Info("convert target is: %v", target)
}

func GetMap(source interface{}) map[string]reflect.Value {

	keys := reflect.TypeOf(source).Elem()
	values := reflect.ValueOf(source).Elem()

	ret := make(map[string]reflect.Value, keys.NumField())

	for index := 0; index < keys.NumField(); index++ {
		ret[keys.Field(index).Name] = values.Field(index)
		Log().Debug("log: key%d is '%v', value is %v \n",
			index, keys.Field(index).Name, values.Field(index))
	}

	return ret
}
