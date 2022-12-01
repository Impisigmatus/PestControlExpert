import {createRouter, createWebHistory} from 'vue-router';
import MainPage from '../views/MainPage.vue';
import PrivacyPage from '../views/PrivacyPage.vue';

const routes = [
  {
    path: '/',
    name: 'main',
    component: MainPage,
  },
  {
    path: '/privacy',
    name: 'privacy',
    component: PrivacyPage,
  },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
