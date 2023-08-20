const CombatEventTypes = {
	kill: 0,
	shot: 1,
} as const;
type CombatEvent = (typeof CombatEventTypes)[keyof typeof CombatEventTypes];

type PlayerCombatEvent = {
	Type: CombatEvent;
	Causer: string;
	Victim: string;
	DmgLoc: string;
	DmgOnHit: number;
	RawDmg: number;
	Wallbang: boolean;
	DetailStr: string;
	SequenceNum: number;
	Weapon: string;
};
