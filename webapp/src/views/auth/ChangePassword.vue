<template>
  <q-page padding>
    <div class="q-pa-md" style="max-width:450px;margin-left:auto;margin-right:auto;">
    <q-toolbar class="q-px-none">
      <q-toolbar-title>
        Change Password
      </q-toolbar-title>
      <q-space />
    </q-toolbar>
      <q-form
        @submit="onChangePassword"
        class="q-gutter-md"
        >

        <q-input
          outlined
          type="password"
          v-model="password_current"
          label="Current Password"
          lazy-rules
          :rules="[
          val => val && val.length > 0 || 'Please type current password',
          val => val && val.length >= 8 || 'please input current password longer then 8'
          ]"
          />

        <q-input
          outlined
          type="password"
          v-model="password_new"
          label="New Password"
          lazy-rules
          :rules="[
          val => val && val.length > 0 || 'Please type new password',
          val => val && val.length >= 8 || 'please input new password longer then 8'
          ]"
          />

          <div>
            <q-btn unelevated label="Submit" type="submit" color="primary"/>
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
      password_current: '',
      password_new: '',
    };
  },
  mounted() {
  },
  methods: {
    onChangePassword() {
      const postData = {
        password_current: this.password_current,
        password: this.password_new,
      };
      if (this.$route.query.redirect) {
        postData.redirect = this.$route.query.redirect;
      }
      this.$store.dispatch('auth/change_password', postData).then((data) => {
        if (data) {
          const { query } = this.$route;
          if (query && query.redirect) {
            this.$router.replace(query.redirect);
          } else {
            this.$router.push({ name: 'home.index' });
          }
          this.$q.notify({ message: 'Chnage password successfully.' });
        }
      }, (data) => {
        this.$q.notify({ message: 'Change password failed, incorrect current password.' });
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
