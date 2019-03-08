export default {
    install (Vue, options) {
        Vue.prototype.$l = console.log.bind(console)
        Vue.l = console.log.bind(console)
    }
}