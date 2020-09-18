
const STATES = {
		CONNECTED: "CONNECTED",	// messages sent using connection id in the from field
		LOGGED_IN: "LOGGED_IN",	// messages sent using user id in the from field
		DISCONNECTED: "DISCONNECTED",	// messages can not be sent
	}

var Chatter = {
	CurrentState: STATES.DISCONNECTED,

	ConnectionId: null,
	Connection: null,

	UserId: null,
	UserName: null,
	UserPublicKey: null,

	onOpen: "Socket Connection is now open",
	onConnected: "Chatter connection is now established",
	onLogin: "User has been logged in",
	onChatMessage: function(msg){console.log("Chat Message :", msg)},
	onLogout: "User logged out",
	onAddedToGroup: null,
	onRemovedFromGroup: null,
	onClose: "Socket Connection closed",

	GetConnectionUrl: function() {
		return (window.location.protocol.includes("https") ? "wss://" : "ws://") + window.location.host + "/soc/chatter"
	},

	ReqConnection: function(name = null, publicKey = null) {
		var thiz = Chatter
		if(thiz.CurrentState != STATES.DISCONNECTED) {
			return false
		}

		var queryParams = []
		if(name != null) {queryParams[0] = "name=" + name}
		if(publicKey != null) {queryParams[1] = "publicKey=" + publicKey }

		var URL = [thiz.GetConnectionUrl(), queryParams.join("&")].join("?")

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
			executeOnlyAFunctionIfNotNull(thiz.onDisconnected)
		}

		return true
	},

	// a true means a request to create a new user in was sent successfully
	ReqCreateUser: function(name, publicKey) {
		var thiz = Chatter
		if(thiz.CurrentState != STATES.CONNECTED) {
			return false
		}

		thiz.Connection.send(JSON.stringify({
			From: thiz.ConnectionId,
			To: "server-create-and-login-as-chat-user",
			SentAt: new Date(),
			Message: name + "," + publicKey,
		}))

		return true
	},

	// a true means a request to log you in was sent successfully
	ReqLogin: function(name, publicKey) {
		var thiz = Chatter
		if(thiz.CurrentState != STATES.CONNECTED) {
			return false
		}

		thiz.Connection.send(JSON.stringify({
			From: thiz.ConnectionId,
			To: "server-login-as-chat-user",
			SentAt: new Date(),
			Message: name + "," + publicKey,
		}))

		return true
	},

	// a true means a request to send a mesage was successfull
	// thiz does not ensure delivery
	SendMessage: function(to, textMsg) {
		var thiz = Chatter
		if(thiz.CurrentState != STATES.LOGGED_IN) {
			return false
		}
		if(!(typeof(textMsg) === 'string' || textMsg instanceof String)) {
			return false
		}
		sendMessageInternal(thiz.UserId,to,textMsg,"","","")
		return true
	},

	// a true means a request to log you out was sent successfully
	ReqLogout: function() {
		var thiz = Chatter
		if(thiz.CurrentState != STATES.LOGGED_IN) {
			return false
		}
		sendMessageInternal(thiz.ConnectionId,"server-logout","","","","")
		return true
	},

	ReqGetAllUsers: function() {
		var thiz = Chatter
		if(thiz.CurrentState != STATES.LOGGED_IN) {
			return false
		}
		sendMessageInternal(thiz.UserId,"server-get-all-users","","","","")
		return true
	},

	ReqGetAllMyGroups: function() {
		var thiz = Chatter
		if(thiz.CurrentState != STATES.LOGGED_IN) {
			return false
		}
		sendMessageInternal(thiz.UserId,"server-get-all-my-groups","","","","")
		return true
	},

	ReqGetAllMyActiveConnections: function() {
		var thiz = Chatter
		if(thiz.CurrentState != STATES.LOGGED_IN) {
			return false
		}
		sendMessageInternal(thiz.UserId,"server-get-all-my-active-connections","","","","")
		return true
	},
}

export default Chatter;

function GetRandomString(length) {
   var result           = '';
   var characters       = "+-/<[abcdefghijklmnopqrstuvwxyz](ABCDEFGHIJKLMNOPQRSTUVWXYZ){0123456789}>-_^#";
   var charactersLength = characters.length;
   for ( var i = 0; i < length; i++ ) {
      result += characters.charAt(Math.floor(Math.random() * charactersLength));
   }
   return result;
}

function executeOnlyAFunctionIfNotNull(funcN) {
	if(funcN != null && funcN instanceof Function) { 
		funcN()
		return true
	} else {
		console.log("Not a function : ", funcN)
		return false
	}
}

function sendMessageInternal(From,To,Message,Messages,MessageId,ContextId) {
	Chatter.Connection.send(JSON.stringify({
		From: From,To: To,SentAt: new Date(),
		Message: Message,Messages: Messages,
		MessageId: MessageId,ContextId: ContextId
	}))
}

/* restricted access */
/* Code below is meant to allow the chatter client to follow the standard state transitions and protocol for chattering */
/* Access to below source is restricted to only those people who are familiar with chatter protocol */

function ChatterConnectionHandler(chatter, msgEvent) {

	var msg = JSON.parse(msgEvent.data)

	if(msg.From.startsWith("server")) {
		console.log("Server event", msg)
		switch(msg.From){
			case "server-chatterer-created" : {
				if(isChatConnectionId(msg.To)) {
					chatter.ConnectionId = msg.To
					chatter.CurrentState = STATES.CONNECTED
					executeOnlyAFunctionIfNotNull(chatter.onConnected)
				} else if(isChatUserId(msg.To)) {
					// user created
				} else if (isChatGroupId(msg.To)) {
					// group created
				}
				break;
			}
			case "server-create-and-login-as-chat-user" :
			case "server-login-as-chat-user" : {
				if(!isErrorEvent(msg.Message)) {
					chatter.UserId = msg.Message
					chatter.CurrentState = STATES.LOGGED_IN
					executeOnlyAFunctionIfNotNull(chatter.onLogin)
				}
				break;
			}
			case "server-logout" : {
				if(!isErrorEvent(msg.Message)) {
					chatter.UserId = null
					chatter.CurrentState = STATES.CONNECTED
					executeOnlyAFunctionIfNotNull(chatter.onLogout)
				}
				break;
			}
		}
	} else {
		chatter.onChatMessage(msg)
	}
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

function isErrorEvent(textMsg) {
	return textMsg.startsWith("ERROR")
}