package services

type QueryReadmodel interface {
	Query(table string, limit int) ([]map[string]interface{}, error)
}
