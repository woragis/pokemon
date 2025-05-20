# Features

## Game Guide

1.  Search guides by keyword or tag
2.  Featured guides on homepage
3.  Markdown support (rendered content)
4.  Admin-only create/edit permissions
5.  User comments or reactions to guides
6.  Save/bookmark guides

### Main

- Markdown preview
- Search
- Admin only protection

## Shouts

What Twitter has that your shout feature might still be missing:

1. Rich Media Support

   Twitter lets users attach images, videos, GIFs, polls, etc.

   Your current model only supports text content (up to 280 chars).

2. User Mentions & Hashtags

   Twitter parses and links @mentions and #hashtags in tweets.

   Enables notifications, trending topics, and search by hashtags.

   You don’t yet have parsing or storing of mentions or hashtags.

3. Notifications

   Likes, retweets, mentions trigger notifications.

   Your system currently doesn’t handle notifications.

4. Direct Messages

   Twitter has a private messaging system.

   Your feature is public-only so far.

5. Retweet vs Quote Retweet distinction

   You implemented reshouting + quote content, so you have the basic concept.

   Twitter tracks metrics separately for retweets vs quote retweets.

6. Tweet Threads

   Twitter lets users chain tweets into threads.

   You’d want a way to link shouts in sequences (like a parent/child relationship).

7. Tweet Visibility & Privacy

   Twitter supports protected accounts, block/mute, follower-only content.

   Your system has no privacy or permission controls yet.

8. Analytics & Metrics

   Twitter tracks impressions, engagement rates, etc.

   You currently only track likes and comments count via relations.

9. Edit Tweet

   Twitter recently added editing capability.

   Your system doesn’t support editing yet.

10. Spam / Abuse Detection

    Twitter uses rate limiting, spam filters, and content moderation.

    Your system has no rate limiting or moderation.

11. Bookmarking & Saving

    Twitter allows users to save tweets for later.

    Not present in your feature.

12. Advanced Search

    Twitter has a rich search system across all tweets.

    You currently only support simple feed/timeline queries.
