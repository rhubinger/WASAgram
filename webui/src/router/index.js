import {createRouter, createWebHashHistory} from 'vue-router'
import Login from '../views/Login.vue'
import Profile from '../views/Profile.vue'
import Stream from '../views/Stream.vue'
import Post from '../views/Post.vue'
import Followed from '../views/Followed.vue'
import Followers from '../views/Followers.vue'
import Banned from '../views/Banned.vue'
import CreatePost from '../views/CreatePost.vue'
import Search from '../views/Search.vue'

const router = createRouter({
	history: createWebHashHistory(import.meta.env.BASE_URL),
	routes: [
		{path: '/', component: Login},
		{path: '/search', component: Search},
		{path: '/profile/:uid', component: Profile},
		{path: '/stream/:uid', component: Stream},

		{path: '/followers/:uid', component: Followed},
		{path: '/followed/:uid', component: Followers},
		{path: '/banned/:uid', component: Banned},

		{path: '/posts/create', component: CreatePost},
		{path: '/posts/:pid', component: Post},
	]
})

export default router
