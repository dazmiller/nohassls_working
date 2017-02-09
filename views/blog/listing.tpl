  
<div class="row">
<a href="/blog/add">Add an article</a>
</div>
  
   {{range $article := .articles}}
             <article class="row">
                <div class="col s12 m12">
                    <div class="card brown blue lighten-5">
                        <div class="card-content">
                            <p>Created {{date $article.Created "Y-m-d H:i:s"}}</p>
                            <span class="card-title black-text">{{$article.Title}}</span>
                            <p>{{substr $article.Content 0 400 | str2html}}</p>
                        </div>
                        <div class="card-action">
                            <a class="blue-text" href="/blog/article/{{$article.Id}}">View</a>
                            <a class="blue-text" href="/blog/edit/{{$article.Id}}">Edit</a>
                            <a class="blue-text" href="/blog/delete/{{$article.Id}}">Delete</a>
                        </div>
                    </div>
                </div>
            </article>
            {{end}}