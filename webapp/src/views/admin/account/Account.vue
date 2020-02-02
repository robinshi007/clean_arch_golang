<template>
  <q-page class="app-page">
    <q-card v-if="loading == 0">
      <q-card-section>
        <div class="text-h6">
          Account Details
        </div>
      </q-card-section>
      <q-card-section>
        ID: {{account.id}}
      </q-card-section>
      <q-card-section>
        Name: {{account.name}}
      </q-card-section>
      <q-card-section>
        Email: {{account.email}}
      </q-card-section>
    </q-card>
  </q-page>
</template>

<script>
import ADMIN_ACCOUNT from '@/graphql/Account.gql';
// import gql from 'graphql-tag';

export default {
  data() {
    return {
      loading: 0,
    };
  },
  apollo: {
    account: {
      query: ADMIN_ACCOUNT,
      update: data => data.fetchAccount,
      variables() {
        return {
          id: this.$route.params.id,
        };
      },
    },
  },
  computed: {
  },

  methods: {
  },
};
</script>
