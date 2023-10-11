package json

import iter "github.com/json-iterator/go"

// 定义JSON操作
var (
	json          = iter.ConfigFastest
	Marshal       = json.Marshal
	Unmarshal     = json.Unmarshal
	MarshalIndent = json.MarshalIndent
	NewDecoder    = json.NewDecoder
	NewEncoder    = json.NewEncoder
	Get           = json.Get
)

// MarshalToString JSON编码为字符串
func MarshalToString(v interface{}) string {
	s, err := json.MarshalToString(v)
	if err != nil {
		return ""
	}
	return s
}

var (
	// Marshal JSON编码
	_ = Marshal
	// Unmarshal JSON解码
	_ = Unmarshal
	// MarshalIndent JSON编码为格式化的字符串
	_ = MarshalIndent
	// NewDecoder JSON解码器
	_ = NewDecoder
	// NewEncoder JSON编码器
	_ = NewEncoder
	// MarshalToString JSON编码为字符串
	_ = MarshalToString
	// Get JSON获取
	_ = Get
)
