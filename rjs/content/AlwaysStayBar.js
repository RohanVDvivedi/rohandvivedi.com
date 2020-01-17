import React from "react";

export default class AlwaysStayBar extends React.Component {
    render() {
        var container_style = {
            justifyContent: "space-between",
        };
        return (
            <div id="container" class="container_style flex-row-container" style={container_style}>
                <a id="home" class="component_style">Home</a>
                <a id="search_box" class="component_style">
                    Search Box
                </a>
                <div class="flex-row-container" style={{
                    justifyContent: "flex-end",
                }}>
                    <a id="sign_in" class="component_style">Sign in</a>
                    <a id="sign_up" class="component_style">Sign up</a>
                </div>
            </div>
        );
    }
}