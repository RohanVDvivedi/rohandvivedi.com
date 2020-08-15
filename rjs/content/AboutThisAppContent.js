import React from "react";

export default class AboutThisAppContent extends React.Component {
	render() {
		return (
			<div class="content-container content-root-background"
				style={{padding: "5% 15% 2% 15%"}}>
				<div class="set_sub_content_background_color generic-content-box-border"
				style={{
					padding: "5%",
					display: "none"
				}}>
	            	<p>
	            		This app with a mediocre looking front end, has a quite well organized backend.
	            	Equivalent to something, that could be expected from a skilled Backend Developer.
	            	</p>

	            	The tech stack of this application includes
	            	<ul>
		            	<li>ReactJS<i>(with react router for client side routing)</i></li>
						<li>golang<i>(without any framework, builtin go modules are well designed to cater all my needs)</i></li>
						<li>sqlite3<i>(plain single file managed embedded database, Everything you see on this website is served dynamically through api's even my work experinece has a well designed table devoated to itsef)</i></li>
					</ul>

					<p>
					I know right, database for a portfolio application? And, No I do not even have 100 projects under my belt to really need a well designed database to showcase them. 
					And yes as you see here it is not as tough to go through all of my projects, you surely won't need the search bar provided.
					</p>
					
					<p>
					So the question is why over engineer something and so seriously. One of the very reason is that it is just who I am.
					I like building stuff, whether it is software or a hardware. I like intericately designing and solving echnical projects.
					But let's not get into this. I do not really want to make this page about myself (I have one for that here).
					</p>

					<p>
					I built this application not just to fulfill my interest in building a self spotlighting portfolio.
	            	When I started, I did not build/design it to be just for my self, I wanted to make an application that any one can use,
	            	by just changing his/her information details in the database, anyone could turn this application in to their database.
	            	</p>

	            	<p>
	            	I wanted to make an application that would grab the owners details about his/her github repositories using the github public apis,
	            	index the readme files, and build a search engine for only his/her projects.
	            	but I can see, that every developer (be is software developer, game dev, UI/UX etc all have their own needs)
	            	</p>

	            	<p>
	            	I started this project to fit everyone's needs, it does not fit all their needs. Even if linkedin/github could not target professional life of 
	            	individuals how could I. but what I could really do is to automate my life, I do not really like, updating about myself or my projects
	            	I literally hate documenting my projects.
	            	</p>

	            	<p>
	            	So built this app to do something similar for me. It uses all of my github repos, and stores it all in a databse
	            	it even updates my information from linkedin. and my resume from my own google drive.
	            	</p>


	            	
	            	
	            	<p>
	            	Thank you.<br/>
	            	Rohan Dvivedi,<br/>
	            	Creator of rohandvivedi.com.<br/>
	            	</p>
            	</div>
            </div>
        );
    }
}