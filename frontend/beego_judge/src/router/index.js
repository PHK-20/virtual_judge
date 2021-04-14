import Vue from 'vue'
import Router from 'vue-router'
import main from '@/view/main'
import problem from '@/view/problem'
import status from '@/view/status'
import user from '@/view/user'
import matchList from '@/view/matchList'
Vue.use(Router)

export default new Router({
  routes: [
    {
      path: "/",
      name: 'main',
      component: main
    },
  ]
})
