package query

const (
	prefix = "shorturl:"
	KeyURL
)

func GetPrefix(key string) string {
	return prefix + key
}
