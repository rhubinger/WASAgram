# WASAgram

This repository contains the code of the WASAgram webapp. The application allows users to share images much like Instagram. The WASAgram webapp was created as part of the Web And Software Architecture course (WASA) course at Sapienza Università di Roma. This repository is based on the [Fantastic coffee (decaffeinated)](https://github.com/sapienzaapps/fantastic-coffee-decaffeinated).

## Functional design specifications 

Each user will be presented with a stream of photos (images) in reverse chronological order, with
information about when each photo was uploaded (date and time) and how many likes and comments
it has. The stream is composed by photos from “following” (other users that the user follows). Users
can place (and later remove) a “like” to photos from other users. Also, users can add comments to any
image (even those uploaded by themself). Only authors can remove their comments.
Users can ban other users. If user Alice bans user Eve, Eve won’t be able to see any information about
Alice. Alice can decide to remove the ban at any moment.
Users will have their profiles. The personal profile page for the user shows: the user’s photos (in reverse
chronological order), how many photos have been uploaded, and the user’s followers and following.
Users can change their usernames, upload photos, remove photos, and follow/unfollow other users.
Removal of an image will also remove likes and comments.
A user can search other user profiles via username.
A user can log in just by specifying the username. See the “Simplified login” section for details.

Note that these specifications are copied from the WASA project description by Emanuele Panizzi and Enrico Bassetti. This application only implements a dummy bearer authentication as the projects focus was on developing a small webapplication withot focussing too much on the security side.

## How to build

```shell
go build ./cmd/webapi/
```

## License

See [LICENSE](LICENSE).
