<script setup>
import { RouterLink, RouterView } from 'vue-router'
import store from './services/store.js';
</script>
<script>
export default {
	data() {
		return {
			loggedIn: false,
		}
	},

	methods: {
		async login() {
			try {
				uid = document.getElementById("uid").value;
				let response = await this.$axios.post("/login", {
					userId: uid
				});
				store.identifier = response.data.identifier;
				store.userId = uid;
				this.loggedIn = true;
				document.getElementById("loginForm").reset();
				//this.$router.push('/stream/' + uid);
			} catch (e) {
				console.log(e.toString());
			}
		},

		goToLogin(){
			this.loggedIn = false;
		}
	}
}
</script>

<template>
	<div v-if="this.loggedIn">
		<header class="navbar navbar-dark sticky-top bg-dark flex-md-nowrap p-0 shadow">
			<a class="navbar-brand col-md-3 col-lg-2 me-0 px-3 fs-6" @click="goToLogin()">WASAgram <br> Logged in as: {{ store.userId }}</a>
			<button class="navbar-toggler position-absolute d-md-none collapsed" type="button" data-bs-toggle="collapse" data-bs-target="#sidebarMenu" aria-controls="sidebarMenu" aria-expanded="false" aria-label="Toggle navigation">
				<span class="navbar-toggler-icon"></span>
			</button>
		</header>

		<div class="container-fluid">
			<div class="row">
				<nav id="sidebarMenu" class="col-md-3 col-lg-2 d-md-block bg-light sidebar collapse">
					<div class="position-sticky pt-3 sidebar-sticky">
						<h6 class="sidebar-heading d-flex justify-content-between align-items-center px-3 mt-4 mb-1 text-muted text-uppercase">
							<span>General</span>
						</h6>
						<ul class="nav flex-column">
							<li class="nav-item">
								<RouterLink to="/search" class="nav-link">
									<svg class="feather"><use href="/feather-sprite-v4.29.0.svg#home"/></svg>
									Search
								</RouterLink>
							</li>
							<li class="nav-item">
								<RouterLink :to="'/profile/' + store.userId" class="nav-link">
									<svg class="feather"><use href="/feather-sprite-v4.29.0.svg#home"/></svg>
									Profile
								</RouterLink>
							</li>
							<li class="nav-item">
								<RouterLink :to="'/stream/' + store.userId" class="nav-link">
									<svg class="feather"><use href="/feather-sprite-v4.29.0.svg#layout"/></svg>
									Stream
								</RouterLink>
							</li>
						</ul>

						<h6 class="sidebar-heading d-flex justify-content-between align-items-center px-3 mt-4 mb-1 text-muted text-uppercase">
							<span>Users </span>
						</h6>
						<ul class="nav flex-column">
							<li class="nav-item">
								<RouterLink :to="'/followed/' + store.userId" class="nav-link">
									<svg class="feather"><use href="/feather-sprite-v4.29.0.svg#file-text"/></svg>
									Followed
								</RouterLink>
							</li>
							<li class="nav-item">
								<RouterLink :to="'/followers/' + store.userId" class="nav-link">
									<svg class="feather"><use href="/feather-sprite-v4.29.0.svg#file-text"/></svg>
									Followers
								</RouterLink>
							</li>
							<li class="nav-item">
								<RouterLink :to="'/banned/' + store.userId" class="nav-link">
									<svg class="feather"><use href="/feather-sprite-v4.29.0.svg#file-text"/></svg>
									Banned
								</RouterLink>
							</li>
						</ul>
						<h6 class="sidebar-heading d-flex justify-content-between align-items-center px-3 mt-4 mb-1 text-muted text-uppercase">
							<span>Settings </span>
						</h6>
						<ul class="nav flex-column">
							<li class="nav-item">
								<RouterLink :to="'/changeUsername/' + store.userId" class="nav-link">
									<svg class="feather"><use href="/feather-sprite-v4.29.0.svg#file-text"/></svg>
									Change Username
								</RouterLink>
							</li>
						</ul>
					</div>
				</nav>

				<main class="col-md-9 ms-sm-auto col-lg-10 px-md-4">
					<RouterView />
				</main>
			</div>
		</div>
	</div>
	<div v-else>
        <h2>WASAgram Login</h2>
		<form id="loginForm" onsubmit="return false">
			<label for="User ID">User ID:</label><br>
			<input type="text" id="uid" name="UID" value="@Alan_Turing"><br>
			<input type="submit" value="login" @click="login">
		</form>
	</div>
</template>

<style>
</style>
