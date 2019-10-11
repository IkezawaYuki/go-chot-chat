$(function () {
    $(".messages").animate({ scrollTop: $(document).height() }, "fast");

    $("#profile-img").click(function() {
        $("#status-options").toggleClass("active");
    });

    $(".expand-button").click(function() {
        $("#profile").toggleClass("expanded");
        $("#contacts").toggleClass("expanded");
    });

    $("#status-options ul li").click(function() {
        $("#profile-img").removeClass();
        $("#status-online").removeClass("active");
        $("#status-away").removeClass("active");
        $("#status-busy").removeClass("active");
        $("#status-offline").removeClass("active");
        $(this).addClass("active");

        if($("#status-online").hasClass("active")) {
            $("#profile-img").addClass("online");
        } else if ($("#status-away").hasClass("active")) {
            $("#profile-img").addClass("away");
        } else if ($("#status-busy").hasClass("active")) {
            $("#profile-img").addClass("busy");
        } else if ($("#status-offline").hasClass("active")) {
            $("#profile-img").addClass("offline");
        } else {
            $("#profile-img").removeClass();
        };

        $("#status-options").removeClass("active");
    });

    // function newMessage() {
    //     message = $(".message-input input").val();
    //     if($.trim(message) == '') {
    //         return false;
    //     }
    //     $('<li class="sent"><img src="http://emilcarlsson.se/assets/mikeross.png" alt="" /><p>' + message + '</p></li>').appendTo($('.messages ul'));
    //     $('.message-input input').val(null);
    //     $('.contact.active .preview').html('<span>You: </span>' + message);
    //     $(".messages").animate({ scrollTop: $(document).height() }, "fast");
    // };
    //
    // $('.submit').submit(function() {
    //     console.log("submit")
    //     newMessage();
    //     return false
    // });
    //
    // $(window).on('keydown', function(e) {
    //     if (e.which == 13) {
    //         console.log("13")
    //         newMessage();
    //         return false;
    //     }
    // });

    var socket = null;
    var msgBox = $(".message-input input");
    var messages = $(".messages");
    $(window).on('keydown', function(e) {
        if (e.which == 13) {
            if(!socket){
                alert("エラー：WebSocket接続が行われていません。");
                return false;
            }
            socket.send($(".message-input input").val())
            console.log("send!")
            $(".message-input input").val("");
            return false;
        }
    });

    if(!window["WebSocket"]){
        alert("エラー：WebSocketに対応していないブラウザです。")
    }else{
        socket = new WebSocket("ws://localhost:8080/room");
        socket.onclose = function () {
            alert("接続が終了しました。")
        }
        socket.onmessage = function (e) {
            messages.append(e.data)
            console.log(e.data)
            $('<li class="sent"><img src="http://emilcarlsson.se/assets/mikeross.png" alt="" /><p>' + message + '</p></li>').appendTo($('.messages ul'));
            $('.message-input input').val(null);
            $('.contact.active .preview').html('<span>You: </span>' + message);
            $(".messages").animate({ scrollTop: $(document).height() }, "fast");
        }
    }

});