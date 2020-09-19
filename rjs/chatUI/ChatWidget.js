import React from "react";

import { MessageList, ChatList, Input, Button } from 'react-chat-elements'
import 'react-chat-elements/dist/main.css';
import Icon from '../utility/Icon'

export default class ChatWidget extends React.Component {
	updateState(objNew) {
		super.setState(Object.assign({}, this.state, objNew))
	}
	createMessageWidgetObject(msg) {
		return {
					position: msg.From == this.state.UserId ? 'right' : left,
					type: 'text',
					text: msg.Message,
					date: msg.SentAt,
				}
	}
	createChatWidgetObject(userId, userName, userPublicKey, userConnections, messageWidgetObjects) {
		return {
					userId: userId,
					avatar: null,
					alt: 'https://ui-avatars.com/api/?rounded=true&size=128&name=' + userName,
					title: userName,
					subtitle: 'What are you doing?',
					date: new Date(),
					unread: 0,
					messages: messageWidgetObjects,
				}
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

					avatar: 'https://ui-avatars.com/api/?rounded=true&size=128&name=' + "Jyotirmoy Pain",
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

					avatar: 'https://ui-avatars.com/api/?rounded=true&size=128&name=' + "Parthiv Kativarapu",
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
							status: "read",
						},
						{
							position: 'right',
							type: 'text',
							text: 'Lorem ipsum dolor sit amet, cing el dolor, consec te consec isicing icing el dol elit',
							date: new Date(),
							status: "sent"
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
	onUserJoinClicked() {
		console.log(this.refs.userName.input.value, this.refs.userEmail.input.value)
	}
	render() {
		console.log(this.state)
		return(
		<div class="chat-widget flex-row-container">

			{(!this.state.WindowOpen) ? (<Icon onClick={this.onChatBubbleClicked.bind(this)} iconPath="/icon/chat-bubble.png" height="40px" width="40px" padding="5px"/>) : ""}

			{ this.state.WindowOpen && this.UserId != null && this.state.ActiveChat != null ? 
			(<div class="chat-container">
				<div class="chat-header flex-row-container">
					<div>{this.state.ActiveChat.UserName}</div>
					<Icon onClick={this.onPartyWindowCloseClicked.bind(this)} iconPath="/icon/close.png" height="20px" width="20px" padding="3px"/>
				</div>
				<div class="chat-content">
					<MessageList className='message-list' toBottomHeight={'100%'} dataSource={this.state.ActiveChat.Messages} />
				</div>
				<Input className="chat-input" placeholder="Type here..." multiline={true} rightButtons={<Button className="chat-button" text='Send'/>}/>
			</div>) : ""}

			{ this.state.WindowOpen && this.UserId != null ? 
			(<div class="chat-container">
				<div class="chat-header flex-row-container">
					<div class="identifier">{this.state.UserName}</div>
					<Icon onClick={this.onChatWindowCloseClicked.bind(this)} iconPath="/icon/close.png" height="20px" width="20px" padding="3px"/>
				</div>
				<div class="chat-content">
					<ChatList className='chat-list' dataSource={this.state.Chats} onClick={this.onChatListItemClicked.bind(this)}/>
				</div>
				<Input className="chat-input" placeholder="Search user..." multiline={false} rightButtons={<Button className="chat-button" text='Search'/>}/>
			</div>) : ""}

			{ this.state.WindowOpen && this.UserId == null ? 
			(<div class="chat-container">
				<div class="chat-header flex-row-container">
					<div class="identifier">Join chat</div>
					<Icon onClick={this.onChatWindowCloseClicked.bind(this)} iconPath="/icon/close.png" height="20px" width="20px" padding="3px"/>
				</div>
				<div class="login-signup-content flex-col-container">
					<div class="lbl">Name :</div>
					<Input className="chat-input" placeholder="Name" multiline={false} ref="userName"/>
					<div class="lbl">Email :</div>
					<Input className="chat-input" placeholder="Email (or `gibberish` allowed)" multiline={false} ref="userEmail" />
					<Button className="chat-button" text='Join' onClick={this.onUserJoinClicked.bind(this)}/>
				</div>
			</div>) : ""}

		</div>);
	}
}