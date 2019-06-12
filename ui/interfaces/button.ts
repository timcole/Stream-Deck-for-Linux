export interface IButton {
	image: string
	action: Action
	data: string
}

export enum Action {
	website,
	hotkey,
	open,
}