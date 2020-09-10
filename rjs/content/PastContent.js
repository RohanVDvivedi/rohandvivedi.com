import React from "react";

import ApiComponent from "../utility/ApiComponent";

class TimedEvent extends React.Component {
	render() {
		return (<div class="event-container flex-row-container">
					<div class="event-name">{this.props.revent}</div>
					<div class="event-date">{this.props.rtime}</div>
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
					<div class="past-position">
                        <span>{removeAbbreviationInBrackets(exp.Position)}</span>
                        <span class="hidden-only-mobile">{exp.Position.replace(removeAbbreviationInBrackets(exp.Position), "")}</span>
                    </div>
						{exp.Organizations.map(function(expOrg){
							return (
							<div style={{marginLeft:"10px", marginBottom:"3px"}}>
								<a href={expOrg.OrganizationLink} target="_blank" class="organization">
                                    <span>{removeAbbreviationInBrackets(expOrg.Organization)}</span>
                                    <span class="hidden-only-mobile">{expOrg.Organization.replace(removeAbbreviationInBrackets(expOrg.Organization), "")}</span>
                                </a>
								<div style={{marginLeft:"5px"}}>
									{expOrg.Teams.map(function(work){
										return (<div>
                                                    <TimedEvent revent={
                                                        (<span>
                                                            <span>{shortenByFirstComma(work.Team_or_ResearchTitle)}</span>
                                                            <span class="hidden-only-mobile">{work.Team_or_ResearchTitle.replace(shortenByFirstComma(work.Team_or_ResearchTitle), "")}</span>
                                                        </span>)} 
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

                    <div class="past-container set_sub_content_background_color generic-content-box-border">

                        <div class="past-PAST-title">
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