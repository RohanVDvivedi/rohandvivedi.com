
window.Chatter = {
	STATES: {
		CONNECTED: "CONNECTED",	// messages sent using connection id in the from field
		LOGGED_IN: "LOGGED_IN",	// messages sent using user id in the from field
		DISCONNECTED: "DISCONNECTED",	// messages can not be sent
	},

	CurrentState: "DISCONNECTED",

	Connection: null,
	onOpen: "Socket Connection is now open",
	onClose: "Socket Connection is now closed",

	ConnectionId: null,
	onConnected: "Chatter connection is now established",

	User: null,
	onLogin: "User has been logged in",
	onChatMessage: function(msg){console.log("Chat Message :", msg)},
	onLogout: "User logged out",

	onChangeUsersList: function(userList){console.log("Users list :", userList)},
	onUserNotification: function(user){console.log("User :", user)},
	onSearchResultsReady: function(results){console.log("SearchResults :", results)},

	getConnectionUrl() {
		return (window.location.protocol.includes("https") ? "wss://" : "ws://") + window.location.host + "/soc/chatter"
	},

	ReqConnection(name = null, publicKey = null) {
		if(this.CurrentState != this.STATES.DISCONNECTED) {
			return false
		}

		var queryParams = []
		if(name != null) {queryParams.push("name=" + name)}
		if(publicKey != null) {queryParams.push("publicKey=" + publicKey)}

		var URL = [this.getConnectionUrl(), queryParams.join("&")].join("?")

		this.Connection = new WebSocket(URL);

		this.Connection.onopen = function() {
			this.CurrentState = this.STATES.CONNECTED
			executeOnlyAFunctionIfNotNull(this.onOpen)
		}.bind(this)

		this.Connection.onmessage = function(msgEvent) {
			this.chatterConnectionHandler(msgEvent)
		}.bind(this)

		this.Connection.onerror = function(error) {
			console.log("Error: " + error.message)
		}.bind(this)

		this.Connection.onclose = function() {
			this.CurrentState = this.STATES.DISCONNECTED
			executeOnlyAFunctionIfNotNull(this.onClose)
		}.bind(this)

		return true
	},

	sendMessageRaw(From,To,Message = null,Messages = null,ContextId = null) {
		var msg = {
			From: From,To: To,SentAt: new Date(),
			Message: Message,Messages: Messages,
			MessageId: getRandomString(64),ContextId: ContextId
		}
		this.Connection.send(JSON.stringify(msg))
		return msg
	},

	ReqCreateUser(name, publicKey) {
		if(this.CurrentState != this.STATES.CONNECTED) {
			return null
		}
		return this.sendMessageRaw(this.ConnectionId,"server-create-and-login-as-chat-user","",[name,publicKey])
	},
	ReqLogin(name, publicKey) {
		if(this.CurrentState != this.STATES.CONNECTED) {
			return null
		}
		return this.sendMessageRaw(this.ConnectionId,"server-login-as-chat-user","",[name,publicKey])
	},
	ReqLogout() {
		if(this.CurrentState != this.STATES.LOGGED_IN) {
			return null
		}
		return this.sendMessageRaw(this.ConnectionId,"server-logout-from-chat-user")
	},

	SendMessage(to, textMsg, contextId = null) {
		if(this.CurrentState != this.STATES.LOGGED_IN) {
			return null
		}
		if(!(typeof(textMsg) === 'string' || textMsg instanceof String)) {
			return null
		}
		return this.sendMessageRaw(this.User.Id,to,textMsg,null,contextId)
	},

	RequestGenericQuery(serverReceiverName, queryParam = "") {
		if(this.CurrentState != this.STATES.LOGGED_IN) {
			return null
		}
		return this.sendMessageRaw(this.User.Id,serverReceiverName, queryParam)
	},

	ReqGetAllUsers() {
		return this.RequestGenericQuery("server-get-all-users")
	},

	ReqGetAllGroups() {
		return this.RequestGenericQuery("server-get-all-groups")
	},

	ReqGetAllOnlineUsers() {
		return this.RequestGenericQuery("server-get-all-online-users")
	},

	ReqGetAllMyGroups() {
		return this.RequestGenericQuery("server-get-all-my-groups")
	},

	ReqGetAllMyActiveConnections() {
		return this.RequestGenericQuery("server-get-all-my-active-connections")
	},

	ReqSearch(query) {
		return this.RequestGenericQuery("server-search-chatter-box",query)
	},

	/* restricted access */
	/* Code below is meant to allow the chatter client to follow the standard state transitions and protocol for chattering */
	/* Access to below source is restricted to only those people who are familiar with chatter protocol */

	chatterConnectionHandler(msgEvent) {

		var msg = JSON.parse(msgEvent.data)
		msg.SentAt = Date.parse(msg.SentAt)

		if(isServerEvent(msg)) {
			console.log("Server event", msg)
			switch(msg.From){
				case "server-chatters-creator" : {
					if(isChatConnectionId(msg.To)) {
						this.ConnectionId = msg.To
						this.CurrentState = this.STATES.CONNECTED
						executeOnlyAFunctionIfNotNull(this.onConnected)
					} else if (isChatUserId(msg.To)) {
						// user created
					} else if (isChatGroupId(msg.To)) {
						// group created
					}
					break;
				}
				case "server-create-and-login-as-chat-user" :
				case "server-login-as-chat-user" : {
					if(!isErrorEvent(msg)) {
						this.User = GetUserFromString(msg.Message)
						this.CurrentState = this.STATES.LOGGED_IN
						executeOnlyAFunctionIfNotNull(this.onLogin)
					}
					break;
				}
				case "server-logout-from-chat-user" : {
					if(!isErrorEvent(msg)) {
						this.User = null
						this.CurrentState = this.STATES.CONNECTED
						executeOnlyAFunctionIfNotNull(this.onLogout)
					}
					break;
				}
				case "server-get-all-users" : {
					if(!isErrorEvent(msg)) {
						var results = msg.Messages.map(function(userStr){
							return GetUserFromString(userStr)
						}).filter(function(user){
							return user != null && user.Id != this.User.Id
						}.bind(this))
						executeOnlyAFunctionIfNotNull(this.onChangeUsersList, results)
					}
					break;
				}
				case "server-event-update" : {
					if(!isErrorEvent(msg)) {
						var user = GetUserFromString(msg.Message)
						if(user != null && user.Id != this.User.Id) {
							executeOnlyAFunctionIfNotNull(this.onUserNotification, user)
						}
					}
					break;
				}
				case "server-search-chatter-box" : {
					if(!isErrorEvent(msg)) {
						var results = msg.Messages.map(function(userStr){
							return GetUserFromString(userStr)
						}).filter(function(user){
							return user != null
						})
						executeOnlyAFunctionIfNotNull(this.onSearchResultsReady, results)
					}
					break;
				}
			}
		} else {
			this.onChatMessage(msg)
		}
	},
}

function getRandomString(length) {
	var result           = '';
	var characters       = "+-/<[abcdefghijklmnopqrstuvwxyz](ABCDEFGHIJKLMNOPQRSTUVWXYZ){0123456789}>-_^#";
	var charactersLength = characters.length;
	for ( var i = 0; i < length; i++) {
		result += characters.charAt(Math.floor(Math.random() * charactersLength));
	}
	return result;
}

function executeOnlyAFunctionIfNotNull(funcN, ...args) {
	if(funcN != null && funcN instanceof Function) { 
		funcN.apply(null, args);
	} else {
		console.log("Not a function : ", funcN)
	}
}



function GetUserFromString(userStr) {
	var userData = userStr.split(',')
	if(userData.length == 3 && isChatUserId(userData[0])) {
		return {
			Id: userData[0],
			Name: userData[1],
			ConnectionCount: parseInt(userData[2], 10),
		}
	}
	return null
}

function isChatConnectionId(Id) {
	return Id.startsWith("CHAT_CONN-")
}

function isChatUserId(Id) {
	return Id.startsWith("CHAT_USER-")
}

function isChatGroupId(Id) {
	return Id.startsWith("CHAT_GRUP-")
}

function isServerEvent(Msg) {
	return Msg.From.startsWith("server")
}

function isErrorEvent(Msg) {
	return Msg.Message != null && Msg.Message.startsWith("ERROR")
}