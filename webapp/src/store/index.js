import Vue from 'vue';
import Vuex from 'vuex';

import authState from './auth';

Vue.use(Vuex);

export default new Vuex.Store({
  modules: {
    auth: authState,
  },
  state: {
  },
  getters: {
  },
  mutations: {
  },
  actions: {
  },
});
