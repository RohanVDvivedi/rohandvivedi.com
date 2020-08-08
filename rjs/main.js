import React from "react";
import ReactDOM from "react-dom";
import {BrowserRouter, Route, Switch, Redirect} from 'react-router-dom';

import NavBar from "./nav/NavBar";
import ContentHash from "./ContentHash";

class Root extends React.Component {
    render() {
    	const defaultContent = "about";
    	console.log("Don't you poke around in my source!!");
        return (
            <BrowserRouter>
                <NavBar/>
	            <Switch>
	            	<Redirect exact from="/" to="/pages/about" />
	                {/*<Route path="/" exact component={ContentHash[defaultContent]["component"]} />*/}
	                {Object.keys(ContentHash).map((buttonName) => {     
		           		return (<Route path={ContentHash[buttonName]["route_path"]} component={ContentHash[buttonName]["component"]}/>)
		        	})}
	            </Switch>
            </BrowserRouter>
        );
    }
}

// ================================= >>>>

ReactDOM.render(<Root />, document.getElementById("root"));