<script type="text/javascript" src="/static/js/jquery.min.js"></script>
<script src="/static/js/bootstrap.min.js"></script>
<script src="/static/js/fileinput.min.js" type="text/javascript" charset="utf-8"></script>
<script src="/static/js/bootstrap-markdown.js" type="text/javascript" charset="utf-8"></script>
<script src="/static/js/markdown.js" type="text/javascript" charset="utf-8"></script>
<script src="/static/js/to-markdown.js" type="text/javascript" charset="utf-8"></script>
<script src="/static/js/bootstrap-tagsinput.min.js" type="text/javascript" charset="utf-8"></script>
<script type="text/javascript">
	 {{if not .Context.IsLogin}}
	 $(function(){
			$('.form-signin').submit(function(){
			 $.post($(this).attr('action'), $(this).serialize(), function(json) {
				 if(json.Ret == "ok") {
					$('.signin').modal('hide')
				 	window.location = '/'
				 } else {
				 	$('.signintips').html(json.Ret)
				 }
			 }, 'json');
			 return false;
			});
		   $('.form-signup').submit(function(){
		     $.post($(this).attr('action'), $(this).serialize(), function(json) {
				 if(json.Ret == "ok") {
					$('.signup').modal('hide')
				 	window.location = '/'
				 } else {
				 	$('.signuptips').html(json.Ret)
				 }
		     }, 'json');
		     return false;
		   });
	});
	{{end}}
	$("#inputGameLogo").fileinput({'showUpload':false, 'previewFileType':'any'});
</script>