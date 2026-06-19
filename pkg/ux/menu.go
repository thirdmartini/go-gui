package ux

type OnMenuSelection func(m *MenuItem) bool
type OnMenuOption func(m *MenuItem) bool

type MenuItem struct {
	ID          string
	Name        string
	Description string

	Items []*MenuItem

	idx    int
	parent *MenuItem

	// Callbacks
	OnMenuSelection OnMenuSelection
	OnMenuOption    OnMenuOption
}

func (m *MenuItem) Next() {
	m.idx++
	if m.idx >= len(m.Items) {
		m.idx = 0
	}
}

func (m *MenuItem) Prev() {
	m.idx--
	if m.idx < 0 {
		m.idx = len(m.Items) - 1
	}
}

func (m *MenuItem) AtOffset(ofs int) *MenuItem {
	if len(m.Items) == 0 {
		return m
	}

	cidx := (m.idx + ofs) % len(m.Items)
	if cidx >= 0 {
		return m.Items[cidx]
	} else {
		cidx = len(m.Items) + cidx
	}

	return m.Items[cidx]
}

func (m *MenuItem) Get() *MenuItem {
	if len(m.Items) == 0 {
		return m
	}

	m.Items[m.idx].parent = m
	return m.Items[m.idx]
}

func (m *MenuItem) Parent() *MenuItem {
	return m.parent
}

func (m *MenuItem) HasChildren() bool {
	return len(m.Items) > 0
}
