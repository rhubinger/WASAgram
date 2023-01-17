<script>
import store from '../services/store.js'
export default {
    props: ['pid', 'allowDelete'],

	data: function() {
		return {
            uid: null,
            username: null,
            datetime: null,
            caption: null,
            pictureId: null,
            pictureBlob: null,
            likes: 0,
            liked: false,
            comments: 0,
            deleted: false,
		}
	},
	methods: {
		async like() {
            if (!this.liked) {
			    this.likes++;
                this.liked = true;
            }
            else {
                this.likes--;
                this.liked = false;
            }
		},

        openPost() {
			this.$router.push('/posts/' + this.pid);
        },

        async deletePost() {
            if(!confirm("Are you sure that you want to delete this post?")){
                return
            }
			try {
				let response = await this.$axios.delete("/posts/" + this.pid, { headers: {
					'Authorization': `Bearer ${store.identifier}` ,
					},
				});
			} catch (e) {
				console.log(e.toString());
			}
			this.$router.push('/profile/' + store.userId);
            this.deleted = true;
        }
	},

    async created() {
		try {
			let post = await this.$axios.get("/posts/" + this.pid, { headers: {
				'Authorization': `Bearer ${store.identifier}`,
				},
			});
			this.uid = post.data["userId"];
            this.username = post.data["username"];
            this.datetime = post.data["date-time"];
            this.caption = post.data["caption"];
            this.pictureId = post.data["pictureId"];
            this.likes = post.data["likes"];
            this.comments = post.data["comments"];
		} catch (e) {
			console.log(e.toString());
		}

        try {
            let picture = await this.$axios.get("/pictures/" + this.pictureId, { headers: {
                'Authorization': `Bearer ${store.identifier}`,
                'content-type': `image/png`,
                },
                responseType: 'blob',
            });
            this.pictureBlob = URL.createObjectURL(picture.data);
        } catch (e) {
            console.log(e.toString());
	    }
    }
}
</script>

<template>
	<div v-if="!deleted" class="post" @click="openPost()">
        <div>
            <div> {{username}} </div>
            <div> {{uid}}</div>
        </div>
        <div>
            <button v-on:click.stop="deletePost()">
                delete
            </button>
        </div>
        <div>
            <img class="picture" :src="this.pictureBlob"/><br>
        </div>
        <div>
            <button @click="like()">
                Like ({{ likes }})
            </button> 
            #Comments: {{ comments }}
        </div>
        <div>
            <div> {{datetime}} </div>
            <div> {{caption}} </div>
        </div>
    </div>
    <div v-else>
        Post deleted
    </div>
</template>

<style>
.post{
    background-color: lightgrey;
    border-radius: 8px;
}
.picture {
	border-radius: 8px;
    width: 80%;
    height: auto;
    margin-left: auto;
    margin-right: auto;
}
</style>