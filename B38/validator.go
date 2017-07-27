package B38

type Validator struct {
}

func (v *Validator) CheckImportValue(c Combination) bool {
	switch {
	case c.LegSum == 0,
		c.HeadSum == 0,
		c.CraneLeg == 0 && c.TurtleLeg == 0,
		c.LegSum < c.TurtleLeg,
		c.LegSum < c.CraneLeg,
		c.LegSum < (c.CraneLeg + c.TurtleLeg): return false
	default: return true
	}
}