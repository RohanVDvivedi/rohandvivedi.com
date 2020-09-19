import React from "react";

import { MessageList, ChatList, Input, Button } from 'react-chat-elements'
import 'react-chat-elements/dist/main.css';
import Icon from '../utility/Icon'
import Chatter from '../chatter/Chatter'

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
		latestMessage = messageWidgetObjects == null || messageWidgetObjects.Length == 0 ? null : messageWidgetObjects[messageWidgetObjects.Length - 1]
		return {
					userId: userId,
					userName: userName,
					userPublicKey: userPublicKey,
					isOnline: userConnections > 0,
					avatar: null,
					alt: 'https://ui-avatars.com/api/?rounded=true&size=128&name=' + userName,
					title: userName,
					subtitle: latestMessage == null ? "" : latestMessage.text,
					date: latestMessage == null ? "" : latestMessage.date,
					unread: 0,
					messages: messageWidgetObjects,
				}
	}
	constructor(props) {
		super(props)
		Chatter.ReqConnection()
		this.state = {
			WindowOpen: false,
			UserId: null,
			UserName: null,
			ActiveChat: null,
			Chats : null
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

			{ this.state.WindowOpen && this.state.UserId != null && this.state.ActiveChat != null ? 
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

			{ this.state.WindowOpen && this.state.UserId != null ? 
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

			{ this.state.WindowOpen && this.state.UserId == null ? 
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