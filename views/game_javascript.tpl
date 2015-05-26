<script src="/static/js/jquery.min.js" type="text/javascript" ></script>
<script src="/static/js/bootstrap.min.js" type="text/javascript" ></script>
<script src="/static/js/bootstrap-markdown.js" type="text/javascript" charset="utf-8"></script>
<script src="/static/js/markdown.js" type="text/javascript" charset="utf-8"></script>
<script src="/static/js/to-markdown.js" type="text/javascript" charset="utf-8"></script>
<script type="text/javascript">
	// var text = "Type **Markdown** here.";
	$("#gamedetail").html(markdown.toHTML({{.Game.Detail}}))
</script>