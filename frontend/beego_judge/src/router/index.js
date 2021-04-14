import Vue from 'vue'
import Router from 'vue-router'
import main from '@/view/main'
import matchDetail from '@/view/matchDetail'
Vue.use(Router)

export default new Router({
  routes: [
    {
      path: "/",
      name: 'main',
      component: main
    },

    {
      path: "/match/:id",
      name: 'match',
      component: matchDetail
    },

  ]
})
