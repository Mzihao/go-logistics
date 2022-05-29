package utils

// Rev 切片反转
func Rev(slice []map[string]string) []map[string]string {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
	return slice
}
