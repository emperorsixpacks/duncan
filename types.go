package duncan

type Connection interface {
	GetConnectionName() string
	ConnectionString() string
}
