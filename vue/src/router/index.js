import Vue from 'vue'
import Router from 'vue-router'
Vue.use(Router)
export default new Router({
  routes: [
    {
      path: '/',
      name: 'index',
      component: resolve => require(['@/views/index.vue'], resolve),
      children: [
        {
          path: '/content/:pageId',
          name: 'content',
          //component: () => import('@/components/markdown.vue')
          component: resolve => require(['@/views/index.vue'], resolve),
        }]
    }
  ]
})
