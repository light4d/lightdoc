import Vue from 'vue'
import Router from 'vue-router'
Vue.use(Router)
export default new Router({
  routes: [
    {
      path: '/',
      name: 'index',
      component: resolve => require(['@/views/index.vue'], resolve)
    }
  ]
})
