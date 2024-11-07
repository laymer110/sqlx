package sqlx

import (
	"strings"
	"unsafe"
)

/*
主要增加了这一页的功能，
Oracle的字段需要转换为小写
结构体默认转换为蛇形
将多余的字段自动忽略，只绑定已经枚举的结果
保留sql.Null格式的字段，需要进一步参考gorm如何处理的
*/

func b2s(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func isUpper(c byte) bool {
	return 'A' <= c && c <= 'Z'
}

// user_code -> UserCode
func titleCasedName(name string) string {
	newstr := make([]byte, 0, len(name))
	upNextChar := true

	name = strings.ToLower(name)

	for i := 0; i < len(name); i++ {
		c := name[i]
		switch {
		case upNextChar:
			upNextChar = false
			if 'a' <= c && c <= 'z' {
				c -= 'a' - 'A'
			}
		case c == '_':
			upNextChar = true
			continue
		}

		newstr = append(newstr, c)
	}

	kk := b2s(newstr)
	return kk
}

func sliceToLower(ss []string) []string {
	for k, v := range ss {
		ss[k] = strings.ToLower(v)
	}
	return ss
}

// UserCode -> user_code
func snakeCasedName(name string) string {
	newstr := make([]byte, 0, len(name)+1)
	for i := 0; i < len(name); i++ {
		c := name[i]
		//默认最大，不添加下划线
		var s byte = 'Z'
		if i > 1 {
			s = name[i-1]
		}
		if isUpper(c) {

			if i > 0 && !isUpper(s) {
				//大写字母前面插入_
				newstr = append(newstr, '_')
			}
			//大写转小写
			c += 'a' - 'A'
		}
		newstr = append(newstr, c)
	}
	dd := b2s(newstr)
	return dd
}
