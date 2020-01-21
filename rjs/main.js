import React from "react";
import ReactDOM from "react-dom";

import NavBar from "./nav/NavBar";
import ContentContainer from "./content/ContentContainer";

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