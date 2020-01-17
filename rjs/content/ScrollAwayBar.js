import React from "react";

export default class ScrollAwayBar extends React.Component {
    render() {
        var container_style = {
            padding: "0 10px",
            justifyContent: "flex-end",
        };
        var component_style = {
            width: "10%",
        };
        return (
            <div id="container" class="container_style flex-row-container" style={container_style}>
                <a id="about" class="component_style" style={component_style}>About</a>
                <a id="contact" class="component_style" style={component_style}>Contact</a>
                <a id="chat" class="component_style" style={component_style}>Chat</a>
                <a id="email" class="component_style" style={component_style}>Email</a>
            </div>
        );
    }
}