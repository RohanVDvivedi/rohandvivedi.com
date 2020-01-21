import React from "react";
import ReactDOM from "react-dom";

import NavBar from "./nav/NavBar.js";
import ContentArea from "./content/ContentArea.js";

class Root extends React.Component {
    render() {
        return (
            <div>
                <NavBar />
                <ContentArea />
            </div>
        );
    }
}

// ================================= >>>>

ReactDOM.render(<Root />, document.getElementById("root"));