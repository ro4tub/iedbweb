		<form class="form-gamecreate" method="post" action="/editgame" enctype="multipart/form-data">
		  <div class="form-group">
		    <label for="inputGameName">游戏名字</label>
		    <input type="text" class="form-control" id="inputGameName" name="gamename" value="{{.Game.Name}}">
		  </div>
		  <div class="form-group">
		    <label for="inputGameGenre">游戏类型</label>
			<select class="form-control" id="inputGameGenre" name="gamegenre">
			  <option value="1" {{if eq .Game.Genre "1"}}selected{{end}}>动作</option>
			  <option value="2" {{if eq .Game.Genre "2"}}selected{{end}}>动作冒险</option>
			  <option value="3" {{if eq .Game.Genre "3"}}selected{{end}}>冒险解密</option>
			  <option value="4" {{if eq .Game.Genre "4"}}selected{{end}}>休闲益智</option>
			  <option value="5" {{if eq .Game.Genre "5"}}selected{{end}}>角色扮演</option>
			  <option value="6" {{if eq .Game.Genre "6"}}selected{{end}}>竞速</option>
			  <option value="7" {{if eq .Game.Genre "7"}}selected{{end}}>模拟</option>
			  <option value="8" {{if eq .Game.Genre "8"}}selected{{end}}>体育竞技</option>
			  <option value="9" {{if eq .Game.Genre "9"}}selected{{end}}>策略</option>
		      <option value="99"{{if eq .Game.Genre "99"}}selected{{end}}>其它</option>
			</select>
		  </div>
		  <div class="form-group">
		    <label for="inputGamePlatform">游戏平台</label>
			<select class="form-control" id="inputGamePlatform" name="gameplatform">
			  <option value="1" {{if eq .Game.Platform "1"}}selected{{end}}>iOS</option>
			  <option value="2" {{if eq .Game.Platform "2"}}selected{{end}}>Android</option>
			  <option value="3" {{if eq .Game.Platform "3"}}selected{{end}}>Xbox 360</option>
			  <option value="4" {{if eq .Game.Platform "4"}}selected{{end}}>Xbox One</option>
			  <option value="5" {{if eq .Game.Platform "5"}}selected{{end}}>PS3</option>
			  <option value="6" {{if eq .Game.Platform "6"}}selected{{end}}>PS4</option>
			  <option value="7" {{if eq .Game.Platform "7"}}selected{{end}}>PC</option>
			  <option value="8" {{if eq .Game.Platform "8"}}selected{{end}}>浏览器</option>
			  <option value="99" {{if eq .Game.Platform "99"}}selected{{end}}>其它</option>
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