	  <div class="game-title">
	  	<h1 >{{.Game.Name}}<span><a class="btn btn-default" type="submit" href="/game/{{.Game.Name}}/edit">编辑</a></span>
		</h1>
	</p>
	  </div>
	  <div class="row">
		  <div class="col-md-3 game-images item">
			  <img src="{{.Game.Logo}}" class="game" alt="{{.Game.Name}}">
			  <div>
			  		<b class="variablecolor">平台:</b> <a href="#/search/p={{.Game.Platform}}">{{.Game.Platform}}</a><br/>
			  		<b class="variablecolor">类型:</b> <a href="#/search/g={{.Game.Genre}}">{{.Game.Genre}}</a><br/>
			  		<b class="variablecolor">标签:</b> {{.Game.Tags}}<br/>
					<b class="variablecolor">简介:</b> {{.Game.SimpleDesc}}<br/>
			  </div>
		  </div>
		  <div id="gamedetail" class="col-md-6" >
			 
		  </div>
		  <div class="editstat col-md-offset-9">
			  <p>编辑信息</p>
			  <p>浏览次数: 9999次</p>
  			  <p>编辑次数: 3次</p>
  			  <p>更新时间: 2015年4月26日</p>
			  <p>创建者: huangyaoshi</p>
		  </div>
	  </div>