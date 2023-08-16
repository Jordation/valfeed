const CombatEventTypes = {
    died: 0,
    killed: 1,
    tookDmg: 2,
    dealtDmg: 3,
} as const;
type CombatEvent = (typeof CombatEventTypes)[keyof typeof CombatEventTypes];

type PlayerCombatEvent = {
    Type: CombatEvent;
    DetailStr: string;
    Causer: string;
    Victim: string;
    DmgLoc: string;
    DmgDone: number;
    SequenceNum: number; 
}
