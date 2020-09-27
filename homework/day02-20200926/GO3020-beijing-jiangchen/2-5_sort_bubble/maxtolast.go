package main

//MaxToLast ...
func MaxToLast(idxmax int, max int, idxend int, origin *[]int) {
	copy(*origin, (*origin)[:idxmax])
	copy((*origin)[idxmax:idxend], (*origin)[idxmax+1:])
	(*origin)[idxend] = max
}
