import Vue from 'vue'
import Router from 'vue-router'

import Signup from '@/components/Signup'
import Signin from '@/components/Signin'
import HelloWorld from '@/components/HelloWorld'

import firebase from 'firebase'

Vue.use(Router)

const router = new Router({
    routes: [
        {
            path: '*',
            redirect: 'signin'
        },
        {
            name: 'HelloWorld',
            component: HelloWorld,
            path: '/',
            meta: {
                requiresAuth: true
            }
        },
        {
            name: 'signup',
            component: Signup,
            path: '/signup'
        },
        {
            name: 'signin',
            component: Signin,
            path: '/signin'
        }
    ]
})

router.beforeEach((to, from, next) => {
    let currentUser = firebase.auth().currentUser
    let requiresAuth = to.matched.some(record => record.meta.requiresAuth)
    if (requiresAuth && !currentUser) next('signin')
    else if (!requiresAuth && currentUser) next()
    else next()
})

export default router
