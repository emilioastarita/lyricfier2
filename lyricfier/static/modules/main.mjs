import Lyricfier from "./Lyricfier.mjs";
import Vue from './vue.mjs';
import VueRouter from './vue-router.mjs'

Vue.use(VueRouter);

let mainApp = null;
async function setup() {
    mainApp = new Vue({
        el: '#app',
        components: {
            Lyricfier,
        }
    });
}
document.addEventListener('DOMContentLoaded', setup, false);
