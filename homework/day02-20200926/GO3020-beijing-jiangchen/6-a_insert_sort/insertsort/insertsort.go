package insertsort

//InsertSort ...
func InsertSort(origin *[]int) {
	var i, j, tmp int
	for i = 1; i < len(*origin); i++ {
		tmp = (*origin)[i]
		for j = i - 1; j >= 0 && (*origin)[j] > tmp; j-- {
			(*origin)[j+1] = (*origin)[j]
		}
		(*origin)[j+1] = tmp
	}
}
