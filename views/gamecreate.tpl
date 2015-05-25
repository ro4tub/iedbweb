		<form class="form-gamecreate">
		  <div class="form-group">
		    <label for="inputGameName">游戏名字</label>
		    <input type="text" class="form-control" id="inputGameName" placeholder="输入游戏名字">
		  </div>
		  <div class="form-group">
		    <label for="inputGameGenre">游戏类型</label>
			<select class="form-control" id="inputGameGenre" >
			  <option value="1">动作</option>
			  <option value="2">动作冒险</option>
			  <option value="3">冒险解密</option>
			  <option value="4">休闲益智</option>
			  <option value="5">角色扮演</option>
			  <option value="6">竞速</option>
			  <option value="7">模拟</option>
			  <option value="8">体育竞技</option>
			  <option value="9">策略</option>
		      <option value="99">其它</option>
			</select>
		  </div>
		  <div class="form-group">
		    <label for="inputGamePlatform">游戏平台</label>
			<select class="form-control" id="inputGamePlatform">
			  <option value="1">iOS</option>
			  <option value="2">Android</option>
			  <option value="3">Xbox 360</option>
			  <option value="4">Xbox One</option>
			  <option value="5">PS3</option>
			  <option value="6">PS4</option>
			  <option value="7">PC</option>
			  <option value="8">浏览器</option>
			  <option value="99">其它</option>
			</select>
		  </div>
		  <div class="form-group">
		    <label for="inputGameTag">自定义标签</label>
		    <input type="text" class="form-control" id="inputGameTag" data-role="tagsinput" placeholder="输入">
		  </div>
		  <div class="form-group">
		    <label for="inputGameLogo">游戏Logo</label>
		    <input type="file" id="inputGameLogo" class="file">
		  </div>
		  <div class="form-group">
		    <label for="inputGameSimpleDesc">游戏简单描述(140字)</label>
		    <textarea class="form-control" id="inputGameSimpleDesc" rows="3" placeholder="概述">
			</textarea>
		  </div>
		  <div class="form-group">
		    <label for="inputGameDetail">游戏详细描述</label>
		    <textarea class="form-control" id="inputGameSimpleDesc" rows="20" data-provide="markdown" placeholder="markdown格式">
			</textarea>
		  </div>
		  <button id="form-gamecreate-submit" type="submit" class="btn btn-default">递交</button>
		</form>