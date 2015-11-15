jQuery(document).ready(function($){

	$('.js-btn').click(function(e){
		e.preventDefault();
		var action = $(this).attr('data-action');
		wsSend({ action: action, source: 'client' });
	});


	var wsUri = 'ws://localhost:8000/ws';
	// var wsUri = 'ws://echo.websocket.org';

	var $output = $('#output');

	var ws = new WebSocket(wsUri);

	ws.onopen = function(evt) {
		wsLog('<div class="alert alert-danger"><strong>CONNECTED</strong></div>');
		wsSend({ action: 'init', source: 'client' });
	};

	ws.onclose = function(evt) {
		wsLog('<div class="alert alert-danger"><strong>DISCONNECTED</strong></div>');
	};

	ws.onmessage = function(evt) {
		wsLog('<div class="alert alert-success"><strong>RESPONSE:</strong> ' + evt.data + '</div>');
		// ws.close();
	};

	ws.onerror = function(evt) {
		wsLog('<div class="alert alert-warning"><strong>ERROR:</strong> ' + evt.data + '</div>');
	};

	function wsSend(payload) {
		var msg = JSON.stringify(payload);
		wsLog('<div class="alert alert-info"><strong>SENT:</strong> ' + msg + '</div>');
		ws.send(msg);
	}

	function wsLog(msg) {
		$output.append($('<div/>').html(msg));
	}


});
