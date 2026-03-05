package priest

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
)

var TalentTreeSizes = [3]int{22, 21, 21}

type Priest struct {
	core.Character
	SelfBuffs
	Talents *proto.PriestTalents

	SurgeOfLight bool

	Latency float64

	ShadowfiendAura *core.Aura
	// ShadowfiendPet  *Shadowfiend
	InnerFocus     *core.Spell
	HolyFire       *core.Spell
	Smite          *core.Spell
	ShadowWordPain *core.Spell
	Shadowfiend    *core.Spell
	VampiricTouch  *core.Spell
}

type TargetDoTInfo struct {
	Swp time.Duration
	VT  time.Duration
}

type SelfBuffs struct {
	UseShadowfiend bool
}

func (priest *Priest) GetCharacter() *core.Character {
	return &priest.Character
}

func (priest *Priest) AddPartyBuffs(_ *proto.PartyBuffs) {
}

func (priest *Priest) Initialize() {
	// priest.registerShadowWordPainSpell()
	// priest.registerShadowfiendSpell()
	// priest.registerVampiricTouchSpell()

	// priest.registerDispersionSpell()

	// priest.registerPowerInfusionSpell()
}

func (priest *Priest) ApplyTalents() {
}

func (priest *Priest) Reset(_ *core.Simulation) {

}

func (priest *Priest) OnEncounterStart(sim *core.Simulation) {
}

func New(char *core.Character, selfBuffs SelfBuffs, talents string) *Priest {
	priest := &Priest{
		Character: *char,
		SelfBuffs: selfBuffs,
		Talents:   &proto.PriestTalents{},
	}

	core.FillTalentsProto(priest.Talents.ProtoReflect(), talents, TalentTreeSizes)
	priest.EnableManaBar()
	// priest.ShadowfiendPet = priest.NewShadowfiend()

	return priest
}

// Agent is a generic way to access underlying priest on any of the agents.
type PriestAgent interface {
	GetPriest() *Priest
}

func NewPriest(character *core.Character, options *proto.Player) *Priest {
	//priestOptions := options.GetPriest()

	selfBuffs := SelfBuffs{
		UseShadowfiend: true,
	}

	basePriest := New(character, selfBuffs, options.TalentsString)
	basePriest.Latency = float64(basePriest.ChannelClipDelay.Milliseconds())
	/*	priest := &Priest{
			Priest:  basePriest,
			options: priestOptions.Options,
		}

		return priest*/
	return basePriest
}

func RegisterPriest() {
	core.RegisterAgentFactory(
		proto.Player_Priest{},
		proto.Spec_SpecPriest,
		func(character *core.Character, options *proto.Player) core.Agent {
			return NewPriest(character, options)
		},
		func(player *proto.Player, spec interface{}) {
			playerSpec, ok := spec.(*proto.Player_Priest)
			if !ok {
				panic("Invalid spec value for Priest!")
			}
			player.Spec = playerSpec
		},
	)
}

const (
	PriestSpellFlagNone  int64 = 0
	PriestSpellArchangel int64 = 1 << iota
	PriestSpellDevouringPlague
	PriestSpellDevouringPlagueDoT
	PriestSpellDevouringPlagueHeal
	PriestSpellHolyFire
	PriestSpellHolyNova
	PriestSpellInnerFocus
	PriestSpellManaBurn
	PriestSpellMindBlast
	PriestSpellMindFlay
	PriestSpellPowerInfusion
	PriestSpellShadowWordDeath
	PriestSpellShadowWordPain
	PriestSpellShadowFiend
	PriestSpellSmite
	PriestSpellVampiricEmbrace
	PriestSpellVampiricTouch
	PriestSpellFade

	PriestSpellLast
	PriestSpellsAll    = PriestSpellLast<<1 - 1
	PriestSpellDoT     = PriestSpellDevouringPlague | PriestSpellHolyFire | PriestSpellMindFlay | PriestSpellShadowWordPain | PriestSpellVampiricTouch
	PriestSpellInstant = PriestSpellDevouringPlague |
		PriestSpellFade |
		PriestSpellHolyNova |
		PriestSpellPowerInfusion |
		PriestSpellShadowWordDeath |
		PriestSpellShadowWordPain |
		PriestSpellVampiricEmbrace
	PriestShadowSpells = PriestSpellDevouringPlague |
		PriestSpellShadowWordDeath |
		PriestSpellShadowWordPain |
		PriestSpellMindFlay |
		PriestSpellMindBlast |
		PriestSpellVampiricTouch
)
