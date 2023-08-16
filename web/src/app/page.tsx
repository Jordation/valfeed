"use client";
import { AccessibleForwardOutlined, Dangerous, GpsFixed, GpsNotFixed } from "@mui/icons-material";
import { Button, Slider } from "@mui/material";
import { useState } from "react";
import { FilterCombatEvents } from "./utils";
import useWebSocket from "react-use-websocket";
export default function Home() {
	const [value, setValue] = useState<number>(1);
	const [fetchedEvents, setFetchedEvents] = useState<PlayerCombatEvent[]>([]);
  const [activeEvents, setactiveEvents] = useState<PlayerCombatEvent[]>([])
	const handleChange = (event: Event, newValue: number | number[]) => {
		setValue(newValue as number);
    setactiveEvents(fetchedEvents.slice(0, newValue as number));
	};
  const GetEvents=()=>{
    fetch("http://localhost:8080/events", {
      method: "POST",
      body: JSON.stringify({ seq: 6200}),
    })
      .then((res) => res.json())
      .then((data) => {
        console.log(data);
        setFetchedEvents(data);
      });
  }
  
  
     const { sendMessage, lastMessage, readyState } = useWebSocket('ws://localhost:8080');
	return (
		<main>

			<div>

				<AccessibleForwardOutlined />
				<div>
          <Button 
          variant="contained"
          className="bg-green-500"
          onClick={GetEvents}

          >GET SUM</Button>
					<Slider
						defaultValue={1}
						valueLabelDisplay="auto"
						min={0}
						max={fetchedEvents.length}
						value={value}
						onChange={handleChange}
					/>
				</div>
				{CombatEvents(FilterCombatEvents(2, activeEvents))}
        <div>
            <button onClick={() => sendMessage('Hello from client')}>
                Send Message
            </button>
            {lastMessage ? <p>{lastMessage.data}</p> : null}
        </div>
			</div>
		</main>
	);
}


function CombatEvents(events: PlayerCombatEvent[]) {
	return <>{events.map((event, i) => {
    return (
      <div key={i+"-ce"}>
        {event.Causer} {DamageLocationToIcon(event.DmgLoc)} {EventType(event.Type)} {event.Victim} for {event.DmgDone} 
      </div>
    )
  })}</>;
}

function DamageLocationToIcon(loc: string){
  switch(loc){
    case "headshot":
      return <GpsFixed />
    case "aoe damage":
      return <Dangerous />
    default:
      return <GpsNotFixed />
  }
}

function EventType(eventType: CombatEvent) {
  switch (eventType) {
    case 0:
      return "died"
    case 1:
      return "killed"
    case 2:
      return "damaged"
    case 3:
      return "did damage"
  }
}