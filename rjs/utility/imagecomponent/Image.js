import React from "react";

export default class Image extends React.Component {
    constructor(props) {
        super(props)
    }
    render() {
        var imageStyle = {
            verticalAlign: "middle",
            display: "inherit",
        };
        if(this.props.hasOwnProperty('borderRadius')) {
            imageStyle.borderRadius = this.props.borderRadius;
        }
        if(this.props.hasOwnProperty('height')) {
            imageStyle.height = this.props.height;
        }
        if(this.props.hasOwnProperty('width')) {
            imageStyle.width = this.props.width;
        }
        return (
            <img className={"no-padding-and-no-margin"} src={this.props.imgUrl} style={imageStyle}></img>
        );
    }
}