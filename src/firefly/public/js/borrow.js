$(document).ready(function(){
    var total = $(".total")
	$("button.remove").click(function () {
		if (confirm("确定删除吗?")) {
			var id = $(this).attr("id");
			console.log('book/'+ id + '/delete')
			$.getJSON('/book/'+ id + '/delete', function(data) {
				if (data) {
					$("tr#" + id).fadeOut("slow", function () {
						$("tr#" + id).remove();
					});
                    total.text(parseInt(total.text())-1);
				};
			});
		};
	});

	$("button.preret").click(function () {
		if (confirm("确定要还书吗?")) {
			var id = $(this).attr("id");
			console.log('book/'+ id + '/return')
			$.getJSON('/book/'+ id + '/return', function(data) {
				if (data) {
					$("tr#" + id).fadeOut("slow", function () {
						$("tr#" + id).remove();
					});
                    total.text(parseInt(total.text())-1);
				};
			});
		};
	});
});
