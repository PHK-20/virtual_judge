import Vue from 'vue'
import Router from 'vue-router'
import problem from '@/view/problem'
Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/problem',
      name: 'problem',
      component: problem
    }
  ]
})
