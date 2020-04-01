export default {
    props: ['data'],
    template: `
    <div>
    
    <div class="full-vertical-flex">
              <div class="not-playing-view">
                  <img src="/static/img/waves.svg" alt="Waveform">
                  <h3>Looking for a Song on Spotify</h3>
                  
                   <p>Is spotify running and playing a song?</p>
                    
                    <template v-if="data.inSnap">
<p>
    Looks like you are running the "snap" version of Lyricfier.<br />
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
