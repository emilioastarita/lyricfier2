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
        <div  id="lyricBox" class="lyrics">{{ song.lyric }}</div>
        <SongFooter :song="song" />
    </div>        
`,
    computed: {
        song() {
            return this.data.song;
        }
    }

}
