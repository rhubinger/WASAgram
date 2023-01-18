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
        openPost() {
			this.$router.push('/posts/' + this.pid);
        },

        getLikes(){
			this.$router.push('/posts/' + this.pid + '/likes');
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
            this.deleted = true;
			this.$router.push('/profile/' + store.userId);
        },
        async likeBtnHandler(){
            // Not liked by user
            if(!this.liked) {
                try {
                    let response = await this.$axios.put("/posts/" + this.pid + "/likes/" + store.userId, {}, { headers: {
                        'Authorization': `Bearer ${store.identifier}` ,
                        },
                    });
                } catch (e) {
                    console.log(e.toString());
                }
                this.likes++;
            } else {
                try {
                    let response = await this.$axios.delete("/posts/" + this.pid + "/likes/" + store.userId, { headers: {
                        'Authorization': `Bearer ${store.identifier}` ,
                        },
                    });
                } catch (e) {
                    console.log(e.toString());
                }
                this.likes--;
            }
            this.liked = !this.liked;
        },
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
            <div> <b>{{username}}</b> </div>
            <div> {{uid}}</div>
        </div>
        <div>
            <button class="btn btn-sm btn-danger" v-on:click.stop="deletePost()">
                delete
            </button>
        </div>
        <div>
            <img class="picture" :src="this.pictureBlob"/><br>
        </div>
        <div>
            <button class="btn btn-sm btn-primary" v-on:click.stop="likeBtnHandler()">
                <div v-if="!this.liked"> Like </div>
                <div v-else> Remove Like </div>
            </button> 
            <a v-on:click.stop="getLikes()">
                #Likes: {{ likes }}
            </a> 
            #Comments: {{ comments }}
        </div>
        <div>
            <div class="caption"> {{caption}} </div>
            <div> {{datetime}} </div>
        </div>
    </div>
    <div v-else class=deleted>
        Post deleted
    </div>
</template>

<style>
.post{
    border-color: rgb(20, 120, 212);
    border-style: solid;
    border-radius: 8px;
    text-align: center;
}
.picture {
	border-radius: 8px;
    max-width: 95%;
    max-height: 450px * 1,3;
	min-width: 0;
	min-height: 0;
}
.deleted{
    background-color: rgb(230, 120, 120);
    border-radius: 8px;
}
.caption{
    margin: 10px
}

</style>