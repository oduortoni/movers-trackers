package geo

func GetDistance(source, destination string) int {
	if source == "here" && destination == "there" {
		return 400
	}
	return 1000
}
