
class Service {
    baseUrl = '';
    socket = null;

    constructor() {
        this.baseUrl = document.location.protocol + '//' + document.location.host;
    }

    getSocket() {
        if (!this.socket) {
            this.socket = new WebSocket("ws://" + document.location.host + "/ws");
        }
        return this.socket;
    }

    url(endPoint) {
        return `${this.baseUrl}${endPoint}`;
    }

    async saveSong(songData) {
        const url = this.url('/save-song');
        const res = await fetch(url, {
            method: 'POST',
            body: JSON.stringify(songData),
            headers:{
                'Content-Type': 'application/json'
            }
        });
        return await res.json();
    }

    connectUpdates(update) {
        const conn = this.getSocket();
        conn.onclose = evt => {
            console.log('Connection error', evt)
        };
        conn.onmessage = update;
    }

    async getLyricfierStatus() {
        const response = await fetch('/status');
        if (response.status !== 200) {
            console.error('Response error', response);
            return new Error('Error lyricfier status');
        }
        const data = await response.json();
        return data;
    }

    async saveSettings(settings) {
        const url = document.location.protocol + '//' + document.location.host + '/save-settings';
        const res = await fetch(url, {
            method: 'POST',
            body: JSON.stringify(settings),
            headers:{
                'Content-Type': 'application/json'
            }
        });
        const data = await res.json();
    }
}


export const service = new Service();