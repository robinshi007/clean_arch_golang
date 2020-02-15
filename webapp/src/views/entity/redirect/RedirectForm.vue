<template>
  <q-page class="app-page">
    <q-card v-if="loading == 0">
      <q-card-section>
        <div class="text-h6">
          Redirect {{redirect.id ? "Edit" : "Create"}} Form
        </div>
      </q-card-section>
      <q-card-section>
        <q-form
          @submit="onSubmit"
          class='q-gutter-md'
          >
          <q-input
            outlined
            v-model="redirect.url"
            label="URL"
            lazy-rules
            :rules="[
            val => val && val.length > 0 || 'Please input url',
            val => val && checkUrl(val) || 'Please use url https?://xxx.xxx.xxx',
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
import validateUrl from '../../../utils/validateUrl';
import ADMIN_REDIRECT from '@/graphql/Redirect.gql';
import ADMIN_REDIRECT_CREATE from '@/graphql/RedirectCreate.gql';
// import gql from 'graphql-tag';

export default {
  data() {
    return {
      loading: 0,
      redirect: {
        id: '',
        url: '',
      },
    };
  },
  created() {
    // fetch data when edit item
    if (this.$route.params.id) {
      this.$apollo.addSmartQuery('redirect', {
        query: ADMIN_REDIRECT,
        update: data => data.fetchRedirect,
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
      this.$apollo.mutate({
        mutation: ADMIN_REDIRECT_CREATE,
        variables: {
          url: this.redirect.url,
        },
        update: (cache, { data: { createRedirect } }) => {
          this.$q.notify({
            message: `Redirect #${createRedirect.id} created successfully.`,
            actions: [
              { label: 'Dismiss', handler: () => { /* ... */ } },
            ],
          });
          this.$router.go(-1);
          // this.$router.push({ name: 'admin.redirect.get', params: { id: this.redirect.id } });
        },
      }).catch((error) => {
        const { graphQLErrors } = error;
        if (graphQLErrors && graphQLErrors.some(e => e.message === 'requested item is not changed')) {
          this.$q.notify({ message: 'Item is not changed' });
        } else if (graphQLErrors && graphQLErrors.some(e => e.message.includes('pq: duplicate key value violates unique constraint "user_redirects_name_key'))) {
          this.$q.notify({ message: `The name "${this.redirect.name}" has been used` });
        } else if (graphQLErrors && graphQLErrors.some(e => e.message.includes('pq: duplicate key value violates unique constraint "user_redirects_email_key'))) {
          this.$q.notify({ message: `The email "${this.redirect.email}" has been used` });
        }
      });
    },
    onCancel() {
      this.$router.go(-1);
    },
    checkUrl(val) {
      const res = validateUrl(val);
      return res;
    },
  },
};
</script>
