package safe_gob

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"reflect"
	"sort"
	"strings"
)

type Pack struct {
	data      interface{}
	signature string
}

func NewPack(data interface{}, secretKey string) *Pack {
	signature := makeSignature(data, secretKey)
	return &Pack{
		data:      data,
		signature: signature,
	}
}

func (p *Pack) GetData() interface{} {
	return p.data
}

func (p *Pack) GetSignature() string {
	return p.signature
}

func makeSignature(data interface{}, secretKey string) string {

	if val := reflect.ValueOf(data); val.Kind() == reflect.Struct {
		keys := make([]string, 0)
		numberOfFields := val.NumField()
		for i := 0; i < numberOfFields; i++ {
			fieldName := val.Type().Field(i).Name
			keys = append(keys, fieldName)
		}
		sort.Strings(keys)
		fields := make([]string, 0)
		for _, key := range keys {
			field := reflect.Indirect(val).FieldByName(key)
			fieldBytes := field.Bytes()
			fields = append(fields, fmt.Sprintf("%v=%v", key, fieldBytes))
		}
		joinedFields := strings.Join(fields, "&")
		signature := getSigned(joinedFields, secretKey)
		return signature
	}
	dataString := fmt.Sprintf("%v", data)
	signature := getSigned(dataString, secretKey)
	return signature
}

func getSigned(param string, secretKey string) string {
	sig := hmac.New(sha256.New, []byte(secretKey))
	sig.Write([]byte(param))
	signature := hex.EncodeToString(sig.Sum(nil))
	return signature
}
