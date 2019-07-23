package services

type Item struct {
	Value float64
	Year  int
	Month int
	Index int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int {
	return len(pq)
}

// We'll inverse this to get newest dates first
func (pq PriorityQueue) Less(i, j int) bool {
	ii := pq[i]
	ij := pq[j]
	if ij.Year < ii.Year {
		return true
	} else if ij.Year > ii.Year {
		return false
	}
	return ij.Month < ii.Month
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.Index = -1
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.Index = n
	*pq = append(*pq, item)
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}
