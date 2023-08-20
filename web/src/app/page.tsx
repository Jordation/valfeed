"use client";
import {
	AccessibleForwardOutlined,
	Dangerous,
	GpsFixed,
	GpsNotFixed,
} from "@mui/icons-material";
import { useState } from "react";
import { FilterCombatEvents } from "./utils";
import useWebSocket from "react-use-websocket";
export default function Home() {
	const [fetchedEvents, setFetchedEvents] = useState<PlayerCombatEvent[]>([]);
	const GetEvents = () => {
		fetch("http://localhost:8080/events", {
			method: "POST",
			body: JSON.stringify({ seq: 6200 }),
		})
			.then((res) => res.json())
			.then((data) => {
				console.log(data);
				setFetchedEvents(data);
			});
	};

	const { sendMessage, lastMessage, readyState } = useWebSocket(
		"ws://localhost:9090/event-ws"
	);

	return (
		<main>
			<div>
				<AccessibleForwardOutlined />
				<div>
					<button onClick={() => sendMessage("Hello from client")}>
						Send Message
					</button>
					{lastMessage ? <p>{lastMessage.data}</p> : null}
				</div>
			</div>
		</main>
	);
}

function CombatEvents(events: PlayerCombatEvent[]) {
	return (
		<>
			{events.map((event, i) => {
				return (
					<div key={i + "-ce"}>
						{event.Causer} {DamageLocationToIcon(event.DmgLoc)}{" "}
						{EventType(event.Type)} {event.Victim} for{" "}
						{event.DmgDone}
					</div>
				);
			})}
		</>
	);
}

function DamageLocationToIcon(loc: string) {
	switch (loc) {
		case "headshot":
			return <GpsFixed />;
		case "aoe damage":
			return <Dangerous />;
		default:
			return <GpsNotFixed />;
	}
}

function EventType(eventType: CombatEvent) {
	switch (eventType) {
		case 0:
			return "died";
		case 1:
			return "killed";
		case 2:
			return "damaged";
		case 3:
			return "did damage";
	}
}
