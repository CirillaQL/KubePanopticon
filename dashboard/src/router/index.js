import { createWebHistory, createRouter } from "vue-router";

const routes = [
    {
        path: "/about",
        name: "About",
        component: () => import("@/pages/About/index.vue")
    },
    {
        path: "/nodes",
        name: "nodes",
        component: () => import("@/pages/Node/index.vue")
    }
]

const router = createRouter({
    history: createWebHistory(),
    routes,
});

export default router;