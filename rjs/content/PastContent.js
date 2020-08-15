import React from "react";

import ApiComponent from "../utility/ApiComponent";

class TimedEvent extends React.Component {
	render() {
		return (<div class="flex-row-container" style={{justifyContent: "space-between"}}>
					<div style={{fontSize:"17px", fontWeight: "550"}}>{this.props.revent}</div>
					<div style={{fontStyle:"italic"}}>{this.props.rtime}</div>
				</div>);
	}
}

class Experience extends React.Component {
	render() {
		var exp = this.props.exp;
		var dateFormat = {month:"short", year:"2-digit"}
		return (<div style={{marginTop: "15px"}}>
					<div style={{fontSize:"22px",fontWeight: "600"}}>{exp.Position}</div>
						{exp.Organizations.map(function(expOrg){
							return (
							<div style={{marginLeft:"10px", marginBottom:"3px"}}>
								<a href={expOrg.OrganizationLink} target="_blank" style={{fontSize:"18px"}}>{expOrg.Organization}</a>
								<div style={{marginLeft:"5px"}}>
									{expOrg.Teams.map(function(work){
										return (<div>
													<TimedEvent revent={work.Team_or_ResearchTitle} rtime={work.FromDate.toLocaleDateString("en-US", dateFormat) + " - " + work.ToDate.toLocaleDateString("en-US", dateFormat)} />
													{work.PastType == "RESEARCH" ? (<div style={{fontWeight: "600"}}>Research paper: <a href={work.ResearchPaperLink} target="_blank">DOI link here</a></div>) : ""}
													<div>{work.Descr == null ? "" : work.Descr}</div>
												</div>);
									})}
								</div>
							</div>)
						})}
				</div>);
	}
}

export default class PastContent extends ApiComponent {
	apiPath() {
        return "/api/owner?get_pasts=true";
    }
    bodyDataBeforeApiFirstResponds() {
    	return {Fname:"Firstname",Lname:"Lastname",Email:"loading email id","Socials":null,"Pasts":[]};
    }
    render() {
    	var owner = this.state.api_response_body;
    	var pasts = owner["Pasts"]	.map((pasti) => {
    									pasti.FromDate = new Date(pasti.FromDate)
    									pasti.ToDate = new Date(pasti.ToDate)
    									return pasti
    								}).sort((past1, past2) => {
    									var compare = compareDates(past1.FromDate, past2.FromDate)
    									if(compare == 0){
    										compare = compareDates(past1.ToDate, past2.ToDate)
    									}
    									return -compare;
    								});
    	var pastsCombine = []
    	pasts.forEach((past) => {
    		if(pastsCombine.length > 0 && past.Position == pastsCombine[pastsCombine.length-1].Position) {
    			var Organizations = pastsCombine[pastsCombine.length-1].Organizations
				if(past.Organization == Organizations[Organizations.length-1].Organization){
					var Teams = Organizations[Organizations.length-1].Teams
    				Teams.push(past)
    			} else {
    				Organizations.push({
    										Organization: past.Organization,
    										OrganizationLink: past.OrganizationLink,
    										Teams: [past]
    									})
    			}
    		} else {
    			pastsCombine.push({
    									Position: past.Position,
    									Organizations: [
    														{
    															Organization: past.Organization,
    															OrganizationLink: past.OrganizationLink,
    															Teams: [past]
    														}
    													]
    								})
    		}
    	})

    	pasts = pastsCombine

        return (
            <div class="content-root-background content-screen-widthed content-screen-heighted flex-col-container"
                style={{justifyContent: "center",
                        alignItems: "center",}}>

                    <div class="set_sub_content_background_color generic-content-box-border"
                    	style={{ width: "60%", padding: "1.5%",
                    		color: "var(--font_color_lighter)",
                    		}}>

                        <div style={{
                        	textAlign: "center",
							fontFamily: "lato, sans-serif",
							fontSize: "25px",
							fontWeight: "700",
							fontStyle: "italic"
                        }}>
                            Past
                        </div>

                        <div>
	                        {pasts.map(function(past){
	                        	return <Experience exp={past}/>
	                        })}
	                    </div>

                    </div>

            </div>
        );
    }
}

function compareDates(a, b) {
	return (a < b) ? -1 : ((a > b) ? 1 : 0)
}

    /*
    Past
Work Experience
OYO
Position: SDE1 (Software Development Engineer I)
OYO Vacation Homes, Amsterdam, Netherlands
Jul’19-Feb’20
Finance Tech. Team, Gurgaon, India
Dec’18-Jun’19
Supply Tech. Team, Gurgaon, India
Aug’18-Dec’18
Job Description: SDE1, Backend Developer contributing in java (spring boot) and ruby on rails techstack to Finance and Supply Technology teams at OYO. I also contributed in website revamp and feature development for OYO Vacation Homes' subsidiary entity Belvilla B.V..
Research Experience
Flexible Processor Architecture Design
Jul’17-Dec’17
DOI: 10.1109/DISCOVER47552.2019.9008052
Authors : D. R. Vipulkumar, P. V. Bhanu and J. Soumya
Abstract problem statement: To Design a processor that can execute any instruction set (custom designed or commercial). To try and encourage programmers to come up with their own custom instructions to target higher efficiencies for their applications.
Undergraduate Education
BITS Pilani
B.E. (Hons.) in Mechanical Engineering
Jul’14-Jul’18
Description: I graduated from BITS Pilani, Hyderabad Campus, majoring in Mechanical Engineering. My exposure to Computer Science, during my undergraduate education, gave me confidence to explore the field further which led me to learn to analyze problems and code their way out.
*/