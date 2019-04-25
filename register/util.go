package register

import (
	"fmt"
	"io"
	"os"
	"time"
	"bytes"
	"strconv"
	"strings"
	"runtime"
	"path/filepath"
	"crypto/md5"
	"encoding/json"
	"encoding/binary"
	"log"
)

func JsonEncode(data interface{}) string {
	json, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err.Error())
	}
	return string(json)
}

func JsonDecode(str string, fields interface{}) {
	err := json.Unmarshal([]byte(str), &fields)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func Implode(split string, array map[string]string) string {
	var str string
	for _, v := range array {
		str += v + split
	}
	return strings.Trim(str, split)
}

func Md5File(path string) string {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	hash := md5.New()
	io.Copy(hash, file)

	return fmt.Sprintf("%x", hash.Sum(nil))
}

func Md5(str string) string {
	hash := md5.New()
	hash.Write([]byte(str))

	return fmt.Sprintf("%x", hash.Sum(nil))
}

func Time() int64 {
	return time.Now().Unix()
}

func IntToStr(data interface{}) string {
	switch value := data.(type) {
	case int:
		return strconv.Itoa(value) // int to str
	case int64:
		return strconv.FormatInt(value, 10) // int64 to str
	case uint32:
		return strconv.FormatUint(uint64(value), 10)
	case uint64:
		return strconv.FormatUint(value, 10)
	default:
		return ""
	}
}

func StrToInt(str string) int {
	data, _ := strconv.Atoi(str)
	return data
}

func StrToByte(str string) byte {
	return []byte(str)[0]
}

func StrToInt64(str string) int64 {
	data, _ := strconv.ParseInt(str, 10, 64)
	return data
}

func PathExist(_path string) bool {
	_, err := os.Stat(_path)
	if err != nil && os.IsNotExist(err) {
		return false
	}
	return true
}

func Min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func Substr(s string, pos int, length int) string {
	runes := []rune(s)
	l := pos + length
	if l > len(runes) {
		l = len(runes)
	}
	return string(runes[pos:l])
}

/**
BytesCombine 多个[]byte数组合并成一个[]byte
 */
func BytesCombine(pBytes ...[]byte) []byte {
	return bytes.Join(pBytes, []byte(""))
}

func GetFuncName() string {
	pc, _, _, _ := runtime.Caller(1)
	funcName := runtime.FuncForPC(pc).Name()
	funcName = filepath.Ext(funcName)
	funcName = strings.TrimPrefix(funcName, ".")
	return funcName
}

func IsExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

func pack(w io.Writer, data interface{}) error {
	return binary.Write(w, binary.BigEndian, data)
}

func getGID() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)
	return n
}

// 默认基础目录
func defaultBaseDir() string {
	dir, err := filepath.Abs(filepath.Dir(""))
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}

func checkExist(path []string) []string {
	var exist []string
	for _, p := range path {
		if !IsExist(p) {
			continue
		}
		exist = append(exist, p)
	}
	return exist
}