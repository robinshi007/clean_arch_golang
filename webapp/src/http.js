import Vue from 'vue';
import Axios from 'axios';

Vue.axios = Axios;

Object.defineProperties(Vue.prototype, {
  axios: {
    get() {
      return Axios;
    },
  },
  $http: {
    get() {
      return Axios;
    },
  },
});

// init token for axios
const token = localStorage.getItem('token') || '';
Vue.axios.defaults.headers.common.Authorization = `Bearer ${token}`;
