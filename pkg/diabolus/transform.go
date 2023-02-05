package diabolus

type Matrix44 [16]float64

func (m Matrix44) Adj() Matrix44 {
	m00 := m[m.indexOf(1, 1)]*m[m.indexOf(2, 2)]*m[m.indexOf(3, 3)] +
		m[m.indexOf(1, 2)]*m[m.indexOf(2, 3)]*m[m.indexOf(3, 1)] +
		m[m.indexOf(1, 3)]*m[m.indexOf(2, 1)]*m[m.indexOf(3, 2)] -
		m[m.indexOf(1, 3)]*m[m.indexOf(2, 2)]*m[m.indexOf(3, 1)] -
		m[m.indexOf(1, 2)]*m[m.indexOf(2, 1)]*m[m.indexOf(3, 3)] -
		m[m.indexOf(1, 1)]*m[m.indexOf(2, 3)]*m[m.indexOf(3, 2)]
	m01 := -m[m.indexOf(0, 1)]*m[m.indexOf(2, 2)]*m[m.indexOf(3, 3)] -
		m[m.indexOf(0, 2)]*m[m.indexOf(2, 3)]*m[m.indexOf(3, 1)] -
		m[m.indexOf(0, 3)]*m[m.indexOf(2, 1)]*m[m.indexOf(3, 2)] +
		m[m.indexOf(0, 3)]*m[m.indexOf(2, 2)]*m[m.indexOf(3, 1)] +
		m[m.indexOf(0, 2)]*m[m.indexOf(2, 1)]*m[m.indexOf(3, 3)] +
		m[m.indexOf(0, 1)]*m[m.indexOf(2, 3)]*m[m.indexOf(3, 2)]
	m02 := m[m.indexOf(0, 1)]*m[m.indexOf(1, 2)]*m[m.indexOf(3, 3)] +
		m[m.indexOf(0, 2)]*m[m.indexOf(1, 3)]*m[m.indexOf(3, 1)] +
		m[m.indexOf(0, 3)]*m[m.indexOf(1, 1)]*m[m.indexOf(3, 2)] -
		m[m.indexOf(0, 3)]*m[m.indexOf(1, 2)]*m[m.indexOf(3, 1)] -
		m[m.indexOf(0, 2)]*m[m.indexOf(1, 1)]*m[m.indexOf(3, 3)] -
		m[m.indexOf(0, 1)]*m[m.indexOf(1, 3)]*m[m.indexOf(3, 2)]
	m03 := -m[m.indexOf(0, 1)]*m[m.indexOf(1, 2)]*m[m.indexOf(2, 3)] -
		m[m.indexOf(0, 2)]*m[m.indexOf(1, 3)]*m[m.indexOf(2, 1)] -
		m[m.indexOf(0, 3)]*m[m.indexOf(1, 1)]*m[m.indexOf(2, 2)] +
		m[m.indexOf(0, 3)]*m[m.indexOf(1, 2)]*m[m.indexOf(2, 1)] +
		m[m.indexOf(0, 2)]*m[m.indexOf(1, 1)]*m[m.indexOf(2, 3)] +
		m[m.indexOf(0, 1)]*m[m.indexOf(1, 3)]*m[m.indexOf(2, 2)]

	m10 := -m[m.indexOf(1, 0)]*m[m.indexOf(2, 2)]*m[m.indexOf(3, 3)] -
		m[m.indexOf(1, 2)]*m[m.indexOf(2, 3)]*m[m.indexOf(3, 0)] -
		m[m.indexOf(1, 3)]*m[m.indexOf(2, 0)]*m[m.indexOf(3, 2)] +
		m[m.indexOf(1, 3)]*m[m.indexOf(2, 2)]*m[m.indexOf(3, 0)] +
		m[m.indexOf(1, 2)]*m[m.indexOf(2, 0)]*m[m.indexOf(3, 3)] +
		m[m.indexOf(1, 0)]*m[m.indexOf(2, 3)]*m[m.indexOf(3, 2)]
	m11 := m[m.indexOf(0, 0)]*m[m.indexOf(2, 2)]*m[m.indexOf(3, 3)] +
		m[m.indexOf(0, 2)]*m[m.indexOf(2, 3)]*m[m.indexOf(3, 0)] +
		m[m.indexOf(0, 3)]*m[m.indexOf(2, 0)]*m[m.indexOf(3, 2)] -
		m[m.indexOf(0, 3)]*m[m.indexOf(2, 2)]*m[m.indexOf(3, 0)] -
		m[m.indexOf(0, 2)]*m[m.indexOf(2, 0)]*m[m.indexOf(3, 3)] -
		m[m.indexOf(0, 0)]*m[m.indexOf(2, 3)]*m[m.indexOf(3, 2)]
	m12 := -m[m.indexOf(0, 0)]*m[m.indexOf(1, 2)]*m[m.indexOf(3, 3)] -
		m[m.indexOf(0, 2)]*m[m.indexOf(1, 3)]*m[m.indexOf(3, 0)] -
		m[m.indexOf(0, 3)]*m[m.indexOf(1, 0)]*m[m.indexOf(3, 2)] +
		m[m.indexOf(0, 3)]*m[m.indexOf(1, 2)]*m[m.indexOf(3, 0)] +
		m[m.indexOf(0, 2)]*m[m.indexOf(1, 0)]*m[m.indexOf(3, 3)] +
		m[m.indexOf(0, 0)]*m[m.indexOf(1, 3)]*m[m.indexOf(3, 2)]
	m13 := m[m.indexOf(0, 0)]*m[m.indexOf(1, 2)]*m[m.indexOf(2, 3)] +
		m[m.indexOf(0, 2)]*m[m.indexOf(1, 3)]*m[m.indexOf(2, 0)] +
		m[m.indexOf(0, 3)]*m[m.indexOf(1, 0)]*m[m.indexOf(2, 2)] -
		m[m.indexOf(0, 3)]*m[m.indexOf(1, 2)]*m[m.indexOf(2, 0)] -
		m[m.indexOf(0, 2)]*m[m.indexOf(1, 0)]*m[m.indexOf(2, 3)] -
		m[m.indexOf(0, 0)]*m[m.indexOf(1, 3)]*m[m.indexOf(2, 2)]

	m20 := m[m.indexOf(1, 0)]*m[m.indexOf(2, 1)]*m[m.indexOf(3, 3)] +
		m[m.indexOf(1, 1)]*m[m.indexOf(2, 3)]*m[m.indexOf(3, 0)] +
		m[m.indexOf(1, 3)]*m[m.indexOf(2, 0)]*m[m.indexOf(3, 1)] -
		m[m.indexOf(1, 3)]*m[m.indexOf(2, 1)]*m[m.indexOf(3, 0)] -
		m[m.indexOf(1, 1)]*m[m.indexOf(2, 0)]*m[m.indexOf(3, 3)] -
		m[m.indexOf(1, 0)]*m[m.indexOf(2, 3)]*m[m.indexOf(3, 1)]
	m21 := -m[m.indexOf(0, 0)]*m[m.indexOf(2, 1)]*m[m.indexOf(3, 3)] -
		m[m.indexOf(0, 1)]*m[m.indexOf(2, 3)]*m[m.indexOf(3, 0)] -
		m[m.indexOf(0, 3)]*m[m.indexOf(2, 0)]*m[m.indexOf(3, 1)] +
		m[m.indexOf(0, 3)]*m[m.indexOf(2, 1)]*m[m.indexOf(3, 0)] +
		m[m.indexOf(0, 1)]*m[m.indexOf(2, 0)]*m[m.indexOf(3, 3)] +
		m[m.indexOf(0, 0)]*m[m.indexOf(2, 3)]*m[m.indexOf(3, 1)]
	m22 := m[m.indexOf(0, 0)]*m[m.indexOf(1, 1)]*m[m.indexOf(3, 3)] +
		m[m.indexOf(0, 1)]*m[m.indexOf(1, 3)]*m[m.indexOf(3, 0)] +
		m[m.indexOf(0, 3)]*m[m.indexOf(1, 0)]*m[m.indexOf(3, 1)] -
		m[m.indexOf(0, 3)]*m[m.indexOf(1, 1)]*m[m.indexOf(3, 0)] -
		m[m.indexOf(0, 1)]*m[m.indexOf(1, 0)]*m[m.indexOf(3, 3)] -
		m[m.indexOf(0, 0)]*m[m.indexOf(1, 3)]*m[m.indexOf(3, 1)]
	m23 := -m[m.indexOf(0, 0)]*m[m.indexOf(1, 1)]*m[m.indexOf(2, 3)] -
		m[m.indexOf(0, 1)]*m[m.indexOf(1, 3)]*m[m.indexOf(2, 0)] -
		m[m.indexOf(0, 3)]*m[m.indexOf(1, 0)]*m[m.indexOf(2, 1)] +
		m[m.indexOf(0, 3)]*m[m.indexOf(1, 1)]*m[m.indexOf(2, 0)] +
		m[m.indexOf(0, 1)]*m[m.indexOf(1, 0)]*m[m.indexOf(2, 3)] +
		m[m.indexOf(0, 0)]*m[m.indexOf(1, 3)]*m[m.indexOf(2, 1)]

	m30 := -m[m.indexOf(1, 0)]*m[m.indexOf(2, 1)]*m[m.indexOf(3, 2)] -
		m[m.indexOf(1, 1)]*m[m.indexOf(2, 2)]*m[m.indexOf(3, 0)] -
		m[m.indexOf(1, 2)]*m[m.indexOf(2, 0)]*m[m.indexOf(3, 1)] +
		m[m.indexOf(1, 2)]*m[m.indexOf(2, 1)]*m[m.indexOf(3, 0)] +
		m[m.indexOf(1, 1)]*m[m.indexOf(2, 0)]*m[m.indexOf(3, 2)] +
		m[m.indexOf(1, 0)]*m[m.indexOf(2, 2)]*m[m.indexOf(3, 1)]
	m31 := m[m.indexOf(0, 0)]*m[m.indexOf(2, 1)]*m[m.indexOf(3, 2)] +
		m[m.indexOf(0, 1)]*m[m.indexOf(2, 2)]*m[m.indexOf(3, 0)] +
		m[m.indexOf(0, 2)]*m[m.indexOf(2, 0)]*m[m.indexOf(3, 1)] -
		m[m.indexOf(0, 2)]*m[m.indexOf(2, 1)]*m[m.indexOf(3, 0)] -
		m[m.indexOf(0, 1)]*m[m.indexOf(2, 0)]*m[m.indexOf(3, 2)] -
		m[m.indexOf(0, 0)]*m[m.indexOf(2, 2)]*m[m.indexOf(3, 1)]
	m32 := -m[m.indexOf(0, 0)]*m[m.indexOf(1, 1)]*m[m.indexOf(3, 2)] -
		m[m.indexOf(0, 1)]*m[m.indexOf(1, 2)]*m[m.indexOf(3, 0)] -
		m[m.indexOf(0, 2)]*m[m.indexOf(1, 0)]*m[m.indexOf(3, 1)] +
		m[m.indexOf(0, 2)]*m[m.indexOf(1, 1)]*m[m.indexOf(3, 0)] +
		m[m.indexOf(0, 1)]*m[m.indexOf(1, 0)]*m[m.indexOf(3, 2)] +
		m[m.indexOf(0, 0)]*m[m.indexOf(1, 2)]*m[m.indexOf(3, 1)]
	m33 := m[m.indexOf(0, 0)]*m[m.indexOf(1, 1)]*m[m.indexOf(2, 2)] +
		m[m.indexOf(0, 1)]*m[m.indexOf(1, 2)]*m[m.indexOf(2, 0)] +
		m[m.indexOf(0, 2)]*m[m.indexOf(1, 0)]*m[m.indexOf(2, 1)] -
		m[m.indexOf(0, 2)]*m[m.indexOf(1, 1)]*m[m.indexOf(2, 0)] -
		m[m.indexOf(0, 1)]*m[m.indexOf(1, 0)]*m[m.indexOf(2, 2)] -
		m[m.indexOf(0, 0)]*m[m.indexOf(1, 2)]*m[m.indexOf(2, 1)]

	return Matrix44{
		m00, m01, m02, m03,
		m10, m11, m12, m13,
		m20, m21, m22, m23,
		m30, m31, m32, m33,
	}
}

func (m Matrix44) Det() float64 {
	a0 := m[m.indexOf(0, 0)] * (m[m.indexOf(1, 1)]*m[m.indexOf(2, 2)]*m[m.indexOf(3, 3)] +
		m[m.indexOf(1, 2)]*m[m.indexOf(2, 3)]*m[m.indexOf(3, 1)] +
		m[m.indexOf(1, 3)]*m[m.indexOf(2, 1)]*m[m.indexOf(3, 2)] -
		m[m.indexOf(1, 3)]*m[m.indexOf(2, 2)]*m[m.indexOf(3, 1)] -
		m[m.indexOf(1, 2)]*m[m.indexOf(2, 1)]*m[m.indexOf(3, 3)] -
		m[m.indexOf(1, 1)]*m[m.indexOf(2, 3)]*m[m.indexOf(3, 2)])

	a1 := m[m.indexOf(1, 0)] * (m[m.indexOf(0, 1)]*m[m.indexOf(2, 2)]*m[m.indexOf(3, 3)] +
		m[m.indexOf(0, 2)]*m[m.indexOf(2, 3)]*m[m.indexOf(3, 1)] +
		m[m.indexOf(0, 3)]*m[m.indexOf(2, 1)]*m[m.indexOf(3, 2)] -
		m[m.indexOf(0, 3)]*m[m.indexOf(2, 2)]*m[m.indexOf(3, 1)] -
		m[m.indexOf(0, 2)]*m[m.indexOf(2, 1)]*m[m.indexOf(3, 3)] -
		m[m.indexOf(0, 1)]*m[m.indexOf(2, 3)]*m[m.indexOf(3, 2)])

	a2 := m[m.indexOf(2, 0)] * (m[m.indexOf(0, 1)]*m[m.indexOf(1, 2)]*m[m.indexOf(3, 3)] +
		m[m.indexOf(0, 2)]*m[m.indexOf(1, 3)]*m[m.indexOf(3, 1)] +
		m[m.indexOf(0, 3)]*m[m.indexOf(1, 1)]*m[m.indexOf(3, 2)] -
		m[m.indexOf(0, 3)]*m[m.indexOf(1, 2)]*m[m.indexOf(3, 1)] -
		m[m.indexOf(0, 2)]*m[m.indexOf(1, 1)]*m[m.indexOf(3, 3)] -
		m[m.indexOf(0, 1)]*m[m.indexOf(1, 3)]*m[m.indexOf(3, 2)])

	a3 := m[m.indexOf(3, 0)] * (m[m.indexOf(0, 1)]*m[m.indexOf(1, 2)]*m[m.indexOf(2, 3)] +
		m[m.indexOf(0, 2)]*m[m.indexOf(1, 3)]*m[m.indexOf(2, 1)] +
		m[m.indexOf(0, 3)]*m[m.indexOf(1, 1)]*m[m.indexOf(2, 2)] -
		m[m.indexOf(0, 3)]*m[m.indexOf(1, 2)]*m[m.indexOf(2, 1)] -
		m[m.indexOf(0, 2)]*m[m.indexOf(1, 1)]*m[m.indexOf(2, 3)] -
		m[m.indexOf(0, 1)]*m[m.indexOf(1, 3)]*m[m.indexOf(2, 2)])

	return a0 - a1 + a2 - a3
}

func (m Matrix44) DivFloat(f float64) Matrix44 { return Matrix44.MulFloat(m, 1/f) }

func (m Matrix44) Inverse() Matrix44 {
	return Matrix44.DivFloat(m.Adj(), m.Det())
}

func (m Matrix44) Mul(n Matrix44) Matrix44 {
	m00 := m[m.indexOf(0, 0)]*n[n.indexOf(0, 0)] +
		m[m.indexOf(0, 1)]*n[n.indexOf(1, 0)] +
		m[m.indexOf(0, 2)]*n[n.indexOf(2, 0)] +
		m[m.indexOf(0, 3)]*n[n.indexOf(3, 0)]
	m01 := m[m.indexOf(0, 0)]*n[n.indexOf(0, 1)] +
		m[m.indexOf(0, 1)]*n[n.indexOf(1, 1)] +
		m[m.indexOf(0, 2)]*n[n.indexOf(2, 1)] +
		m[m.indexOf(0, 3)]*n[n.indexOf(3, 1)]
	m02 := m[m.indexOf(0, 0)]*n[n.indexOf(0, 2)] +
		m[m.indexOf(0, 1)]*n[n.indexOf(1, 2)] +
		m[m.indexOf(0, 2)]*n[n.indexOf(2, 2)] +
		m[m.indexOf(0, 3)]*n[n.indexOf(3, 2)]
	m03 := m[m.indexOf(0, 0)]*n[n.indexOf(0, 3)] +
		m[m.indexOf(0, 1)]*n[n.indexOf(1, 3)] +
		m[m.indexOf(0, 2)]*n[n.indexOf(2, 3)] +
		m[m.indexOf(0, 3)]*n[n.indexOf(3, 3)]

	m10 := m[m.indexOf(1, 0)]*n[n.indexOf(0, 0)] +
		m[m.indexOf(1, 1)]*n[n.indexOf(1, 0)] +
		m[m.indexOf(1, 2)]*n[n.indexOf(2, 0)] +
		m[m.indexOf(1, 3)]*n[n.indexOf(3, 0)]
	m11 := m[m.indexOf(1, 0)]*n[n.indexOf(0, 1)] +
		m[m.indexOf(1, 1)]*n[n.indexOf(1, 1)] +
		m[m.indexOf(1, 2)]*n[n.indexOf(2, 1)] +
		m[m.indexOf(1, 3)]*n[n.indexOf(3, 1)]
	m12 := m[m.indexOf(1, 0)]*n[n.indexOf(0, 2)] +
		m[m.indexOf(1, 1)]*n[n.indexOf(1, 2)] +
		m[m.indexOf(1, 2)]*n[n.indexOf(2, 2)] +
		m[m.indexOf(1, 3)]*n[n.indexOf(3, 2)]
	m13 := m[m.indexOf(1, 0)]*n[n.indexOf(0, 3)] +
		m[m.indexOf(1, 1)]*n[n.indexOf(1, 3)] +
		m[m.indexOf(1, 2)]*n[n.indexOf(2, 3)] +
		m[m.indexOf(1, 3)]*n[n.indexOf(3, 3)]

	m20 := m[m.indexOf(2, 0)]*n[n.indexOf(0, 0)] +
		m[m.indexOf(2, 1)]*n[n.indexOf(1, 0)] +
		m[m.indexOf(2, 2)]*n[n.indexOf(2, 0)] +
		m[m.indexOf(2, 3)]*n[n.indexOf(3, 0)]
	m21 := m[m.indexOf(2, 0)]*n[n.indexOf(0, 1)] +
		m[m.indexOf(2, 1)]*n[n.indexOf(1, 1)] +
		m[m.indexOf(2, 2)]*n[n.indexOf(2, 1)] +
		m[m.indexOf(2, 3)]*n[n.indexOf(3, 1)]
	m22 := m[m.indexOf(2, 0)]*n[n.indexOf(0, 2)] +
		m[m.indexOf(2, 1)]*n[n.indexOf(1, 2)] +
		m[m.indexOf(2, 2)]*n[n.indexOf(2, 2)] +
		m[m.indexOf(2, 3)]*n[n.indexOf(3, 2)]
	m23 := m[m.indexOf(2, 0)]*n[n.indexOf(0, 3)] +
		m[m.indexOf(2, 1)]*n[n.indexOf(1, 3)] +
		m[m.indexOf(2, 2)]*n[n.indexOf(2, 3)] +
		m[m.indexOf(2, 3)]*n[n.indexOf(3, 3)]

	m30 := m[m.indexOf(3, 0)]*n[n.indexOf(0, 0)] +
		m[m.indexOf(3, 1)]*n[n.indexOf(1, 0)] +
		m[m.indexOf(3, 2)]*n[n.indexOf(2, 0)] +
		m[m.indexOf(3, 3)]*n[n.indexOf(3, 0)]
	m31 := m[m.indexOf(3, 0)]*n[n.indexOf(0, 1)] +
		m[m.indexOf(3, 1)]*n[n.indexOf(1, 1)] +
		m[m.indexOf(3, 2)]*n[n.indexOf(2, 1)] +
		m[m.indexOf(3, 3)]*n[n.indexOf(3, 1)]
	m32 := m[m.indexOf(3, 0)]*n[n.indexOf(0, 2)] +
		m[m.indexOf(3, 1)]*n[n.indexOf(1, 2)] +
		m[m.indexOf(3, 2)]*n[n.indexOf(2, 2)] +
		m[m.indexOf(3, 3)]*n[n.indexOf(3, 2)]
	m33 := m[m.indexOf(3, 0)]*n[n.indexOf(0, 3)] +
		m[m.indexOf(3, 1)]*n[n.indexOf(1, 3)] +
		m[m.indexOf(3, 2)]*n[n.indexOf(2, 3)] +
		m[m.indexOf(3, 3)]*n[n.indexOf(3, 3)]

	return Matrix44{
		m00, m01, m02, m03,
		m10, m11, m12, m13,
		m20, m21, m22, m23,
		m30, m31, m32, m33,
	}
}

func (m Matrix44) MulFloat(f float64) Matrix44 {
	n := Matrix44{}
	for i := 0; i < len(m); i++ {
		n[i] = m[i] * f
	}
	return n
}

func (m Matrix44) Transpose() Matrix44 {
	return Matrix44{
		m[m.indexOf(0, 0)], m[m.indexOf(1, 0)], m[m.indexOf(2, 0)], m[m.indexOf(3, 0)],
		m[m.indexOf(0, 1)], m[m.indexOf(1, 1)], m[m.indexOf(2, 1)], m[m.indexOf(3, 1)],
		m[m.indexOf(0, 2)], m[m.indexOf(1, 2)], m[m.indexOf(2, 2)], m[m.indexOf(3, 2)],
		m[m.indexOf(0, 3)], m[m.indexOf(1, 3)], m[m.indexOf(2, 3)], m[m.indexOf(3, 3)],
	}
}

func (m Matrix44) indexOf(i, j int) int {
	return i*4 + j
}

func Identity() Matrix44 { return Matrix44{0: 1, 5: 1, 10: 1, 15: 1} }

type Transform struct{ M, Inv Matrix44 }

func (t Transform) Inverse() Transform {
	return Transform{
		M:   t.Inv,
		Inv: t.M,
	}
}

func (t Transform) Mul(u Transform) Transform {
	return Transform{
		M:   Matrix44.Mul(t.M, u.M),
		Inv: Matrix44.Mul(u.Inv, t.Inv),
	}
}

func (t Transform) Transpose() Transform {
	return Transform{
		M:   t.M.Transpose(),
		Inv: t.Inv.Transpose(),
	}
}

func IdentityTransform() Transform {
	return Transform{
		M:   Identity(),
		Inv: Identity(),
	}
}

func ScaleTransform(x, y, z float64) Transform {
	return Transform{
		M: Matrix44{
			x, 0, 0, 0,
			0, y, 0, 0,
			0, 0, z, 0,
			0, 0, 0, 1,
		},
		Inv: Matrix44{
			1 / x, 0, 0, 0,
			0, 1 / y, 0, 0,
			0, 0, 1 / z, 0,
			0, 0, 0, 1,
		},
	}
}

func TranslateTransform(x, y, z float64) Transform {
	return Transform{
		M: Matrix44{
			1, 0, 0, x,
			0, 1, 0, y,
			0, 0, 1, z,
			0, 0, 0, 1,
		},
		Inv: Matrix44{
			1, 0, 0, -x,
			0, 1, 0, -y,
			0, 0, 1, -z,
			0, 0, 0, 1,
		},
	}
}

func (i Interaction) Transform(t Transform) Interaction {
	return Interaction{
		P:      i.P.Transform(t),
		N:      i.N.Transform(t).Normalize(),
		Wo:     i.Wo.Transform(t).Normalize(),
		T:      i.T,
		Object: i.Object,
	}
}

func (n Normal3) Transform(t Transform) Normal3 {
	inv := t.Inv
	return Normal3{
		X: inv[inv.indexOf(0, 0)]*n.X + inv[inv.indexOf(1, 0)]*n.Y + inv[inv.indexOf(2, 0)]*n.Z,
		Y: inv[inv.indexOf(0, 1)]*n.X + inv[inv.indexOf(1, 1)]*n.Y + inv[inv.indexOf(2, 1)]*n.Z,
		Z: inv[inv.indexOf(0, 2)]*n.X + inv[inv.indexOf(1, 2)]*n.Y + inv[inv.indexOf(2, 2)]*n.Z,
	}
}

func (p Point3) Transform(t Transform) Point3 {
	m := t.M
	q := Point3{
		X: m[m.indexOf(0, 0)]*p.X + m[m.indexOf(0, 1)]*p.Y + m[m.indexOf(0, 2)]*p.Z + m[m.indexOf(0, 3)],
		Y: m[m.indexOf(1, 0)]*p.X + m[m.indexOf(1, 1)]*p.Y + m[m.indexOf(1, 2)]*p.Z + m[m.indexOf(1, 3)],
		Z: m[m.indexOf(2, 0)]*p.X + m[m.indexOf(2, 1)]*p.Y + m[m.indexOf(2, 2)]*p.Z + m[m.indexOf(2, 3)],
	}
	w := m[m.indexOf(3, 0)]*p.X + m[m.indexOf(3, 1)]*p.Y + m[m.indexOf(3, 2)]*p.Z + m[m.indexOf(3, 3)]
	if !EqualFloat(w, 1) {
		q.DivAssignFloat(w)
	}
	return q
}

func (r Ray) Transform(t Transform) Ray {
	return Ray{
		O:    r.O.Transform(t),
		D:    r.D.Transform(t).Normalize(),
		TMax: r.TMax,
	}
}

func (v Vector3) Transform(t Transform) Vector3 {
	m := t.M
	return Vector3{
		X: m[m.indexOf(0, 0)]*v.X + m[m.indexOf(0, 1)]*v.Y + m[m.indexOf(0, 2)]*v.Z,
		Y: m[m.indexOf(1, 0)]*v.X + m[m.indexOf(1, 1)]*v.Y + m[m.indexOf(1, 2)]*v.Z,
		Z: m[m.indexOf(2, 0)]*v.X + m[m.indexOf(2, 1)]*v.Y + m[m.indexOf(2, 2)]*v.Z,
	}
}
