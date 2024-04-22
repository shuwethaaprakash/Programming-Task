package count

import (
	"sort"
	"sync"
)

type SafeMap struct {
	sync.Map
	Mu sync.Mutex
}

func IncreaseCount(m *SafeMap, key string) {
	m.Mu.Lock()
	defer m.Mu.Unlock()
	
	// Load current value
	value, ok := m.Load(key)
	if !ok {
		m.Store(key, 0)
		value = 0
	}
	newValue := value.(int) + 1

	// Store the updated value
	m.Store(key, newValue)
}

func GetTopThree(m *SafeMap) []string {
	// Slice to store key-value pairs
	var pairs []struct {
		key   string
		value int
	}

	// Load all key-value pairs from the sync.Map
	m.Range(func(key, value interface{}) bool {
		pairs = append(pairs, struct {
			key   string
			value int
		}{key.(string), value.(int)})
		return true
	})

	// Sort key-value pairs by value in descending order
	sort.SliceStable(pairs, func(i, j int) bool {
		return pairs[i].value > pairs[j].value
	})

	// Retrieve top 3 keys
	var topThreeKeys []string
	for i := 0; i < len(pairs) && i < 3; i++ {
		topThreeKeys = append(topThreeKeys, pairs[i].key)
	}

	return topThreeKeys
}

func GetMapLength(m *SafeMap) int {
	length := 0
	m.Range(func(_, _ interface{}) bool {
		length++
		return true
	})
	return length
}
