import SongHeader from "./SongHeader.mjs";

export default {
    components: {
        SongHeader
    },
    props: ['data'],
    template: `
    <div class="full-vertical-flex">
        <SongHeader :song="song" v-if="song" />
        <form @submit.prevent="submit" method="post" action="#" class="settings-container">
            <h4>
                 <a href="/" @click.prevent="cancel">Back</a> |
                 Editing lyric |
                 <a target="_blank" :href="externalSearchUrl" v-if="song">Search for lyric in the web</a>
            </h4>
            <div v-if="song">
                <input type="hidden" name="artist" :value="song.artist" />
                <input type="hidden" name="title" :value="song.title" />
                <textarea  class="lyrics" rows="20" name="lyric" v-model="song.lyric"></textarea>
                <p>
                    <button type="submit">Save</button>
                    <a href="/" @click.prevent="cancel">Cancel</a>
                </p>            
            </div>
        </form>
    </div>    
`,
    data: function() {
        return {

        }
    },
    beforeRouteEnter (to, from, next) {
        next(vm => {
            if (!vm.song) {
                vm.$router.push({'name': 'view'})
            }
        })
    },
    computed: {
        song() {
          return this.data.editSong;
        },
        externalSearchUrl() {
            const {artist, title} = this.song;
            return 'https://duckduckgo.com/?q=\\' + encodeURIComponent('lyrics "' + artist + '" "' + title + '"');
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


