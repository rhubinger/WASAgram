<script>
import store from '../services/store.js'

export default {
	data() {
		return {
			postId: null,
			comments: null,
		}
	},
	methods: {
		async postComment() {
			try {
				let post = await this.$axios.post("/posts/" + this.$route.params.pid + "/comments", {
					commentId: null,
					postId: this.postId,
					userId: store.userId,
					username: null,
					datetime: null,
					comment: document.getElementById("commentInput").value,

				}, { headers: {
				'Authorization': `Bearer ${store.identifier}`,
				},
			});
			document.getElementById("commentInput").reset(); 
			} catch (e) {
				console.log(e.toString());
			}
		}
	},

	async created() {
		try {
			let response = await this.$axios.get("/posts/" + this.$route.params.pid + "/comments", { headers: {
				'Authorization': `Bearer ${store.identifier}`,
				},
			});
			this.comments = response.data.comments;
			console.log(response.data);
		} catch (e) {
			console.log(e.toString());
		}
    }
}
</script>

<template>
	<div>
		<div
			class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
			<h1 class="h2">Post</h1>
			<div class="btn-toolbar mb-2 mb-md-0">
				<div class="btn-group me-2">
					<button type="button" class="btn btn-sm btn-outline-secondary" @click="refresh">
						Refresh
					</button>
					<button type="button" class="btn btn-sm btn-outline-secondary" @click="exportList">
						Export
					</button>
				</div>
				<div class="btn-group me-2">
					<button type="button" class="btn btn-sm btn-outline-primary" @click="newItem">
						New
					</button>
				</div>
			</div>
		</div>

		<div>
			<Post :pid="$route.params.pid" />
		</div>
		<div>
			<div v-for="comment in this.comments">
				<Comment :cid="comment.commentId" :pid="comment.postId" :uid="comment.userId" :username="comment.username" :datetime="comment['date-time']" :comment="comment.comment"/>
			</div>
		</div>
		<div>
			<form id="commentForm" onsubmit="return false">
				<label for="commentInput">Comment here:</label><br>
				<input type="text" id="commentInput" name="commentInput" value="Nice Picture!"><br>
				<input type="submit" value="comment" @click="postComment()">
			</form>
		</div>
	</div>
</template>

<style>
</style>