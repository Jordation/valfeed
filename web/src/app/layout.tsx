import "./globals.css";
import type { Metadata } from "next";

import "@fontsource/roboto/300.css";
import "@fontsource/roboto/400.css";
import "@fontsource/roboto/500.css";
import "@fontsource/roboto/700.css";

export const metadata: Metadata = {
	title: "Create Next App",
	description: "Generated by create next app",
};

export default function RootLayout({
	children,
}: {
	children: React.ReactNode;
}) {
	return (
		<html lang="en">
			<body className="bg-slate-800 h-screen w-screen grid grid-cols-3 grid-rows-3">
				<div className="row-start-1 col-start-1 col-span-3">header</div>
				<div className="row-start-1 col-start-1 row-span-3">left</div>
				<div className="row-start-1 col-start-3 row-span-3">right</div>
				<div className="row-start-2 col-start-2 p-5">
					{children}
				</div>
				<div className="row-start-3 col-start-1 col-span-3">bottom</div>
			</body>
		</html>
	);
}
