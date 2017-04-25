package archeage

type Character struct {
	Name              string
	UUID              string
	Server            string
	Level             int
	Nation            string
	Expedition        *Expedition
	ExpeditionHistory []*Expedition
	Stat              *Stat
	Equipments        *Equipments
	Class             *Class
	Actabilities      []*Actabilitie
	EastLanguage      int
	WestLanguage      int
}

type Stat struct {
	Health   int // 생명력
	Vitality int // 활력

	Strength     int // 힘
	Spirit       int // 정신
	Intelligence int // 지능
	Stamina      int // 체력
	Agility      int // 민첩

	AttackSpeed float64 // 공격 속도(%)

	// 기본
	Defense      int // 물리 방어도
	MagicDefense int // 마법 저항도
}

type Equipments struct {
	Costume    *Equipment
	InnerArmor *Equipment

	Head   *Equipment
	Chest  *Equipment
	Waist  *Equipment
	Wrists *Equipment
	Hands  *Equipment
	Legs   *Equipment
	Feet   *Equipment

	Cloak *Equipment

	Hand1Weapon *Equipment
	Hand2Weapon *Equipment
	Shield      *Equipment
	Bow         *Equipment
	Instruments *Equipment

	Necklace *Equipment
	Earring1 *Equipment
	Earring2 *Equipment
	Ring1    *Equipment
	Ring2    *Equipment
}

type Equipment struct {
	Name  string
	Grade string
	Score int
}

type Class struct {
	Name   string
	Skills []*Skills
}

type Skills struct {
	Level       int
	IsHatred    bool
	HatredLevel int
	SkillList   []string
}

type Actabilitie struct {
	Name       string
	Level      int
	Score      int
	Percentage int
}

func (a *ArcheAge) SearchCharactor(server, name string) (c Character, err error) {
	return
}
