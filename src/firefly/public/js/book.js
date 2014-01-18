$(document).ready(function(){
    var user =$('#user_menu'); 
    var isLogin;
    if(user.length > 0) {
        isLogin = true;
    } else {
        isLogin = false;
    }
	var id = $(".book").attr("id");

	$("button#collect").unbind("click").click(function () {
        if(!isLogin) {
            alert("请先登录，谢谢！");
            return;
        }

        var button = $(this);
        var num = $(".collected>.num")
        if(button.hasClass("uncollect")) {
		    $.getJSON(id + '/collect', function(data) {
                if(data) {
                    button.removeClass("uncollect").addClass("collect");
                    button.children("i").removeClass().addClass("icon-remove");
                    button.children("span").text("取消收藏");
                    num.text(parseInt(num.text())+1);
                } 
                });
        } else {
		    $.getJSON(id + '/uncollect', function(data) {
                if(data) {
                    button.removeClass("collect").removeClass("btn-danger").addClass("uncollect");
                    button.children("i").removeClass().addClass("icon-plus");
                    button.children("span").text("加入收藏");
                    num.text(parseInt(num.text())-1);
                } 
		    	});

            };
	}).mouseover(function(event) {
        if($(this).hasClass("collect")) {
        $(this).removeClass("btn-success").addClass("btn-danger");
        $(this).children("i").removeClass().addClass("icon-white icon-remove");
        $(this).children("span").text("取消收藏");
        };
  }).mouseout(function(event) {
        if($(this).hasClass("collect")) {
        $(this).removeClass("btn-danger").addClass("btn-success");
        $(this).children("i").removeClass().addClass("icon-white icon-ok");
        $(this).children("span").text("已经收藏");
        };
  });


	$("button#borrow").unbind("click").click(function () {
        if(!isLogin) {
            alert("请先登录，谢谢！");
            return;
        }

        var button = $(this);
        var num = $(".booked>.num");
        if(button.hasClass("unbooked")) {
		    $.getJSON(id + '/booking', function(data) {
                if(data) {
                    button.removeClass("unbooked").addClass("booked");
                    button.children("i").removeClass().addClass("icon-remove");
                    button.children("span").text("取消预借");
                    num.text(parseInt(num.text())+1);
                } else {
                    alert("该书已经被借完了！");
                };
		    	});
        } else if(button.hasClass("booked")) {
		    $.getJSON(id + '/unbooking', function(data) {
                if(data) {
                    button.removeClass("booked").removeClass("btn-danger").addClass("unbooked");
                    button.children("i").removeClass().addClass("icon-plus");
                    button.children("span").text("加入预借");
                    num.text(parseInt(num.text())-1);
		    	};
            })
            };
	}).mouseover(function(event) {
        if($(this).hasClass("booked")) {
        $(this).removeClass("btn-success").addClass("btn-danger");
        $(this).children("i").removeClass().addClass("icon-white icon-remove");
        $(this).children("span").text("取消预借");
        };
  }).mouseout(function(event) {
        if($(this).hasClass("booked")) {
        $(this).removeClass("btn-danger").addClass("btn-success");
        $(this).children("i").removeClass().addClass("icon-white icon-ok");
        $(this).children("span").text("已经预借");
        };
  });

});
