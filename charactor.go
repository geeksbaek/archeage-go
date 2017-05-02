package archeage

import (
	"fmt"
	"regexp"
	"strings"

	"sync"

	"github.com/PuerkitoBio/goquery"
)

const (
	searchCharactorFormat = `https://archeage.xlgames.com/search?dt=characters&keyword=%s&server=%s`
	charactorURLFormat    = "https://archeage.xlgames.com/characters/%s"
)

// query
const (
	searchCharactorQuery       = `#container-common > div > div > div.view > div > ul > li`
	searchCharactorURLQuery    = `.character_card > a`
	searchCharactorNameQuery   = `.character_name`
	searchCharactorServerQuery = `.character_server`

	characterNameQuery   = `#container > div > section > div.head > div.character_card > a.character_name > strong`
	characterUUIDQuery   = `#container > div > section > div.head > div.character_card > a.character_name`
	characterServerQuery = `#container > div > section > div.head > div.character_card > span.character_server`
	characterLevelQuery  = `#container > div > section > div.head > div.character_card > span.character_sub > span.level > span:nth-child(2)`
	characterRaceQuery   = `#container > div > section > div.head > div.character_card > span.character_sub > span.race > span:nth-child(2)`

	expeditionNameQuery = `#container > div > section > div.head > div.character_card > span.character_exped > a > span`

	statHealthQuery              = `#container > div > section > div.body > div.info_stat > div.left > dl:nth-child(1) > dd`
	statVitalityQuery            = `#container > div > section > div.body > div.info_stat > div.left > dl:nth-child(2) > dd`
	statStrengthQuery            = `#container > div > section > div.body > div.info_stat > div.left > div > dl:nth-child(1) > dd`
	statSpiritQuery              = `#container > div > section > div.body > div.info_stat > div.left > div > dl:nth-child(2) > dd`
	statIntelligenceQuery        = `#container > div > section > div.body > div.info_stat > div.left > div > dl:nth-child(3) > dd`
	statStaminaQuery             = `#container > div > section > div.body > div.info_stat > div.left > div > dl:nth-child(4) > dd`
	statAgilityQuery             = `#container > div > section > div.body > div.info_stat > div.left > div > dl:nth-child(5) > dd`
	statSpeedQuery               = `#container > div > section > div.body > div.info_stat > div.left > dl:nth-child(4) > dd`
	statCastTimeQuery            = `#container > div > section > div.body > div.info_stat > div.left > dl:nth-child(5) > dd`
	statAttackSpeedQuery         = `#container > div > section > div.body > div.info_stat > div.left > dl:nth-child(6) > dd`
	statMeleeDPSQuery            = `#container > div > section > div.body > div.info_stat > div.right > div.wrap.characterStatTabContent.characterStatTabContent_1 > dl:nth-child(1) > dd`
	statRangeDPSQuery            = `#container > div > section > div.body > div.info_stat > div.right > div.wrap.characterStatTabContent.characterStatTabContent_1 > dl:nth-child(2) > dd`
	statMagicDPSQuery            = `#container > div > section > div.body > div.info_stat > div.right > div.wrap.characterStatTabContent.characterStatTabContent_1 > dl:nth-child(3) > dd`
	statHealingPowerQuery        = `#container > div > section > div.body > div.info_stat > div.right > div.wrap.characterStatTabContent.characterStatTabContent_1 > dl:nth-child(4) > dd`
	statDefenseQuery             = `#container > div > section > div.body > div.info_stat > div.right > div.wrap.characterStatTabContent.characterStatTabContent_1 > dl:nth-child(5) > dd`
	statMagicDefenseQuery        = `#container > div > section > div.body > div.info_stat > div.right > div.wrap.characterStatTabContent.characterStatTabContent_1 > dl:nth-child(6) > dd`
	statGearScoreQuery           = `#container > div > section > div.body > div.info_stat > div.right > div.wrap.characterStatTabContent.characterStatTabContent_1 > dl:nth-child(7) > dd`
	statMeleeCriticalRateQuery   = `#container > div > section > div.body > div.info_stat > div.right > div.wrap.characterStatTabContent.characterStatTabContent_2 > dl:nth-child(3) > dd`
	statRangeCriticalRateQuery   = `#container > div > section > div.body > div.info_stat > div.right > div.wrap.characterStatTabContent.characterStatTabContent_2 > dl:nth-child(6) > dd`
	statAccuracyQuery            = `#container > div > section > div.body > div.info_stat > div.right > div.wrap.characterStatTabContent.characterStatTabContent_2 > dl:nth-child(7) > dd`
	statMagicCriticalRateQuery   = `#container > div > section > div.body > div.info_stat > div.right > div.wrap.characterStatTabContent.characterStatTabContent_2 > dl:nth-child(9) > dd`
	statHealingCriticalRateQuery = `#container > div > section > div.body > div.info_stat > div.right > div.wrap.characterStatTabContent.characterStatTabContent_2 > dl:nth-child(10) > dd`
	statResilienceQuery          = `#container > div > section > div.body > div.info_stat > div.right > div.wrap.characterStatTabContent.characterStatTabContent_3 > dl:nth-child(4) > dd`
	statToughnessQuery           = `#container > div > section > div.body > div.info_stat > div.right > div.wrap.characterStatTabContent.characterStatTabContent_3 > dl:nth-child(5) > dd`

	classNameQuery = `#container > div > section > div.head > div.character_card > div > a > strong`
)

var (
	overSpaceRe = regexp.MustCompile(`\s{2,}`)
)

type Character struct {
	Name       string
	UUID       string
	Server     string
	Level      string
	Race       string
	Expedition *Expedition
	Stat       *Stat
	// Equipments   *Equipments
	Class *Class
	// Actabilities []*Actabilitie
}

func (c *Character) String() string {
	return strings.Join(
		[]string{
			fmt.Sprintf(charactorURLFormat, c.UUID),
			fmt.Sprintf("[%v %v] 세력: %v / 원정대: %v / 직업: %v / 장비점수: %v",
				c.Name, c.Server, c.Race, c.Expedition.Name, c.Class.Name, c.Stat.GearScore),
			fmt.Sprintf("생명력: %v / 활력: %v / 힘: %v / 정신: %v / 지능: %v / 체력: %v / 민첩: %v",
				c.Stat.Health, c.Stat.Vitality, c.Stat.Strength, c.Stat.Spirit, c.Stat.Intelligence, c.Stat.Stamina, c.Stat.Agility),
			fmt.Sprintf("근접 공격력 : %v / 원거리 공격력: %v / 주문력: %v / 치유력: %v",
				c.Stat.MeleeDPS, c.Stat.RangeDPS, c.Stat.MagicDPS, c.Stat.HealingPower),
			fmt.Sprintf("물리 방어도: %v / 마법 저항도: %v / 적중도: %v / 유연성: %v / 전투 숙련: %v",
				c.Stat.Defense, c.Stat.MagicDefense, c.Stat.Accuracy, c.Stat.Resilience, c.Stat.Toughness),
		},
		"\n",
	)
}

type Characters []*Character

type Stat struct {
	Health   string // 생명력
	Vitality string // 활력

	Strength     string // 힘
	Spirit       string // 정신
	Intelligence string // 지능
	Stamina      string // 체력
	Agility      string // 민첩

	Speed       string // 이동 속도(%)
	CastTime    string // 시전 시간(%)
	AttackSpeed string // 공격 속도(%)

	// 기본
	MeleeDPS     string // 근접 공격력
	RangeDPS     string // 원거리 공격력
	MagicDPS     string // 주문력
	HealingPower string // 치유력
	Defense      string // 물리 방어도
	MagicDefense string // 마법 저항도
	GearScore    string // 장비점수

	// 공격
	// 근접 공격속도
	// 근접 성공률
	MeleeCriticalRate string // 근접 치명타율
	// 원거리 공격속도
	// 원거리 성공률
	RangeCriticalRate string // 원거리 치명타율
	Accuracy          string // 적중도
	// 마법 성공률
	MagicCriticalRate   string // 마법 치명타율
	HealingCriticalRate string // 치유 극대화율

	// 방어
	// 무기 막기 확률
	// 방패막기 확률
	// 회피율
	Resilience string // 유연성
	Toughness  string // 전투 숙련

	// 회복
	// 생명력 회복
	// 지속적인 생명력 회복
	// 활력 회복
	// 기술 사용 후 활력 회복
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
	Name string
	// Skills []*Skills
}

// type Skills struct {
// 	Level       string
// 	IsHatred    bool
// 	HatredLevel string
// 	SkillList   []string
// }

// type Actabilitie struct {
// 	Name       string
// 	Level      int
// 	Score      int
// 	Percentage int
// }

func (a *ArcheAge) SearchCharactor(server, name string) (cs Characters, err error) {
	doc, err := a.get(fmt.Sprintf(searchCharactorFormat, name, server))
	if err != nil {
		return
	}

	cs = make(Characters, doc.Find(searchCharactorQuery).Length())
	wg := sync.WaitGroup{}
	doc.Find(searchCharactorQuery).Each(func(i int, s *goquery.Selection) {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			parsedName := strings.TrimSpace(s.Find(searchCharactorNameQuery).Text())
			if strings.Contains(parsedName, name) {
				url, ok := s.Find(searchCharactorURLQuery).Attr("href")
				if !ok {
					return
				}
				c, err := a.fetchCharactorByURL(url)
				if err != nil {
					return
				}
				cs[i] = c
			}
		}(i)
	})
	wg.Wait()
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
	uuid, _ := doc.Find(characterUUIDQuery).Attr("href")
	c = &Character{
		Name:   doc.Find(characterNameQuery).Text(),
		UUID:   strings.TrimLeft(uuid, "/characters/"),
		Server: doc.Find(characterServerQuery).Text(),
		Level:  doc.Find(characterLevelQuery).Text(),
		Race:   doc.Find(characterRaceQuery).Text(),
		Expedition: &Expedition{
			Name: doc.Find(expeditionNameQuery).Text(),
		},
		Stat: &Stat{
			Health:              doc.Find(statHealthQuery).Text(),
			Vitality:            doc.Find(statVitalityQuery).Text(),
			Strength:            doc.Find(statStrengthQuery).Text(),
			Spirit:              doc.Find(statSpiritQuery).Text(),
			Intelligence:        doc.Find(statIntelligenceQuery).Text(),
			Stamina:             doc.Find(statStaminaQuery).Text(),
			Agility:             doc.Find(statAgilityQuery).Text(),
			Speed:               doc.Find(statSpeedQuery).Text(),
			CastTime:            doc.Find(statCastTimeQuery).Text(),
			AttackSpeed:         doc.Find(statAttackSpeedQuery).Text(),
			MeleeDPS:            doc.Find(statMeleeDPSQuery).Text(),
			RangeDPS:            doc.Find(statRangeDPSQuery).Text(),
			MagicDPS:            doc.Find(statMagicDPSQuery).Text(),
			HealingPower:        doc.Find(statHealingPowerQuery).Text(),
			Defense:             doc.Find(statDefenseQuery).Text(),
			MagicDefense:        doc.Find(statMagicDefenseQuery).Text(),
			GearScore:           doc.Find(statGearScoreQuery).Text(),
			MeleeCriticalRate:   doc.Find(statMeleeCriticalRateQuery).Text(),
			RangeCriticalRate:   doc.Find(statRangeCriticalRateQuery).Text(),
			Accuracy:            doc.Find(statAccuracyQuery).Text(),
			MagicCriticalRate:   doc.Find(statMagicCriticalRateQuery).Text(),
			HealingCriticalRate: doc.Find(statHealingCriticalRateQuery).Text(),
			Resilience:          doc.Find(statResilienceQuery).Text(),
			Toughness:           doc.Find(statToughnessQuery).Text(),
		},
		Class: &Class{
			Name: doc.Find(classNameQuery).Text(),
		},
	}
	c.Stat.MeleeDPS = overSpaceRe.ReplaceAllString(c.Stat.MeleeDPS, "")
	c.Stat.RangeDPS = overSpaceRe.ReplaceAllString(c.Stat.RangeDPS, "")
	c.Stat.Defense = overSpaceRe.ReplaceAllString(c.Stat.Defense, "")
	c.Stat.MagicDefense = overSpaceRe.ReplaceAllString(c.Stat.MagicDefense, "")
	return
}
