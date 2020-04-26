export default {
    props: ['song'],
    template: `
        <header class="header"  >
            <template v-if="false">
                <span  class="header__background" :style="{'background-image': 'url(' +song.artUrl +')'}"></span>
                <img   class="header__album-art" :src="song.artUrl" :alt="song.artist">            
            </template>
            <template v-if="true">            
                <span class="header__background" style="background-image: url(/static/img/icon.png);"></span>
                <img  class="header__album-art" style="padding: 5px; box-sizing: border-box;" src="/static/img/icon.png" :alt="song.artist" />
            </template>
            <h3>{{ song.title }}</h3>
            <h4>{{ song.artist }}</h4>
        </header>      
`,

}
