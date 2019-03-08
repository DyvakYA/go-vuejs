import Vue from 'vue'
import App from './App.vue'
import router from './router/router'
import log from './components/js/log'

Vue.config.productionTip = false

Vue.use(log)

new Vue({
    el: '#app',
    router,
    render: h => h(App),
});
