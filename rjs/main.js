import React from "react";
import ReactDOM from "react-dom";

import NavBar from "./nav/NavBar.js";
import ContentContainer from "./content/ContentContainer.js";

class Root extends React.Component {
    render() {
        return (
            <div>
                <NavBar />
                <ContentContainer />
            </div>
        );
    }
}

// ================================= >>>>

ReactDOM.render(<Root />, document.getElementById("root"));