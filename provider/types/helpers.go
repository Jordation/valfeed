package types

func (p *Position) TransformWithMapData(xmulti, xscalaradd, ymulti, yscalaradd float64) {
	p.X = p.X * (xmulti * xscalaradd) * 1000
	p.Y = p.Y * (ymulti * yscalaradd) * 1000
}

//func (Round)
