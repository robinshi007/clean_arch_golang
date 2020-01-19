<template>
  <q-page class="app-page">
    <q-card v-if="loading == 0">
      <q-card-section>
        <div class="text-h6">
          User Details
        </div>
      </q-card-section>
      <q-card-section>
        {{user.id}}
      </q-card-section>
      <q-card-section>
        {{user.name}}
      </q-card-section>
    </q-card>
  </q-page>
</template>

<script>
import ADMIN_USER from '@/graphql/User.gql';
// import gql from 'graphql-tag';

export default {
  data() {
    return {
      loading: 0,
    };
  },
  apollo: {
    user: {
      query: ADMIN_USER,
      update: data => data.fetchUser,
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
