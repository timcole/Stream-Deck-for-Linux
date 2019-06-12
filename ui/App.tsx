import { render } from "react-dom";
import * as React from 'react';

import Header from "./components/Header";
import Menu from "./components/Menu";

import Deck from "./components/Deck";
 
import "./styles/colours.scss";
import "./styles/App.scss";
class App extends React.Component {
	private isDevTools: boolean = false;

	constructor(props) {
		super(props);
		
		document.addEventListener('astilectron-ready', function () {
			// @ts-ignore
			window.astilectron.onMessage(function (message) {
				console.log(message);
			});
		});

		document.onkeypress = (ev: KeyboardEvent) => {
			// @ts-ignore
			if (ev.ctrlKey && ev.key === "w") window.astilectron.sendMessage("close", console.log)
			if (ev.ctrlKey && ev.shiftKey && ev.key.toLowerCase() === "i") {
				this.isDevTools = !this.isDevTools;
				// @ts-ignore
				window.astilectron.sendMessage(this.isDevTools ? "openDevTools" : "closeDevTools", console.log)
			}
		}
	}

	render() {
		return (<>
			<div className="router">
				<Header />
				<Deck />
			</div>

			<Menu />
		</>);
	}
}

render(<App />, document.getElementById("root"));