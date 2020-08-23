import React from "react";

import { NavLink } from 'react-router-dom';

export default class AboutThisAppContent extends React.Component {
	render() {
		return (
			<div class="content-container content-root-background"
				style={{
					paddingLeft: "5%",
					paddingRight: "5%"
				}}>
				<div class="behind-nav"></div>
				<div class="set_sub_content_background_color generic-content-box-border"
				style={{
					padding: "3%",
				}}>
					<p>
						Github: <a href="https://github.com/RohanVDvivedi/rohandvivedi.com" target="_blank">rohandvivedi.com</a>
					</p>
	            	<p>
						This app with a mediocre-looking front end has backend which is quite over-engineered in some aspects. Yet, I would consider it equivalent in quality to something, that I would expect in his/her portfolio from a backend developer.
	            	</p>

					The tech stack of this application includes:
	            	<ul>
		            	<li>
		            		<a href="https://reactjs.org/" target="_blank">Reactjs</a>
		            		<ul>
		            			<li>
		            				With React router for client-side routing.
		            			</li>
		            		</ul>
		            	</li>
						<li>
							<a href="https://golang.org/" target="_blank">golang</a>
							<ul>
								<li>
									go (short for golang) is well equipped with built-in packages to fit my requirements for building this application.
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
									I wanted a plain and simple to use, single file storage embedded database. SQLite happens to be one of its kind.
								</li>
							</ul>
						</li>
						<li>
							<a href="https://blevesearch.com/" target="_blank">Bleve</a>
							<ul>
								<li>
									Again I wanted an embedded search index; and Bleve ended up as the first result on my google search.
									<br/><i>(Haa.. ha., saw what I did there!)</i>
								</li>
							</ul>
						</li>
					</ul>

					<p>
						You might be thinking, "Database and search index for a portfolio application !!!". Well, being honest, I am not storing anything of much sense except for details about me and my projects, and I do not have numerous projects or accomplishments that would really require a well-designed database and a search index on top of that, to showcase them on a website such as this.
					</p>
					
					<p>
						I am well aware, that I have extremely over-engineered my portfolio, in trying to simulate conditions for me to allow me to learn to single-handedly manage a small website. At the other end of this extreme, I could build a static portfolio website, which can be easily hosted on an Apache2 (HTTPD) server or Nginx server or use WordPress or any other templated web hosting services but where is the fun in that, where is the problem solving involved. If you are not maintaining the application, not managing all the resources it requires, then you are not learning enough by working on that project and if you are not learning, then this application is not serving you (in this case me) any purpose. Static website and web hosting services do not solve my purpose, I am not a content developer or UI/UX designer. I am a backend developer, a software and hardware developer, I too desire a beautiful and aesthetic frontend for my portfolio, but elegance and simplicity is what I need.
					</p>

					<p>
						I built this application not just as my portfolio. At least when I started, I wanted to build/design it to fit needs for just about anyone and that is one of the few reasons why this application does not have a trivial or a non-existent backend. The backend (built-in go) serves to cater to APIs for the frontend from the database (SQLite3). Even my name on the 
						<i> <NavLink to="/pages/about">about page</NavLink> </i>
						and my past experiences on the
						<i> <NavLink to="/pages/past">past page</NavLink> </i>
						need to be queried in from the persons and pasts table in the database. The frontend caches every HTTP API in the local storage, for at least 15 minutes to reduce the load on the server.
					</p>

					<p>
						The backend is devised to provide a meaning full search (using Bleve), to anyone who wants to surf through projects of the owner. There are cron jobs and HTTP APIs built and setup (authorizing only the owner) to execute and to fetch
						<i> README.md </i>
						files of various projects of the owner from corresponding Github repositories and to rebuild this search index. The backend also provides many powerful core APIs authorizing only the owner to control the core functionality of this portfolio. The owner has access to APIs to get the session values of users.
	            		<i> (yes, I can know exactly when you clicked to open this page, LOL) </i>
	            		This might sound creepy, but such data collection helps the owner, to reconsider the design and routing of the website and optimize it for user-friendliness at every point.
					</p>

					<p>
						In the end, the whole point of this ridiculous portfolio is not to just showcase my projects, but to showcase this application, which is meant for anyone who wants, to come up, and host their awesome low maintenance automatically updatable portfolio within minutes.
					</p>

					<p>
						Moreover, the backend of this application is essential enough (and not over-engineered) for what I wanted to do in my portfolio.
					</p>
	            	
	            	
	            	<p>
		            	Thank you,<br/>
		            	Rohan Dvivedi,<br/>
		            	Creator of rohandvivedi.com.<br/>
	            	</p>
            	</div>
            </div>
        );
    }
}