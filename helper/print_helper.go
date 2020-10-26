package helper

const (
	PRINT_MAX_LEN = 1024
)

func TruncateData(data []byte) []byte {

	if len(data) <= PRINT_MAX_LEN {
		return data
	} else {
		return data[:PRINT_MAX_LEN]
	}
}
