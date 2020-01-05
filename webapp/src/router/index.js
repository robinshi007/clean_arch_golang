import Vue from 'vue';
import VueRouter from 'vue-router';
import { Notify } from 'quasar';

import store from '../store';
import Home from '../views/Home.vue';
import ViewEmpty from '../views/ViewEmpty.vue';

Vue.use(VueRouter);


const loginRoutes = [
  {
    name: 'auth.login',
    path: '/login',
    component: () => import('../views/Login.vue'),
    meta: {
      isPublic: true,
      layout: 'empty',
    },
  },
];
const routes = [
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
    name: '404',
    path: '/404',
    component: () => import('../views/NotFound.vue'),
    meta: {
      isPublic: true,
    },
  },
];
const adminRoutes = [
  {
    path: '/admin',
    component: ViewEmpty,
    children: [
      {
        name: 'admin.account.new',
        path: 'accounts/create',
        component: () => import('../views/admin/AccountForm.vue'),
      },
      {
        name: 'admin.account.get',
        path: 'accounts/:id',
        component: () => import('../views/admin/Account.vue'),
      },
      {
        name: 'admin.account.edit',
        path: 'accounts/:id/edit',
        component: () => import('../views/admin/AccountForm.vue'),
      },
      {
        name: 'admin.account.list',
        path: 'accounts',
        component: () => import('../views/admin/Accounts.vue'),
      },
    ],
  },
];

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes: [
    ...loginRoutes,
    ...adminRoutes,
    ...routes,
  ],
});

// router hooks before
router.beforeEach((to, from, next) => {
  // console.log('router before:');
  // if (to.meta && to.meta.isPublic && !store.getters.isLoggedIn) {
  new Promise((resolve, reject) => {
    // add some code
    resolve(0);
  }).then((res) => {
    // if no to mached, go to global 404 page
    if (to.matched.length === 0) {
      router.push({ name: '404' });
      return;
    }
    if ((!to.meta || !to.meta.isPublic) && !store.getters.isLoggedIn) {
      // need auth.login
      next({ name: 'auth.login', query: { redirect: to.path } });
      Notify.create({ message: 'Please login' });
    } else if ((!to.meta || !to.meta.isPublic)
      && (store.getters.tokenExpiresAt - Date.now() / 1000 < 0)) {
      // session expired
      next({ name: 'auth.login', query: { redirect: to.path } });
      Notify.create({ message: 'The session is expired, please login.' });
    } else if (!to.meta || !to.meta.isPublic) {
      if (store.getters.tokenExpiresAt - (Date.now() / 1000) <= 300) {
        // need refresh token
        console.log('token refreshed');
        store.dispatch('refreshToken').then((data) => {
          if (data) {
            next();
          }
        }, (data) => {
          this.$q.notify({ message: data });
        }).catch((err) => {
          console.log('refresh token catched err:', err);
        });
      }
      next();
    } else {
      // normal route
      next();
    }
  }).catch((err) => {
    console.log('router.beforeEach catched err:', err);
    next(false);
  });
});

// router hooks after
router.afterEach((to, from) => {
});

export default router;
