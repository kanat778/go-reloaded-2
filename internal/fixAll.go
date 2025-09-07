package internal

func FixAll(b []byte) []string {
	arr := BytesToStrArray(b)
	arr = CommandFix(arr)
	arr = ArticleFix(arr)
	return arr
}
