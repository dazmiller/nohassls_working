<!DOCTYPE html>
  <html>
    <head>
    <title>{{.title}}</title>
      <!--Import Google Icon Font-->
      <link href="http://fonts.googleapis.com/icon?family=Material+Icons" rel="stylesheet">
      <!--Import materialize.css-->
       <link rel="stylesheet" href="/static/css/materialize.min.css">

      <!--Let browser know website is optimized for mobile-->
      <meta name="viewport" content="width=device-width, initial-scale=1.0"/>
    </head>




<body>

    <nav class="nav-wrapper light-blue darken-4">
        <a href="/blog" class="brand-logo">My blog</a>
        <ul id="nav-mobile" class="right side-nav">
            <li>
                <a href="/blog/add">Add an article</a>
            </li>
        </ul>
        <a class="button-collapse" href="#" data-activates="nav-mobile">
            <i class="mdi-navigation-menu"></i>
        </a>
    </nav>

    <br /><!-- Not pretty :( -->

    <div class="container">



        <div class="row">
            {{if .flash.notice}}
            <div id="toast-container">
                <div class="toast">{{.flash.notice}}</div>
            </div>
            {{end}}
            {{.LayoutContent}}
        </div>

        <div class="valign-demo valign-wrapper">
            Powered by Beego (a Go framework)
        </div>

    </div>

    <script src="http://code.jquery.com/jquery-2.1.3.min.js"></script>
    <script src="/static/js/materialize.min.js"></script>
    <script>
        $(document).ready(function(){
            $("input[type='text']")
                .addClass("validate")
                .wrapAll('<div class="input-field"></div>');
            $("textarea")
                .addClass("materialize-textarea validate")
                .wrapAll('<div class="input-field"></div>');
            $(".button-collapse").sideNav({edge: "left"});
        });
    </script>

</body>
</html>