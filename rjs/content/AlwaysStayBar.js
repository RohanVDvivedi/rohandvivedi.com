import React from "react";

export default class AlwaysStayBar extends React.Component {
    render() {
        return (
            <div id="container" class="container_style">
                <a id="home" class="component_style float-left">Home Button</a>
                <a id="search_box" class="component_style float-left">Search Box</a>
                <a id="sign_in" class="component_style float-left">Sign in</a>
                <a id="sign_up" class="component_style float-left">Sign up</a>
                <div class="float-clear"></div>
            </div>
        );
    }
}