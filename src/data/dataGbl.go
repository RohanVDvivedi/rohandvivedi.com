package data

import (
    "database/sql"
    "strings"
    "reflect"
)

var Db *sql.DB = nil;

type Row interface {
    Scan(...interface{}) error
}

func getRepeatedQueryParamHolders(n int) string {
	return getRepeatedwithDelimeter("?", ",", n);
}

func getRepeatedwithDelimeter(repeatString string, delimeterString string, n int) string {
    if(n == 0) {
        return "";
    }
    return repeatString + strings.Repeat(delimeterString + repeatString, n-1);
}

func convertToInterfaceSlice(slice interface{}) []interface{} {
    s := reflect.ValueOf(slice)
    if s.Kind() != reflect.Slice {
        panic("InterfaceSlice() given a non-slice type")
    }

    ret := make([]interface{}, s.Len())

    for i:=0; i<s.Len(); i++ {
        ret[i] = s.Index(i).Interface()
    }

    return ret
}