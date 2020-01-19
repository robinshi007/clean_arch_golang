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
];
export default authRoutes;
