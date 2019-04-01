const template = _.template(`
<% if (song.title) { %>
    <div v-if="song" class="full-vertical-flex">
        <header class="header" >
            <% if (song.artUrl) { %>
                <span  class="header__background" style="background-image: url('<%- song.artUrl %>');"></span>
                <img   class="header__album-art" src="<%- song.artUrl %>" alt="<%- song.artist %>">
            <% } else { %>
                <span class="header__background" style="background-image: url(/static/img/icon.png);"></span>
                <img  class="header__album-art" style="padding: 5px; box-sizing: border-box;" src="/static/img/icon.png" alt="<%- song.artist %>" />
            <% } %>
            <h3><%- song.title %></h3>
            <h4><%- song.artist %></h4>
        </header>
        <div  id="lyricBox" class="lyrics"><%- song.lyric %></div>
    
        <div class="credits-source">Source:
            <a rel="noopener" target="_blank" href="#<%- song.source %>" >
                <%- song.source %>
            </a>
        </div>
    </div>    
<% } else { %>
    <div class="full-vertical-flex">
          <div class="not-playing-view">
              <img src="/static/img/waves.svg" alt="Waveform">
              <h3>Looking for a Song on Spotify</h3>
              <p>Connecting...</p>
          </div>
    </div>
<% } %>
`);
let currentSong = null;
let app = null;
function update() {
    fetch('/status')
        .then(
            function (response) {
                if (response.status !== 200) {
                    return;
                }
                response.json().then(function (data) {
                        if (JSON.stringify(currentSong) !== JSON.stringify(data)) {
                            currentSong = data;
                            app.innerHTML = template(data);
                        }
                });
            }
        )
        .catch(function (err) {
            console.error('Fetch Error ', err);
        });
}
function setup() {
    app = document.getElementById('app');
    update();
    setInterval(update, 2000);
}
document.addEventListener('DOMContentLoaded', setup, false);
