export default {
    props: ['data'],
    template: `
    <div>
    <div v-if="song.title" class="full-vertical-flex">
        <header class="header" >
            <template v-if="song.artUrl">
                <span  class="header__background" :style="{'background-image': 'url(' +song.artUrl +')'}"></span>
                <img   class="header__album-art" :src="song.artUrl" :alt="song.artist">            
            </template>
            <template v-if="!song.artUrl">            
                <span class="header__background" style="background-image: url(/static/img/icon.png);"></span>
                <img  class="header__album-art" style="padding: 5px; box-sizing: border-box;" src="/static/img/icon.png" :alt="song.artist" />
            </template>
            <h3>{{ song.title }}</h3>
            <h4>{{ song.artist }}</h4>
        </header>
        <div  id="lyricBox" class="lyrics">{{ song.lyric }}</div>
    
        <div class="credits-source">Source:
                {{ song.source }} | <a href="#edit" @click.prevent="edit" >Edit lyric</a>
        </div>
    </div>    
    
    <div v-if="!song.title" class="full-vertical-flex">
              <div class="not-playing-view">
                  <img src="/static/img/waves.svg" alt="Waveform">
                  <h3>Looking for a Song on Spotify</h3>
                  
                   <p>Is spotify running and playing a song?</p>
                    
                    <template v-if="data.inSnap">
<p>
    Looks like you are running the "snap" version of Lyrcifer.<br />
                        Have you allowed Spotify connection?<br />
                        Run in terminal:                    

<pre style="background: #444; color: #15ff1c; padding: 5px; text-align: center">sudo snap connect lyricfier:mpris spotify:spotify-mpris</pre>
</p>                    
                    
                    </template>

                  <p>Look for help at <a href="https://github.com/emilioastarita/lyricfier2">homepage</a></p>  
                  
              </div>
    </div>
    
    </div>
    
`,
    methods: {
        edit() {
            this.$emit('edit', this.song);
        }
    },
    computed: {
        song() {
            return this.data.song;
        }
    }

}
