import React from "react";

export default class AboutThisAppContent extends React.Component {
	render() {
		return (
			<div class="content-container content-root-background"
				style={{padding: "5% 15% 2% 15%"}}>
				<div class="set_sub_content_background_color generic-content-box-border"
				style={{
					padding: "5%",
				}}>
	            	<p>
	            		This app with a mediocre looking front end, has backend which is quite over engineered in some sense.
	            	Yet, I would consider it equivalent in quality to something, that I would expect in a portfolio of a backend developer.
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
									go (short for golang) is well equipped with built-in packages to fit my requirements for building this application
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
						<li>
							<a href="https://blevesearch.com/" target="_blank">Bleve</a>
							<ul>
								<li>
									Again I wanted an embedded search index; and Bleve ended up as the first result on my google search. <i>(ha.. ha. saw what I did there)</i>
								</li>
							</ul>
						</li>
					</ul>

					<p>
						You might be thinking, "Database and search index for portfolio application !!!".
						Well being honest, I am not storing anything of much sense except for details about me and my projects,
						and I do not have numerous projects or accomplishments that I would really require 
						a well designed database (and a search index on top of that), to showcase them on a website such as this.
					</p>
					
					<p>
						I an well aware, that I have extremely over engineered my portfolio, 
						I wanted it simulate conditions for me to allow me
						to learn about single handedly managing a small website.
					</p>

					<p>
						At the other end of this extreme, I could build static portfolio website, 
						which can be easily hosted on an apache/nginx server, or use wordpress or any other 
						templated webhosting services. But where is the fun in that, where is the problem solving involved.
						If you are not maintaining the application, not managing all the resources it requires,
						then you are not learning enough from working on that project and if you are not learning, 
						then this application is not serving you (in this case me) any purpose.
					</p>

					<p>
						I built this application not just as my portfolio. Atleast when I started, 
						I wanted to build/design it to fit needs for just about any one.

						And that is one of the few reason why this application does not have a trivial or a non-existent backend.

						The backend (built in go) serves to cater apis for the frontend from the database (SQLite3).
						Even my name on the about page and my past experiences on the pasts page, needs to be queried in from the persons and pasts table in the database.
						The frontend caches every api in the local storage, for atleast 15 minutes to reduce the load on the server.
						The backend is devised to provide a meaning full search (using Bleve), to any one who wants to surf through projects of the owner.
						There are cron jobs and HTTP APIs (authorizing only the owner) built to fetch readme files of various projects of the owner from corresponding Github repositories and to rebuild this search index.
						The backend provides authentication support to owner, providing owner (and only the owner) with powerfull API's, to access sessions of users.
						(yes, I know exactly when you clicked to open this page, LOL)
						This might sound creepy, but such data collection helps the owner, to reconsider the design and routing of the website and optimize it for user friendliness at every point.

						In the end, the whole point of this ridiculous portfolio is not to showcase my projects, but to showcase this application.
	            		Which is meant for anyone who wants, to come up, and host their awesome low maintenance automatically updatable protfolio within minutes. (I made that up, I am short of words to describe this application.)
	            	</p>

	            	<p>
	            		Moreover, the backend of this application is essential enough (and not over engineered) for what I wanted to do in my portfolio.
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