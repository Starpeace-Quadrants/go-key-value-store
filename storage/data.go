package storage

import (
	"encoding/json"
	"sync"
)

type KeyValueStoreAccess interface {
	Get(key string) interface{}
	GetOrDefault(key string, defaultVal interface{}) interface{}
	Exists(key string) bool

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
	Delete(key string)
	Flush()
}

type Storage struct {
	data map[string]interface{}
	mu   sync.RWMutex
}

func New() KeyValueStoreAccess {
	return &Storage{data: make(map[string]interface{})}
}

func NewFromJsonString(jsonString string) KeyValueStoreAccess {
	store := New()

	var jsonKeyValues map[string]interface{}
	if err := json.Unmarshal([]byte(jsonString), &jsonKeyValues); err != nil {
		panic(err)
	}

	for key, value := range jsonKeyValues {
		store.Set(key, value)
	}

	return store
}

func (s *Storage) Get(key string) interface{} {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.data[key]
}

func (s *Storage) GetOrDefault(key string, defaultVal interface{}) interface{} {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if val, ok := s.data[key]; ok {
		return val
	}
	return defaultVal
}

func (s *Storage) Exists(key string) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	_, exists := s.data[key]
	return exists
}

func (s *Storage) GetBytes(key string) []byte {
	if val, ok := s.Get(key).([]byte); ok {
		return val
	}
	return nil
}

func (s *Storage) GetString(key string) string {
	if val, ok := s.Get(key).(string); ok {
		return val
	}
	return ""
}

func (s *Storage) GetStrings(key string) []string {
	if val, ok := s.Get(key).([]string); ok {
		return val
	}
	return nil
}

func (s *Storage) GetBool(key string) bool {
	if val, ok := s.Get(key).(bool); ok {
		return val
	}
	return false
}

func (s *Storage) GetInt(key string) int {
	switch v := s.Get(key).(type) {
	case int:
		return v
	case float64:
		return int(v)
	}
	return 0
}

func (s *Storage) GetInt8(key string) int8 {
	switch v := s.Get(key).(type) {
	case int8:
		return v
	case float64:
		return int8(v)
	}
	return 0
}

func (s *Storage) GetInt16(key string) int16 {
	switch v := s.Get(key).(type) {
	case int16:
		return v
	case float64:
		return int16(v)
	}
	return 0
}

func (s *Storage) GetInt32(key string) int32 {
	switch v := s.Get(key).(type) {
	case int32:
		return v
	case float64:
		return int32(v)
	}
	return 0
}

func (s *Storage) GetInt64(key string) int64 {
	switch v := s.Get(key).(type) {
	case int64:
		return v
	case float64:
		return int64(v)
	}
	return 0
}

func (s *Storage) GetFloat32(key string) float32 {
	switch v := s.Get(key).(type) {
	case float32:
		return v
	case float64:
		return float32(v)
	}
	return 0
}

func (s *Storage) GetFloat64(key string) float64 {
	if val, ok := s.Get(key).(float64); ok {
		return val
	}
	return 0
}

func (s *Storage) Set(key string, val interface{}) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data[key] = val
}

func (s *Storage) Delete(key string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.data, key)
}

func (s *Storage) Flush() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data = make(map[string]interface{})
}
