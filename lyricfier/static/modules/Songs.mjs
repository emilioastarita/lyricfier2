import {Bus, SAVED_SETTINGS, SONGS_LIST} from "./Events.mjs";


export default {
    props: ['data'],
    template: `
    <div class="full-vertical-flex">
        <div class="songs-container">
            <h3>
                 <a href="/" @click.prevent="cancel">Back</a> |
                 Songs
            </h3>

        </div>
    </div>    
`,
    beforeRouteEnter(to, from, next) {
      next((vm) => {
          if (!vm.data.songs) {
              Bus.$emit(SONGS_LIST);
          }
      })
    },
    methods: {
        cancel() {
            this.$router.push({ name: `view`});
        }
    }

}


