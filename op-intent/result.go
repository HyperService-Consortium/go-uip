package opintent

import (
	"encoding/json"
	"github.com/tidwall/gjson"
	"reflect"
	"time"
)

type ResultsI interface {
	Index(int) ResultI
	Len() int
}

type ResultI interface {
	// String returns a string representation of the value.
	String() string
	// Bool returns an boolean representation.
	Bool() bool
	// Int returns an integer representation.
	Int() int64
	// Uint returns an unsigned integer representation.
	Uint() uint64
	// Float returns an float64 representation.
	Float() float64
	// Time returns a time.Time representation.
	Time() time.Time
	// Array returns back an array of values.
	// If the result represents a non-existent value, then an empty array will be
	// returned. If the result is not a JSON array, the return value will be an
	// array containing one result.
	Array() ResultsI
	// IsObject returns true if the result value is a JSON object.
	IsObject() bool
	// IsArray returns true if the result value is a JSON array.
	IsArray() bool
	// ForEach iterates through values.
	// If the result represents a non-existent value, then no values will be
	// iterated. If the result is an Object, the iterator will pass the key and
	// value of each item. If the result is an Array, the iterator will only pass
	// the value of each item. If the result is not a JSON array or object, the
	// iterator will pass back one value equal to the result.
	//ForEach(iterator func(key, value ResultI) bool)
	// Map returns back an map of values. The result should be a JSON array.
	//Map() map[string]ResultI
	// Get searches result for the specified path.
	// The result should be a JSON array or object.
	Get(path string) ResultI
	// Exists returns true if value exists.
	//
	//  if gjson.Get(json, "name.last").Exists(){
	//		println("value exists")
	//  }
	Exists() bool
	// Value returns one of these types:
	//
	//	bool, for JSON booleans
	//	float64, for JSON numbers
	//	Number, for JSON numbers
	//	string, for JSON string literals
	//	nil, for JSON null
	//	map[string]interface{}, for JSON objects
	//	[]interface{}, for JSON arrays
	//
	Value() interface{}
	RawBytes() ([]byte, error)
	// Less return true if a token is less than another token.
	// The caseSensitive paramater is used when the tokens are Strings.
	// The order when comparing two different type is:
	//
	//  Null < False < Number < String < True < JSON
	//
	//Less(token ResultI, caseSensitive bool) bool
}

type GJSONResult struct {
	gjson.Result
}

func (G GJSONResult) MarshalJSON() ([]byte, error) {
	return G.RawBytes()
}

func (G *GJSONResult) UnmarshalJSON(b []byte) error {
	G.Result = gjson.ParseBytes(b)
	return nil
}

func (G GJSONResult) RawBytes() ([]byte, error) {
	return []byte(G.Raw), nil
}

type GJSONResults []gjson.Result

func (G GJSONResults) Index(i int) ResultI {
	return GJSONResult{Result: G[i]}
}

func (G GJSONResults) Len() int {
	return len(G)
}

func (G GJSONResult) Array() ResultsI {
	return GJSONResults(G.Result.Array())
}

func (G GJSONResult) Get(path string) ResultI {
	return GJSONResult{Result: G.Result.Get(path)}
}

func NewGJSONResult(contents []byte) (ResultI, error) {
	return GJSONResult{Result: gjson.ParseBytes(contents)}, nil
}

type mObj map[string]interface{}
type MapResult struct {
	m mObj
	e interface{}
}

func (m MapResult) String() string {
	if m.e == nil {
		return ""
	}
	return m.e.(string)
}

func (m MapResult) Bool() bool {
	if m.e == nil {
		return false
	}
	return m.e.(bool)
}

var iT = reflect.TypeOf(int64(0))

func (m MapResult) Int() int64 {
	if m.e == nil {
		return 0
	}
	return reflect.ValueOf(m.e).Convert(iT).Int()
}

var uiT = reflect.TypeOf(uint64(0))

func (m MapResult) Uint() uint64 {
	if m.e == nil {
		return 0
	}
	return reflect.ValueOf(m.e).Convert(uiT).Uint()
}

func (m MapResult) Float() float64 {
	if m.e == nil {
		return 0
	}
	return reflect.ValueOf(m.e).Float()
}

func (m MapResult) Time() time.Time {
	if m.e == nil {
		return time.Time{}
	}
	return m.e.(time.Time)
}

func (m MapResult) Array() ResultsI {
	return NewMapResults(m.e.([]interface{}))
}

func (m MapResult) IsObject() bool {
	return m.m != nil
}

func (m MapResult) IsArray() (ok bool) {
	_, ok = m.e.([]interface{})
	return
}

func (m MapResult) Get(path string) ResultI {
	i, _ := m.m[path]
	r, _ := NewMapResult(i)
	return r
}

func (m MapResult) Exists() bool {
	return reflect.ValueOf(m.e).IsValid()
}

func (m MapResult) Value() interface{} {
	return m.e
}

func (m MapResult) RawBytes() ([]byte, error) {
	return json.Marshal(m.m)
}

func NewMapResult(m interface{}) (MapResult, error) {
	return MapResult{
		m: maybeMap(m),
		e: m,
	}, nil
}

type MapResults []MapResult

func (m MapResults) Index(i int) ResultI {
	return m[i]
}

func (m MapResults) Len() int {
	return len(m)
}

func NewMapResults(is []interface{}) ResultsI {
	mps := make(MapResults, len(is))
	for i := range is {
		mps[i], _ = NewMapResult(is[i])
	}
	return mps
}

func maybeMap(m interface{}) mObj {
	if m, ok := m.(mObj); ok {
		return m
	} else {
		return nil
	}
}
