import { render } from "react-dom";
import { BrowserRouter as Router, Route } from "react-router-dom";
import * as React from 'react';

import Header from "./components/Header";
import Menu from "./components/Menu";

import Home from "./routes/Home";
 
import "./styles/colours.scss";
import "./styles/App.scss";
class App extends React.Component {
	render() {
		return (
			<Router>
				<div className="router">
					<Header />
					<Route exact path="/" component={Home} />
				</div>

				<Menu />
			</Router>
		);
	}
}

render(<App />, document.getElementById("root"));