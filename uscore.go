package uscore

import (
	"strings"
	"strconv"
)

var Needle string
var found bool

var stringCB func(str string, k string)

var values []string


/**
 * HasVal: looks recursively deep to see if the map/array has the value
 */

func HasVal(haystack map[string]interface{}, needle string) bool {
	Needle = needle
	found = false
	stringCB = handleString
	handleMap(haystack)
	return found
}


/**
 * HasKey: looks recursively and returns an array of values for
 *         provided key
 */
func HasKey(haystack interface{}, needle string) ([]string, bool) {

	str, strOk := haystack.(string)
	mpp, mppOk := haystack.(map[string]interface{})
	arr, arrOk := haystack.([]interface{})

	Needle = needle
	found = false
	values = []string{}
	stringCB = addValuesForKey

	if strOk {
		return []string{str}, false
	} else if mppOk {
		handleMap(mpp)
	} else if arrOk {
		handleArray(arr)
	}

	return values, found
}


func handleMap(mp map[string]interface{}){
	for key, val := range mp {
		str, strOk := val.(string)
		mpp, mppOk := val.(map[string]interface{})
		arr, arrOk := val.([]interface{})

		if strOk {
			stringCB(str, key)
		} else if mppOk {
			handleMap(mpp)
		} else if arrOk {
			handleArray(arr)
		}
	}
}

func handleArray(arr []interface{}){
	for x:=0; x<len(arr); x++ {
		n := arr[x]

		str, strOk := n.(string)
		mpp, mppOk := n.(map[string]interface{})
		arr, arrOk := n.([]interface{})

		if strOk {
			stringCB(str, strconv.Itoa(x))
		} else if mppOk {
			handleMap(mpp)
		} else if arrOk {
			handleArray(arr)
		}

	}
}

func handleString(s string, k string){
	if strings.Contains(strings.ToLower(s), strings.ToLower(Needle)) {
		found = true
	}
}

func addValuesForKey(s string, k string) {
	if strings.EqualFold(k, Needle) {
		found = true
		values = append(values, s)
	}
}