package storage

type schema struct {
	tableName  string
	domainList []domain
}

//return sizes of this schema in bytes
func (schema *schema) SchemaSizeInBytes() int {
	return 0
}
