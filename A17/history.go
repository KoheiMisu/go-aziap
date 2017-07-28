package A17

import "../util"

type History struct {
	coordinates []int
}

func NewHistory(screenWidth int) *History {
	h := new(History)

	for i := 0; i < screenWidth; i++ {
		h.coordinates = append(h.coordinates, 0)
	}

	return h
}

func (h *History) GetHeight(f FallenObject) int {
	y := h.coordinates[f.GetOffset()]

	// return max width of range
	for x := f.GetOffset(); x < (f.GetOffset() + f.GetWidth()); x++ {
		if h.coordinates[x] > y {
			y = h.coordinates[x]
		}
	}

	return y
}

func (h *History) SetHeight(f FallenObject) {
	rangeLimit := f.GetOffset() + f.GetWidth()
	arr := h.coordinates[f.GetOffset():rangeLimit]

	baseValue := util.Max(arr)

	for x := f.GetOffset(); x < rangeLimit; x++ {
		h.coordinates[x] = baseValue + f.GetHeight()
	}
}