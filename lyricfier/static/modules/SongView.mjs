import SongHeader from "./SongHeader.mjs";

export default {
    components: {
        SongHeader
    },
    props: ['data'],
    template: `
    <div class="full-vertical-flex">
        <SongHeader :song="song" />
        <div  id="lyricBox" class="lyrics">{{ song.lyric }}</div>
        <div class="credits-source">Source:
                {{ song.source }} | 
                <a href="#edit" @click.prevent="edit" >Edit lyric</a>
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
