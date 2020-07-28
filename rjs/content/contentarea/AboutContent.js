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
                fontWeight: "500",
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
    render() {
        return (
            <div class="content-root-background content-screen-widthed content-screen-heighted flex-col-container"
                style={{justifyContent: "center",
                        alignItems: "center",}}>

                    <div class="set_sub_content_background_color" style={{
                    	height: "20%", 
                    	minHeight: "180px", 
                    	padding: "2%",
                    }}>
                        <div class="flex-row-container" style={{height: "100%"}}>

                            <img src={"/img/me_500h.jpg"} style={{height: "100%"}}/>
                            
                        	<div class="flex-col-container" style={{justifyContent: "space-between", marginLeft: "30px"}}>
                            	<AboutParagraph size={20}>Hi, I am <span style={{ fontSize: "22px", fontWeight: "bold"}}>Rohan Dvivedi</span>.</AboutParagraph>
                            	<AboutParagraph size={20}>I am a Software and Hardware Developer.</AboutParagraph>
                            	<AboutParagraph size={20}>Predominantly a <ColoredBoldWord color="var(--color6)">Backend Developer</ColoredBoldWord>,<br/> who also indulges in building crappy <ColoredBoldWord color="var(--color6)">Frontend</ColoredBoldWord>s like this one.</AboutParagraph>
                            	<AboutParagraph size={20}>My interests also include <ColoredBoldWord color="var(--color6)">Systems Programming, Databases,<br/>Computer Vision, Embedded Systems, Robotics, </ColoredBoldWord> and <ColoredBoldWord color="var(--color6)">FPGAs</ColoredBoldWord>.</AboutParagraph>
                            	<AboutParagraph size={20}>Find my Curriculum Vitae <a href={"https://drive.google.com/file/d/12hE5q84en4QAsGkIlOcPEjlFL4kzgHxw/view?usp=sharing"} target="_blank">here</a>.</AboutParagraph>
                            </div>
                        </div>
                    </div>

            </div>
        );
    }
}