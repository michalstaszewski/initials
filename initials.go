package initials

func GetInitials(firstName string, lastName string) string {
	fInitial := string([]rune(firstName)[0:1])
	lInitial := string([]rune(lastName)[0:1])

	return fInitial + lInitial
}