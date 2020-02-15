import ViewEmpty from '../views/ViewEmpty.vue';

const adminRoutes = [
  {
    path: '/admin',
    component: ViewEmpty,
    children: [
      {
        name: 'admin.account.new',
        path: 'accounts/create',
        component: () => import('../views/admin/account/AccountForm.vue'),
      },
      {
        name: 'admin.account.get',
        path: 'accounts/:id',
        component: () => import('../views/admin/account/Account.vue'),
      },
      {
        name: 'admin.account.edit',
        path: 'accounts/:id/edit',
        component: () => import('../views/admin/account/AccountForm.vue'),
      },
      {
        name: 'admin.account.list',
        path: 'accounts',
        component: () => import('../views/admin/account/Accounts.vue'),
      },
    ],
  },
];

export default adminRoutes;
