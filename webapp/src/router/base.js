import Home from '../views/Home.vue';

const baseRoutes = [
  {
    name: 'home.index',
    path: '/',
    component: Home,
    meta: {
      isPublic: true,
    },
  },
  {
    name: 'home.about',
    path: '/about',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import(/* webpackChunkName: "about" */ '../views/About.vue'),
    meta: {
      isPublic: true,
    },
  },
  // 404 error page
  {
    name: 'notfound',
    path: '/404',
    component: () => import('../views/NotFound.vue'),
    meta: {
      isPublic: true,
    },
  },
  {
    name: '404',
    path: '*',
    component: () => import('../views/NotFound.vue'),
    meta: {
      isPublic: true,
    },
  },
];

export default baseRoutes;
