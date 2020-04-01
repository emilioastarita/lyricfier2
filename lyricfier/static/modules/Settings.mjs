import {Bus, SAVED_SETTINGS, EDIT_SETTINGS} from "./Events.mjs";
import {service} from "./Service.mjs";


export default {
    props: ['data'],
    template: `
    <div class="full-vertical-flex">
        <form @submit.prevent="submit" method="post" action="#" class="settings-container">
            <h3>
                 <a href="/" @click.prevent="cancel">Back</a> |
                 Settings
            </h3>
            <div v-if="data.editSettings">
                
                <h5>Theme:</h5>
                <input type="radio" name="theme" id="theme_default" value="default" v-model="data.editSettings.theme">
                <label for="theme_default">Default</label>
                <input type="radio" name="theme" id="theme_dark" value="dark" v-model="data.editSettings.theme">
                <label for="theme_dark">Dark</label>

                <h5>Font size:</h5>
                <input type="number" class="fontSize" name="fontSize" v-model.number="data.editSettings.fontSize" min="8" max="90" /> px

                <h5>Text align:</h5>
                <input type="radio" name="textAlign" id="textAlignLeft" value="left" v-model="data.editSettings.textAlign">
                <label for="textAlignLeft">Left</label>
                <input type="radio" name="textAlign" if="textAlignCenter" value="center" v-model="data.editSettings.textAlign">
                <label for="textAlignCenter">Center</label>

                
                <hr />
                <p>
                    <button type="submit">Save</button>
                    <a href="/" @click.prevent="cancel">Cancel</a>
                </p>            
            </div>
        </form>
    </div>    
`,
    beforeRouteEnter(to, from, next) {
      next((vm) => {
          if (!vm.data.editSettings) {
              Bus.$emit(EDIT_SETTINGS);
          }
      })
    },
    methods: {
        async submit() {
            const data = await service.saveSettings(this.data.editSettings);
            Bus.$emit(SAVED_SETTINGS, data.settings);
        },
        cancel() {
            Bus.$emit(SAVED_SETTINGS, null);
        }
    }

}


