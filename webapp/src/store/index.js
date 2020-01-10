import Vue from 'vue';
import Vuex from 'vuex';

Vue.use(Vuex);

export default new Vuex.Store({
  modules: {
  },
  state: {
    status: '',
    email: localStorage.getItem('email') || '',
    name: localStorage.getItem('name') || '',
    token: localStorage.getItem('token') || '',
    token_expires_at: localStorage.getItem('token_expires_at') || '',
  },
  getters: {
    email: state => state.email,
    name: state => state.name,
    isLoggedIn: state => !!state.token && (+state.token_expires_at > (Date.now() / 1000)),
    tokenExpiresAt: state => +state.token_expires_at,
  },
  mutations: {
    auth_request(state) {
      state.status = 'login requesting';
    },
    auth_success(state, data) {
      state.status = 'login success';
      state.email = data.email;
      state.name = data.name;
      state.token = data.token;
      state.token_expires_at = data.expiresAt;
      localStorage.setItem('email', state.email);
      localStorage.setItem('name', state.name);
      localStorage.setItem('token', state.token);
      localStorage.setItem('token_expires_at', state.token_expires_at);
    },
    auth_error(state) {
      state.status = 'login error';
      state.email = '';
      state.name = '';
      state.token = '';
      state.token_expires_at = '';
      localStorage.removeItem('email');
      localStorage.removeItem('name');
      localStorage.removeItem('token');
      localStorage.removeItem('token_expires_at');
    },
    auth_token_refresh(state, data) {
      state.status = 'token refresh success';
      state.token = data.token;
      state.token_expires_at = data.expiresAt;
      localStorage.setItem('token', state.token);
      localStorage.setItem('token_expires_at', state.token_expires_at);
    },
    auth_logout(state) {
      state.status = '';
      state.token = '';
      state.token_expires_at = '';
      state.email = '';
      state.name = '';
      localStorage.removeItem('email');
      localStorage.removeItem('name');
      localStorage.removeItem('token');
      localStorage.removeItem('token_expires_at');
    },
  },
  actions: {
    login({ commit }, data) {
      commit('auth_request');
      return Vue.axios.post('/api/v1/auth/login', data).then((response) => {
        if (response.data.success) {
          const { name, token, expiresAt } = response.data.data;
          commit('auth_success', {
            name, email: data.email, token, expiresAt,
          });
          Vue.axios.defaults.headers.common.Authorization = `Bearer ${token}`;
          const postData = response.data;
          // add redirect info
          if (data.redirect) {
            postData.redirect = data.redirect;
          }
          return Promise.resolve(postData);
        }
        commit('auth_error');
        return Promise.reject(response.data);
      }).catch((err) => {
        commit('auth_error');
        return Promise.reject(err);
      });
    },
    refreshToken({ commit }) {
      return Vue.axios.get('/api/v1/auth/refresh_token').then((response) => {
        if (response.data.data.token) {
          const { token, expiresAt } = response.data.data;
          commit('auth_token_refresh', { token, expiresAt });
          Vue.axios.defaults.headers.common.Authorization = `Bearer ${token}`;
          return Promise.resolve(response);
        }
        return Promise.reject(response.data);
      }).catch(err => Promise.reject(err));
    },
    logout({ commit }) {
      return new Promise((resolve, reject) => {
        commit('auth_logout');
        delete Vue.axios.defaults.headers.common.Authorization;
        resolve();
      });
    },
    setUsername({ commit }) {
      return Vue.axios.get('/api/v1/auth/hello').then((response) => {
        commit('set_username', response.data.userName);
        localStorage.setItem('username', response.data.userName);
      });
    },
  },
});
