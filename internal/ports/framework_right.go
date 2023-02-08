package ports

type DBPort interface {
	CloseDB()
	AddtoHistory(result int32, operation string) error
}
