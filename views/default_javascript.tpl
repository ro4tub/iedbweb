<script type="text/javascript" src="/static/js/jquery.min.js"></script>
<script src="/static/js/bootstrap.min.js"></script>
<script type="text/javascript">
	 {{if not .Context.IsLogin}}
		// $('#form-signup-submit').click(function(e){
		//       // e.preventDefault();
		//
		//       $.post('/signup',
		//          $('#form-signup').serialize(),
		//          function(data, status, xhr){
		//              if(data.Ret == "ok") {
		//              	// location.reload();
		//              } else {
		//              	alert(data.Ret)
		//              }
		//          });
		// });
	{{end}}
</script>