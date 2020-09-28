import React from "react";

import { MessageList, ChatList, Input, Button } from 'react-chat-elements'
import 'react-chat-elements/dist/main.css';

import Icon from '../utility/Icon'

var Chatter = window.Chatter

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
			ChatUsersById : null,	//{ User.Id => {User, Unread, ChatMessageQueue, ChatMessagesById}
		}

		Chatter.onLogin = (function() {this.updateState({User: Chatter.User}); Chatter.ReqGetAllUsers()}).bind(this)
		Chatter.onLogout = (function() {this.updateState({User: null,})}).bind(this)
		Chatter.onClose = (function() {this.updateState({WindowOpen:false,User: null})}).bind(this)

		Chatter.onChangeUsersList = (function(userList) {
			var ChatUsersById = {}
			var oldChatUsersById = Object.assign({}, this.state.ChatUsersById)
			userList.forEach(function(user){
				ChatUsersById[user.Id] = {User: user, Unread: 0, ChatMessageQueue: [], ChatMessagesById: {}}
				if(oldChatUsersById[user.Id] != null) {
					ChatUsersById[user.Id].Unread = oldChatUsersById[user.Id].Unread
					ChatUsersById[user.Id].ChatMessageQueue = oldChatUsersById[user.Id].ChatMessageQueue
					ChatUsersById[user.Id].ChatMessagesById = oldChatUsersById[user.Id].ChatMessagesById
				}
			})
			this.updateState({ChatUsersById: ChatUsersById})
		}).bind(this)
		Chatter.onUserNotification = (function(user) {
			var ChatUsersById = Object.assign({}, this.state.ChatUsersById)
			if(ChatUsersById[user.Id] == null) {
				ChatUsersById[user.Id] = {User: user, Unread: 0, ChatMessageQueue: [], ChatMessagesById: {}}
			} else {
				ChatUsersById[user.Id].User = user;
			}
			this.updateState({ChatUsersById: ChatUsersById})
		}).bind(this)
		
		Chatter.onChatMessage = (function(msg) {
			var ChatUsersById = Object.assign({}, this.state.ChatUsersById)
			if(msg.ContextId == null) { // normal chat message
				ChatUsersById[msg.From].ChatMessagesById[msg.MessageId] = msg
				ChatUsersById[msg.From].ChatMessageQueue.push(msg.MessageId)
				// send recv receipt here
				Chatter.SendMessage(msg.From, msg.MessageId + "-RECV", msg.MessageId)
				if(msg.From != this.state.ActiveChatUserId) {
					ChatUsersById[msg.From].Unread += 1
				} else {
					// send read receipt directly here
					Chatter.SendMessage(msg.From, msg.MessageId + "-READ", msg.MessageId)
				}
			} else { // is a sent, received or read receipt
				if(msg.Message == msg.ContextId + "-SENT") {
					ChatUsersById[msg.From].ChatMessagesById[msg.ContextId].Status = "sent"
				} else if(msg.Message == msg.ContextId + "-RECV") {
					ChatUsersById[msg.From].ChatMessagesById[msg.ContextId].Status = "received"
				} else if(msg.Message == msg.ContextId + "-READ") {
					ChatUsersById[msg.From].ChatMessagesById[msg.ContextId].Status = "read"
				} else { // this message is a reply message

				}
			}
			this.updateState({ChatUsersById: ChatUsersById})
		}).bind(this)
	}
	componentDidMount() {
		if(window.screen.width < 600) {
			return
		}
		Chatter.ReqConnection()
	}
	onChatBubbleClicked() {
		this.updateState({WindowOpen: true})
	}
	onChatMessagesWindowCloseClicked() {
		this.updateState({ActiveChatUserId: null})
	}
	onChatWindowCloseClicked() {
		this.updateState({WindowOpen: false})
	}
	onChatListItemClicked(c) {
		var ChatUsersById = Object.assign({}, this.state.ChatUsersById)

		// send read receipts until the last message
		for (var i = 0; i < ChatUsersById[c.Id].Unread; i++) {
			var msgId = ChatUsersById[c.Id].ChatMessageQueue[ChatUsersById[c.Id].ChatMessageQueue.length - 1 - i];
			Chatter.SendMessage(c.Id, msgId + "-READ", msgId)
		}
		ChatUsersById[c.Id].Unread = 0
		this.updateState({ActiveChatUserId: c.Id, ChatUsersById: ChatUsersById})
	}
	onMessageSend() {
		var msg = Chatter.SendMessage(this.state.ActiveChatUserId, this.refs.userMessage.input.value)
		msg.Status = "waiting"
		if(msg != null) {
			var ChatUsersById = Object.assign({}, this.state.ChatUsersById)
			ChatUsersById[this.state.ActiveChatUserId].ChatMessagesById[msg.MessageId] = msg
			ChatUsersById[this.state.ActiveChatUserId].ChatMessageQueue.push(msg.MessageId)
			this.updateState({ChatUsersById: ChatUsersById})
		}
		console.log(msg)
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

		var showChatwindow = this.state.WindowOpen && this.state.User != null && this.state.ActiveChatUserId != null && this.state.ChatUsersById[this.state.ActiveChatUserId] != null
		var showUsersWindow = this.state.WindowOpen && this.state.User != null
		var showLoginWindow = this.state.WindowOpen && this.state.User == null

		var messagesArray = []
		if(showChatwindow) {
			console.log("show 1")
			messagesArray = this.state.ChatUsersById[this.state.ActiveChatUserId].ChatMessageQueue.map(function(MessageId) {
				return createMessageWidgetObject(this.state.ChatUsersById[this.state.ActiveChatUserId].ChatMessagesById[MessageId])
			}.bind(this))
			console.log(messagesArray)
		}

		var chatsArray = []
		if(showUsersWindow) {
			console.log("show 2", this.state.ChatUsersById)
			for (const userId in this.state.ChatUsersById) {
				console.log("print build", userId)
				chatsArray.push(createChatWidgetObject(this.state.ChatUsersById[userId]))
			}
			console.log(chatsArray)
		}

		return(
		<div class="chat-widget flex-row-container">

			{(!this.state.WindowOpen) ? (<Icon onClick={this.onChatBubbleClicked.bind(this)} iconPath="/icon/chat-bubble.png" height="40px" width="40px" padding="5px"/>) : ""}

			{ showChatwindow ? 
			(<div class="chat-container">
				<div class="chat-header flex-row-container">
					<div>{this.state.ChatUsersById[this.state.ActiveChatUserId].User.Name}</div>
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
					<Button className="chat-button" text='Sign out' onClick={this.onUserSignoutClicked.bind(this)}/>
					<Icon onClick={this.onChatWindowCloseClicked.bind(this)} iconPath="/icon/close.png" height="20px" width="20px" padding="3px"/>
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
		position: (msg.From == Chatter.User.Id) ? "right" : "left",
		type: 'text',
		text: msg.Message,
		date: msg.SentAt,
		status: msg.Status,
		//replyButton: true,
	}
}

function createChatWidgetObject(chat) {
	var latestMessage = chat.ChatMessageQueue.length == 0 ? null : chat.ChatMessagesById[chat.ChatMessageQueue[chat.ChatMessageQueue.length - 1]]
	return {
		Id: chat.User.Id,
		avatar: 'https://ui-avatars.com/api/?rounded=true&size=128&bold=true&name=' + chat.User.Name,
		alt: chat.User.Name,
		title: chat.User.Name,
		subtitle: latestMessage == null ? "" : latestMessage.Message,
		date: latestMessage == null ? "" : latestMessage.SentAt,
		unread: chat.Unread,
		statusColor: chat.User.ConnectionCount > 0 ? "#4CAF50" : "#f1f1f1",
	}
}