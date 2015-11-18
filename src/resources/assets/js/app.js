jQuery(document).ready(function($){

	var wsUri = 'ws://localhost:' + window.location.port + '/ws';
	var ws = new WebSocket(wsUri);
	var $output = $('#output');

	$('.js-btn').click(function(e){
		e.preventDefault();
		var action = $(this).attr('data-action');
		wsSend({ action: action, source: 'client' });
	});

	ws.onopen = function(evt) {
		wsLog('<div class="alert alert-danger"><strong>CONNECTED</strong></div>');
		wsSend({ code: 200, action: 'init', data: 'hello, world' });
	};

	ws.onclose = function(evt) {
		wsLog('<div class="alert alert-danger"><strong>DISCONNECTED</strong></div>');
	};

	ws.onmessage = function(evt) {
		wsLog('<div class="alert alert-success"><strong>RESPONSE:</strong> ' + evt.data + '</div>');
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
