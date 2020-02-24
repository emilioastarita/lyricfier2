import SongView from "./SongView.mjs";
import SongEdit from "./SongEdit.mjs";

export default {
    components: {
        SongView,
        SongEdit,
    },
    template: `
            <div>
                <SongView v-if="currentView === 'SongView'" :song="song" v-on:edit="edit" ></SongView>
                <SongEdit v-if="currentView === 'SongEdit'" :song="editSong" v-on:song-saved="saved"  ></SongEdit>
            </div>
    
    `,
    data: function () {
        return {
            currentView: 'SongView',
            song: {
                title: '',
                artist: '',
                lyric: '',
                artUrl: '',
                source: '',
            },
            editSong: null,
        }
    },
    mounted() {
        this.update();
        const conn = new WebSocket("ws://" + document.location.host + "/ws");
        conn.onclose = evt => {
            console.log('Connection error', evt)
        };
        conn.onmessage = () => {
            this.update();
        };
    },
    methods: {
        async update() {
            const response = await fetch('/status');
            if (response.status !== 200) {
                return;
            }
            const data = await response.json();
            this.song = data.song;
        },
        edit(song) {
            this.editSong = song;
            this.currentView = 'SongEdit';
        },
        saved(song) {
            if (song) {
                this.song.lyric = song.lyric;
            }
            this.currentView = 'SongView'
        }
    }
}