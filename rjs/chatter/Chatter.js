
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
		if(this.CurrentState != STATES.DISCONNECTED) {
			return false
		}

		var queryParams = []
		if(name != null) {queryParams[0] = "name=" + name}
		if(publicKey != null) {queryParams[1] = "publicKey=" + publicKey }

		var URL = [this.GetConnectionUrl(), queryParams.join("&")].join("?")

		this.Connection = new WebSocket(URL);

		this.Connection.onopen = function() {
			this.CurrentState = STATES.CONNECTED
			executeOnlyAFunctionIfNotNull(this.onOpen)
		}

		this.Connection.onmessage = function(msgEvent) {
			ChatterConnectionHandler(this, msgEvent)
		}

		this.Connection.onerror = function(error) {
			console.log("Error: " + error.message)
		}

		this.Connection.onclose = function() {
			this.CurrentState = STATES.DISCONNECTED
			executeOnlyAFunctionIfNotNull(this.onDisconnected)
		}

		return true
	},

	// a true means a request to create a new user in was sent successfully
	ReqCreateUser: function(name, publicKey) {
		if(this.CurrentState != STATES.CONNECTED) {
			return false
		}

		this.Connection.send(JSON.stringify({
			From: this.ConnectionId,
			To: "server-create-chat-user",
			SentAt: new Date(),
			Message: name + "," + publicKey,
		}))

		return true
	},

	// a true means a request to log you in was sent successfully
	ReqLogin: function(name, publicKey) {
		if(this.CurrentState != STATES.CONNECTED) {
			return false
		}

		this.Connection.send(JSON.stringify({
			From: this.ConnectionId,
			To: "server-login-as-chat-user",
			SentAt: new Date(),
			Message: name + "," + publicKey,
		}))

		return true
	},

	// a true means a request to send a mesage was successfull
	// this does not ensure delivery
	SendMessage(to, textMsg) {
		if(this.CurrentState == STATES.DISCONNECTED) {
			return false
		}

		if(!(typeof(s) === 'string' || s instanceof String)) {
			return false
		}

		var From = "" 
		if(this.CurrentState == STATES.CONNECTED) {
			From = this.ConnectionId;
		} else if(this.CurrentState == STATES.LOGGED_IN) {
			From = this.UserId;
		}

		this.Connection.send(JSON.stringify({
			From: From,
			To: to,
			SentAt: new Date(),
			Message: textMsg,
		}))

		return true
	},

	// a true means a request to log you out was sent successfully
	ReqLogout: function() {
		if(this.CurrentState != STATES.LOGGED_IN) {
			return false
		}

		this.Connection.send(JSON.stringify({
			From: this.ConnectionId,
			To: "server-logout",
			SentAt: new Date(),
			Message: "",
		}))

		return true
	},
}

export default Chatter;

function executeOnlyAFunctionIfNotNull(funcN) {
	if(funcN != null && funcN instanceof Function) { 
		funcN()
		return true
	} else {
		console.log(funcN)
		return false
	}
}

/* restricted access */
/* Code below is meant to allow the chatter client to follow the standard state transitions and protocol for chattering */
/* Access to below source is restricted to only those people who are familiar with chatter protocol */

function ChatterConnectionHandler(chatter ,msgEvent) {

	var msg = JSON.parse(msgEvent.data)

	if(msg.From.startsWith("server")) {
		console.log("Server event", msg)
	} else {
		chatter.onChatMessage(msg)
	}
}