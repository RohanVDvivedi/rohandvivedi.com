import React from "react";

import ApiComponent from "../utility/ApiComponent";

class TimedEvent extends React.Component {
	render() {
		return (<div class="flex-row-container" style={{justifyContent: "space-between", alignItems: "baseline"}}>
					<div style={{fontSize:"16px", fontWeight: "550"}}>{this.props.revent}</div>
					<div style={{fontSize:"14px", fontStyle:"italic"}}>{this.props.rtime}</div>
				</div>);
	}
}

function removeAbbreviationInBrackets(str) {
    return (str.includes("(") && str.includes(")")) ? 
    (str.replace(str.slice(str.search("\\("), str.search("\\)") + 1), "").trim()) : str;
}

function shortenByFirstComma(str) {
    return str.includes(",") ? (str.slice(0, str.search(",")).trim()) : str;
}

class Experience extends React.Component {
	render() {
		var exp = this.props.exp;
		var dateFormat = {month:"short", year:"2-digit"}
		return (<div style={{marginTop: "15px"}}>
					<div style={{fontSize:"18px",fontWeight: "600"}}>
                        <span>{removeAbbreviationInBrackets(exp.Position)}</span>
                        <span class="hidden-only-mobile">{exp.Position.replace(removeAbbreviationInBrackets(exp.Position), "")}</span>
                    </div>
						{exp.Organizations.map(function(expOrg){
							return (
							<div style={{marginLeft:"10px", marginBottom:"3px"}}>
								<a href={expOrg.OrganizationLink} target="_blank" style={{fontSize:"15px"}}>{expOrg.Organization}</a>
								<div style={{marginLeft:"5px"}}>
									{expOrg.Teams.map(function(work){
										return (<div>
                                                    <TimedEvent revent={(<span><span>{shortenByFirstComma(work.Team_or_ResearchTitle)}</span>
                                                        <span class="hidden-only-mobile">{work.Team_or_ResearchTitle.replace(shortenByFirstComma(work.Team_or_ResearchTitle), "")}</span></span>)} 
                                                    rtime={work.FromDate.toLocaleDateString("en-US", dateFormat) + " - " + work.ToDate.toLocaleDateString("en-US", dateFormat)} />
													{work.PastType == "RESEARCH" ? (<div style={{fontWeight: "600"}}>Research paper: <a href={work.ResearchPaperLink} target="_blank">DOI link here</a></div>) : ""}
													<div class="hidden-only-mobile">{work.Descr == null ? "" : work.Descr}</div>
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
            <div class="content-root-container content-root-background flex-col-container"
                style={{justifyContent: "center",
                        alignItems: "center",}}>

                    <div class="set_sub_content_background_color generic-content-box-border"
                    	style={{ padding: "1.5%",
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