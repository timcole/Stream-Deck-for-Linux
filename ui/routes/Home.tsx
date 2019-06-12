import * as React from 'react';

import "../styles/streamdeck.scss";
class Home extends React.Component {
	private keys;

	constructor(props) {
		super(props);
		this.keys = 15;

		document.addEventListener('astilectron-ready', function () {
			// This will listen to messages sent by GO
			// @ts-ignore
			window.astilectron.onMessage(function (message) {
				console.log(message);
			});
		})
	}

	render() {
		return (
			<div className="streamdeck">
				{new Array(this.keys).fill(null).map(d => (
					<div className="button" key={Math.random()}></div>
				))}
			</div>
		);
	}
}

export default Home;