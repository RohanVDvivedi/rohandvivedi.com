import React from "react";

import AboutContent from "./contentarea/AboutContent";
import ProjectsContent from "./contentarea/ProjectsContent";
import ContactContent from "./contentarea/ContactContent";
import SocialContent from "./contentarea/SocialContent";
import ChatContent from "./contentarea/ChatContent";

export default class ContentArea extends React.Component {
    render() {
        return (
            <div id="content-container">
                <AboutContent />
                <ProjectsContent />
                <ContactContent />
                <SocialContent />
                <ChatContent />
            </div>
        );
    }
}