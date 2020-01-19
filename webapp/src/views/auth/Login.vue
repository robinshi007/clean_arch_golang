<template>
  <q-page padding>
    <div class="q-pa-md" style="max-width:450px;margin-left:auto;margin-right:auto;">
    <q-toolbar class="q-px-none">
      <q-toolbar-title>
        Login
      </q-toolbar-title>
      <q-space />
    </q-toolbar>
      <q-form
        @submit="onLogin"
        class="q-gutter-md"
        >
        <q-input
          outlined
          v-model="email"
          label="Email"
          lazy-rules
          :rules="[
          val => val && val.length > 0 || 'Please type user email',
          val => val && checkEmail(val) || 'Please use email xxx@xxx.xxx',
          ]"
          />

        <q-input
          outlined
          type="password"
          v-model="password"
          label="Password"
          lazy-rules
          :rules="[
          val => val && val.length > 0 || 'Please type password',
          val => val && val.length >= 8 || 'please input password longer then 8'
          ]"
          />

          <div>
            <q-btn unelevated label="Login" type="submit" color="primary"/>
          </div>
      </q-form>
    </div>
  </q-page>
</template>

<script>
import validateEmail from '@/utils/validateEmail';

export default {
  name: 'PageLogin',
  data() {
    return {
      email: '',
      password: '',
    };
  },
  mounted() {
  },
  methods: {
    onLogin() {
      const postData = {
        email: this.email,
        password: this.password,
      };
      if (this.$route.query.redirect) {
        postData.redirect = this.$route.query.redirect;
      }
      this.$store.dispatch('auth/login', postData).then((data) => {
        if (data) {
          const { query } = this.$route;
          if (query && query.redirect) {
            this.$router.replace(query.redirect);
          } else {
            this.$router.push({ name: 'home.index' });
          }
          this.$q.notify({ message: 'Login successfully.' });
        }
      }, (data) => {
        this.$q.notify({ message: 'Login failed, incorrect email or password.' });
      }).catch((err) => {
        console.log('auth catched err:', err);
      });
    },
    checkEmail(val) {
      return validateEmail(val);
    },
  },
};
</script>

<style>
</style>
