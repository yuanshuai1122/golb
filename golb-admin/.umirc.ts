import { defineConfig } from '@umijs/max';

export default defineConfig({
  antd: {},
  access: {},
  model: {},
  initialState: {},
  request: {},
  layout: {
    title: '@umijs/max',
  },
  routes: [
    {
      path: '/',
      redirect: '/home',
    },
    {
      name: '首页',
      path: '/home',
      component: './Home',
    },
    {
      name: '文章管理',
      path: '/articles',
      routes: [
        {
          name: '文章发布',
          path: '/articles/publish',
          component: './Articles/Publish',
        },
      ],
    },

  ],
  npmClient: 'pnpm',
});

