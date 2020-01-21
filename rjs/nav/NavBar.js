import React from "react";
import AboutNav from "./navbuttons/AboutNav";
import ProjectsNav from "./navbuttons/ProjectsNav";
import ContactNav from "./navbuttons/ContactNav";
import SocialNav from "./navbuttons/SocialNav";
import ChatNav from "./navbuttons/ChatNav";

export default class NavBar extends React.Component {
    showNavNowAndHideAfterTms(Tms){
        // clear the previous time out
        if(this.timeoutNavContainer != null){
            clearTimeout(this.timeoutNavContainer);
            this.timeoutNavContainer = null;
        }
        // hide the nav container currently
        document.getElementById("nav-container").classList.remove("hide-nav");

        // set a timeout, to remove the NavContainer after T milliseconds
        var hideNavBar = function() {
            document.getElementById("nav-container").classList.add("hide-nav");
        }
        this.timeoutNavContainer = setTimeout(hideNavBar, Tms);
    }
    componentDidMount() {
        this.timeoutNavContainer = null;
        window.addEventListener('scroll', this.showNavNowAndHideAfterTms.bind(this, 2000));
        window.addEventListener('mousemove', this.showNavNowAndHideAfterTms.bind(this, 500));
    }
    render() {
        return (
            <div id="nav-container" class="nav-container-style flex-row-container"
            style={{
                justifyContent: "flex-end",
            }}>
                <AboutNav />
                <ProjectsNav />
                <ContactNav />
                <SocialNav />
                <ChatNav />
            </div>
        );
    }
}