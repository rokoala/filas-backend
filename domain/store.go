package domain

// Store - Store Domain
// Store contains an ordered consumer queue
type Store struct {
	ID      string      `bson:"_id,omitempty"`
	Name    string      `bson:"name,omitempty"`
	URLName string      `bson:"urlname,omitempty"`
	Queue   []*Consumer `bson:"queue,omitempty"`
}

// filas.app/outback
