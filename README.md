# Go Key Value Store

A super simple package created for adding data to events initially.

Basically a simple `map[string]interface{}`.

While useful for storing, I needed to mix types and retrieve values back in their correct type.

Now correct type based on knowing the type, the package does not somehow decide what type it is
and cast, it simply provides accessors to cast, so it's ready for use.

The following methods are available:

```go
Get(key string) interface{}
GetBytes(key string) []byte
GetString(key string) string
GetStrings(key string) []string
GetBool(key string) bool
GetInt(key string) int
GetInt8(key string) int8
GetInt16(key string) int16
GetInt32(key string) int32
GetInt64(key string) int64
GetFloat32(key string) float32
GetFloat64(key string) float64
Set(key string, val interface{})
Flush()
Delete(key string)
```

## Usage

```go
import (
	"github.com/ronappleton/go-key-value-store/storage"
)
```

You can then use one of the two constructors:

```go
New()
NewFromJsonString(jsonString string)
```

Both return an instance of `github.com/ronappleton/go-key-value-store/storage.Storage` that implements the
`KeyValueStoreAccess` interface.

## Extension

To make it more usable for your purpose you may wish to add a new cast for example `map[string]interface{}` so
you can nest.

```go
package your_package

import (
	"github.com/ronappleton/go-key-value-store/storage"
)

type AddsMap interface {
	GetMap(key string) map[string]interface{}
}

type StorageWithMap struct {
	storage.Storage
	AddsMap
}

func New() StorageWithMap {
	return StorageWithMap{}
}

func (e *StorageWithMap) GetMap(key string) map[string]interface{} {
	return e.Get(key).(map[string]interface{})
}
```

Then you just use your own constructor, so:

```go
var somevar your_package.StorageWithMap = your_package.New() 
```

And you can then store and access your maps and everything else.
