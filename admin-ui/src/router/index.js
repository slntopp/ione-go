import Vue from "vue";
import VueRouter from "vue-router";

Vue.use(VueRouter);

const routes = [
  {
    path: "/",
    name: "Home",
    redirect: { name: "Dashboard" },
  },
  {
    path: "/dashboard",
    name: "Dashboard",
    component: () => import("../views/Dashboard.vue"),
    meta: {
      requireLogin: true,
    },
  },
  {
    path: "/namespaces",
    name: "Namespaces",
    component: () => import("../views/Namespaces.vue"),
    meta: {
      requireLogin: true,
    },
  },
  {
    path: "/accounts",
    name: "Accounts",
    component: () => import("../views/Accounts.vue"),
    meta: {
      requireLogin: true,
    },
  },
  {
    path: "/accounts/:accountId",
    name: "Account",
    component: () => import("../views/AccountPage.vue"),
    meta: {
      requireLogin: true,
    },
  },
  {
    path: "/sp",
    name: "ServicesProviders",
    component: () => import("../views/ServicesProviders.vue"),
    meta: {
      requireLogin: true,
    },
  },
  {
    path: "/sp/create",
    name: "ServicesProviders create",
    component: () => import("../views/ServicesProvidersCreate.vue"),
    meta: {
      requireLogin: true,
    },
  },
  {
    path: "/sp/:uuid",
    name: "ServicesProvider",
    component: () => import("../views/ServicesProvidersPage.vue"),
    meta: {
      requireLogin: true,
    },
  },
  {
    path: "/dns",
    name: "DNS manager",
    component: () => import("../views/dnsManager.vue"),
    meta: {
      requireLogin: true,
    },
  },
  {
    path: "/dns/:dnsname",
    name: "Zone manager",
    component: () => import("../views/ZoneManager.vue"),
    meta: {
      requireLogin: true,
    },
  },
  {
    path: "/settigns",
    name: "Settings",
    component: () => import("../views/Settings.vue"),
    meta: {
      requireLogin: true,
    },
  },
  {
    path: "/settigns/app",
    name: "AppSetting",
    component: () => import("../views/AppSettings.vue"),
    meta: {
      requireLogin: true,
    },
  },
  {
    path: "/services",
    name: "Services",
    component: () => import("../views/Services.vue"),
    meta: {
      requireLogin: true,
    },
  },
  {
    path: "/services/create",
    name: "Service create",
    component: () => import("../views/ServiceCreate.vue"),
    meta: {
      requireLogin: true,
    },
  },
  {
    path: "/services/:serviceId",
    name: "Service",
    component: () => import("../views/ServicePage.vue"),
    meta: {
      requireLogin: true,
    },
  },
  {
    path: "/instances/:instanceId/vnc",
    component: () => import("../views/Vnc.vue"),
    name: "Vnc",
    meta: {
      requireLogin: true,
    },
  },
  {
    path: "/instances/:instanceId/dns",
    component: () => import("../views/InstanceDNS.vue"),
    name: "InstanceDns",
    meta: {
      requireLogin: true,
    },
  },
  {
    path: "/login",
    name: "Login",
    component: () =>
      import(/* webpackChunkName: "about" */ "../views/login.vue"),
    meta: {
      requireUnlogin: true,
    },
  },
  {
    path: "/plans",
    name: "Plans",
    component: () => import("../views/Plans.vue"),
    meta: {
      requireLogin: true,
    },
  },
  {
    path: "/plans/create",
    name: "Plans create",
    component: () => import("../views/PlansCreate.vue"),
    meta: {
      requireLogin: true,
    },
  },
  {
    path: "/plans/:planId",
    name: "Plan",
    component: () => import("../views/PlanPage.vue"),
    meta: {
      requireLogin: true,
    },
  },
  {
    path: "/transactions",
    name: "Transactions",
    component: () => import("../views/Transactions.vue"),
    meta: {
      requireLogin: true,
    },
  },
  {
    path: "/transactions/create",
    name: "Transactions create",
    component: () => import("../views/TransactionsCreate.vue"),
    meta: {
      requireLogin: true,
    },
  },
];

const router = new VueRouter({
  // mode: 'history',
  base: process.env.BASE_URL,
  routes,
});

export default router;
