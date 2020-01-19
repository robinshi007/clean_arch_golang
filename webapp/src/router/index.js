import Vue from 'vue';
import VueRouter from 'vue-router';
import { Notify } from 'quasar';

import store from '../store';
import authRoutes from './auth';
import adminRoutes from './admin';
import baseRoutes from './base';

Vue.use(VueRouter);

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes: [
    ...authRoutes,
    ...adminRoutes,
    ...baseRoutes,
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
    if ((!to.meta || !to.meta.isPublic) && !store.getters['auth/isLoggedIn']) {
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
