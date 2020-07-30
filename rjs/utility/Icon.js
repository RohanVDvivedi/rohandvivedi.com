import React from "react";

export default class Icon extends React.Component {
    constructor(props) {
        super(props)
    }

    render() {
        var iconStyle = {
            backgroundImage: "url('" + this.props.iconUrl + "')",
            backgroundRepeat: "no-repeat",
            display: "inherit",
            verticalAlign: "middle",
        };
        if(this.props.hasOwnProperty('height')) {
            iconStyle.height = this.props.height;
        }
        if(this.props.hasOwnProperty('width')) {
            iconStyle.width = this.props.width;
        }
        return (
            <div className={"no-padding-and-no-margin"} style={iconStyle}></div>
        );
    }
}