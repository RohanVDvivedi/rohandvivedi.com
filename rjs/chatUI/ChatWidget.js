import React from "react";

import { MessageList, ChatList, Input, Button } from 'react-chat-elements'
import 'react-chat-elements/dist/main.css';

export default class ChatWidget extends React.Component {
	constructor(props) {
		super(props)
		this.state = {
			User: {

			},
			Chats : [
				{
					avatar: 'https://facebook.github.io/react/img/logo.svg',
					alt: 'Reactjs',
					title: 'Facebook',
					subtitle: 'What are you doing?',
					date: new Date(),
					unread: 0,
				},
				{
					avatar: 'https://facebook.github.io/react/img/logo.svg',
					alt: 'Reactjs',
					title: 'Facebook',
					subtitle: 'What are you doing?',
					date: new Date(),
					unread: 0,
				}
			],
			Party: {
				Messages : [
					{
						position: 'right',
						type: 'text',
						text: 'Lorem ipsum dolor sit amet, consectetur adipisicing elit',
						date: new Date(),
					},
				],
			}
		}
	}
	render() {
		return(
		<div class="chat-widget">

			<div class="messages">
				<MessageList className='message-list'
					toBottomHeight={'100%'} dataSource={this.state.Party.Messages} />

				<Input placeholder="Type here..." multiline={true} 
					rightButtons={<Button color='white' backgroundColor='black' text='Send'/>}/>
			</div>

			<div class="chats">
				<ChatList className='chat-list' dataSource={this.state.Chats} />
			</div>

		</div>);
	}
}