package src

type Set[T comparable] struct {
	_map map[T]bool
}

func (this *Set[T]) append(val T) {
	if !this.contains(val) {
		this._map[val] = true
	}
}
func (this *Set[T]) contains(val T) bool {
	_, exist := this._map[val]
	return exist
}

func NewSet[T comparable]() Set[T] {
	_map := make(map[T]bool)
	return Set[T]{_map: _map}
}

func removeDuplicates(nums []int) int {
	tmp := make([]int, len(nums))
	set := NewSet[int]()
	i := 0
	for _, num := range nums {
		if !set.contains(num) {
			tmp[i] = num
			i++

			set.append(num)
		}
	}
	nums = tmp
	return i

}
