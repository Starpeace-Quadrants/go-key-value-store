package storage_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/ronappleton/go-key-value-store/storage"
)

func TestStorage_SetAndGet(t *testing.T) {
	s := storage.New()

	s.Set("string", "hello")
	s.Set("int", 42)
	s.Set("bool", true)
	s.Set("float64", 3.14)
	s.Set("slice", []string{"a", "b", "c"})

	assert.Equal(t, "hello", s.GetString("string"))
	assert.Equal(t, 42, s.GetInt("int"))
	assert.Equal(t, true, s.GetBool("bool"))
	assert.InDelta(t, 3.14, s.GetFloat64("float64"), 0.001)
	assert.Equal(t, []string{"a", "b", "c"}, s.GetStrings("slice"))
}

func TestStorage_GetWithWrongTypeReturnsZeroValue(t *testing.T) {
	s := storage.New()

	s.Set("value", "not-a-number")

	assert.Equal(t, 0, s.GetInt("value"))
	assert.Equal(t, 0.0, s.GetFloat64("value"))
	assert.Equal(t, false, s.GetBool("value"))
	assert.Equal(t, "", s.GetString("missing"))
	assert.Nil(t, s.GetStrings("value"))
}

func TestStorage_ExistsAndDelete(t *testing.T) {
	s := storage.New()

	s.Set("exists", "yes")
	assert.True(t, s.Exists("exists"))

	s.Delete("exists")
	assert.False(t, s.Exists("exists"))
}

func TestStorage_Flush(t *testing.T) {
	s := storage.New()

	s.Set("one", 1)
	s.Set("two", 2)

	s.Flush()

	assert.False(t, s.Exists("one"))
	assert.False(t, s.Exists("two"))
}

func TestStorage_GetOrDefault(t *testing.T) {
	s := storage.New()

	assert.Equal(t, "default", s.GetOrDefault("missing", "default"))
	s.Set("present", "here")
	assert.Equal(t, "here", s.GetOrDefault("present", "default"))
}

func TestStorage_FromJson(t *testing.T) {
	jsonStr := `{
		"name": "ron",
		"age": 47,
		"enabled": true,
		"height": 6.1
	}`

	s := storage.NewFromJsonString(jsonStr)

	assert.Equal(t, "ron", s.GetString("name"))
	assert.Equal(t, 47, s.GetInt("age"))
	assert.Equal(t, true, s.GetBool("enabled"))
	assert.InDelta(t, 6.1, s.GetFloat64("height"), 0.001)
}
