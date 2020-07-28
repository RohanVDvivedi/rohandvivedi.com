import React from "react";
import ReactDOM from "react-dom";

import NavBar from "./nav/NavBar";
import ContentContainer from "./content/ContentContainer";

class Root extends React.Component {
	constructor(props) {
        super(props);
        this.state = {
			selected: "about",
		};
    }
	ifNavButtonClicked(new_selection){
		this.setState({
			selected: new_selection,
		});
	}
    render() {
    	console.log("Don't you poke around in my source!!");
        return (
            <div>
                <NavBar selected={this.state.selected} ifNavButtonClicked={this.ifNavButtonClicked.bind(this)}/>
                <ContentContainer selected={this.state.selected}/>
            </div>
        );
    }
}

// ================================= >>>>

ReactDOM.render(<Root />, document.getElementById("root"));