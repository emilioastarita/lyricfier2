export default {
    props: ['data'],
    template: `
    <div class="full-vertical-flex settings-container">
        <form @submit.prevent="submit" method="post" action="#">
            <header >
                <h4>
                Editing {{ song.artist }} - {{ song.title }}
                | <a target="_blank" :href="externalSearchUrl">Search for lyric</a>
                </h4>
                
            </header>
            
            <input type="hidden" name="artist" :value="song.artist" />
            <input type="hidden" name="title" :value="song.title" />
            <textarea  class="lyrics" rows="20" name="lyric" v-model="song.lyric"></textarea>
            <p>
                <button type="submit">Save</button>
                <a href="/" @click.prevent="cancel">Cancel</a>
            </p>
        </form>
    </div>    
`,
    data: function() {
        return {

        }
    },
    computed: {
        song() {
          return this.data.song;
        },
        externalSearchUrl() {
            const {artist, title} = this.song;
            return 'https://duckduckgo.com/?q=' + encodeURIComponent('lyrics "' + artist + '" "' + title + '"');
        },
    },
    methods: {
        async submit() {
            const url = document.location.protocol + '//' + document.location.host + '/save-song';
            const res = await fetch(url, {
                method: 'POST',
                body: JSON.stringify(this.song),
                headers:{
                    'Content-Type': 'application/json'
                }
            });
            const data = await res.json();
            this.$emit('song-saved', data.song);
        },
        cancel() {
            this.$emit('song-saved', null);
        }
    }

}


