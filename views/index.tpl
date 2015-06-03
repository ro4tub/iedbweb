<div class="row">				 
 <ul class="img-list">
	{{range .SearchResult}}
		<li>
		     <a href="/game/{{.Name}}">
		       <img src="{{.Logo}}" class="logo"/>
		       <span class="text-content"><span>{{.Name}}</span></span>
		     </a>
	 	</li>
	{{end}}
 </ul>
</div>