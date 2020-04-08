package document

import (
	"encoding/json"
	"fmt"
	"github.com/tidwall/gjson"
	"reflect"
	"strconv"
	"time"
)

type Document interface {
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
	Array() Documents
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
	//ForEach(iterator func(key, value Document) bool)
	// Map returns back an map of values. The result should be a JSON array.
	//Map() map[string]Document
	// Get searches result for the specified path.
	// The result should be a JSON array or object.
	Get(path string) Document
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
	//Less(token Document, caseSensitive bool) bool
}

type GJSONDocument struct {
	gjson.Result
}

func (G GJSONDocument) MarshalJSON() ([]byte, error) {
	return G.RawBytes()
}

func (G *GJSONDocument) UnmarshalJSON(b []byte) error {
	G.Result = gjson.ParseBytes(b)
	return nil
}

func (G GJSONDocument) RawBytes() ([]byte, error) {
	return []byte(G.Raw), nil
}

type GJSONDocuments []gjson.Result

func (G GJSONDocuments) Index(i int) Document {
	return GJSONDocument{Result: G[i]}
}

func (G GJSONDocuments) Len() int {
	return len(G)
}

func (G GJSONDocument) Array() Documents {
	return GJSONDocuments(G.Result.Array())
}

func (G GJSONDocument) Get(path string) Document {
	return GJSONDocument{Result: G.Result.Get(path)}
}

func NewGJSONDocument(contents []byte) (GJSONDocument, error) {
	return GJSONDocument{Result: gjson.ParseBytes(contents)}, nil
}

type MObj map[string]interface{}
type MapDocument struct {
	m MObj
	e interface{}
}

func (m MapDocument) String() string {
	if m.e == nil {
		return ""
	}
	v := reflect.ValueOf(m.e)
	switch v.Kind() {
	case reflect.String:
		return m.e.(string)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(v.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return strconv.FormatUint(v.Uint(), 10)
	case reflect.Float32, reflect.Float64:
		return strconv.FormatFloat(v.Float(), 'f', 10, 64)
	default:
		panic(fmt.Errorf("bad convertion to string: %v", v.Type()))
	}
}

func (m MapDocument) Bool() bool {
	if m.e == nil {
		return false
	}
	return m.e.(bool)
}

var iT = reflect.TypeOf(int64(0))

func (m MapDocument) Int() int64 {
	if m.e == nil {
		return 0
	}
	return reflect.ValueOf(m.e).Convert(iT).Int()
}

var uiT = reflect.TypeOf(uint64(0))

func (m MapDocument) Uint() uint64 {
	if m.e == nil {
		return 0
	}
	return reflect.ValueOf(m.e).Convert(uiT).Uint()
}

func (m MapDocument) Float() float64 {
	if m.e == nil {
		return 0
	}
	return reflect.ValueOf(m.e).Float()
}

func (m MapDocument) Time() time.Time {
	if m.e == nil {
		return time.Time{}
	}
	return m.e.(time.Time)
}

func (m MapDocument) Array() Documents {
	d, _ := NewMapDocuments(m.e.([]interface{}))
	return d
}

func (m MapDocument) IsObject() bool {
	return m.m != nil
}

func (m MapDocument) IsArray() (ok bool) {
	_, ok = m.e.([]interface{})
	return
}

func (m MapDocument) Get(path string) Document {
	i, _ := m.m[path]
	r, _ := NewMapDocument(i)
	return r
}

func (m MapDocument) Exists() bool {
	return reflect.ValueOf(m.e).IsValid()
}

func (m MapDocument) Value() interface{} {
	return m.e
}

func (m MapDocument) RawBytes() ([]byte, error) {
	return json.Marshal(m.m)
}

func NewMapDocument(m interface{}) (MapDocument, error) {
	return MapDocument{
		m: MaybeMap(m),
		e: m,
	}, nil
}

type MapDocuments []MapDocument

func (m MapDocuments) Index(i int) Document {
	return m[i]
}

func (m MapDocuments) Len() int {
	return len(m)
}

func NewMapDocuments(is []interface{}) (MapDocuments, error) {
	mps := make(MapDocuments, len(is))
	var err error
	for i := range is {
		mps[i], err = NewMapDocument(is[i])
		if err != nil {
			return nil, err
		}
	}
	return mps, err
}

func MaybeMap(m interface{}) MObj {
	if m, ok := m.(MObj); ok {
		return m
	} else {
		return nil
	}
}

type Documents interface {
	Index(int) Document
	Len() int
}
