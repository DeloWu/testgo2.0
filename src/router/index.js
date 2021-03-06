import Vue from 'vue'
import Router from 'vue-router'

Vue.use(Router)


/**
 * Note: sub-menu only appear when route children.length >= 1
 * Detail see: https://panjiachen.github.io/vue-element-admin-site/guide/essentials/router-and-nav.html
 *
 * hidden: true                   if set true, item will not show in the sidebar(default is false)
 * alwaysShow: true               if set true, will always show the root menu
 *                                if not set alwaysShow, when item has more than one children route,
 *                                it will becomes nested mode, otherwise not show the root menu
 * redirect: noRedirect           if set noRedirect will no redirect in the breadcrumb
 * name:'router-name'             the name is used by <keep-alive> (must set!!!)
 * meta : {
    roles: ['admin','editor']    control the page roles (you can set multiple roles)
    title: 'title'               the name show in sidebar and breadcrumb (recommend set)
    icon: 'svg-name'             the icon show in the sidebar
    breadcrumb: false            if set false, the item will hidden in breadcrumb(default is true)
    activeMenu: '/example/list'  if set path, the sidebar will highlight the path you set
  }
 */

/**
 * constantRoutes
 * a base page that does not have permission requirements
 * all roles can be accessed
 */
export const constantRoutes = [
    {
        path: '/login',
        component: () => import('@views/login/loginView.vue'),
        hidden: true
    },
    {
        path: '/',
        component: () => import('@views/indexView.vue'),
        hidden: true
    },
    {
        path: '/pro-index',
        component: () => import('@views/project/proIndexView.vue'),
        hidden: true
    },
    {
        path: '/pro-add',
        component: () => import('@views/project/proAddView.vue'),
        hidden: true
    },
    {
        path: '/pro-update',
        component: () => import('@views/project/proUpdateView.vue'),
        hidden: true
    },
    {
        path: '/env-index',
        component: () => import('@views/environment/envIndexView.vue'),
        hidden: true
    },
    {
        path: '/env-add',
        component: () => import('@views/environment/envAddView.vue'),
        hidden: true
    },
    {
        path: '/env-update',
        component: () => import('@views/environment/envUpdateView.vue'),
        hidden: true
    },
    {
        path: '/api-index',
        component: () => import('@views/api/apiIndexView.vue'),
        hidden: true
    },
    {
        path: '/api-add',
        component: () => import('@views/api/apiAddView.vue'),
        hidden: true
    },
    {
        path: '/api-update',
        component: () => import('@views/api/apiUpdateView.vue'),
        hidden: true
    },
    {
        path: '/case-index',
        component: () => import('@views/case/caseIndexView.vue'),
        hidden: true
    },
    {
        path: '/case-add',
        component: () => import('@views/case/caseAddView.vue'),
        hidden: true
    },
    {
        path: '/case-update',
        component: () => import('@views/case/caseUpdateView.vue'),
        hidden: true
    },
    {
        path: '/suite-index',
        component: () => import('@views/suite/suiteIndexView.vue'),
        hidden: true
    },
    {
        path: '/suite-add',
        component: () => import('@views/suite/suiteAddView.vue'),
        hidden: true
    },
    {
        path: '/suite-update',
        component: () => import('@views/suite/suiteUpdateView.vue'),
        hidden: true
    },
    {
        path: '/report-index',
        component: () => import('@views/report/reportIndexView.vue'),
        hidden: true
    },
    {
        path: '/plan-index',
        component: () => import('@views/plan/planIndexView.vue'),
        hidden: true
    },
    {
        path: '/mock-server-index',
        component: () => import('@views/mockServer/mockServerIndexView.vue'),
        hidden: true
    },
]

const createRouter = () => new Router({
    // mode: 'history', // require service support
    scrollBehavior: () => ({
        y: 0
    }),
    routes: constantRoutes
})

const router = createRouter()

// Detail see: https://github.com/vuejs/vue-router/issues/1234#issuecomment-357941465
export function resetRouter() {
    const newRouter = createRouter()
    router.matcher = newRouter.matcher // reset router
}

export default router
