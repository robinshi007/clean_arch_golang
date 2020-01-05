import Vue from 'vue';

import './registerServiceWorker';
import './quasar';
import { createProvider } from './graphql';
import './http';

import LayoutDefault from './layouts/LayoutDefault.vue';
import LayoutEmpty from './layouts/LayoutEmpty.vue';

import store from './store';
import router from './router';
import App from './App.vue';

Vue.component('default-layout', LayoutDefault);
Vue.component('empty-layout', LayoutEmpty);

Vue.config.productionTip = false;

new Vue({
  router,
  store,
  apolloProvider: createProvider(),
  render: h => h(App),
}).$mount('#app');
