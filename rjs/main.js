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

import "../css_raw/about.css"
import "../css_raw/contact_form.css"
import "../css_raw/content.css"
import "../css_raw/drop_down.css"
import "../css_raw/loading.css"
import "../css_raw/modal.css"
import "../css_raw/nav.css"
import "../css_raw/past.css"
import "../css_raw/projects.css"
import "../css_raw/search.css"
import "../css_raw/tool_tip.css"
import "../css_raw/utility.css"

// initialize the cache that you will be using to cache apis
import EffiCache from "./utility/EffiCache"
EffiCache.Init()

ReactDOM.render(<Root />, document.getElementById("root"));

window.ChatterStart = function(name, publicKey) {
    var queryParams = []
    if(name != null) {queryParams[0] = "name=" + name}
    if(publicKey != null) {queryParams[1] = "publicKey=" + publicKey }
    var baseUrl = (window.location.protocol.includes("https") ? "wss://" : "ws://") + window.location.host
    var path = "/chat"
    var URL = [baseUrl + path, queryParams.join("&")].join("?")
    window.ChatterSocket = new WebSocket(URL);
    window.ChatterSocket.onmessage = function(event){
        var payload = JSON.parse(event.data)
        console.log(payload.From + " -> " + payload.To + " : " + payload.Message)
    }
}

window.ChatterSend = function(frm ,to, message) {
    var payload = {From:frm,To:to,Message:message,SentAt:new Date()}
    window.ChatterSocket.send(JSON.stringify(payload))
}