<script type="text/javascript" src="static/js/jquery.min.js"></script>
<script src="static/js/bootstrap.min.js"></script>
<script src="static/js/fileinput.min.js" type="text/javascript" charset="utf-8"></script>
<script src="static/js/bootstrap-markdown.js" type="text/javascript" charset="utf-8"></script>
<script src="static/js/markdown.js" type="text/javascript" charset="utf-8"></script>
<script src="static/js/to-markdown.js" type="text/javascript" charset="utf-8"></script>
<script src="static/js/bootstrap-tagsinput.min.js" type="text/javascript" charset="utf-8"></script>
<script type="text/javascript">
	//  {{if not .Context.IsLogin}}
	// 	$('#form-signup-submit').click(function(e){
	// 	      // e.preventDefault();
	//
	// 	      $.post('/signup',
	// 	         $('#form-signup').serialize(),
	// 	         function(data, status, xhr){
	// 	             if(data.Ret == "ok") {
	// 	             	location.reload();
	// 	             } else {
	// 	             	alert(data.Ret)
	// 	             }
	// 	         });
	// 	});
	// {{end}}
	// $('#form-gamecreate-submit').click(function(e){
	//       // e.preventDefault();
	// 	  alert("click")
	//       $.post('/creategame',
	//          $('#form-gamecreate').serialize(),
	//          function(data, status, xhr){
	//              	alert(data.Ret)
	//          });
	// });
	$("#inputGameLogo").fileinput({'showUpload':false, 'previewFileType':'any'});
</script>