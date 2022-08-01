import { createRouter, createWebHistory, RouteRecordRaw } from "vue-router";

const routes: Array<RouteRecordRaw> = [
  {
    path: "/",
    name: "home",
    component: () => import("@/modules/core/pages/DashBoard.vue"),
  },
  {
    path: "/users",
    name: "users",
    component: () => import("@/modules/core/pages/UsersPage.vue"),
  },
  {
    path: "/contacts",
    name: "contacts",
    component: () => import("@/modules/core/pages/ContactsPage.vue"),
  },
  {
    path: "/accounts",
    name: "accounts",
    component: () => import("@/modules/core/pages/AccountsPage.vue"),
  },
  {
    path: "/settings",
    name: "settings",
    component: () => import("@/modules/core/pages/SettingsPage.vue"),
  },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
