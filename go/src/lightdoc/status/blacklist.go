package status

var Folders = []string{
	".git",
	".vsocde",
	".idea",
}

func Exist(str string, arr []string) bool {
	for _, s := range arr {
		if s == str {
			return true
		}
	}
	return false
}
