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

    async post(path, data) {
        const url = this.url(path);
        const res = await fetch(url, {
            method: 'POST',
            body: JSON.stringify(data),
            headers: {
                'Content-Type': 'application/json'
            }
        });
        return await res.json();
    }

    async get(path, data) {
        const url = this.url(path);
        const res = await fetch(url, {
            method: 'GET',
            headers: {
                'Content-Type': 'application/json'
            }
        });

        if (res.status !== 200) {
            console.error('Response error', res);
            return new Error(`Error GET fetching: ${url}.`);
        }
        return await res.json();
    }

    connectUpdates(update) {
        const conn = this.getSocket();
        conn.onclose = evt => {
            console.log('Connection error', evt)
        };
        conn.onmessage = update;
    }

    async saveSong(songData) {
        return this.post(`/save-song`, songData);
    }

    async getLyricfierStatus() {
        return this.get(`/status`)
    }

    async saveSettings(settings) {
        return this.post(`/save-settings`, settings);
    }
}


export const service = new Service();