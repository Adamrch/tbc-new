package shadow

import (
	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/priest"
)

func RegisterShadowPriest() {
	core.RegisterAgentFactory(
		proto.Player_ShadowPriest{},
		proto.Spec_SpecShadowPriest,
		func(character *core.Character, options *proto.Player) core.Agent {
			return NewShadowPriest(character, options)
		},
		func(player *proto.Player, spec interface{}) {
			playerSpec, ok := spec.(*proto.Player_ShadowPriest)
			if !ok {
				panic("Invalid spec value for Shadow Priest!")
			}
			player.Spec = playerSpec
		},
	)
}

const MaxShadowOrbs = 3

func NewShadowPriest(character *core.Character, options *proto.Player) *ShadowPriest {
	shadowOptions := options.GetShadowPriest()

	selfBuffs := priest.SelfBuffs{
		UseShadowfiend: true,
		UseInnerFire:   shadowOptions.Options.ClassOptions.Armor == proto.PriestOptions_InnerFire,
	}

	basePriest := priest.New(character, selfBuffs, options.TalentsString)
	basePriest.Latency = float64(basePriest.ChannelClipDelay.Milliseconds())
	spriest := &ShadowPriest{
		Priest:  basePriest,
		options: shadowOptions.Options,
	}

	return spriest
}

type ShadowPriest struct {
	*priest.Priest
	options *proto.ShadowPriest_Options

	// Shadow Spells
	DevouringPlague *core.Spell
	MindBlast       *core.Spell
}

func (spriest *ShadowPriest) GetPriest() *priest.Priest {
	return spriest.Priest
}

func (spriest *ShadowPriest) Initialize() {
	spriest.Priest.Initialize()

	// spriest.AddStat(stats.HitRating, -spriest.GetBaseStats()[stats.Spirit])
	// spriest.AddStatDependency(stats.Spirit, stats.HitRating, 1)
	// spriest.registerMindBlastSpell()
	// spriest.registerDevouringPlagueSpell()
	// spriest.registerMindSpike()
	// spriest.registerShadowWordDeathSpell()
	// spriest.registerMindFlaySpell()
	// spriest.registerShadowyRecall() // Mastery
	// spriest.registerShadowyApparition()
}

func (spriest *ShadowPriest) Reset(sim *core.Simulation) {
	spriest.Priest.Reset(sim)
}

func (spriest *ShadowPriest) ApplyTalents() {
	spriest.Priest.ApplyTalents()

	core.MakePermanent(spriest.RegisterAura(core.Aura{
		Label: "Shadowform",
		ActionID: core.ActionID{
			SpellID: 15473,
		},
	}))

}
