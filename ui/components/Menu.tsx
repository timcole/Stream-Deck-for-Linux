import * as React from 'react';

import "../styles/Menu";
class Menu extends React.Component<any> {
	render() {
		return (
			<div className="menu">
				<div className="section">
					<h3>System</h3>
					<ul>
						<li><i class="material-icons">language</i> Website</li>
						<li><i class="material-icons">keyboard</i> Hotkey</li>
						<li><i class="material-icons">launch</i> Open</li>
					</ul>
				</div>
			</div>
		);
	}
}

export default Menu;