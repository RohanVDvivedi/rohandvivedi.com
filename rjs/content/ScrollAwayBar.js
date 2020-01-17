import React from "react";

export default class ScrollAwayBar extends React.Component {
    render() {
        return (
            <div id="container" class="container_style flex-row-container">
                <a id="about" class="component_style">About</a>
                <a id="contact" class="component_style">Contact</a>
                <a id="chat" class="component_style">Chat</a>
                <a id="email" class="component_style">Email</a>
            </div>
        );
    }
}