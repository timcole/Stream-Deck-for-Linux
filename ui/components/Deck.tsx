import * as React from 'react';

import { IButton, Action } from "../interfaces/button";

interface State {
	selected?: number
}

import "../styles/streamdeck.scss";
class Home extends React.Component<{}, State> {
	private keys: number;
	private buttons: IButton[];

	constructor(props: {}) {
		super(props);
		this.keys = 15;
		this.buttons = [];

		this.state = {
			selected: -1
		};

		this.initButtons()
	}

	private initButtons(): void {
		for (let i = 0; i < this.keys; i++) {
			let imageSize = 200 + i
			let key: IButton = {
				image: `https://www.placecage.com/${imageSize}/${imageSize}`,
				action: Action.open,
				data: "https://timcole.me/"
			}
			this.buttons.push(key)
		}
	}

	private onDragOver(e): void {
		e.target.style.border = "1px solid var(--sidebar-background)";
		e.target.style.opacity = 0.5;
		e.preventDefault();
	}

	private onDragLeave(e): void {
		e.target.style.border = "";
		e.target.style.opacity = 1;
	}

	private onDrop(e): void {
		e.preventDefault();
		this.onDragLeave(e);

		let key: number = Number(e.target.dataset.key);
		this.buttons[key] = {
			image: `https://www.placecage.com/${220 + key}/${220 + key}`,
			action: 0,
			data: "",
		}

		this.onClick(e)
	}

	private onClick(e): void {
		let key: number = Number(e.target.dataset.key);
		if (this.state.selected === key) {
			this.setState({ selected: -1 });
			return;
		}

		this.setState({ selected: key });
	}

	render() {
		const { selected } = this.state;

		return (<>
			<div className="streamdeck">
				{this.buttons.map((d: IButton, i: number) => (
					<div
						onDragOver={this.onDragOver}
						onDragLeave={this.onDragLeave}
						onDrop={this.onDrop.bind(this)}
						onClick={this.onClick.bind(this)}
						className={`button ${selected == i ? "selected" : ""}`}
						key={i}
						data-key={i}
						style={{ backgroundImage: `url(${d.image})` }}>
					</div>
				))}
			</div>
			{selected != -1 && <div className="options">
				<div className="button" style={{ backgroundImage: `url(${this.buttons[selected].image})`}}></div>
				<div className="right">
					<input type="text" placeholder="data" value={this.buttons[selected].data} />
				</div>
			</div>}
		</>);
	}
}

export default Home;