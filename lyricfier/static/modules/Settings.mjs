import {Bus, SAVED_SETTINGS} from "./Events.mjs";

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
                <input type="radio" name="theme" value="dark" v-model="data.editSettings.theme">
                <label for="theme_dark">Dark</label>
                <hr />
                <p>
                    <button type="submit">Save</button>
                    <a href="/" @click.prevent="cancel">Cancel</a>
                </p>            
            </div>
        </form>
    </div>    
`,
    methods: {
        async submit() {
            const url = document.location.protocol + '//' + document.location.host + '/save-settings';
            const res = await fetch(url, {
                method: 'POST',
                body: JSON.stringify(this.data.editSettings),
                headers:{
                    'Content-Type': 'application/json'
                }
            });
            const data = await res.json();
            Bus.$emit(SAVED_SETTINGS, this.data.editSettings);
        },
        cancel() {
            Bus.$emit(SAVED_SETTINGS, null);
        }
    }

}


