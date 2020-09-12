import React from "react";

// the parent of this component must have class loading-able-container
export default class Loading extends React.Component {
	render() {
		return (<div class={"loading-div " + ((this.props.loading)?"":"hidden")}>
					<div className="no-padding-and-no-margin loading-icon"></div>
				</div>);
	}
}