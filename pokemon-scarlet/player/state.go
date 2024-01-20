package player

type State interface {
	Name() StateName
}

type StateName string

const (
	IN_FREE_WORLD                 StateName = "IN_FREE_WORLD"
	CHALLENGING_A_POKEMON                   = "CHALLENGING_A_POKEMON"
	CHOOSING_A_POKEMON                      = "CHOOSING_A_POKEMON"
	USING_A_SPELL                           = "USING_A_SPELL"
	RUNNING_AWAY                            = "RUNNING_AWAY"
	RESTORING_SPELL_PP                      = "RESTORING_SPELL_PP"
	USING_SPELL_APPLICABLE_ITEM             = "USING_SPELL_APPLICABLE_ITEM"
	USING_POKEMON_APPLICABLE_ITEM           = "USING_POKEMON_APPLICABLE_ITEM"
	USING_HOLDABLE_ITEM                     = "USING_HOLDABLE_ITEM"
)
