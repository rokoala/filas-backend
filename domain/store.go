package domain

// Store - Domínio do estabelecimento
// O establecimento contém um fila de consumidores em ordem;
type Store struct {
	ID      string
	Name    string
	URLName string
	Queue   []*Consumer
}

// filas.app/outback
