package pokemon

type Spell struct {
	name     string
	pp       int
	maxPP    int
	ppPerUse int
}

func NewSpell(
	name string,
	pp int,
	maxPP int,
	ppPerUse int,
) *Spell {
	return &Spell{
		name:     name,
		pp:       pp,
		maxPP:    maxPP,
		ppPerUse: ppPerUse,
	}
}

func (s *Spell) PP() int {
	return s.pp
}

func (s *Spell) MaxPP() int {
	return s.maxPP
}

func (s *Spell) HasEnoughPP() bool {
	return s.pp > 0
}

func (s *Spell) DecreasePP() {
	if s.pp-1*s.ppPerUse < 0 {
		s.pp = 0
		return
	}

	s.pp = s.pp - 1*s.ppPerUse
}

func (s *Spell) IncreasePP() {
	if s.pp+1*s.ppPerUse > s.maxPP {
		s.pp = s.maxPP
		return
	}

	s.pp = s.pp + 1*s.ppPerUse
}

func (s *Spell) SetPP(quantity int) {
	if quantity > s.maxPP {
		s.pp = s.maxPP
		return
	}

	s.pp = quantity
}

func (s *Spell) Name() string {
	return s.name
}
