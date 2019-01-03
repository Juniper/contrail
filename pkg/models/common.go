package models

func boolMap(ss []string) map[string]bool {
	r := map[string]bool{}
	for _, s := range ss {
		r[s] = true
	}
	return r
}
