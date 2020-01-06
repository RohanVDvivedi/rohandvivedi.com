import React from "react";
import ReactDOM from "react-dom";

import ApiCalledComponent from "./content/ApiCalledContent.js";

class Root extends React.Component {
    render() {
        return (
            <div>
                Hello World!!
                <ApiCalledComponent />
            </div>
        );
    }
}

// ================================= >>>>

ReactDOM.render(<Root />, document.getElementById("root"));