const authRoutes = [
  {
    name: 'auth.login',
    path: '/login',
    component: () => import('../views/auth/Login.vue'),
    meta: {
      isPublic: true,
      layout: 'empty',
    },
  },
  {
    name: 'auth.change_password',
    path: '/change_password',
    component: () => import('../views/auth/ChangePassword.vue'),
    meta: {
      layout: 'empty',
    },
  },
];
export default authRoutes;
