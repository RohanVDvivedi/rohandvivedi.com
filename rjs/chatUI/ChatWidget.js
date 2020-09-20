import React from "react";

import { MessageList, ChatList, Input, Button } from 'react-chat-elements'
import 'react-chat-elements/dist/main.css';
import Icon from '../utility/Icon'
import Chatter from '../chatter/Chatter'

function createMessageWidgetObject(msg) {
	return {
		position: msg.From == this.state.UserId ? 'right' : left,
		type: 'text',
		text: msg.Message,
		date: msg.SentAt,
	}
}
function createChatWidgetObject(user, messageWidgetObjects) {
	var latestMessage = messageWidgetObjects == null || messageWidgetObjects.Length == 0 ? null : messageWidgetObjects[messageWidgetObjects.Length - 1]
	return {
		userId: user.Id,
		userName: user.Name,
		userPublicKey: user.PublicKey,
		isOnline: user.ConnectionCount > 0,
		isActive: false,
		avatar: 'https://ui-avatars.com/api/?rounded=true&size=128&name=' + user.Name,
		alt: user.Name,
		title: user.Name,
		subtitle: latestMessage == null ? "" : latestMessage.text,
		date: latestMessage == null ? "" : latestMessage.date,
		unread: 0,
		messages: messageWidgetObjects == null ? [] : messageWidgetObjects,
	}
}

export default class ChatWidget extends React.Component {
	updateState(objNew) {
		super.setState(Object.assign({}, this.state, objNew))
	}
	constructor(props) {
		super(props)
		this.state = {
			WindowOpen: false,
			UserId: null,
			UserName: null,
			ActiveChat: null,
			ChatsById : null
		}
		Chatter.onLogin = (function() {
			this.updateState({
				UserId: Chatter.UserId,
				UserName: Chatter.UserName,
			})
		}).bind(this)
		Chatter.onChangeUsersList = (function(userList) {
			var ChatsById = {}
			var oldState = Object.assign({}, this.state)
			userList.filter(function(user){
				return user.Id != oldState.UserId
			}).forEach(function(user){
				var oldMessageWidgets = (oldState.ChatsById == null || oldState.ChatsById[user.Id] == null) ? [] : oldState.ChatsById[user.Id].messages
				ChatsById[user.Id] = createChatWidgetObject(user,oldMessageWidgets)
			})
			this.updateState({
				ChatsById: ChatsById,
			})
		}).bind(this)
		Chatter.onUserNotification = (function(user) {
			if(user.Id == this.state.UserId) {
				return
			}
			var ChatsById = Object.assign({}, this.state.ChatsById)
			var oldMessageWidgets = (this.state.ChatsById == null || this.state.ChatsById[user.Id] == null) ? [] : this.state.ChatsById[user.Id].messages
			ChatsById[user.Id] = createChatWidgetObject(user,oldMessageWidgets)
			this.updateState({
				ChatsById: ChatsById,
			})
		}).bind(this)
		Chatter.onLogout = (function() {
			this.updateState({
				UserId: null,
				UserName: null,
			})
		}).bind(this)
	}
	componentDidMount() {
		Chatter.ReqConnection()
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
		var ChatsById = Object.assign({}, this.state.ChatsById)
		ChatsById[c.userId].unread = 0
		this.updateState({ActiveChat: c, ChatsById: ChatsById})
	}
	onMessageSend() {

	}
	onUserSearch() {

	}
	onUserSignInClicked() {
		Chatter.ReqLogin(this.refs.userName.input.value, this.refs.userEmail.input.value)
	}
	onUserSignUpClicked() {
		Chatter.ReqCreateUser(this.refs.userName.input.value, this.refs.userEmail.input.value)
	}
	onUserSignoutClicked() {
		Chatter.ReqLogout()
	}
	render() {
		console.log(this.state)
		var chatsArray = []
		for (const userId in this.state.ChatsById) {
			chatsArray.push(this.state.ChatsById[userId])
		}
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
					<ChatList className='chat-list' dataSource={chatsArray} onClick={this.onChatListItemClicked.bind(this)}/>
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
					<div className="chat-button-container flex-row-container">
						<Button className="chat-button" text='Sign in' onClick={this.onUserSignInClicked.bind(this)}/>
						<Button className="chat-button" text='Sign up' onClick={this.onUserSignUpClicked.bind(this)}/>
					</div>
				</div>
			</div>) : ""}

		</div>);
	}
}