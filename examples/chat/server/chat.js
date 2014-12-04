var backoff_attempts = 0;
function backoff_ok() {
    backoff_attempts = 0;
}
function backoff_compute() {
    var SECOND = 1000;
    var backoff_initial = 0.5 * SECOND;
    var backoff_max = 60 * SECOND;
    var backoff_multiplier = 1.5;
    var backoff_randomization = 0.5;

    var retry = backoff_initial * Math.pow(backoff_multiplier, backoff_attempts);
    var limited = Math.min(retry, backoff_max);

    // randomness is backoff_randomization wide centered around 1.0
    var rand = Math.random() * backoff_randomization;
    rand = rand + 1 - backoff_randomization/2;
    var delay = limited * rand;
    console.log({retry: retry, limited: limited, rand: rand, delay: delay});

    backoff_attempts += 1;
    var pretty = delay/1000;
    console.log("Retrying websocket in " + pretty.toFixed(2) + "s...");
    return delay;
}

var ws;
function connect() {
    var div = document.getElementById("msg");
    var line = document.createElement("p");
    line.appendChild(document.createTextNode("Connecting..."));
    div.appendChild(line);

    if (ws != null) {
	ws.close();
	ws = null;
    }
    var scheme = "ws";
    if (window.location.protocol == "https") {
	scheme = "wss";
    }
    ws = new WebSocket(scheme + "://" + window.location.host + "/sock");
    // Initializing these event handlers is not racy as long as we do
    // it before we let the execution here return to the event loop.
    // https://www.w3.org/Bugs/Public/show_bug.cgi?id=12510
    ws.onopen = function () {
	backoff_ok();
	var line = document.createElement("p");
	line.appendChild(document.createTextNode("Opened."));
	div.appendChild(line);
    };
    ws.onmessage = function (e) {
	// TODO can throw SyntaxError
	var rpc = JSON.parse(e.data);
	console.log("rpc in:", rpc);

	if (rpc.fn === undefined
	    || rpc.fn == "") {
	    // responses; we're currently ignoring errors
	    return;
	}

	// kludge dispatch for now
	if (rpc.fn != "Chat.Message") {
	    rpc.error = {rpc: "No such function."};
	    delete rpc.fn;
	    delete rpc.args;
	    delete rpc.result;
	    ws.send(rpc)
	    return;
	}

	// TODO catch exceptions

	var line = document.createElement("p");

	var time = document.createElement("span");
	time.setAttribute("class", "time");
	time.appendChild(document.createTextNode(rpc.args.time));
	line.appendChild(time);

	var from = document.createElement("span");
	from.setAttribute("class", "from");
	from.appendChild(document.createTextNode(rpc.args.from));
	line.appendChild(from);

	var text = document.createElement("span");
	text.setAttribute("class", "message");
	text.appendChild(document.createTextNode(rpc.args.message));
	line.appendChild(text);

	div.appendChild(line);

	delete rpc.fn;
	delete rpc.args;
	rpc.result = {};
	delete rpc.error;
	console.log("rpc out:", rpc);
	ws.send(JSON.stringify(rpc));
    };
    ws.onclose = function (e) {
	var line = document.createElement("p");
	line.appendChild(document.createTextNode("Closed."));
	div.appendChild(line);

	// If this websocket is still the main one, reconnect;
	// otherwise, let it die. This allows easy fiddling from js
	// console with just ``connect()``. Ordering of the status
	// messages gets weird, though.
	if (this === ws) {
	    var delay = backoff_compute();
	    setTimeout(connect, delay);
	}
    };
}
function init() {
    document.msgform.message.focus();
    connect();
};

function send() {
    var rpc = {
	// TODO
	id: "0",
	fn: "Chat.Message",
	args: {
	    from: document.msgform.name.value,
	    message: document.msgform.message.value
	}
    };
    console.log("rpc out:", rpc);
    ws.send(JSON.stringify(rpc));
    document.msgform.message.value = "";
    return false;
};
