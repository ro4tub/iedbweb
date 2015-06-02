<!DOCTYPE html>
<html>
<head>
    <title>{{.PageContext.TitleName}}</title>
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
	{{.Css}}
</head>
<body>
    <div class="container">
		<!-- toolbar -->
		<nav class="navbar navbar-default navbar-fixed-top">
		  <div class="container-fluid">
		    <!-- Brand and toggle get grouped for better mobile display -->
		    <div class="navbar-header">
		      <button type="button" class="navbar-toggle collapsed" data-toggle="collapse" data-target="#bs-example-navbar-collapse-1">
		        <span class="sr-only">Toggle navigation</span>
		        <span class="icon-bar"></span>
		        <span class="icon-bar"></span>
		        <span class="icon-bar"></span>
		      </button>
		      <a class="navbar-brand" href="#">IEDB</a>
		    </div>

		    <!-- Collect the nav links, forms, and other content for toggling -->
		    <div class="collapse navbar-collapse" id="bs-example-navbar-collapse-1">
		      <form class="navbar-form navbar-left" role="search" method="get" action="/search">
		        <div class="form-group">
		          <input type="text" class="form-control" name="t" placeholder="搜索，比如天天魔斗士">
		        </div>
		        <button type="submit" class="btn btn-default">搜索</button>
		      </form>
		      <ul class="nav navbar-nav navbar-right">
				  {{if .Context.IsLogin}}
	  			  <li><a href="/creategame">新建游戏</a></li>
		  	      <li class="dropdown">
		  	          <a href="#" class="dropdown-toggle" data-toggle="dropdown" role="button" aria-expanded="false">{{.Context.NickName}} <span class="caret"></span></a>
		  	          <ul class="dropdown-menu" role="menu">
		  	            <li><a href="#">个人信息</a></li>
		  	            <li><a href="#">待审核</a></li>
		  	            <li><a href="/signout">退出</a></li>
		  	          </ul>
		  	       </li>
				   {{else}}
		  			  <li><a data-toggle="modal" data-target=".signin">登录</a></li>
		  			  <li><a data-toggle="modal" data-target=".signup">注册</a></li>
				   {{end}}
		      </ul>
		    </div><!-- /.navbar-collapse -->
		  </div><!-- /.container-fluid -->
		</nav>
		
		
	 	{{if not .IsLogin}}
		<!-- signin -->
		<div class="modal fade signin" tabindex="-1" role="dialog" aria-labelledby="myLargeModalLabel" aria-hidden="true">
		  <div class="modal-dialog">
		    <div class="modal-content">
			  <button type="button" class="close" data-dismiss="modal" aria-label="Close"><span aria-hidden="true">&times;</span></button>
	          <form class="form-signin" method="post" action="/signin">
	            <h2 class="form-signin-heading">输入帐号</h2>
	            <label for="inputEmail" class="sr-only">电子邮件</label>
	            <input type="email" id="inputEmail" name="email"  class="form-control" placeholder="电子邮件" required autofocus>
	            <label for="inputPassword" class="sr-only">密码</label>
	            <input type="password" id="inputPassword" name="password" class="form-control" placeholder="密码" required>
	            <div class="checkbox">
	              <label>
	                <input type="checkbox" value="remember-me"> 记住我
	              </label>
	            </div>
				<div class="label label-warning text-left last signintips"></div>
	            <button class="btn btn-lg btn-primary btn-block signbutton" type="submit">登录</button>
	          </form>
		    </div>
		  </div>
		</div>
		
		<!-- signup -->
		<div class="modal fade signup" tabindex="-1" role="dialog" aria-labelledby="myLargeModalLabel" aria-hidden="true">
		  <div class="modal-dialog">
		    <div class="modal-content">
			  <button type="button" class="close" data-dismiss="modal" aria-label="Close"><span aria-hidden="true">&times;</span></button>
			  <form class="form-signup" method="post" action="/signup">
				  <h2 class="form-signup-heading">注册帐号</h2>
				  <label for="inputEmail" class="sr-only">电子邮箱</label>
				  <input type="email" id="inputEmail" name="email" class="form-control" placeholder="电子邮箱" required autofocus>
				  <label for="inputNickName" class="sr-only">昵称</label>
				  <input type="text" id="inputNickName" name="nickname" class="form-control" placeholder="昵称" required>
				  <label for="inputInviteCode" class="sr-only">邀请码</label>
				  <input type="text" id="inputInviteCode" name="invitecode" class="form-control" placeholder="邀请码" required>
				  <label for="inputPassword" class="sr-only">密码</label>
				  <input type="password" id="inputPassword" name="password" class="form-control" placeholder="密码" required>
				  <label for="inputPassword2" class="sr-only">重复密码</label>
				  <input type="password" id="inputPassword2" class="form-control" placeholder="重复密码" required>
				  <div class="label label-warning text-left last signuptips"></div>
				  <button id="form-signup-submit" class="btn btn-lg btn-primary btn-block signbutton" type="submit">注册</button>
			  </form>
		    </div>
		  </div>
		</div>
		{{end}}
		
		{{.LayoutContent}}
		
		<div class="footer">
				<ul class="list-inline text-right">
				  <li><p class="text-muted">2015 all rights reserved</p></li>
				  <li><a class="btn btn-link" href="#" role="button">关于我们</a></li>
				  <li><a class="btn btn-link" href="#" role="button">联系我们</a></li>
				  <li><a class="btn btn-link" href="#" role="button">免责声明</a></li>
				  <li><a class="btn btn-link" href="#" role="button">开发者</a></li>
				</ul>
		</div>
    </div>
	{{.JavaScript}}
</body>
</html>