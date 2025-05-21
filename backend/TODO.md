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

# To implement later

1.  Achievements System

    Track user milestones like:

        Number of Pokémon caught

        Living Dex completion

        Number of shiny Pokémon

        Participation in chats/events

    Store achievements per user in Postgres.

    Display badges or trophies on user profiles.

2.  Friend System & Social Feed

    Allow users to:

        Add friends, remove friends

        See friend’s recent activity: catches, achievements, posts

    Social feed for posting short updates (text, images)

    Notifications for friend requests, posts, likes, comments

3.  Nature, Abilities & Moves Data in Pokedex

    Extend PokémonSpecies model:

        Nature list or default (maybe link to a nature table)

        Abilities (primary, secondary, hidden)

        Moves learned (by level-up, TM, egg moves)

    Show these in game guide and pokedex details

4.  Search & Filter

    Powerful filtering/search on:

        Pokémon species by name, type, region, generation

        Trainer or user search by username, achievements

        Feed and posts search by keywords or tags

5.  Mobile App / PWA

    Build a React Native or Flutter app or a PWA version of your frontend

    Support push notifications for:

        Friend activity

        Achievements unlocked

        New posts or comments

    Sync with backend API you’re building

6.  Localization & Multilanguage Support

    Structure all frontend text and content for easy translations

    Backend APIs support locale param (if applicable)

    Initially support English + maybe one or two others (like Portuguese, Spanish)

    Later add more languages as community grows

## Blog

Everything
