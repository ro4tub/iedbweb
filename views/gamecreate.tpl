		<form class="form-gamecreate" method="post" action="/creategame" enctype="multipart/form-data">
		  <div class="form-group">
		    <label for="inputGameName">游戏名字</label>
		    <input type="text" class="form-control" id="inputGameName" name="gamename" placeholder="输入游戏名字">
		  </div>
		  <div class="form-group">
		    <label for="inputGameGenre">游戏类型</label>
			<select class="form-control" id="inputGameGenre" name="gamegenre">
			  <option value="动作">动作</option>
			  <option value="动作冒险">动作冒险</option>
			  <option value="冒险解密">冒险解密</option>
			  <option value="休闲益智">休闲益智</option>
			  <option value="角色扮演">角色扮演</option>
			  <option value="竞速">竞速</option>
			  <option value="模拟">模拟</option>
			  <option value="体育竞技">体育竞技</option>
			  <option value="策略">策略</option>
		      <option value="其它">其它</option>
			</select>
		  </div>
		  <div class="form-group">
		    <label for="inputGamePlatform">游戏平台</label>
			<select class="form-control" id="inputGamePlatform" name="gameplatform">
			  <option value="iOS">iOS</option>
			  <option value="Adnroid">Android</option>
			  <option value="Xbox 360">Xbox 360</option>
			  <option value="Xbox One">Xbox One</option>
			  <option value="PS3">PS3</option>
			  <option value="PS4">PS4</option>
			  <option value="PC">PC</option>
			  <option value="浏览器">浏览器</option>
			  <option value="其它">其它</option>
			</select>
		  </div>
		  <div class="form-group">
		    <label for="inputGameTag">自定义标签</label>
		    <input type="text" class="form-control" id="inputGameTag" name="gametags" data-role="tagsinput" placeholder="输入">
		  </div>
		  <div class="form-group">
		    <label for="inputGameLogo">游戏Logo</label>
		    <input type="file" id="inputGameLogo" name="gamelogo" class="file">
		  </div>
		  <div class="form-group">
		    <label for="inputGameSimpleDesc">游戏简单描述(140字)</label>
		    <textarea class="form-control" id="inputGameSimpleDesc" name="gamesimpledesc" rows="3" placeholder="概述"></textarea>
		  </div>
		  <div class="form-group">
		    <label for="inputGameDetail">游戏详细描述</label>
		    <textarea class="form-control" id="inputGameDetail" name="gamedetail" rows="20" data-provide="markdown" placeholder="markdown格式"></textarea>
		  </div>
		  <button id="form-gamecreate-submit" type="submit" class="btn btn-default">递交</button>
		</form>