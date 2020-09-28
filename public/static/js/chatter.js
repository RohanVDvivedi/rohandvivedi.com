
const STATES = {
		CONNECTED: "CONNECTED",	// messages sent using connection id in the from field
		LOGGED_IN: "LOGGED_IN",	// messages sent using user id in the from field
		DISCONNECTED: "DISCONNECTED",	// messages can not be sent
	}

function getConnectionUrl() {
	return (window.location.protocol.includes("https") ? "wss://" : "ws://") + window.location.host + "/soc/chatter"
}

window.Chatter = {
	CurrentState: STATES.DISCONNECTED,

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

	ReqConnection: function(name = null, publicKey = null) {
		var thiz = Chatter
		if(thiz.CurrentState != STATES.DISCONNECTED) {
			return false
		}

		var queryParams = []
		if(name != null) {queryParams.push("name=" + name)}
		if(publicKey != null) {queryParams.push("publicKey=" + publicKey)}

		var URL = [getConnectionUrl(), queryParams.join("&")].join("?")

		thiz.Connection = new WebSocket(URL);

		thiz.Connection.onopen = function() {
			thiz.CurrentState = STATES.CONNECTED
			executeOnlyAFunctionIfNotNull(thiz.onOpen)
		}

		thiz.Connection.onmessage = function(msgEvent) {
			ChatterConnectionHandler(thiz, msgEvent)
		}

		thiz.Connection.onerror = function(error) {
			console.log("Error: " + error.message)
		}

		thiz.Connection.onclose = function() {
			thiz.CurrentState = STATES.DISCONNECTED
			executeOnlyAFunctionIfNotNull(thiz.onClose)
		}

		return true
	},

	ReqCreateUser: function(name, publicKey) {
		var thiz = Chatter
		if(thiz.CurrentState != STATES.CONNECTED) {
			return null
		}
		return sendMessageInternal(thiz.ConnectionId,"server-create-and-login-as-chat-user","",[name,publicKey])
	},
	ReqLogin: function(name, publicKey) {
		var thiz = Chatter
		if(thiz.CurrentState != STATES.CONNECTED) {
			return null
		}
		return sendMessageInternal(thiz.ConnectionId,"server-login-as-chat-user","",[name,publicKey])
	},
	ReqLogout: function() {
		var thiz = Chatter
		if(thiz.CurrentState != STATES.LOGGED_IN) {
			return null
		}
		return sendMessageInternal(thiz.ConnectionId,"server-logout-from-chat-user")
	},

	// a true means a request to send a mesage was successfull
	// thiz does not ensure delivery
	SendMessage: function(to, textMsg) {
		var thiz = Chatter
		if(thiz.CurrentState != STATES.LOGGED_IN) {
			return null
		}
		if(!(typeof(textMsg) === 'string' || textMsg instanceof String)) {
			return null
		}
		return sendMessageInternal(thiz.User.Id,to,textMsg)
	},

	RequestGenericQuery: function(serverReceiverName, queryParam = "") {
		var thiz = Chatter
		if(thiz.CurrentState != STATES.LOGGED_IN) {
			return null
		}
		return sendMessageInternal(thiz.User.Id,serverReceiverName, queryParam)
	},

	ReqGetAllUsers: function() {
		var thiz = Chatter
		return thiz.RequestGenericQuery("server-get-all-users")
	},

	ReqGetAllGroups: function() {
		var thiz = Chatter
		return thiz.RequestGenericQuery("server-get-all-groups")
	},

	ReqGetAllOnlineUsers: function() {
		var thiz = Chatter
		return thiz.RequestGenericQuery("server-get-all-online-users")
	},

	ReqGetAllMyGroups: function() {
		var thiz = Chatter
		return thiz.RequestGenericQuery("server-get-all-my-groups")
	},

	ReqGetAllMyActiveConnections: function() {
		var thiz = Chatter
		return thiz.RequestGenericQuery("server-get-all-my-active-connections")
	},

	ReqSearch: function(query) {
		var thiz = Chatter
		return thiz.RequestGenericQuery("server-search-chatter-box",query)
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

function sendMessageInternal(From,To,Message = null,Messages = null,ContextId = null) {
	var msg = {
		From: From,To: To,SentAt: new Date(),
		Message: Message,Messages: Messages,
		MessageId: getRandomString(64),ContextId: ContextId
	}
	Chatter.Connection.send(JSON.stringify(msg))
	return msg
}

/* restricted access */
/* Code below is meant to allow the chatter client to follow the standard state transitions and protocol for chattering */
/* Access to below source is restricted to only those people who are familiar with chatter protocol */

function ChatterConnectionHandler(chatter, msgEvent) {

	var msg = JSON.parse(msgEvent.data)
	msg.SentAt = Date.parse(msg.SentAt)

	if(isServerEvent(msg)) {
		console.log("Server event", msg)
		switch(msg.From){
			case "server-chatters-creator" : {
				if(isChatConnectionId(msg.To)) {
					chatter.ConnectionId = msg.To
					chatter.CurrentState = STATES.CONNECTED
					executeOnlyAFunctionIfNotNull(chatter.onConnected)
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
					chatter.User = GetUserFromString(msg.Message)
					chatter.CurrentState = STATES.LOGGED_IN
					executeOnlyAFunctionIfNotNull(chatter.onLogin)
				}
				break;
			}
			case "server-logout-from-chat-user" : {
				if(!isErrorEvent(msg)) {
					chatter.User = null
					chatter.CurrentState = STATES.CONNECTED
					executeOnlyAFunctionIfNotNull(chatter.onLogout)
				}
				break;
			}
			case "server-get-all-users" : {
				if(!isErrorEvent(msg)) {
					chatter.AllUsersById = {}
					var results = msg.Messages.map(function(userStr){
						return GetUserFromString(userStr)
					}).filter(function(user){
						return user != null && user.Id != Chatter.User.Id
					})
					executeOnlyAFunctionIfNotNull(chatter.onChangeUsersList, results)
				}
				break;
			}
			case "server-event-update" : {
				if(!isErrorEvent(msg)) {
					var user = GetUserFromString(msg.Message)
					if(user != null && user.Id != Chatter.User.Id) {
						executeOnlyAFunctionIfNotNull(chatter.onUserNotification, user)
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
					executeOnlyAFunctionIfNotNull(chatter.onSearchResultsReady, results)
				}
				break;
			}
		}
	} else {
		chatter.onChatMessage(msg)
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