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
                padding: "10px",
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
                    <AboutParagraph size={26}>Hi, I am Rohan.</AboutParagraph>
                    <AboutParagraph size={26}>I am a Software and Hardware Developer.</AboutParagraph>
                    <AboutParagraph size={24}>Predominantly a <ColoredBoldWord color="var(--color6)">Backend Developer</ColoredBoldWord>.</AboutParagraph>
                    <AboutParagraph size={24}>Tinkering with <ColoredBoldWord color="var(--color6)">Embedded</ColoredBoldWord> and <ColoredBoldWord color="var(--color6)">FPGA systems</ColoredBoldWord> is my hobby</AboutParagraph>
                    <AboutParagraph size={24}>and I also build some crappy <ColoredBoldWord color="var(--color6)">Frontend</ColoredBoldWord>s like this one</AboutParagraph>
                    <AboutParagraph size={20}>It is pleasure meeting you..</AboutParagraph>

                </div>
                <div class="flex-col-container"
                    style={{
                        width: "30%",
                        justifyContent: "center",
                    }}
                >
                    <AboutParagraph size={22}>Curriculum Vitae</AboutParagraph>
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