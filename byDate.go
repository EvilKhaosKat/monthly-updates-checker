package main

type ByDate []*Result

func (a ByDate) Len() int {
	return len(a)
}

func (a ByDate) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a ByDate) Less(i, j int) bool {
	iYear := a[i].year
	jYear := a[j].year

	if iYear < jYear {
		return true
	}

	if iYear > jYear {
		return false
	}

	iMonth := a[i].month
	jMonth := a[j].month

	if iMonth < jMonth {
		return true
	}

	if iMonth > jMonth {
		return false
	}

	return true
}