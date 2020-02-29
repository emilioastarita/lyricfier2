import SongView from "./SongView.mjs";
import SongEdit from "./SongEdit.mjs";
import VueRouter from "./vue-router.mjs";


const routes = [
    { path: '/', component: SongView },
    { path: '/edit', component: SongEdit }
]

const router = new VueRouter({
    routes // short for `routes: routes`
});


export default {
    router,
    components: {
        SongView,
        SongEdit,
    },
    template: `
            <div>
                <router-view :data="data" v-on:edit="edit" v-on:song-saved="saved" ></router-view>
            </div>
    `,
    data: function () {
        return {
            data: {
                song: {
                    title: '',
                    artist: '',
                    lyric: '',
                    artUrl: '',
                    source: '',
                },
                inSnap: false,
                editSong: null,
            }
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
            this.data.song = data.song;
            this.data.inSnap = data.inSnap;
        },
        edit(song) {
            this.data.editSong = song;
            this.$router.push({ path: `/edit`});
        },
        saved(song) {
            if (song) {
                this.data.song.lyric = song.lyric;
            }
            this.$router.push({ path: `/`});
        }
    }
}