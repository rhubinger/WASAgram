<script>
import store from '../services/store';

export default {
    props: ['uid', 'username', 'posts', 'followed', 'followers'],

	data: function() {
		return {
            followedUser: false,
            banned: false,
		}
	},
	methods: {
        openProfile() {
            this.$router.push('/profile/' + this.uid);
        },

		async followBtnHandler(){
            // Not followed by user
            if(!this.followedUser) {
                try {
                    let response = await this.$axios.put("/users/" + this.uid + "/followers/" + store.userId, {}, { headers: {
                        'Authorization': `Bearer ${store.identifier}` ,
                        },
                    });
                } catch (e) {
                    console.log(e.toString());
                }
            } else {
                try {
                    let response = await this.$axios.delete("/users/" + this.uid + "/followers/" + store.userId, { headers: {
                        'Authorization': `Bearer ${store.identifier}` ,
                        },
                    });
                } catch (e) {
                    console.log(e.toString());
                }
            }
            this.followedUser = !this.followedUser;
        },
        async banBtnHandler(){
            // Not banned by user
            if(!this.banned) {
                if (!confirm("Are your sure that you want to ban this user?")) {
                    return;
                }
                try {
                    let response = await this.$axios.put("/users/" + store.userId + "/banned/" + this.uid, {}, { headers: {
                        'Authorization': `Bearer ${store.identifier}` ,
                        },
                    });
                } catch (e) {
                    console.log(e.toString());
                }
            } else {
                if (!confirm("Are your sure that you want to unban this user?")) {
                    return;
                }
                try {
                    let response = await this.$axios.delete("/users/" + store.userId + "/banned/" + this.uid, { headers: {
                        'Authorization': `Bearer ${store.identifier}` ,
                        },
                    });
                } catch (e) {
                    console.log(e.toString());
                }
            }
            this.banned = !this.banned;
        },
    },

    async created() {
        try {
			let response = await this.$axios.get("/users/" + this.uid + "/isFollowedBy/" + store.userId, { headers: {
				'Authorization': `Bearer ${store.identifier}` ,
				},
			});
			this.followed = response.data.exists;
		} catch (e) {
			console.log(e.toString());
		}

        try {
			let response = await this.$axios.get("/users/" + store.userId + "/hasBanned/" + this.uid, { headers: {
				'Authorization': `Bearer ${store.identifier}` ,
				},
			});
			this.banned = response.data.exists;
		} catch (e) {
			console.log(e.toString());
		}
    }
}
</script>

<template>
	<div class="user" @click="openProfile">
        <div>
            <p class="left"> <b>{{ this.username }}</b> - {{ this.uid }} </p>
            <p class="right"> Posts: {{ this.posts }} | Followers: {{ this.followed }} | Followed: {{ this.followers }} </p>
        </div>
        <div style="clear: both;"></div>
        <div>
            <button type="button" class="btn btn-sm btn-primary" v-on:click.stop="followBtnHandler()"> 
                <div v-if="!this.followedUser"> Follow </div>
				<div v-else> Followed </div>
            </button>
            <button type="button" class="btn btn-sm btn-danger" v-on:click.stop="banBtnHandler()"> 
                <div v-if="!this.banned"> Ban </div>
				<div v-else> Banned </div>
            </button>
        </div>
    </div>
</template>

<style>
.user{
    background-color: lightgrey;
    border-radius: 8px;
}
.left{
    float: left;
}
.right{
    float: right;
}
</style>