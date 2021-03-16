import Vue from 'vue'
import Router from 'vue-router'
import problem from '@/view/problem'
import status from '@/view/status'
import menu from '@/components/menu'
Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/problem',
      name: 'problem',
      component: problem
    },
    {
      path: '/status',
      name: 'status',
      component: status
    },
    {
      path: "/",
      name: 'menu',
      component: menu
    },
  ]
})
