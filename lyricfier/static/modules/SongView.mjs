import SongHeader from "./SongHeader.mjs";
import SongFooter from "./SongFooter.mjs";

export default {
    components: {
        SongHeader,
        SongFooter
    },
    props: ['data'],
    template: `
    <div class="full-vertical-flex">
        <SongHeader :song="song" />
        <div  id="lyricBox" class="lyrics" :style="{'font-size': fontSize, 'text-align': textAlign}">{{ song.lyric }}</div>
        <SongFooter :song="song" />
    </div>        
`,
    computed: {
        song() {
            return this.data.song;
        },
        fontSize() {
            return this.data.settings.fontSize + 'px';
        },
        textAlign() {
            return this.data.settings.textAlign;
        }
    }

}
