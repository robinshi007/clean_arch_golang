<template>
  <q-page class="app-page">
    <q-card v-if="loading == 0">
      <q-card-section>
        <div class="text-h6">
          User {{user.id ? "Edit" : "Create"}} Form
        </div>
      </q-card-section>
      <q-card-section>
        <q-form
          @submit="onSubmit"
          class='q-gutter-md'
          >
          <q-input
            outlined
            v-model="user.name"
            label="Name"
            lazy-rules
            :rules="[val => val && val.length > 0 || 'Please input name']"
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
import ADMIN_USER from '@/graphql/User.gql';
import ADMIN_USER_CREATE from '@/graphql/UserCreate.gql';
import ADMIN_USER_UPDATE from '@/graphql/UserUpdate.gql';
// import gql from 'graphql-tag';

export default {
  data() {
    return {
      loading: 0,
      user: {
        id: '',
        name: '',
      },
    };
  },
  created() {
    // fetch data when edit item
    if (this.$route.params.id) {
      this.$apollo.addSmartQuery('user', {
        query: ADMIN_USER,
        update: data => data.fetchUser,
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
      if (this.user.id) {
        this.$apollo.mutate({
          mutation: ADMIN_USER_UPDATE,
          variables: {
            id: this.user.id,
            name: this.user.name,
          },
          update: () => {
            this.$q.notify({ message: `User #${this.user.id} updated successfully.` });
            this.$router.go(-1);
            // this.$router.push({ name: 'admin.user.get', params: { id: this.user.id } });
          },
        }).catch((error) => {
          const { graphQLErrors } = error;
          if (graphQLErrors && graphQLErrors.some(e => e.message === 'requested item is not changed')) {
            this.$q.notify({ message: 'Item is not changed' });
          } else if (graphQLErrors && graphQLErrors.some(e => e.message.includes('pq: duplicate key value violates unique constraint "user_users_name_key'))) {
            this.$q.notify({ message: `The name "${this.user.name}" has been used` });
          }
        });
      } else {
        this.$apollo.mutate({
          mutation: ADMIN_USER_CREATE,
          variables: {
            name: this.user.name,
          },
          update: (cache, { data: { createUser } }) => {
            this.$q.notify({
              message: `User #${createUser.id} created successfully.`,
              actions: [
                { label: 'Dismiss', handler: () => { /* ... */ } },
              ],
            });
            this.$router.go(-1);
            // this.$router.push({ name: 'admin.user.get', params: { id: this.user.id } });
          },
        }).catch((error) => {
          const { graphQLErrors } = error;
          if (graphQLErrors && graphQLErrors.some(e => e.message === 'requested item is not changed')) {
            this.$q.notify({ message: 'Item is not changed' });
          } else if (graphQLErrors && graphQLErrors.some(e => e.message.includes('pq: duplicate key value violates unique constraint "user_users_name_key'))) {
            this.$q.notify({ message: `The name "${this.user.name}" has been used` });
          }
        });
      }
    },
    onCancel() {
      this.$router.go(-1);
    },
  },
};
</script>
