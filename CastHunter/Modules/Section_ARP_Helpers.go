package CastHunter

func ExterminateExtraVals(cl []string) []string {
	k := make(map[string]bool)
	l := []string{}
	for _, y := range cl {
		if _, u := k[y]; !u {
			k[y] = true
			l = append(l, y)
		}
	}
	return l
}
