package calc

type Memory struct {
	Store map[string]float64
}

// Initialise la mémoire
func NewMemory() *Memory {
	return &Memory{
		Store: make(map[string]float64),
	}
}

// Stocker un résultat en mémoire
func (m *Memory) StoreResult(key string, value float64) {
	m.Store[key] = value
}

// Récupérer un résultat depuis la mémoire
func (m *Memory) RecallResult(key string) (float64, bool) {
	val, exists := m.Store[key]
	return val, exists
}
