import React from "react";

class AboutParagraph extends React.Component {
    constructor(props) {
        super(props)
    }
    render() {
        return(
            <div style={{
                fontSize: this.props.size + "px",
                fontFamily: "lato,sans-serif",
				fontStyle: "italic",
                padding: "6px",
                fontWeight: 500,
                color: "#323232"
            }}>
                    {this.props.children}
            </div>
        );
    }
}

class ColoredBoldWord extends React.Component {
    constructor(props) {
        super(props)
    }
    render() {
        return(
            <span style={{
                fontSize: "inherit",
                fontFamily: "inherit",
                color: this.props.color,
                fontWeight: "bold",
            }}>
                    {this.props.children}
            </span>
        );
    }
}

export default class AboutContent extends React.Component {
    constructor(props) {
        super(props);
        this.name = "about";
        this.contentTitle = "About Me"
    }
    render() {
        return (
            <div id={this.constructor.name + "-content"} class="content-component flex-row-container"
            style={{
                justifyContent: "center",
                alignItems: "center",
            }}>
                <div class={"no-padding-and-no-margin"} 
                    style={{
                            height: "100%",
                            width: "200px",

                            backgroundImage: 'url(/img/me.png)',
                            backgroundPosition: "center",
                            backgroundRepeat: "no-repeat",
                            backgroundSize: "cover",
                        }}>
                </div>
            	<div class="flex-col-container"
            	style={{
                justifyContent: "center",
            	}}>
                	<AboutParagraph size={26}>Hi, I am Rohan.</AboutParagraph>
                	<AboutParagraph size={26}>I am a Software and Hardware Developer.</AboutParagraph>
                	<AboutParagraph size={24}>Predominantly a <ColoredBoldWord color="var(--color6)">Backend Developer</ColoredBoldWord>,<br/> who also indulges in building crappy <ColoredBoldWord color="var(--color6)">Frontend</ColoredBoldWord>s like this one.</AboutParagraph>
                	<AboutParagraph size={24}>My interests also include <ColoredBoldWord color="var(--color6)"> Embedded Systems, Databases,<br/> Computer Vision, Robotics</ColoredBoldWord> and <ColoredBoldWord color="var(--color6)">FPGAs</ColoredBoldWord>.</AboutParagraph>
                	<AboutParagraph size={22}>Find my Curriculum Vitae <a href={window.location.origin.toString() + "/pdf/rohan_cv.pdf"} target="_blank">here</a>.</AboutParagraph>
                </div>
            </div>
        );
    }
}