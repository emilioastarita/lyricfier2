import SongView from "./SongView.mjs";
import SongEdit from "./SongEdit.mjs";
import Connecting from "./Connecting.mjs";
import VueRouter from "./vue-router.mjs";


const routes = [
    { path: '/', component: Connecting, name: 'connecting' },
    { path: '/view', component: SongView, name: 'view' },
    { path: '/edit', component: SongEdit, name: 'edit' }
];

const router = new VueRouter({
    routes
});


export default {
    router,
    components: {
        SongView,
        SongEdit,
    },
    template: `
            <div :class="{'main-view': true, 'dark': data.settings.theme === 'dark'}">
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
                settings: {
                    theme: 'default',
                }
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
            if (this.$router.currentRoute.name === 'connecting' && this.data.song.title) {
                this.$router.push({ name: `view`});
            }
        },
        edit(song) {
            this.data.editSong = song;
            this.$router.push({ name: `edit`});
        },
        saved(song) {
            if (song) {
                this.data.song.lyric = song.lyric;
            }
            this.$router.push({ name: `view`});
        }
    }
}