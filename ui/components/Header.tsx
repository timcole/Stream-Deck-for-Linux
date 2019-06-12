import * as React from 'react';

// @ts-ignore
import logo from "../assets/logo.png";

import "../styles/Header";
class Header extends React.Component {
	render() {
		return (
			<div className="header">
				<div>
					<img src={logo} alt="Stream Deck for Linux" />
				</div>
				<i className="material-icons">settings</i>
			</div>
		);
	}
}

export default Header;