import ViewEmpty from '../views/ViewEmpty.vue';

const entityRoutes = [
  {
    path: '/entity',
    component: ViewEmpty,
    children: [
      // user
      {
        name: 'entity.user.new',
        path: 'users/create',
        component: () => import('../views/entity/user/UserForm.vue'),
      },
      {
        name: 'entity.user.get',
        path: 'users/:id',
        component: () => import('../views/entity/user/User.vue'),
      },
      {
        name: 'entity.user.edit',
        path: 'users/:id/edit',
        component: () => import('../views/entity/user/UserForm.vue'),
      },
      {
        name: 'entity.user.list',
        path: 'users',
        component: () => import('../views/entity/user/Users.vue'),
      }, // redirect
      {
        name: 'entity.redirect.new',
        path: 'redirects/create',
        component: () => import('../views/entity/redirect/RedirectForm.vue'),
      },
      {
        name: 'entity.redirect.get',
        path: 'redirects/:code',
        component: () => import('../views/entity/redirect/Redirect.vue'),
      },
      {
        name: 'entity.redirect.list',
        path: 'redirects',
        component: () => import('../views/entity/redirect/Redirects.vue'),
      },
    ],
  },
];
export default entityRoutes;
