$(document).ready(function(){
	var id, button;
	$("button.remove").click(function () {
		if (confirm("确定删除吗?")) {
			var id = $(this).attr("id");
			console.log('book/'+ id + '/delete')
			$.getJSON('book/'+ id + '/delete', function(data) {
                    alert(data);
				if (data) {
					$("tr#" + id).fadeOut("slow", function () {
						$("tr#" + id).remove();
					});
				};
			});
		};
	});
});
