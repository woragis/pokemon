# Redis Key Schema – Pokémon Social Platform

This document outlines the Redis key structure used across various platform features for caching, fast access, and user interactions.

---

## 🔑 General Format

| Key Pattern                   | Type   | Description                       |
| ----------------------------- | ------ | --------------------------------- |
| `<entity>:<id>:<field>`       | String | General entity fields             |
| `<entity>:<id>:<field>:<sub>` | String | Nested field (e.g. views by date) |
| `comment:<id>`                | Hash   | Stores individual comment data    |
| `comment:<id>:replies`        | List   | Stores reply comment IDs          |
| `user:<id>:followers`         | Set    | User follower IDs                 |
| `user:<id>:following`         | Set    | Users followed by the user        |

---

## 👥 Users

| Key                   | Type | Description              | TTL        |
| --------------------- | ---- | ------------------------ | ---------- |
| `user:<id>`           | Hash | Cached user profile info | 30 mins    |
| `user:<id>:followers` | Set  | IDs of followers         | Persistent |
| `user:<id>:following` | Set  | IDs of followed users    | Persistent |
| `user:<id>:comments`  | List | Comment IDs on profile   | Optional   |

---

## 🛡 Teams

| Key | Type | Description | TTL |
| --- | --- | --- | --- |
| `team:<id>` | Hash | Team data (cached) | 10 mins |
| `team:<id>:likes` | Set | IDs of users who liked | Persistent |
| `team:<id>:comments` | List | Comment IDs | Persistent |
| `team:<id>:views` | String | Total view counter | Optional |
| `team:<id>:views:<yyyy-mm-dd>` | String | Views per day | Optional |

---

## 📸 Snapdex

| Key                  | Type   | Description           | TTL        |
| -------------------- | ------ | --------------------- | ---------- |
| `snap:<id>`          | Hash   | Snap metadata         | 10 mins    |
| `snap:<id>:likes`    | Set    | User IDs who liked it | Persistent |
| `snap:<id>:comments` | List   | Comment IDs           | Persistent |
| `snap:<id>:views`    | String | View counter          | Optional   |

---

## 🐦 Shouts (like Twitter)

| Key                   | Type   | Description            | TTL        |
| --------------------- | ------ | ---------------------- | ---------- |
| `shout:<id>`          | Hash   | Shout content          | 5 mins     |
| `shout:<id>:likes`    | Set    | User IDs who liked it  | Persistent |
| `shout:<id>:retweets` | Set    | User IDs who retweeted | Persistent |
| `shout:<id>:comments` | List   | Comment IDs            | Persistent |
| `shout:<id>:views`    | String | View counter           | Optional   |

---

## 📝 Blog Posts

| Key                  | Type | Description        | TTL        |
| -------------------- | ---- | ------------------ | ---------- |
| `blog:<id>`          | Hash | Blog post metadata | 15 mins    |
| `blog:<id>:likes`    | Set  | Likes              | Persistent |
| `blog:<id>:comments` | List | Comments           | Persistent |

---

## 💬 Forum

| Key                    | Type | Description     | TTL        |
| ---------------------- | ---- | --------------- | ---------- |
| `thread:<id>`          | Hash | Thread metadata | 10 mins    |
| `thread:<id>:likes`    | Set  | Likes           | Persistent |
| `thread:<id>:comments` | List | Comments        | Persistent |

---

## 📰 News

| Key                            | Type   | Description           | TTL        |
| ------------------------------ | ------ | --------------------- | ---------- |
| `news:<id>`                    | Hash   | News article metadata | 15 mins    |
| `news:<id>:likes`              | Set    | Likes                 | Persistent |
| `news:<id>:comments`           | List   | Comments              | Persistent |
| `news:<id>:views`              | String | View counter          | Optional   |
| `news:<id>:views:<yyyy-mm-dd>` | String | Daily view counter    | Optional   |

---

## 💬 Comments & Replies

| Key                    | Type | Description         | TTL        |
| ---------------------- | ---- | ------------------- | ---------- |
| `comment:<id>`         | Hash | Single comment data | Persistent |
| `comment:<id>:replies` | List | Nested replies      | Persistent |

---

## 🧠 Optional: Viewers (Optional Per-Day Unique)

If you want to track **unique views per user per day**, you can use:

| Key | Type | Description |
| --- | --- | --- |
| `viewed:<user_id>:<entity>:<id>` | String | Set with `EXPIRE 86400` (1 day TTL) |

This helps prevent counting duplicate views in a 24-hour period.

---

## 🔁 Background Sync (Queue)

You may also define a Redis **queue** to sync metrics or write-batched data to PostgreSQL:

| Key                   | Type | Description                     |
| --------------------- | ---- | ------------------------------- |
| `queue:sync:likes`    | List | Push batched like changes       |
| `queue:sync:views`    | List | Push batched view count updates |
| `queue:sync:comments` | List | Push batched comments           |

---
