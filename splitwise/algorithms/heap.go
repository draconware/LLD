package algorithm

import "github.com/mastik5h/LLD/splitwise/models"

type MaxHeap []models.UserBalanceMap
type MinHeap []models.UserBalanceMap

func GetMaxHeap(elems []models.UserBalanceMap) *MaxHeap {
	h := MaxHeap{}
	for _, elem := range elems {
		h = append(h, elem)
	}
	return &h
}

func GetMinHeap(elems []models.UserBalanceMap) *MinHeap {
	h := MinHeap{}
	for _, elem := range elems {
		h = append(h, elem)
	}
	return &h
}

func (h MaxHeap) Len() int {
	return len(h)
}

func (h MaxHeap) Less(i, j int) bool {
	return h[i].Amount.Amount < h[j].Amount.Amount
}

func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MaxHeap) Push(x any) {
	*h = append(*h, x.(models.UserBalanceMap))
}

func (h *MaxHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	return h[i].Amount.Amount > h[j].Amount.Amount
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x any) {
	*h = append(*h, x.(models.UserBalanceMap))
}

func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
