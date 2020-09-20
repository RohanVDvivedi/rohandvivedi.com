import React from "react";

import { MessageList, ChatList, Input, Button } from 'react-chat-elements'
import 'react-chat-elements/dist/main.css';

import Icon from '../utility/Icon'

import Chatter from '../chatter/Chatter'

export default class ChatWidget extends React.Component {
	updateState(objNew) {
		super.setState(Object.assign({}, this.state, objNew))
	}
	constructor(props) {
		super(props)
		this.state = {
			WindowOpen: false,
			User: null,	// {}
			ActiveChatUserId: null,
			ChatUsersById : null,	//{ User.Id => {User, Unread, ChatMessageQueue}
		}

		Chatter.onLogin = (function() {this.updateState({User: Chatter.User}); Chatter.ReqGetAllUsers()}).bind(this)
		Chatter.onLogout = (function() {this.updateState({User: null,})}).bind(this)
		Chatter.onClose = (function() {this.updateState({WindowOpen:false,User: null})}).bind(this)

		Chatter.onChangeUsersList = (function(userList) {
			var ChatUsersById = {}
			userList.forEach(function(user){
				ChatUsersById[user.Id].User = user;
			})
			this.updateState({ChatUsersById: ChatUsersById})
		}).bind(this)
		Chatter.onUserNotification = (function(user) {
			var ChatUsersById = Object.assign({}, this.state.ChatUsersById)
			ChatUsersById[user.Id].User = user;
			this.updateState({ChatUsersById: ChatUsersById})
		}).bind(this)
		
		Chatter.onChatMessage = (function(msg) {
			var ChatUsersById = Object.assign({}, this.state.ChatUsersById)
			ChatUsersById[msg.From].ChatMessageQueue.push(msg)
			if(msg.From != this.state.ActiveChatUserId) {
				ChatUsersById[msg.From].Unread += 1
				// send read receipt directly here
			}
			this.updateState({ChatUsersById: ChatUsersById})
		}).bind(this)
	}
	componentDidMount() {
		Chatter.ReqConnection()
	}
	onChatBubbleClicked() {
		this.updateState({WindowOpen: true})
	}
	onChatMessagesWindowCloseClicked() {
		this.updateState({ActiveChatUserId: null})
	}
	onChatUsersWindowCloseClicked() {
		this.updateState({WindowOpen: false})
	}
	onChatListItemClicked(c) {
		var ChatUsersById = Object.assign({}, this.state.ChatUsersById)
		ChatUsersById[c.userId].Unread = 0
		this.updateState({ActiveChatUserId: c.User.Id, ChatUsersById: ChatUsersById})

		// send read receipts until the last message
	}
	onMessageSend() {
		var msg = Chatter.SendMessage(this.state.ActiveChatUserId, this.refs.userMessage.value)
		if(msg != null) {
			var ChatUsersById = Object.assign({}, this.state.ChatUsersById)
			ChatUsersById[this.state.ActiveChatUserId].ChatMessageQueue.push(msg)
			this.updateState({ChatsById: ChatsById})
		}
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

		var showChatwindow = this.state.WindowOpen && this.state.User != null && this.state.ActiveChatUserId != null && this.state.ChatsById[this.state.ActiveChatUserId] != null
		var showUsersWindow = this.state.WindowOpen && this.state.User != null
		var showLoginWindow = this.state.WindowOpen && this.state.UserId == null

		var messagesArray = []
		if(showChatwindow) {
			messagesArray = this.state.ChatsById[this.state.ActiveChatUserId].ChatMessageQueue.map(createMessageWidgetObject)
		}

		var chatsArray = []
		if(showUsersWindow) {
			for (const userId in this.state.ChatsById) {
				chatsArray.push(createChatWidgetObject(this.state.ChatsById[userId]))
			}
		}

		return(
		<div class="chat-widget flex-row-container">

			{(!this.state.WindowOpen) ? (<Icon onClick={this.onChatBubbleClicked.bind(this)} iconPath="/icon/chat-bubble.png" height="40px" width="40px" padding="5px"/>) : ""}

			{ showChatwindow ? 
			(<div class="chat-container">
				<div class="chat-header flex-row-container">
					<div>{this.state.ChatsById[this.state.ActiveChatUserId].User.Name}</div>
					<Icon onClick={this.onChatMessagesWindowCloseClicked.bind(this)} iconPath="/icon/close.png" height="20px" width="20px" padding="3px"/>
				</div>
				<div class="chat-content">
					<MessageList className='message-list' toBottomHeight={'100%'} dataSource={messagesArray} />
				</div>
				<Input ref="userMessage" className="chat-input" placeholder="Type here..." multiline={true} rightButtons={<Button className="chat-button" text='Send' onClick={this.onMessageSend.bind(this)}/>}/>
			</div>) : ""}

			{ showUsersWindow ? 
			(<div class="chat-container">
				<div class="chat-header flex-row-container">
					<div class="identifier">{this.state.User.Name}</div>
					<Icon onClick={this.onChatUsersWindowCloseClicked.bind(this)} iconPath="/icon/close.png" height="20px" width="20px" padding="3px"/>
				</div>
				<div class="chat-content">
					<ChatList className='chat-list' dataSource={chatsArray} onClick={this.onChatListItemClicked.bind(this)}/>
				</div>
				<Input ref="userSearch" className="chat-input" placeholder="Search user..." multiline={false} rightButtons={<Button className="chat-button" text='Search' onClick={this.onUserSearch.bind(this)}/>}/>
			</div>) : ""}

			{ showLoginWindow ? 
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

function createMessageWidgetObject(msg) {
	return {
		Id: msg.MessageId,
		position: (msg.From == Chatter.UserId) ? "right" : "left",
		type: 'text',
		text: msg.Message,
		date: msg.SentAt,
		status: null,
	}
}

function createChatWidgetObject(chat) {
	var latestMessage = chat.ChatMessageQueue.Length == 0 ? null : chat.ChatMessageQueue[chat.ChatMessageQueue.Length - 1]
	return {
		Id: chat.User.Id,
		avatar: 'https://ui-avatars.com/api/?rounded=true&size=128&name=' + chat.User.Name,
		alt: user.Name,
		title: user.Name,
		subtitle: latestMessage == null ? "" : latestMessage.Message,
		date: latestMessage == null ? "" : latestMessage.SentAt,
		unread: chat.Unread,
	}
}