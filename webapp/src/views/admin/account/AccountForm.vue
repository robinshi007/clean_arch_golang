<template>
  <q-page class="app-page">
    <q-card v-if="loading == 0">
      <q-card-section>
        <div class="text-h6">
          Account {{account.id ? "Edit" : "Create"}} Form
        </div>
      </q-card-section>
      <q-card-section>
        <q-form
          @submit="onSubmit"
          class='q-gutter-md'
          >
          <q-input
            outlined
            v-model="account.name"
            label="Name"
            lazy-rules
            :rules="[val => val && val.length > 0 || 'Please input name']"
            />
          <q-input
            outlined
            v-model="account.email"
            label="Email"
            lazy-rules
            :rules="[
            val => val && val.length > 0 || 'Please input email',
            val => val && checkEmail(val) || 'Please use email xxx@xxx.xxx',
            ]"
            v-if="!account.id"
            />
          <q-input
            outlined
            type="password"
            label="Password"
            v-model="account.password"
            v-if="!account.id"
            lazy-rules
            :rules="[
            val => val && val.length > 0 || 'Please input password',
            val => val && val.length >= 8 || 'please input password longer then 8'
            ]"
            />
          <div class='q-gutter-md'>
            <q-btn flat label="Cancel"  @click="onCancel" />
            <q-btn unelevated label="Save" type="submit" color="primary" />
          </div>
        </q-form>
      </q-card-section>
    </q-card>
  </q-page>
</template>

<script>
import validateEmail from '../../../utils/validateEmail';
import ADMIN_ACCOUNT from '@/graphql/Account.gql';
import ADMIN_ACCOUNT_CREATE from '@/graphql/AccountCreate.gql';
import ADMIN_ACCOUNT_UPDATE from '@/graphql/AccountUpdate.gql';
// import gql from 'graphql-tag';

export default {
  data() {
    return {
      loading: 0,
      account: {
        id: '',
        name: '',
        email: '',
        password: '',
      },
    };
  },
  created() {
    // fetch data when edit item
    if (this.$route.params.id) {
      this.$apollo.addSmartQuery('account', {
        query: ADMIN_ACCOUNT,
        update: data => data.fetchAccount,
        variables() {
          return {
            id: this.$route.params.id,
          };
        },
      });
    }
  },
  computed: {
  },

  methods: {
    onSubmit() {
      if (this.account.id) {
        this.$apollo.mutate({
          mutation: ADMIN_ACCOUNT_UPDATE,
          variables: {
            id: this.account.id,
            name: this.account.name,
          },
          update: () => {
            this.$q.notify({ message: `Account #${this.account.id} updated successfully.` });
            this.$router.go(-1);
            // this.$router.push({ name: 'admin.account.get', params: { id: this.account.id } });
          },
        }).catch((error) => {
          const { graphQLErrors } = error;
          if (graphQLErrors && graphQLErrors.some(e => e.message === 'requested item is not changed')) {
            this.$q.notify({ message: 'Item is not changed' });
          } else if (graphQLErrors && graphQLErrors.some(e => e.message.includes('pq: duplicate key value violates unique constraint "user_accounts_name_key'))) {
            this.$q.notify({ message: `The name "${this.account.name}" has been used` });
          }
        });
      } else {
        this.$apollo.mutate({
          mutation: ADMIN_ACCOUNT_CREATE,
          variables: {
            name: this.account.name,
            email: this.account.email,
            password: this.account.password,
          },
          update: (cache, { data: { createAccount } }) => {
            this.$q.notify({
              message: `Account #${createAccount.id} created successfully.`,
              actions: [
                { label: 'Dismiss', handler: () => { /* ... */ } },
              ],
            });
            this.$router.go(-1);
            // this.$router.push({ name: 'admin.account.get', params: { id: this.account.id } });
          },
        }).catch((error) => {
          const { graphQLErrors } = error;
          if (graphQLErrors && graphQLErrors.some(e => e.message === 'requested item is not changed')) {
            this.$q.notify({ message: 'Item is not changed' });
          } else if (graphQLErrors && graphQLErrors.some(e => e.message.includes('pq: duplicate key value violates unique constraint "user_accounts_name_key'))) {
            this.$q.notify({ message: `The name "${this.account.name}" has been used` });
          } else if (graphQLErrors && graphQLErrors.some(e => e.message.includes('pq: duplicate key value violates unique constraint "user_accounts_email_key'))) {
            this.$q.notify({ message: `The email "${this.account.email}" has been used` });
          }
        });
      }
    },
    onCancel() {
      this.$router.go(-1);
    },
    checkEmail(val) {
      const res = validateEmail(val);
      return res;
    },
  },
};
</script>
