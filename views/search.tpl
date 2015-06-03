<div class="row">
	{{range .SearchResult}}
	  <div class="col-md-5 item">
		  <div class="itemlogo">
			  <a href="/game/{{.Name}}">
				  <img src="{{.Logo}}" class="logo" alt="{{.Name}}">
			  </a>
		  </div>
		  <div class="itemtitle">
			  <a href="/game/{{.Name}}">{{.Name}}</a>
		  </div>
		  <div class="itemtags">
			  <ul class="list-inline">
				<li><a href="/search?t={{.Genre}}">{{.Genre}}</a></li>
				<li><a href="/search?t={{.Platform}}">{{.Platform}}</a></li>
				{{range .Tags}}
			  		<li><a href="/search?t={{.}}">{{.}}</a></li>	
				{{end}}
			  <ul>
		  </div>
		  <div class="itemtext">
			    <p>{{.SimpleDesc}}</p>
		  </div>
	  </div>
	{{end}}
</div>