      {{range $error := .errors}}
                {{$error.Key}}
                {{$error.Message}}
            {{end}}
            <div class="row">
                <form id="articles" method="POST">
                    {{.Form | renderform}}
                    <button type="submit" class="waves-effect waves-light btn">Send</button>
                </form>
            </div>