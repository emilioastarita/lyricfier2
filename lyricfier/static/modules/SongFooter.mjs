import {Bus, EDIT_SONG, EDIT_SETTINGS, SONGS_LIST} from "./Events.mjs";


export default {
    props: ['song'],
    template: `
        <footer class="footer" >
            <div class="credits-source">Source:
                {{ song.source }} | 
                <a href="#edit" @click.prevent="editSong" >Edit lyric</a> |
                <a href="#settings" @click.prevent="editSettings" >Settings</a> |
                <a href="#songs" @click.prevent="songsList" >Songs</a> |
                <a href="https://github.com/emilioastarita/lyricfier2" target="_blank" rel="noopener" >About</a> |
            </div>
        </footer>      
`,
    methods: {
        editSong() {
            Bus.$emit(EDIT_SONG, this.song);
        },
        editSettings() {
            Bus.$emit(EDIT_SETTINGS);
        },
        songsList() {
            Bus.$emit(SONGS_LIST);
        }
    }
}
