import React from "react";

import ApiComponent from "../utility/ApiComponent";
import Icon from "../utility/Icon";
import CopyToClipboard from "../utility/Clipboard";

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

function checkIfSocialIsA_CV(social)
{
	return social.Descr.includes("CV");
}

export default class AboutContent extends ApiComponent {
	apiPath() {
        return "/api/owner?get_socials=true";
    }
    bodyDataBeforeApiFirstResponds() {
    	return {Fname:"Firstname",Lname:"Lastname",Email:"loading email id","Socials":[],"Pasts":null};
    }
    render() {
    	var owner = this.state.api_response_body;
    	var cv = owner.Socials.find(checkIfSocialIsA_CV);
    	if(cv != null) {
    		cv = (<Icon path={cv.ProfileLink} iconPath={"/icon/" + cv.LinkType + ".png"} height="35px" width="35px" padding="5px" />);
    	} else {
    		cv = "";
    	}
        return (
            <div class="content-root-background content-screen-widthed content-screen-heighted flex-col-container"
                style={{justifyContent: "center",
                        alignItems: "center",}}>

                    <div class="set_sub_content_background_color generic-content-box-border"
                    	style={{
                    		width: "65%",
	                    	padding: "1.5%",
                    	}}>

	                        <div class="flex-row-container" style={{width: "100%"}}>

	                            <img src={"/img/me_500h.jpg"} style={{width: "25%"}}/>
	                            
	                        	<div class="flex-col-container" style={{justifyContent: "space-between", marginLeft: "30px"}}>
	                            	<div class="about-paragraph">Hi, I am <span class="owner-name">{owner.Fname + " " + owner.Lname}</span>.</div>
	                            	<div class="about-paragraph">I am a Software and Hardware Developer.</div>
	                            	<div class="about-paragraph">Predominantly a <span class="skills-bolden-coloren">Backend Developer</span>, who also indulges in building crappy <span class="skills-bolden-coloren">Frontend</span>s like this one.</div>
	                            	<div class="about-paragraph">My interests also include <span class="skills-bolden-coloren">Systems Programming, Databases, Computer Vision, Embedded Systems, Robotics, </span> and <span class="skills-bolden-coloren">FPGAs</span>.</div>
	                            </div>
	                        </div>

	                    <div class="flex-row-container"
		                    style={{
		                    	paddingTop: "2%",
		                    	justifyContent: "space-between",
		                    }}>

		                    <ContactSubContent title="CV" width="8%">
						        {cv}
							</ContactSubContent>

							<ContactSubContent title="Email" width="40%">
								<a class="generic-content-box-hovering-emboss-border tooltip-container" href="#" onClick={()=>{CopyToClipboard("rohandvivedi@gmail.com")}} style={{display: "block",padding:"10px"}}>
						        	<div style={{display:"inline-block", fontSize: "17px"}}>{owner.Email}</div>
						        	<div class="tooltip-content">Click to copy</div>
						        </a>
							    <Icon path={"mailto:" + owner.Email} iconPath="/icon/mail.png" infoBoxText="Use mail client" height="35px" width="35px" padding="5px" />
							</ContactSubContent>
						        
			                <ContactSubContent title="Online presence" width="35%">
			                	{owner.Socials.filter(function(social){return !checkIfSocialIsA_CV(social);}).map(function(social){
					            	return <Icon path={social.ProfileLink} iconPath={"/icon/" + social.LinkType + ".png"} infoBoxText={"my " + social.LinkType + " profile"} height="35px" width="35px" padding="5px" />
					        	})}
					        </ContactSubContent>

				    	</div>

                	</div>

            </div>
        );
    }
}