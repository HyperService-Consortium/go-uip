package opintent

import (
	"github.com/tidwall/gjson"
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
