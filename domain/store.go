package domain

// Store - Store Domain
// Store contains an ordered consumer queue
type Store struct {
	ID      string      `bson:"_id,omitempty" json:"_id"`
	Name    string      `bson:"name,omitempty" json:"name"`
	URLName string      `bson:"urlname,omitempty" json:"urlName"`
	Queue   []*Consumer `bson:"queue,omitempty" json:"queue"`
}

// filas.app/outback
