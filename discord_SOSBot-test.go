package main
 
import "fmt"

// Creating structure
type discordRoom_name struct {
    name   string
    branch string
    areaID   int
}
 
// Creating nested structure
type discordRoom_area struct {
    name    string
    subject string
    roomID     int
    details discordRoom_name
}

func main() {
	
discordRoom_list := discordRoom_area{
        name: "Welcome Zone",
        subject:  "Doe",
        roomID:       30,
        details: discordRoom_name{
            name:     "Welcome Board",
            branch:       "Anytown",
            areaID: 12345,
        },
    }
    
    fmt.Println(discordRoom_list)
    fmt.Println("roomID:", discordRoom_list.roomID)
    
    }
    
    
    {
  "id": "80351110224678912",
  "username": "Nelly",
  "discriminator": "1337",
  "avatar": "8342729096ea3675442027381ff50dfe",
  "verified": true,
  "email": "nelly@discord.com",
  "flags": 64,
  "banner": "06c16474723fe537c283b8efa61a30c8",
  "accent_color": 16711680,
  "premium_type": 1,
  "public_flags": 64,
  "avatar_decoration_data": {
    "sku_id": "1144058844004233369",
    "asset": "a_fed43ab12698df65902ba06727e20c0e"
  }
}

{
  "op": 0,
  "d": {},
  "s": 42,
  "t": "GATEWAY_EVENT_NAME"
}

{
  "op": 2,
  "d": {
    "token": "my_token",
    "properties": {
      "os": "linux",
      "browser": "disco",
      "device": "disco"
    },
    "compress": true,
    "large_threshold": 250,
    "shard": [0, 1],
    "presence": {
      "activities": [{
        "name": "Cards Against Humanity",
        "type": 0
      }],
      "status": "dnd",
      "since": 91879201,
      "afk": false
    },
    // This intent represents 1 << 0 for GUILDS, 1 << 1 for GUILD_MEMBERS, and 1 << 2 for GUILD_BANS
    // This connection will only receive the events defined in those three intents
    "intents": 7
  }
}


0	Playing	Playing {name}	"Playing Rocket League"
1	Streaming	Streaming {details}	"Streaming Rocket League"
2	Listening	Listening to {name}	"Listening to Spotify"
3	Watching	Watching {name}	"Watching YouTube Together"
4	Custom	{emoji} {state}	":smiley: I am cool"
5	Competing	Competing in {name}	"Competing in Arena World Champions"

{
  "name": "Rocket League",
  "type": 0,
  "application_id": "379286085710381999",
  "state": "In a Match",
  "details": "Ranked Duos: 2-1",
  "timestamps": {
    "start": 15112000660000
  },
  "party": {
    "id": "9dd6594e-81b3-49f6-a6b5-a679e6a060d3",
    "size": [2, 2]
  },
  "assets": {
    "large_image": "351371005538729000",
    "large_text": "DFH Stadium",
    "small_image": "351371005538729111",
    "small_text": "Silver III"
  },
  "secrets": {
    "join": "025ed05c71f639de8bfaa0d679d7c94b2fdce12f",
    "spectate": "e7eb30d2ee025ed05c71ea495f770b76454ee4e0",
    "match": "4b2fdce12f639de8bfa7e3591b71a0d679d7c93f"
  }
}

user_id	snowflake	ID of the user
channel_id	snowflake	ID of the channel
message_id	snowflake	ID of the message
guild_id?	snowflake	ID of the guild
answer_id	integer	ID of the answer

Subscription status should not be used to grant perks. Use entitlements as an indication of whether a user should have access to a specific SKU. See our guide on Implementing App Subscriptions for more information.

{
  "details": "24H RL Stream for Charity",
  "state": "Rocket League",
  "name": "Twitch",
  "type": 1,
  "url": "https://www.twitch.tv/discord"
}


guild_id	snowflake	ID of the guild
role	a role object	Role that was created

guild_id	snowflake	ID of the guild
roles	array of snowflakes	User role ids
user	a user object	User
nick?	?string	Nickname of the user in the guild
avatar	?string	Member's guild avatar hash
joined_at	?ISO8601 timestamp	When the user joined the guild
premium_since?	?ISO8601 timestamp	When the user starting boosting the guild
deaf?	boolean	Whether the user is deafened in voice channels
mute?	boolean	Whether the user is muted in voice channels
pending?	boolean	Whether the user has not yet passed the guild's Membership Screening requirements
communication_disabled_until?	?ISO8601 timestamp	When the user's timeout will expire and the user will be able to communicate in the guild again, null or a time in the past if the user is not timed out
flags?	integer	Guild member flags represented as a bit set, defaults to 0
avatar_decoration_data?	?avatar decoration data	Data for the member's guild avatar decoration

Hello	Defines the heartbeat interval
Ready	Contains the initial state information
Resumed	Response to Resume
Reconnect	Server is going away, client should reconnect to gateway and resume
Invalid Session	Failure response to Identify or Resume or invalid active session
Application Command Permissions Update	Application command permission was updated
Auto Moderation Rule Create	Auto Moderation rule was created
Auto Moderation Rule Update	Auto Moderation rule was updated
Auto Moderation Rule Delete	Auto Moderation rule was deleted
Auto Moderation Action Execution	Auto Moderation rule was triggered and an action was executed (e.g. a message was blocked)
Channel Create	New guild channel created
Channel Update	Channel was updated
Channel Delete	Channel was deleted
Channel Pins Update	Message was pinned or unpinned
Thread Create	Thread created, also sent when being added to a private thread
Thread Update	Thread was updated
Thread Delete	Thread was deleted
Thread List Sync	Sent when gaining access to a channel, contains all active threads in that channel
Thread Member Update	Thread member for the current user was updated
Thread Members Update	Some user(s) were added to or removed from a thread
Entitlement Create	Entitlement was created
Entitlement Update	Entitlement was updated or renewed
Entitlement Delete	Entitlement was deleted
Guild Create	Lazy-load for unavailable guild, guild became available, or user joined a new guild
Guild Update	Guild was updated
Guild Delete	Guild became unavailable, or user left/was removed from a guild
Guild Audit Log Entry Create	A guild audit log entry was created
Guild Ban Add	User was banned from a guild
Guild Ban Remove	User was unbanned from a guild
Guild Emojis Update	Guild emojis were updated
Guild Stickers Update	Guild stickers were updated
Guild Integrations Update	Guild integration was updated
Guild Member Add	New user joined a guild
Guild Member Remove	User was removed from a guild
Guild Member Update	Guild member was updated
Guild Members Chunk	Response to Request Guild Members
Guild Role Create	Guild role was created
Guild Role Update	Guild role was updated
Guild Role Delete	Guild role was deleted
Guild Scheduled Event Create	Guild scheduled event was created
Guild Scheduled Event Update	Guild scheduled event was updated
Guild Scheduled Event Delete	Guild scheduled event was deleted
Guild Scheduled Event User Add	User subscribed to a guild scheduled event
Guild Scheduled Event User Remove	User unsubscribed from a guild scheduled event
Integration Create	Guild integration was created
Integration Update	Guild integration was updated
Integration Delete	Guild integration was deleted
Interaction Create	User used an interaction, such as an Application Command
Invite Create	Invite to a channel was created
Invite Delete	Invite to a channel was deleted
Message Create	Message was created
Message Update	Message was edited
Message Delete	Message was deleted
Message Delete Bulk	Multiple messages were deleted at once
Message Reaction Add	User reacted to a message
Message Reaction Remove	User removed a reaction from a message
Message Reaction Remove All	All reactions were explicitly removed from a message
Message Reaction Remove Emoji	All reactions for a given emoji were explicitly removed from a message
Presence Update	User was updated
Stage Instance Create	Stage instance was created
Stage Instance Update	Stage instance was updated
Stage Instance Delete	Stage instance was deleted or closed
Subscription Create	Premium App Subscription was created
Subscription Update	Premium App Subscription was updated
Subscription Delete	Premium App Subscription was deleted
Typing Start	User started typing in a channel
User Update	Properties about the user changed
Voice Channel Effect Send	Someone sent an effect in a voice channel the current user is connected to
Voice State Update	Someone joined, left, or moved a voice channel
Voice Server Update	Guild's voice server was updated
Webhooks Update	Guild channel webhook was created, update, or deleted
Message Poll Vote Add	User voted on a poll
Message Poll Vote Remove	User removed a vote on a poll

The HTTP API is a REST API that lets you interact and modify core Discord resources like channels, servers, users, and messages. The collection includes all of the public endpoints, organized by resource.
The Gateway API is a WebSocket API that lets you send and receive realtime events. The collection includes the main WebSocket events for connecting to the Gateway.

import "github.com/bwmarrin/discordgo"

The easiest way for an application to share media to a channel or DM is to use the openShareMomentDialog command. This command accepts a Discord CDN mediaUrl (eg https://cdn.discordapp.com/attachments/...) and opens a dialog on the discord client that allows the user to select channels, DMs, and GDMs to share to.

You can link directly to a specific SKU using our Application Directory Store URL scheme:

https://discord.com/application-directory/:appID/store/:skuID

When used in chat, it will render as a rich embed that allows users to launch a modal to view either the SKU details or checkout flow
When used as a direct URL in a browser, it will take the user to your product in the Application Directory on web

return new JsonResponse({
    type: 4, // InteractionResponseType.CHANNEL_MESSAGE_WITH_SOURCE
    data: {
        content: "This command requires Nelly Premium! Upgrade now to get access to these features!",
        components: [{
            type: MessageComponentTypes.ACTION_ROW,
            components: [
                {
                    type: MessageComponentTypes.BUTTON,
                    style: 6, // ButtonStyleTypes.PREMIUM
                    sku_id: '1234965026943668316',
                },
            ],
        }]
    },
});


When a user purchases a subscription, an entitlement is created. Entitlements represent the user's access to your consumable or durable item.

Depending on your app's features, you can use a combination of Gateway events, the Entitlement HTTP API, and interaction payloads to keep track of user and guild entitlements and grant perks to users who have subscribed to your app.

For apps requiring background processing or not solely reliant on interactions, keeping track of entitlements is essential. You can utilize the List Entitlements endpoint to list active and expired entitlements. Your app can filter entitlements by a specific user or guild by using the ?user_id=XYZ or ?guild_id=XYZ query params.

For example, you might keep track of our entitlements in a database and check if a user still has access to a specific SKU before performing a cron job or other task.

Responding with a premium button gives you the ability to prompt users to subscribe to your app when they attempt to use a premium feature without a subscription.

You can do this by sending a message with a button with a premium style and a sku_id that allows the user to upgrade to your premium offering.

You can use the Subscription API to check on the status of your app subscriptions. This API allows you to list all subscriptions for your app for reporting purposes and to check on the status of subscriptions without having to access entitlements directly.

List SKU Subscriptions: List all subscriptions for a specific SKU in your app.
Get SKU Subscription: Get a specific subscription in your app.
Subscription Gateway events: Discord will emit gateway events when a subscription is created, updated, and very rarely, deleted.

If you'd like to test the full payment flow for your app, you can do so by interacting with your Store page or a premium styled button. Any team members associated with your app will automatically see a 100% discount on the price of the subscription, allowing you to purchase without the use of live payment method.

After checkout, you will have a live subscription that includes a starts_at and ends_at value. If you cancel this subscription, it will remain an active entitlement until the ends_at timestamp. This subscription will renew until canceled and can be used in testing subscription renewals in your app.

{
  "id": "1278078770116427839", 
  "user_id": "1088605110638227537", 
  "sku_ids": ["1158857122189168803"], 
  "entitlement_ids": [], 
  "current_period_start": "2024-08-27T19:48:44.406602+00:00", 
  "current_period_end": "2024-09-27T19:48:44.406602+00:00", 
  "status": 0, 
  "canceled_at": null
}

