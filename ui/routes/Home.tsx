import * as React from 'react';

import "../styles/streamdeck.scss";
class Home extends React.Component {
	private keys;

	constructor(props) {
		super(props);
		this.keys = 15;
	}

	render() {
		return (
			<div className="streamdeck">
				{new Array(this.keys).fill(null).map(d => (
					<div className="button"></div>
				))}
			</div>
		);
	}
}

export default Home;