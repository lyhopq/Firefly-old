$(document).ready(function(){
    /*
  $("button.collect").mouseover(function(event) {
    $(this).removeClass("btn-success").addClass("btn-danger");
    $(this).children("i").removeClass().addClass("icon-white icon-remove");
    $(this).children("span").text("取消收藏");
  });
  
  $("button.collect").mouseout(function(event) {
    $(this).removeClass("btn-danger").addClass("btn-success");
    $(this).children("i").removeClass().addClass("icon-white icon-ok");
    $(this).children("span").text("已经收藏");
  });
  */
  
  $("#unbook").mouseover(function(event) {
    $(this).removeClass("btn-success").addClass("btn-danger");
    $(this).children("i").removeClass().addClass("icon-white icon-remove");
    $(this).children("a").text("取消预借");
  });
  
  $("#unbook").mouseout(function(event) {
    $(this).removeClass("btn-danger").addClass("btn-success");
    $(this).children("i").removeClass().addClass("icon-white icon-ok");
    $(this).children("a").text("已经预借");
  });

	$("button#collect").unbind("click").click(function () {
        var button = $(this);
		var id = $(".book").attr("id");
        var feed = $(".feedback>#collect")
        if(button.hasClass("uncollect")) {
		    $.getJSON(id + '/collect', function(data) {
                if(data) {
                    button.removeClass("uncollect").addClass("collect");
                    button.children("i").removeClass().addClass("icon-remove");
                    button.children("span").text("取消收藏");
                    feed.text(parseInt(feed.text())+1);
                } else {
                    alert("请先登录，谢谢！");
                };
		    	});
        } else {
		    $.getJSON(id + '/uncollect', function(data) {
                if(data) {
                    button.removeClass("collect").addClass("uncollect");
                    button.children("i").removeClass().addClass("icon-plus");
                    button.children("span").text("加入收藏");
                    feed.text(parseInt(feed.text())-1);
                } else {
                    alert("请先登录，谢谢！");
                };
		    	});

            };
	});
});
