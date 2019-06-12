import * as React from 'react';

import "../styles/Menu";
class Menu extends React.Component<any> {
	private onDragStart(e): void {
		e.dataTransfer.setData("action", e.target.dataset.actions);
	}

	render() {
		return (
			<div className="menu">
				<div className="section">
					<h3>System</h3>
					<ul>
						<li
							draggable={true}
							onDragStart={this.onDragStart}
							data-action={0}
						><i className="material-icons">language</i> Website</li>
						<li
							draggable={true}
							onDragStart={this.onDragStart}
							data-action={1}
						><i className="material-icons">keyboard</i> Hotkey</li>
						<li
							draggable={true}
							onDragStart={this.onDragStart}
							data-action={2}
						><i className="material-icons">launch</i> Open</li>
					</ul>
				</div>
			</div>
		);
	}
}

export default Menu;