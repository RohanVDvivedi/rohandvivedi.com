import React from "react";
import AbstractContent from "./AbstractContent";

class AboutParagraph extends React.Component {
    constructor(props) {
        super(props)
    }
    render() {
        return(
            <div style={{
                fontSize: this.props.size + "px",
                fontFamily: "'Open Sans',Arial,sans-serif",
                padding: "20px",
            }}>
                    {this.props.children}
            </div>
        );
    }
}

export default class AboutContent extends AbstractContent {
    constructor(props) {
        super(props);
        this.name = "about";
        this.contentTitle = "About Me"
    }
    render() {
        return (
            <div id={this.getContentId()} class="content-component flex-row-container"
            style={{
                justifyContent: "space-evenly",
            }}>
                <div class="flex-col-container"
                    style={{
                        minWidth: "40%",
                        justifyContent: "center",
                    }}
                >
                    <AboutParagraph size={22}>Hi, I am Rohan.</AboutParagraph>
                    <AboutParagraph size={22}>I am a Software and Hardware Developer.</AboutParagraph>
                    <AboutParagraph size={22}>predominantly a Backend Developer.</AboutParagraph>
                    <AboutParagraph size={22}>But, I also build some crappy Frontends like this one.</AboutParagraph>
                    <AboutParagraph size={22}>Tinkering with Embedded System and FPGAs is my hobby.</AboutParagraph>
                    <AboutParagraph size={22}>It is pleasure meeting you.</AboutParagraph>
                </div>
                <div class="flex-col-container"
                    style={{
                        width: "30%",
                        justifyContent: "center",
                    }}
                >
                    <AboutParagraph size={22}></AboutParagraph>
                    <AboutParagraph size={22}></AboutParagraph>
                    <AboutParagraph size={22}></AboutParagraph>
                    <AboutParagraph size={22}></AboutParagraph>
                    <AboutParagraph size={22}></AboutParagraph>
                    <AboutParagraph size={22}></AboutParagraph>
                </div>
            </div>
        );
    }
}