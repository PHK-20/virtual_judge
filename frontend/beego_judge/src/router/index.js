import Vue from 'vue'
import Router from 'vue-router'
import main from '@/view/main'
import status from '@/view/status'
Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'main',
      component: main
    },
    {
      path: '/status',
      name: 'status',
      component: status
    }
  ]
})
