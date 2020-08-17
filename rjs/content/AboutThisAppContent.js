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
	            		This app with a mediocre looking front end, has backend which is quite complex (and over engineered in some sense).
	            	It is equivalent to something, that I would be expected from a portfolio of a backend developer.
	            	</p>

	            	The tech stack of this application includes
	            	<ul>
		            	<li>
		            		<a href="https://reactjs.org/" target="_blank">Reactjs</a>
		            		<ul>
		            			<li>
		            				With React router for client side routing
		            			</li>
		            		</ul>
		            	</li>
						<li>
							<a href="https://golang.org/" target="_blank">golang</a>
							<ul>
								<li>
									I discovered go to be a very friendly to my needs for building this application
								</li>
								<li>
									It appealed to me as it is one of the few systems programming languages with built-in web-dev modules.
								</li>
							</ul>
						</li>
						<li>
							<a href="https://www.sqlite.org/index.html" target="_blank">SQLite3</a>
							<ul>
								<li>
									I wanted a plain and simple to use, single file storage embedded database. SQLite is one of its kind.
								</li>
							</ul>
						</li>
					</ul>

					<p>
						You might be thinking, "Database for portfolio application? What is he even storing in there? His bank details or nuclear codes??".
						Well being honest I am not storing anything of much sense except for details about me and my projects,
						and neither do I have such a number of projects and accomplishments to really need 
						a well designed database to showcase them on a website such as this with a tiny minuscule search engine.
					</p>
					
					<p>
						I have extremely over engineered my portfolio, I wanted it simulate conditions for me to allow me
						to learn about single handedly managing a small website.

						At the other end of this extreme, I could build static portfolio website, 
						which can be easily hosted on an apache/nginx server, or use wordpress or any other 
						webhosting services. But where is the fun in that, where is the problem solving involved.
						If you are not maintaining the application, not managing all the resources it requires,
						then you are not learning enough from working on that project and If you are learning, then this application is not serving you any purpose.
					</p>

					<p>
						I built this application not just as my portfolio. Atleast when I started, 
						I wanted to build/design it to fit needs for just about any one.
						And that is one of the few reason why this application does not have a trivial or a non-existent backend.
						The Backend (built in go) serves to cater apis fo the front end from the database (SQLite3).
						Even my name on the about page and my past experiences on the pasts page, needs to be queried in from the persons and pasts table in the database.
						The frontend caches every api for atleast for 15 minutes, which reduces the load on the server.
	            		The whole point of this is that, someone(/anyone) can come up, and by just changing his/her information details in the database, this server could serve as anyone's portfolio.
	            	</p>

	            	<p>
	            		Moreover, the backend of this application is essentialy complex enough for what I wanted to do.
	            		I wanted to make an application that would grab the details about my github repositories using the github public apis,
	            		and index the readme files, and build a search engine for my projects, without needing me to update details to my portfolio everytime I push.
	            		I don't want to document my projects again and again.
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