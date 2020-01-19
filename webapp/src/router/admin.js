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
      }, // user
      {
        name: 'admin.user.new',
        path: 'users/create',
        component: () => import('../views/admin/user/UserForm.vue'),
      },
      {
        name: 'admin.user.get',
        path: 'users/:id',
        component: () => import('../views/admin/user/User.vue'),
      },
      {
        name: 'admin.user.edit',
        path: 'users/:id/edit',
        component: () => import('../views/admin/user/UserForm.vue'),
      },
      {
        name: 'admin.user.list',
        path: 'users',
        component: () => import('../views/admin/user/Users.vue'),
      }, // redirect
      {
        name: 'admin.redirect.new',
        path: 'redirects/create',
        component: () => import('../views/admin/redirect/RedirectForm.vue'),
      },
      {
        name: 'admin.redirect.get',
        path: 'redirects/:code',
        component: () => import('../views/admin/redirect/Redirect.vue'),
      },
      {
        name: 'admin.redirect.list',
        path: 'redirects',
        component: () => import('../views/admin/redirect/Redirects.vue'),
      },
    ],
  },
];

export default adminRoutes;
