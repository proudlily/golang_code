$(document).ready(function () {
    var mainwrap = $("#mainwrap");
    var us = null;

    function banBackSpace(event) {
        var codenum = event.keyCode;
        switch (codenum) {
            case 87:
                us.css("top", "-=10px");
                send(JSON.stringify({type: "move", data: "{x: '"+us.css("left")+"', y: '"+us.css("top")+"'}"}));
                break;
            case 83:
                us.css("top", "+=10px");
                send(JSON.stringify({type: "move", data: "{x: '"+us.css("left")+"', y: '"+us.css("top")+"'}"}));
                break;
            case 65:
                us.css("left", "-=10px");
                send(JSON.stringify({type: "move", data: "{x: '"+us.css("left")+"', y: '"+us.css("top")+"'}"}));
                break;
            case 68:
                us.css("left", "+=10px");
                send(JSON.stringify({type: "move", data: "{x: '"+us.css("left")+"', y: '"+us.css("top")+"'}"}));
                break;
        }
    }

    document.onkeydown = banBackSpace;

    var socket = new WebSocket('ws://' + window.location.host + '/App/Socket')

    socket.onerror = function (e) {
        console.log("竟然出错了")
    };

    socket.onclose = function (e) {
        console.log("断线了")
    };

    socket.onmessage = function (event) {
        console.log("服务器来消息了")
    }

    var send = function (count) {
        socket.send(count);
    }

    socket.onmessage = function (event) {
        data = JSON.parse(event.data)
        switch (data.Type) {
            case "add":
                if (document.getElementById("u_" + data.Uid) == null) {
                mainwrap.append('<div class="tank" id="u_' + data.Uid + '">' + data.Uid + '</div>');
                }
                if (uid == data.Uid) {
                    console.log("添加了自己")
                    us = $("#u_" + data.Uid);
                }
                break;
            case "init":
                if (document.getElementById("u_" + data.Uid) == null) {
                    mainwrap.append('<div class="tank" id="u_' + data.Uid + '">' + data.Uid + '</div>');
                }
                var user = $("#u_" + data.Uid);
                var position = eval("(" + data.Data + ")");
                user.css("top", position.y);
                user.css("left", position.x );
                break;
            case "move":
                console.log(data.Data)
                var position = eval("(" + data.Data + ")");
                $("#u_" + data.Uid).css({
                    "top":position.y,
                    "left": position.x
                })
/*
                $("#u_" + data.Uid).css("top", position.y);
                $("#u_" + data.Uid).css("left", position.x);
 */
                break;
        }
    }

});
