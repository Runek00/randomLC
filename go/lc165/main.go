package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {}

func compareVersion(version1 string, version2 string) int {
	if version1 == version2 {
		return 0
	}
	v1 := strings.Split(version1, ".")
	v2 := strings.Split(version2, ".")
	l1 := len(v1)
	l2 := len(v2)
	maxL := l1
	if l2 > l1 {
		maxL = l2
	}
	fmt.Println(maxL)
	var err error = nil
	for i := 0; i < maxL; i++ {
		vv1 := 0
		if i < l1 {
			t := strings.TrimLeft(v1[i], "0")
			if len(t) == 0 {
				vv1 = 0
			} else {
				vv1, err = strconv.Atoi(t)
			}
		}
		vv2 := 0
		if i < l2 {
			t := strings.TrimLeft(v2[i], "0")

			if len(t) == 0 {

				vv2 = 0

			} else {

				vv2, err = strconv.Atoi(t)
			}

		}

		if vv1 > vv2 {
			return 1
		} else if vv2 > vv1 {
			return -1
		}
		if err != nil {
			fmt.Println(err)
			return 0
		}
	}

	return 0
}

