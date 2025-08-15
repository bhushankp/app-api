package db

import "fmt"

type MockDB struct {
	Healthy bool
}

func (m *MockDB) Ping() error {
	if !m.Healthy {
		return fmt.Errorf("db unhealthy")
	}
	return nil
}
