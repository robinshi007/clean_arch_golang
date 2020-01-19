<template>
  <q-page class="app-page">
    <q-table
      title= "Users"
      :data="users"
      :columns="columns"
      row-key="id"
      v-if="loading == 0"
      >
      <template v-slot:top-right>
        <q-btn
          type="a"
          @click="onItemNew()"
          round
          color="red"
          icon="add"
          >
          <q-tooltip>New</q-tooltip>
        </q-btn>
      </template>
      <template v-slot:body-cell-action="props">
        <q-td :props="props">
          <div>
            <q-btn type="a" @click="onItemView(props.value)"
              flat round color="primary" icon="launch" >
              <q-tooltip>View</q-tooltip>
            </q-btn>
            <q-btn type="a" @click="onItemEdit(props.value)"
              flat round color="primary" icon="edit" >
              <q-tooltip>Edit</q-tooltip>
            </q-btn>
            <q-btn type="a" @click="onDeleteConfirm(props.value)"
              flat round color="primary" icon="delete" >
              <q-tooltip>Delete</q-tooltip>
            </q-btn>
          </div>
        </q-td>
      </template>
    </q-table>
    <q-dialog
      v-model="confirm"
      persistent
    >
      <q-card style="min-width: 450px" class="q-pa-sm">
        <q-card-section>
          <div class="text-h6">Delete</div>
        </q-card-section>
        <q-card-section class="row items-center">
          <q-avatar icon="warning" color="orange" text-color="white" />
            <span class="q-ml-sm">
              Are you sure to delete User #{{selectedId}}?
            </span>
        </q-card-section>

        <q-card-actions align="right">
          <div class="q-gutter-md">
            <q-btn flat label="Cancel" color="primary" v-close-popup />
            <q-btn unelevated label="Delete" color="red" @click="onDelete()" v-close-popup/>
          </div>
        </q-card-actions>
      </q-card>
    </q-dialog>
  </q-page>
</template>

<script>
import ADMIN_USERS from '@/graphql/Users.gql';
import ADMIN_USER_DELETE from '@/graphql/UserDelete.gql';
import formatDate from '@/utils/date';

export default {
  data() {
    return {
      loading: 0,
      confirm: false,
      selectedId: -1,
      selected: [],
      columns: [
        {
          name: 'id',
          label: '#',
          required: true,
          align: 'left',
          field: row => row.id,
          format: val => `${val}`,
          sortable: true,
        },
        {
          name: 'name',
          label: 'Name',
          align: 'left',
          field: row => row.name,
          format: val => `${val}`,
          sortable: true,
        },
        {
          name: 'email',
          label: 'EMail',
          required: true,
          align: 'left',
          field: row => row.email,
          format: val => `${val}`,
          sortable: true,
        },
        {
          name: 'created_at',
          label: 'Created At',
          required: true,
          align: 'left',
          field: row => row.created_at,
          format: val => `${formatDate(new Date(val))}`,
          sortable: true,
        },
        {
          name: 'updated_at',
          label: 'Update At',
          required: true,
          align: 'left',
          field: row => row.updated_at,
          format: val => `${formatDate(new Date(val))}`,
          sortable: true,
        },
        {
          name: 'action',
          label: 'Action',
          align: 'left',
          field: row => row.id,
          format: val => `${val}`,
        },
      ],
    };
  },
  apollo: {
    users: {
      query: ADMIN_USERS,
      error(err) {
        const { graphQLErrors } = err;
        if (graphQLErrors && graphQLErrors.some(e => e.message === 'the action is unauthorized')) {
          this.$q.notify({ message: 'The action "ADMIN_USERS" is unauthorized' });
        }
        // return 0 to avoid global error handler
        return 0;
      },
    },
  },
  computed: {
  },

  methods: {
    onItemNew() {
      return this.$router.push({ name: 'admin.user.new' });
    },
    onItemView(val) {
      return this.$router.push({ name: 'admin.user.get', params: { id: val } });
    },
    onItemEdit(val) {
      return this.$router.push({ name: 'admin.user.edit', params: { id: val } });
    },
    onDeleteConfirm(val) {
      this.selectedId = val;
      this.confirm = true;
    },
    onDelete() {
      this.$apollo.mutate({
        mutation: ADMIN_USER_DELETE,
        variables: {
          id: this.selectedId,
        },
        update: () => {
          this.$apollo.queries.users.refetch();
          this.$q.notify({
            message: `User #${this.selectedId} is deleted successfully.`,
            actions: [
              { label: 'Dismiss', handler: () => { /* ... */ } },
            ],
          });
          this.selectedId = -1;
        },
      });
    },
    onCancel() {
      this.selectedId = -1;
    },
  },
};
</script>
