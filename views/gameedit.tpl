		<form class="form-gamecreate" method="post" action="/editgame" enctype="multipart/form-data">
		  <div class="form-group">
		    <label for="inputGameName">游戏名字</label>
		    <input type="text" class="form-control" id="inputGameName" name="gamename" value="{{.Game.Name}}">
		  </div>
		  <div class="form-group">
		    <label for="inputGameGenre">游戏类型</label>
			<select class="form-control" id="inputGameGenre" name="gamegenre">
			  <option value="动作" {{if eq .Game.Genre "动作"}}selected{{end}}>动作</option>
			  <option value="动作冒险" {{if eq .Game.Genre "动作冒险"}}selected{{end}}>动作冒险</option>
			  <option value="冒险解密" {{if eq .Game.Genre "冒险解密"}}selected{{end}}>冒险解密</option>
			  <option value="休闲益智" {{if eq .Game.Genre "休闲益智"}}selected{{end}}>休闲益智</option>
			  <option value="角色扮演" {{if eq .Game.Genre "角色扮演"}}selected{{end}}>角色扮演</option>
			  <option value="竞速" {{if eq .Game.Genre "竞速"}}selected{{end}}>竞速</option>
			  <option value="模拟" {{if eq .Game.Genre "模拟"}}selected{{end}}>模拟</option>
			  <option value="体育竞技" {{if eq .Game.Genre "体育竞技"}}selected{{end}}>体育竞技</option>
			  <option value="策略" {{if eq .Game.Genre "策略"}}selected{{end}}>策略</option>
		      <option value="其它"{{if eq .Game.Genre "其它"}}selected{{end}}>其它</option>
			</select>
		  </div>
		  <div class="form-group">
		    <label for="inputGamePlatform">游戏平台</label>
			<select class="form-control" id="inputGamePlatform" name="gameplatform">
			  <option value="iOS" {{if eq .Game.Platform "iOS"}}selected{{end}}>iOS</option>
			  <option value="Android" {{if eq .Game.Platform "Android"}}selected{{end}}>Android</option>
			  <option value="Xbox 360" {{if eq .Game.Platform "Xbox 360"}}selected{{end}}>Xbox 360</option>
			  <option value="Xbox One" {{if eq .Game.Platform "Xbox One"}}selected{{end}}>Xbox One</option>
			  <option value="PS3" {{if eq .Game.Platform "PS3"}}selected{{end}}>PS3</option>
			  <option value="PS4" {{if eq .Game.Platform "PS4"}}selected{{end}}>PS4</option>
			  <option value="PC" {{if eq .Game.Platform "PC"}}selected{{end}}>PC</option>
			  <option value="浏览器" {{if eq .Game.Platform "浏览器"}}selected{{end}}>浏览器</option>
			  <option value="其它" {{if eq .Game.Platform "其它"}}selected{{end}}>其它</option>
			</select>
		  </div>
		  <div class="form-group">
		    <label for="inputGameTag">自定义标签</label>
		    <input type="text" class="form-control" id="inputGameTag" name="gametags" data-role="tagsinput" value="{{.Game.Tags}}">
		  </div>
		  <div class="form-group">
		    <label for="inputGameLogo">游戏Logo</label>
		    <input type="file" id="inputGameLogo" name="gamelogo" class="file">
		  </div>
		  <div class="form-group">
		    <label for="inputGameSimpleDesc">游戏简单描述(140字)</label>
		    <textarea class="form-control" id="inputGameSimpleDesc" name="gamesimpledesc" rows="3">{{.Game.SimpleDesc}}
			</textarea>
		  </div>
		  <div class="form-group">
		    <label for="inputGameDetail">游戏详细描述</label>
		    <textarea class="form-control" id="inputGameDetail" name="gamedetail" rows="20" data-provide="markdown">{{.Game.Detail}}
			</textarea>
		  </div>
		  <button id="form-gamecreate-submit" type="submit" class="btn btn-default">递交</button>
		</form>