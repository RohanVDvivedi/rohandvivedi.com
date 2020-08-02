import React from "react";

import Icon from "../../utility/Icon";
import CopyToClipboard from "../../utility/Clipboard";

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

class ContactSubContent extends React.Component {
	render() {
		return (
			<div class="generic-content-box-border" style={{
				padding: "5px",
				width: this.props.width,
			}}>
				<div style={{
					fontFamily: "Capriola, Helvetica, sans-serif",
					textAlign: "center",
				}}>
					{this.props.title}
				</div>
				<div class="flex-row-container"
					style={{
						justifyContent: "space-evenly",
						alignItems: "center",
						padding: "3px",
					}}>
			    	{this.props.children}
				</div>
			</div>);
	}
}

export default class AboutContent extends React.Component {
    render() {
        return (
            <div class="content-root-background content-screen-widthed content-screen-heighted flex-col-container"
                style={{justifyContent: "center",
                        alignItems: "center",}}>

                    <div class="set_sub_content_background_color generic-content-box-border"
                    	style={{
                    		width: "60%",
	                    	padding: "1.5%",
                    	}}>

	                        <div class="flex-row-container" style={{width: "100%"}}>

	                            <img src={"/img/me_500h.jpg"} style={{width: "25%"}}/>
	                            
	                        	<div class="flex-col-container" style={{justifyContent: "space-between", marginLeft: "30px"}}>
	                            	<AboutParagraph size={20}>Hi, I am <span style={{ fontSize: "22px", fontWeight: "bold"}}>Rohan Dvivedi</span>.</AboutParagraph>
	                            	<AboutParagraph size={20}>I am a Software and Hardware Developer.</AboutParagraph>
	                            	<AboutParagraph size={20}>Predominantly a <ColoredBoldWord color="var(--color6)">Backend Developer</ColoredBoldWord>, who also indulges in building crappy <ColoredBoldWord color="var(--color6)">Frontend</ColoredBoldWord>s like this one.</AboutParagraph>
	                            	<AboutParagraph size={20}>My interests also include <ColoredBoldWord color="var(--color6)">Systems Programming, Databases, Computer Vision, Embedded Systems, Robotics, </ColoredBoldWord> and <ColoredBoldWord color="var(--color6)">FPGAs</ColoredBoldWord>.</AboutParagraph>
	                            </div>
	                        </div>

	                    <div class="flex-row-container"
		                    style={{
		                    	paddingTop: "2%",
		                    	justifyContent: "space-between",
		                    }}>

		                    <ContactSubContent title="CV" width="8%">
						        <Icon path="https://drive.google.com/file/d/12hE5q84en4QAsGkIlOcPEjlFL4kzgHxw/view?usp=sharing" iconPath="/icon/pdf.png" height="35px" width="35px" padding="5px" />
							</ContactSubContent>

							<ContactSubContent title="Email" width="40%">
								<a class="generic-content-box-hovering-emboss-border" href="#" onClick={()=>{CopyToClipboard("rohandvivedi@gmail.com")}} style={{display: "block",padding:"5px"}}>
						        	<div style={{display:"inline-block", fontSize: "15px"}}>rohandvivedi@gmail.com</div>
						        </a>
						        <Icon path="mailto:rohandvivedi@gmail.com" iconPath="/icon/mail.png" height="35px" width="35px" padding="5px" />
							</ContactSubContent>
						        
			                <ContactSubContent title="Online presence" width="35%">
					            <Icon path="https://github.com/RohanVDvivedi" iconPath="/icon/github.png" height="35px" width="35px" padding="5px" />
					            <Icon path="https://www.youtube.com/channel/UCgn_REjbUH2Dm8CaOXvajJg?view_as=subscriber" iconPath="/icon/youtube.png" height="35px" width="35px" padding="5px" />
					            <Icon path="https://www.linkedin.com/in/rohan-dvivedi-ab3014128/" iconPath="/icon/linkedin.png" height="35px" width="35px" padding="5px" />
					            <Icon path="https://www.facebook.com/rohan.dvivedi.961" iconPath="/icon/facebook.png" height="35px" width="35px" padding="5px" />
					        </ContactSubContent>

				    	</div>

                	</div>

            </div>
        );
    }
}