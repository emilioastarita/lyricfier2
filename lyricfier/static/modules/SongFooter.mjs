import {Bus, EDIT_SONG, EDIT_SETTINGS} from "./Events.mjs";


export default {
    props: ['song'],
    template: `
        <footer class="footer" >
            <div class="credits-source">Source:
                {{ song.source }} | 
                <a href="#edit" @click.prevent="editSong" >Edit lyric</a> |
                <a href="#edit" @click.prevent="editSettings" >Settings</a> |
            </div>
        </footer>      
`,
    methods: {
        editSong() {
            Bus.$emit(EDIT_SONG, this.song);
        },
        editSettings() {
            Bus.$emit(EDIT_SETTINGS);
        }
    }
}
