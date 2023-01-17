<script>
import store from '../services/store';

export default {
    props: ['uid', 'username', 'posts', 'followed', 'followers'],

	data: function() {
		return {
            followed: false,
            banned: false,
		}
	},
	methods: {
        openProfile() {
            this.$router.push('/profile/' + this.uid);
        },

		async followBtnHandler(){
            // Not followed by user
            if(!this.followed) {
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
            this.followed = !this.followed;
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
            {{ this.username }} - {{ this.uid }}
        </div>
        <div>
            Posts: {{ this.posts }} | Followers: {{ this.followed }} | Followed: {{ this.followers }}
        </div>
        <div>
            <button type="button" v-on:click.stop="followBtnHandler()"> 
                <div v-if="!this.followed"> Follow </div>
				<div v-else> Followed </div>
            </button>
            <button type="button" v-on:click.stop="banBtnHandler()"> 
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
</style>