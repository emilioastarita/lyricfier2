import SongView from "./SongView.mjs";
import Settings from "./Settings.mjs";
import Songs from "./Songs.mjs";
import SongEdit from "./SongEdit.mjs";
import Connecting from "./Connecting.mjs";
import VueRouter from "./vue-router.mjs";
import {service} from "./Service.mjs";
import {Bus, EDIT_SONG, SAVED_SONG, EDIT_SETTINGS, SAVED_SETTINGS, SONGS_LIST} from "./Events.mjs";


const routes = [
    {path: '/', component: Connecting, name: 'connecting'},
    {path: '/view', component: SongView, name: 'view'},
    {path: '/edit', component: SongEdit, name: 'edit'},
    {path: '/settings', component: Settings, name: 'settings'},
    {path: '/songs', component: Songs, name: 'songs'}
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
                <router-view 
                    :data="data" 
                />
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
                editSettings: null,
                settings: {
                    theme: 'default',
                    fontSize: 11,
                    textAlign: 'left',
                },
                songs: null,
            }
        }
    },
    mounted() {
        this.update(['song', 'settings', 'inSnap']);
        service.connectUpdates(() => {
            this.update();
        });
        Bus.$on(EDIT_SONG, this.editSong);
        Bus.$on(SONGS_LIST, this.songsList);
        Bus.$on(SAVED_SONG, this.savedSong);
        Bus.$on(SAVED_SETTINGS, this.savedSettings);
        Bus.$on(EDIT_SETTINGS, this.editSettings);
    },
    methods: {
        async update(props = ['song']) {
            const data = await service.getLyricfierStatus();
            for (let prop of props) {
                this.data[prop] = data[prop];
            }
            if (this.$router.currentRoute.name === 'connecting' && this.data.song.title) {
                this.$router.push({name: `view`});
            }
        },
        editSong(song) {
            this.data.editSong = song;
            this.$router.push({name: `edit`});
        },
        async songsList() {
            this.data.songs = await service.getSongs();
            this.$router.push({name: `songs`});
        },
        editSettings() {
            this.data.editSettings = {...this.data.settings};
            this.$router.push({name: `settings`});
        },
        savedSong(song) {
            if (song) {
                this.data.song.lyric = song.lyric;
            }
            this.$router.push({name: `view`});
        },
        savedSettings(settings) {
            this.data.editSettings = null;
            if (settings) {
                this.data.settings = settings;
            }
            this.$router.push({name: `view`});
        }
    }
}