package action

import (
	"github.com/hectorgimenez/koolo/internal/action/step"
	"github.com/hectorgimenez/koolo/internal/game"
	"github.com/hectorgimenez/koolo/internal/game/stat"
)

type Character interface {
	Buff() Action
	KillCountess() Action
	KillAndariel() Action
	KillSummoner() Action
	KillMephisto() Action
	KillPindle(skipOnImmunities []stat.Resist) Action
	KillNihlathak() Action
	KillCouncil() Action
	KillMonsterSequence(
		monsterSelector func(data game.Data) (game.UnitID, bool),
		skipOnImmunities []stat.Resist,
		opts ...step.AttackOption,
	) *DynamicAction
}
