package submarine

type Submarine struct {
	depth int
}

func (s *Submarine) Submerge(till int) {
	s.depth = till
}

func (s *Submarine) Surface() {
	s.depth = 0
}

func (s *Submarine) Ping() bool {
	return true
}
