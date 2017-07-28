package A17

type Rectangle struct {
	figuredCoordinates []coordinate

	height int
	width int
	offset int
}

func NewRectangle(height, width, offset int) *Rectangle {
	r := new(Rectangle)

	r.height = height
	r.width = width
	r.offset = offset

	return r
}

func (r *Rectangle) GetFiguredCoordinates() []coordinate {
	return r.figuredCoordinates
}

func (r *Rectangle) SetFiguredCoordinates(origin coordinate) {
	var originX int
	for k := range origin {
		originX = k
	}

	originY := origin[originX]

	limitX := originX + r.width - 1
	limitY := originY + r.height - 1

	switch {
	case r.height == 1 && r.width == 1:
		coordinate := coordinate{originX: originY}
		r.figuredCoordinates = append(r.figuredCoordinates, coordinate)
	case r.height == 1:
		for x := originX; x <= limitX; x++ {
			coordinate := coordinate{x: originY}
			r.figuredCoordinates = append(r.figuredCoordinates, coordinate)
		}
	case r.width == 1:
		for y := originY; y <= limitY; y++ {
			coordinate := coordinate{originX: y}
			r.figuredCoordinates = append(r.figuredCoordinates, coordinate)
		}
	default:
		for y := originY; y <= limitY; y++ {
			for x := originX; x <= limitX; x++ {
				coordinate := coordinate{x: y}
				r.figuredCoordinates = append(r.figuredCoordinates, coordinate)
			}
		}
	}
}

func (r *Rectangle) GetWidth() int {
	return r.width
}

func (r *Rectangle) GetHeight() int {
	return r.height
}

func (r *Rectangle) GetOffset() int {
	return r.offset
}



