package characters

type Monster struct {
	Name      string
	Skill     int
	Stamina   int
	Inventory []Inventory
}

func NewMonster(name string, stamina int, skill int) *Monster {
	return &Monster{
		Name:    name,
		Stamina: stamina,
		Skill:   skill,
		// life:    stamina,
	}
}

func (m *Monster) Alive() bool {
	return m.Stamina > 0
}

func (m *Monster) TakeHit() {
	m.Stamina = (m.Stamina - 2)
}

// func (h *Monster) ShowStats() {
// 	h.WriteLine("*** Character Info ***")
// 	h.WriteLine(fmt.Sprintf("Name: %v", h.user.Name))
// 	h.WriteLine("Stats:")
// 	h.WriteLine(fmt.Sprintf(" - Skill: %v", h.Skill))
// 	h.WriteLine(fmt.Sprintf(" - Luck: %v", h.Luck))
// 	h.WriteLine(fmt.Sprintf(" - Stamina: %v of %v", h.life, h.Stamina))

// 	h.ShowInventory()
// 	h.ShowInfo()
// }
