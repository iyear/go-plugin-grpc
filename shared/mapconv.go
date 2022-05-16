package shared

import (
	"google.golang.org/protobuf/types/known/structpb"
	"sync"
)

type MapConv struct {
	spb *structpb.Struct
	mu  sync.RWMutex
}

//	╔════════════════════════╤════════════════════════════════════════════╗
//	║ Go type                │ Conversion                                 ║
//	╠════════════════════════╪════════════════════════════════════════════╣
//	║ nil                    │ stored as NullValue                        ║
//	║ bool                   │ stored as BoolValue                        ║
//	║ int, int32, int64      │ stored as NumberValue                      ║
//	║ uint, uint32, uint64   │ stored as NumberValue                      ║
//	║ float32, float64       │ stored as NumberValue                      ║
//	║ string                 │ stored as StringValue; must be valid UTF-8 ║
//	║ []byte                 │ stored as StringValue; base64-encoded      ║
//	║ map[string]interface{} │ stored as StructValue                      ║
//	║ []interface{}          │ stored as ListValue                        ║
//	╚════════════════════════╧════════════════════════════════════════════╝
//

func NewMapConv(spb *structpb.Struct) *MapConv {
	if spb.GetFields() == nil {
		spb.Fields = make(map[string]*structpb.Value)
	}
	return &MapConv{
		spb: spb,
		mu:  sync.RWMutex{},
	}
}

func (c *MapConv) String() string {
	return c.spb.String()
}

func (c *MapConv) Map() map[string]interface{} {
	return c.spb.AsMap()
}

func (c *MapConv) Get(key string) (*structpb.Value, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	v, ok := c.spb.GetFields()[key]
	return v, ok
}

func (c *MapConv) MustGet(key string) *structpb.Value {
	if value, exists := c.Get(key); exists {
		return value
	}
	panic("Key \"" + key + "\" does not exist")
}

func (c *MapConv) GetInterface(key string) interface{} {
	if value, exists := c.Get(key); exists {
		return value.AsInterface()
	}
	return nil
}

func (c *MapConv) GetBool(key string) (b bool) {
	if val, ok := c.Get(key); ok && val != nil {
		b = val.GetBoolValue()
	}
	return
}

func (c *MapConv) GetString(key string) (s string) {
	if val, ok := c.Get(key); ok && val != nil {
		s = val.GetStringValue()
	}
	return
}

func (c *MapConv) GetFloat64(key string) (f float64) {
	if val, ok := c.Get(key); ok && val != nil {
		f = val.GetNumberValue()
	}
	return
}

func (c *MapConv) GetFloat32(key string) (f float32) {
	return float32(c.GetFloat64(key))
}

func (c *MapConv) GetInt(key string) (i int) {
	return int(c.GetFloat64(key))
}

func (c *MapConv) GetInt64(key string) (i int64) {
	return int64(c.GetFloat64(key))
}

func (c *MapConv) GetInt32(key string) (i int32) {
	return int32(c.GetFloat64(key))
}

func (c *MapConv) GetUint(key string) (u uint) {
	return uint(c.GetFloat64(key))
}

func (c *MapConv) GetUint32(key string) (u uint32) {
	return uint32(c.GetFloat64(key))
}

func (c *MapConv) GetUint64(key string) (u uint64) {
	return uint64(c.GetFloat64(key))
}

// GetSliceIter return false to stop iteration
func (c *MapConv) GetSliceIter(key string, f func(v *structpb.Value) bool) {
	if val, ok := c.Get(key); ok && val != nil {
		if list := val.GetListValue(); list != nil {
			for _, v := range list.GetValues() {
				if !f(v) {
					break
				}
			}
		}
	}
}

func (c *MapConv) GetSlice(key string) []interface{} {
	interfaces := make([]interface{}, 0)
	c.GetSliceIter(key, func(v *structpb.Value) bool {
		interfaces = append(interfaces, v.AsInterface())
		return true
	})

	return interfaces
}
