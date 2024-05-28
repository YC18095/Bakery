package function

import (
	"strconv"
	"strings"
)

func GetParamUint(str string) uint32 {
	res, err := strconv.ParseUint(str, 0, 32)
	if err != nil {
		return 0
	}
	return uint32(res)
}

func GetParamInt(str string) int {
	res, err := strconv.ParseUint(str, 0, 32)
	if err != nil {
		return 0
	}
	return int(res)
}

func GetParamUintArray(str string) string {
	res := ""
	arr := strings.Split(str, ",")
	for _, itm := range arr {
		if GetParamUint(itm) > 0 {
			if len(res) > 0 {
				res += ","
			}
			res += itm
		}
	}
	return res
}
