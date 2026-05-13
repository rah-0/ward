package config

import "sync"

var (
	mu           sync.RWMutex
	ruleRegistry = map[uint32]map[uint32]any{}
)

func RuleSet(typeID, ruleID uint32, rule any) error {
	mu.Lock()
	defer mu.Unlock()
	if ruleRegistry[typeID] == nil {
		ruleRegistry[typeID] = map[uint32]any{}
	}
	if _, exists := ruleRegistry[typeID][ruleID]; exists {
		return ErrRuleAlreadyRegistered
	}
	ruleRegistry[typeID][ruleID] = rule
	return nil
}

func RuleGet(typeID, ruleID uint32) (any, error) {
	mu.RLock()
	defer mu.RUnlock()
	byType, ok := ruleRegistry[typeID]
	if !ok {
		return nil, ErrRuleNotFound
	}
	rule, ok := byType[ruleID]
	if !ok {
		return nil, ErrRuleNotFound
	}
	return rule, nil
}

func RuleList() map[uint32][]uint32 {
	mu.RLock()
	defer mu.RUnlock()
	list := make(map[uint32][]uint32, len(ruleRegistry))
	for typeID, rules := range ruleRegistry {
		ids := make([]uint32, 0, len(rules))
		for ruleID := range rules {
			ids = append(ids, ruleID)
		}
		list[typeID] = ids
	}
	return list
}
