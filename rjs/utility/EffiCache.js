
var Storage = localStorage;
var MaxLife = 60; 	// nothing survives more than 1 hour
var CleanUpInterval = 2; // in minutes

function Init() {
	const interval = setInterval(CleanUp, CleanUpInterval * 60000);
}

function CleanUp() {
	for (var i = 0; i < Storage.length; i++){
		var key = Storage.key(i)
    	var object = JSON.parse(Storage.getItem(key));
		if(
			object.key == null || 
			(!(object.key == key)) ||
			(new Date()) > (new Date((new Date(object.insertedAt)).getTime() + MaxLife * 60000))
		){
			Storage.removeItem(key);
		}
	}
}

function Get(key) {
	var object = JSON.parse(Storage.getItem(key));
	if (object == null) {
		return null;
	}
	if(
		object.key == key && 
		(
			object.expiryMinutes == null ||
			(new Date()) < (new Date((new Date(object.insertedAt)).getTime() + object.expiryMinutes * 60000))
		)
	){
		return object.value;
	}
	Storage.removeItem(key);
	return null
}

function Set(key, value, expiryMinutes) {
	var object = {
		key: key,
		expiryMinutes: expiryMinutes,
		insertedAt: new Date(),
		value: value
	}
	Storage.setItem(key, JSON.stringify(object));
}

Efficache = {
	Init: Init,
	Get: Get,
	Set: Set,
}

export default Efficache;