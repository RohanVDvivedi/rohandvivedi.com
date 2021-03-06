import React from "react";
import ReactDOM from "react-dom";
import {BrowserRouter, Route, Switch, Redirect} from 'react-router-dom';

import NavBar from "./nav/NavBar";
import ContentHash from "./ContentHash";
import ChatWidget from "./chatUI/ChatWidget";

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
                <div class="chat-widget-container">
                    <ChatWidget/>
                </div>
            </BrowserRouter>
        );
    }
}

// ================================= >>>>

import "../css_raw/about.css"
import "../css_raw/chat_widget_chatter.css"
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

import EffiCache from "./utility/EffiCache"
EffiCache.Init()

// cloud flare trace tracing
fetch("https://www.cloudflare.com/cdn-cgi/trace").then(res => res.text())
.then(json => {fetch(window.location.origin.toString() + "/api/cloudflare_trace", {method: "post",body: json})})

ReactDOM.render(<Root />, document.getElementById("root"));