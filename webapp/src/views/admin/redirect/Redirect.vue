<template>
  <q-page class="app-page">
    <q-card v-if="loading == 0 && !!redirect">
      <q-card-section>
        <div class="text-h6">
          Redirect Details
        </div>
      </q-card-section>
      <q-card-section>
        {{redirect.code}}
      </q-card-section>
      <q-card-section>
        {{redirect.url}}
      </q-card-section>
    </q-card>
  </q-page>
</template>

<script>
import ADMIN_REDIRECT from '@/graphql/Redirect.gql';
// import gql from 'graphql-tag';

export default {
  data() {
    return {
      loading: 0,
    };
  },
  apollo: {
    redirect: {
      query: ADMIN_REDIRECT,
      update: data => data.fetchRedirectByCode,
      variables() {
        return {
          code: this.$route.params.code,
        };
      },
      error(err) {
        const { graphQLErrors } = err;
        if (graphQLErrors && graphQLErrors.some(e => e.message === 'requested item is not found')) {
          this.$q.notify({ message: `The item with "${this.$route.params.code}" is not founed` });
        }
        // return 0 to avoid global error handler
        return 0;
      },
    },
  },
  computed: {
  },

  methods: {
  },
};
</script>
