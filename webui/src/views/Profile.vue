<script>
import axios from 'axios';
import store from '../services/store.js';
import dateTime from '../services/datetime.js';

export default {
	data() {
		return {
			ownprofile: false,
			followed: false,
			user: null,
			posts: [],
		}
	},

	methods: {
		async followBtnHandler(){
			// Not followed by user
			if(!this.followed) {
				try {
					let response = await this.$axios.put("/users/" + this.$route.params.uid + "/followers/" + store.userId, {}, { headers: {
						'Authorization': `Bearer ${store.identifier}` ,
						},
					});
				} catch (e) {
					console.log(e.toString());
				}
			} else {
				try {
					let response = await this.$axios.delete("/users/" + this.$route.params.uid + "/followers/" + store.userId, { headers: {
						'Authorization': `Bearer ${store.identifier}` ,
						},
					});
				} catch (e) {
					console.log(e.toString());
				}
			}
			this.followed = !this.followed;
		},
	},

	async created() {
		if(this.$route.params.uid === store.userId){
			this.ownprofile = true;
		} else {
			this.ownprofile = false;
		}

		try {
			let response = await this.$axios.get("/users/" + this.$route.params.uid + "/isFollowedBy/" + store.userId, { headers: {
				'Authorization': `Bearer ${store.identifier}` ,
				},
			});
			this.followed = response.data.isFollowed;
		} catch (e) {
			console.log(e.toString());
		}

		try {
			let response = await this.$axios.get("/users/" + this.$route.params.uid, { headers: {
				'Authorization': `Bearer ${store.identifier}` ,
				},
			});
			this.user = response.data;
		} catch (e) {
			console.log(e.toString());
		}

		try {
			let response = await this.$axios.get("/users/" + this.$route.params.uid + "/posts?dateTime=" + dateTime.now(), { headers: {
				'Authorization': `Bearer ${store.identifier}` ,
				},
			});
			this.posts = response.data;
		} catch (e) {
			console.log(e.toString());
		}
	},
}
</script>

<template>
	<div>
		<div class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
			<h1 class="h2" v-if="ownprofile">Your profile</h1>
			<h1 class="h2" v-else>{{this.user.name}}'s profile ({{this.$route.params.uid}})</h1> 
			<h5> followed: {{ this.user.followed }} | followers: {{ this.user.followers }} | posts: {{ this.user.posts }}</h5> 
			<div class="btn-group me-2">
				<RouterLink v-if="ownprofile" to="/posts/create" class="nav-link">
					<button type="button" class="btn btn-sm btn-outline-primary">
						New Post
					</button>
				</RouterLink>
				<button v-else type="button" class="btn btn-sm btn-outline-primary" @click="followBtnHandler()">
					<div v-if="!this.followed"> Follow </div>
					<div v-else> Followed </div>
				</button>
			</div>
		</div>
		<div class="grid-container">
			<div class="grid-item" v-for="post in posts">
				<Post :pid="post.postId" :allowDelete="true" />
			</div>
		</div>
	</div>
</template>

<style>
	.grid-container {
	display: grid;
	grid-template-columns: auto auto auto;
	padding: 10px;
	}
	.grid-item {
	padding: 20px;
	}
</style>