Search Result:
<ul>

		{{range .SearchResult}}
			<li>
				name: {{.Name}} <br>
				genre: {{.Genre}} <br>
				platform: {{.Platform}} <br>
				tags: {{.Tags}} <br>
				simple desc: {{.SimpleDesc}} <br>
				 <img src="{{.Logo}}" class="game" alt="{{.Name}}">
		 	</li>
		{{end}}
</ul>