package util

import "reflect"

type DataUtil struct{

}
func (this *DataUtil) GetMapData (keyList []string,data interface{})map[string]interface{}{
	params := make (map[string]interface{})
	reflectVal := reflect.ValueOf(data)
	for _,val := range keyList{
		field := reflectVal.Elem().FieldByName(val) //获取指定Field
		switch field.Kind() {
		case reflect.String:
			if field.String() != ""{
				params[val] = field.String()
			}
		case reflect.Int:
			if field.Int() != 0{
				params[val] = field.Int()
			}
		case reflect.Int64:
			if field.Int() != 0{
				params[val] = field.Int()
			}
		case reflect.Bool:
			params[val] = field.Bool()
		case reflect.Slice:
			params[val] = field.Slice(0,-1)
		case reflect.Map:
			mapRange := field.MapRange()
			newMap := make(map[string]interface{})
			for mapRange.Next() {
				newMap[mapRange.Key().String()] = mapRange.Value()
			}
			params[val] = newMap
		}
	}
	return params
}

func (this *DataUtil) StructToMap(structObj interface{},mapObj *map[string]interface{})error{
	obj1 := reflect.TypeOf(structObj)
	obj2 := reflect.ValueOf(structObj)
	for i := 0; i < obj1.NumField(); i++ {
		(*mapObj)[obj1.Field(i).Name] = obj2.Field(i).Interface()
	}
	return nil
}
