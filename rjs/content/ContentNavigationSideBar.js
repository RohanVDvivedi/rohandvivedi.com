import React from "react";

export default class ContentNavigationSideBar extends React.Component {
    render() {
        var container_style = {
            paddingTop: "10px",
        };
        return (
            <div id="container" class="container_side_style" style={container_style}>
                <a id="blogs" class="component_side_style">Blogs</a>
                <a id="projects" class="component_side_style">Projects</a>
            </div>
        );
    }
}