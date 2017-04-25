package archeage

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

type Character struct {
	Name       string
	UUID       string
	Server     string
	Level      int
	Nation     string
	Expedition *Expedition
	Stat       *Stat
	// Equipments   *Equipments
	Class *Class
	// Actabilities []*Actabilitie
}

type Stat struct {
	Health   int // 생명력
	Vitality int // 활력

	Strength     int // 힘
	Spirit       int // 정신
	Intelligence int // 지능
	Stamina      int // 체력
	Agility      int // 민첩

	Speed       int     // 이동 속도(%)
	CastTime    float64 // 시전 시간(%)
	AttackSpeed float64 // 공격 속도(%)

	// 기본
	MeleeDPS     float64 // 근접 공격력
	RangeDPS     float64 // 원거리 공격력
	MagicDPS     float64 // 주문력
	HealingPower float64 // 치유력
	Defense      int     // 물리 방어도
	MagicDefense int     // 마법 저항도
	GearScore    int     // 장비점수

	// 공격

	// 방어

	// 회복
}

// type Equipments struct {
// 	Costume    *Equipment
// 	InnerArmor *Equipment

// 	Head   *Equipment
// 	Chest  *Equipment
// 	Waist  *Equipment
// 	Wrists *Equipment
// 	Hands  *Equipment
// 	Legs   *Equipment
// 	Feet   *Equipment

// 	Cloak *Equipment

// 	Hand1Weapon *Equipment
// 	Hand2Weapon *Equipment
// 	Shield      *Equipment
// 	Bow         *Equipment
// 	Instruments *Equipment

// 	Necklace *Equipment
// 	Earring1 *Equipment
// 	Earring2 *Equipment
// 	Ring1    *Equipment
// 	Ring2    *Equipment
// }

// type Equipment struct {
// 	Name  string
// 	Grade string
// 	Score int
// }

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

// type Actabilitie struct {
// 	Name       string
// 	Level      int
// 	Score      int
// 	Percentage int
// }

const (
	charactorURLFormat = "https://archeage.xlgames.com/characters/%s"
)

func (a *ArcheAge) SearchCharactor(server, name string) (c *Character, err error) {
	return
}

func (a *ArcheAge) fetchCharactorByUUID(uuid string) (c *Character, err error) {
	url := fmt.Sprintf(charactorURLFormat, uuid)
	return a.fetchCharactorByURL(url)
}

func (a *ArcheAge) fetchCharactorByURL(url string) (c *Character, err error) {
	doc, err := a.get(url)
	if err != nil {
		return
	}
	return a.parseCharactor(doc, url)
}

func (a *ArcheAge) parseCharactor(doc *goquery.Document, url string) (c *Character, err error) {
	return
}
