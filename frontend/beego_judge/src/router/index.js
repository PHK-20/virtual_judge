import Vue from 'vue'
import Router from 'vue-router'
import main from '@/view/main'
import matchDetail from '@/view/matchDetail'
import problemPage from '@/components/problemPage.vue'
import rank from '@/components/rank.vue'
Vue.use(Router)

export default new Router({
  routes: [
    {
      path: "/",
      name: 'main',
      component: main
    },

    {
      path: "/match/:matchid",
      name: 'match',
      component: matchDetail
    },
    {
      path: "/match/:matchid/:info",
      name: 'matchProblem',
      component: problemPage,
    },
    {
      path: "/problem/:oj/:pid",
      name: 'problem',
      component: problemPage,
    },
  ]
})
