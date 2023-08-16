export function FilterCombatEvents(type: CombatEvent, events: PlayerCombatEvent[]){
    return events.filter(event => event.Type == type);
}
