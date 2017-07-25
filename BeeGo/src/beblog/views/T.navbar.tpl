{{define "navbar"}}
<a class="navbar-brand" href="/">My Blog</a>
<div>
        <u1 class="nav navbar-nav">
            <li {{if .IsHome}}class="active"{{end}}><a href="/">首页</a></li>
            <li {{if .IsCategory}}class="active"{{end}}><a href="/category.html">分类</a></li>
            <li {{if .IsTopic}}class="active"{{end}}><a href="/topic.html">文章</a></li>
        </u1>
    </div>

    <div class="pull-right">
        <u1 class="nav navbar-nav">
           {{if .IsLogin}}
           <li><a href="/login.html?exit=true">退出</a></li>
           {{else}}
           <li><a href="/login.html">管理员登陆</a></li>
           {{end}}
        </u1>
    </div>
{{end}}