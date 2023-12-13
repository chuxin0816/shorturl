package query

const (
	prefix = "shorturl:"
)

func GetPrefix(key string) string {
	return prefix + key
}
