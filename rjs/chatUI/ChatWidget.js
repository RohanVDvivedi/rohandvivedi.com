import React from "react";

import { MessageList, ChatList, Input, Button } from 'react-chat-elements'
import 'react-chat-elements/dist/main.css';
import Icon from '../utility/Icon'

export default class ChatWidget extends React.Component {
	updateState(objNew) {
		super.setState(Object.assign({}, this.state, objNew))
	}
	constructor(props) {
		super(props)
		this.state = {
			WindowOpen: false,
			UserName: "Rohan",
			UserId: "",
			ActiveChat: null,
			Chats : [
				{
					UserId: "user_id_1",
					UserName: 'Jyotirmoy',

					avatar: '/avatar/avatar.svg',
					alt: 'J',
					title: 'Jyotirmoy',
					subtitle: 'What are you doing?',
					date: new Date(),
					unread: 2,

					Messages: [
						{
							position: 'left',
							type: 'text',
							text: 'etur adip sicing elit',
							date: new Date(),
						},
						{
							position: 'right',
							type: 'text',
							text: 'Lorem ipsum dolor sit ang icing el dol elit',
							date: new Date(),
						},
					],
				},
				{
					UserId: "user_id_1",
					UserName: 'Parthiv',

					avatar: '/avatar/avatar.svg',
					alt: 'P',
					title: 'Parthiv',
					subtitle: 'Lets go goa',
					date: new Date(),
					unread: 0,

					Messages: [
						{
							position: 'left',
							type: 'text',
							text: 'Lorem ipsum tur adipisicing el dolor, consec tetur adipi tetur adip sicing elit',
							date: new Date(),
						},
						{
							position: 'right',
							type: 'text',
							text: 'Lorem ipsum dolor sit amet, cing el dolor, consec te consec isicing icing el dol elit',
							date: new Date(),
						},
					],
				}
			]
		}
	}
	onChatBubbleClicked() {
		this.updateState({WindowOpen: true})
	}
	onPartyWindowCloseClicked() {
		this.updateState({ActiveChat: null})
	}
	onChatWindowCloseClicked() {
		this.updateState({WindowOpen: false})
	}
	onChatListItemClicked(c) {
		this.updateState({ActiveChat: c})
	}
	onMessageSend() {

	}
	onUserSearch() {

	}
	render() {
		console.log(this.state)
		return(
		<div class="chat-widget flex-row-container">

			{(!this.state.WindowOpen) ? (<Icon onClick={this.onChatBubbleClicked.bind(this)} iconPath="/icon/chat-bubble.png" height="40px" width="40px" padding="5px"/>) : ""}

			{ this.state.WindowOpen && this.state.ActiveChat != null ? 
			(<div class="container">
				<div class="header flex-row-container">
					<div>{this.state.ActiveChat.UserName}</div>
					<Icon onClick={this.onPartyWindowCloseClicked.bind(this)} iconPath="/icon/close.png" height="20px" width="20px" padding="3px"/>
				</div>
				<div class="content">
					<MessageList className='message-list' toBottomHeight={'100%'} dataSource={this.state.ActiveChat.Messages} />
				</div>
				<Input placeholder="Type here..." multiline={true} rightButtons={<Button color='white' backgroundColor='black' text='Send'/>}/>
			</div>) : ""}

			{ this.state.WindowOpen ? 
			(<div class="container">
				<div class="header flex-row-container">
					<div class="identifier">{this.state.UserName}</div>
					<Icon onClick={this.onChatWindowCloseClicked.bind(this)} iconPath="/icon/close.png" height="20px" width="20px" padding="3px"/>
				</div>
				<div class="content">
					<ChatList className='chat-list' dataSource={this.state.Chats} onClick={this.onChatListItemClicked.bind(this)}/>
				</div>
				<Input placeholder="Search user..." multiline={false} rightButtons={<Button color='white' backgroundColor='black' text='Search'/>}/>
			</div>) : ""}

		</div>);
	}
}